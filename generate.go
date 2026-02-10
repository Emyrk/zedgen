package zedgen

import (
	"fmt"
	"go/format"
	"regexp"
	"slices"
	"sort"
	"strings"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	core "github.com/authzed/spicedb/pkg/proto/core/v1"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/authzed/spicedb/pkg/schemadsl/generator"
	"github.com/authzed/spicedb/pkg/tuple"
)

type Options struct {
	Package        string
	SchemaFileName string
}

func Generate(schema string, opts Options) (string, error) {
	var prefix string // TODO: What is the prefix for?
	compiled, err := compiler.Compile(compiler.InputSchema{
		Source:       "policy.zed",
		SchemaString: schema,
	}, compiler.ObjectTypePrefix(prefix))
	if err != nil {
		return "", fmt.Errorf("compile schema: %w", err)
	}

	// sort order for consistency
	sort.Slice(compiled.ObjectDefinitions, func(i, j int) bool {
		return compiled.ObjectDefinitions[i].Name < compiled.ObjectDefinitions[j].Name
	})
	// TODO: Handle caveats
	sort.Slice(compiled.CaveatDefinitions, func(i, j int) bool {
		return compiled.CaveatDefinitions[i].Name < compiled.CaveatDefinitions[j].Name
	})

	tpl, err := LoadTemplates()
	if err != nil {
		return "", fmt.Errorf("load templates: %w", err)
	}

	var output strings.Builder
	// Write the file header first
	err = tpl.All.ExecuteTemplate(&output, "header", map[string]interface{}{
		"Package": opts.Package,
	})
	if err != nil {
		return "", fmt.Errorf("execute header template: %w", err)
	}

	// If you do not write this, there is no line break between them.
	output.WriteString("\n")

	definitions := make(map[string]objectDefinition, len(compiled.ObjectDefinitions))
	for _, obj := range compiled.ObjectDefinitions {
		d := newDef(obj)
		d.Filename = opts.SchemaFileName
		definitions[d.Name] = d
	}

	// Add optional relations
	for k := range definitions {
		// For each optional relation, lets add it to the other defintion
		for _, opt := range definitions[k].OptionalRelations {
			if includeOpt, ok := definitions[opt.For]; ok {
				index := slices.IndexFunc(includeOpt.IncludeRelations, func(i includedRelation) bool {
					return i.RelationName == opt.Relation
				})

				if index == -1 {
					includeOpt.IncludeRelations = append(includeOpt.IncludeRelations, includedRelation{
						RelationName: opt.Relation,
						From:         []string{opt.From},
					})
				} else {
					includeOpt.IncludeRelations[index].From = append(includeOpt.IncludeRelations[index].From, opt.From)
				}

				definitions[opt.For] = includeOpt
			}
		}
	}

	sorted := make([]objectDefinition, 0, len(definitions))
	for k := range definitions {
		sorted = append(sorted, definitions[k])
	}

	// Sort for a consistent output order
	slices.SortStableFunc(sorted, func(a, b objectDefinition) int {
		return strings.Compare(a.Name, b.Name)
	})

	for _, obj := range sorted {
		err := tpl.All.ExecuteTemplate(&output, "object", obj)
		if err != nil {
			return "", fmt.Errorf("[%q] execute template: %w", obj.Name, err)
		}
		output.WriteString("\n")
	}

	formatted, err := format.Source([]byte(output.String()))
	if err != nil {
		// Return the failed output for debugging.
		return output.String(), fmt.Errorf("format source: %w", err)
	}
	return string(formatted), nil
}

// objectDefinition is the type used for the template generator.
type objectDefinition struct {
	// The core type
	*core.NamespaceDefinition

	Filename        string
	DirectRelations []objectDirectRelation
	// RelationConstants are the names of the relations that are defined in the schema.
	// This just makes it easier to reference the relations in the code.
	RelationConstants map[string]struct{}
	Permissions       []objectPermission
	// UniqueSubjectTypes holds unique (objectType, optionalRelation) combinations
	// for generating permission check methods. This prevents duplicate methods when
	// multiple relations have the same subject type.
	UniqueSubjectTypes []subjectType
	// OptionalRelations keep track of which optional relations are being
	// referenced in this object definition. We can use this to make helpers
	// to add optional relations when using the object as a subject.
	//
	// OptionalRelations can reference other objects.
	OptionalRelations []optionalRelation
	// IncludeRelations are which to include in the generated file.
	IncludeRelations []includedRelation
}

// subjectType represents a unique subject type for permission checks.
type subjectType struct {
	ObjectType       string
	OptionalRelation string
}

type includedRelation struct {
	RelationName string
	From         []string
}

type optionalRelation struct {
	For      string
	Relation string
	From     string
}

type objectDirectRelation struct {
	LinePos      string
	Comment      string
	RelationName string
	FunctionName string
	Subject      v1.SubjectReference
}

type objectPermission struct {
	LinePos      string
	Comment      string
	Permission   string
	FunctionName string
}

var permissionSchema = regexp.MustCompile(`^[^{]*{\s*(.*)\s}$`)

func newDef(obj *core.NamespaceDefinition) objectDefinition {
	parsedObject := objectDefinition{
		NamespaceDefinition: obj,
		RelationConstants:   make(map[string]struct{}),
	}
	rels := make([]objectDirectRelation, 0)
	perms := make([]objectPermission, 0)
	optRels := make([]optionalRelation, 0)

	objectReference := &v1.ObjectReference{
		ObjectType: obj.Name,
		ObjectId:   "<id>",
	}

	// Each relation is a "relation" or "permission" line in the schema. Example:
	// - relation member: group#membership | user
	// - permission membership = member
	for _, r := range obj.Relation {
		linePos := fmt.Sprintf("%d", r.SourcePosition.ZeroIndexedLineNumber+1)
		// A UsersetRewrite is a permission. In spicedb, permissions are just
		// dynamically created relations.
		if r.UsersetRewrite != nil {
			comment := fmt.Sprintf("Object: %s", tuple.V1StringObjectRef(objectReference))
			// The UsersetRewrite is an expression shown as an AST here.
			// The code to parse the AST is all in the `/internal` pkg, and
			// writing our own AST traversal is not worth it just to place
			// a helpful comment.
			//
			// Instead, we use the 'GenerateSource' to effectively generate the
			// schema block. With a bit of regex, we can extract what we want.
			schema, _, _ := generator.GenerateSource(&core.NamespaceDefinition{
				Name:           obj.Name,
				Relation:       []*core.Relation{r},
				Metadata:       obj.Metadata,
				SourcePosition: obj.SourcePosition,
			}, nil)
			matches := permissionSchema.FindStringSubmatch(schema)
			if len(matches) >= 2 {
				comment += "\nSchema: " + matches[1]
			}

			// For our case, we only care about the permissions name.
			// If we decided to actually parse the AST, we could extract which
			// relations that the permission is mapping to.
			// With that information, we could add some sort of validation or
			// type safety to ensure subjects passed in are in fact valid subjects.
			//
			// Right now, you could pass in any object to the "CanX" function, and
			// it will try even though we know some objects will never have the permission.
			perms = append(perms, objectPermission{
				LinePos:      linePos,
				Comment:      strings.TrimSpace(comment),
				Permission:   r.Name,
				FunctionName: capitalize(r.Name),
			})
			continue
		}

		// Each "AllowedDirectRelations" is a "subject" on the right hand side
		// of a relation. So in the example above, the subjects are:
		//	- group#membership
		//	- user
		multipleSubjects := make([]objectDirectRelation, 0)
		for _, d := range r.TypeInformation.AllowedDirectRelations {
			optRel := ""
			// TODO: I cannot recall what this is
			if d.GetRelation() != "..." {
				optRel = d.GetRelation()
			}

			if optRel != "" {
				// An optional relation means we are not relating the parent "obj" to
				// the "r" directly.
				//
				// The common example is 'group#membership'.
				// If a relationship from "Obj A" -> "Group B" exists.
				// Then checking if "Group B" can read "Obj A" is a valid check, but it
				// will fail. Because the check **should be** if "Group B#membership" can read "Obj A".
				//
				// So in our autogenerated output, we have a "AsSubject()" function. We can do a bit better
				// and add a "AsAnyMembership()" to the "Group" Object.
				from := tuple.V1StringSubjectRef(&v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectType: obj.Name,
						ObjectId:   "<id>",
					},
					OptionalRelation: r.Name,
				})

				optRels = append(optRels, optionalRelation{
					For:      d.Namespace,
					Relation: d.GetRelation(),
					From:     from,
				})
			}

			subj := v1.SubjectReference{
				Object: &v1.ObjectReference{
					ObjectType: d.Namespace,
					// This ObjectID will be overwritten, so just use a placeholder.
					ObjectId: "<id>",
				},
				OptionalRelation: optRel,
			}

			if d.GetPublicWildcard() != nil {
				subj.Object.ObjectId = "*"
			}

			// Generate a comment above the function that is the string representation
			// of the relationship that the function will create.
			// The format is:
			//	<obj_typ>:<obj_id>#<relation>@<subj_typ>:<subj_id>
			comment, _ := tuple.V1StringRelationship(&v1.Relationship{
				Resource:       objectReference,
				Relation:       r.Name,
				Subject:        &subj,
				OptionalCaveat: nil,
			})
			parsedObject.RelationConstants[r.Name] = struct{}{}
			multipleSubjects = append(multipleSubjects, objectDirectRelation{
				LinePos:      linePos,
				Comment:      fmt.Sprintf("Relationship: %s", comment),
				RelationName: r.Name,
				FunctionName: r.Name,
				Subject:      subj,
			})
		}

		// If we have more than 1 subject, we need to suffix the name of the function with the
		// object type. Otherwise, we have 2 functions with the same name.
		//
		// Example:
		//	- group#member -- func Member_Group
		//	- user  	   -- func Member_User
		//
		// If we only have 1 subject, we do not need the suffix.
		// If someone could find a typesafe argument pattern with generics, that would be great.
		if len(multipleSubjects) > 1 {
			for i := range multipleSubjects {
				multipleSubjects[i].FunctionName += "_" + capitalize(multipleSubjects[i].Subject.Object.ObjectType)
			}
		}

		rels = append(rels, multipleSubjects...)
	}
	parsedObject.DirectRelations = rels
	parsedObject.Permissions = perms
	parsedObject.OptionalRelations = optRels

	// Collect unique subject types for permission check generation
	// This prevents duplicate CanX_User methods when multiple relations have user as subject
	seenSubjects := make(map[string]struct{})
	for _, rel := range rels {
		// Skip wildcards - they can't be subjects in checks
		if rel.Subject.Object.ObjectId == "*" {
			continue
		}
		key := rel.Subject.Object.ObjectType + "#" + rel.Subject.OptionalRelation
		if _, exists := seenSubjects[key]; !exists {
			seenSubjects[key] = struct{}{}
			parsedObject.UniqueSubjectTypes = append(parsedObject.UniqueSubjectTypes, subjectType{
				ObjectType:       rel.Subject.Object.ObjectType,
				OptionalRelation: rel.Subject.OptionalRelation,
			})
		}
	}

	return parsedObject
}

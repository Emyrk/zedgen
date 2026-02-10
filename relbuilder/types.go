package relbuilder

import (
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
)

// String is a convenience type that implements fmt.Stringer for string literals.
type String string

func (s String) String() string {
	return string(s)
}

func (b *Build) Object(ref *v1.ObjectReference, optRel string) Object {
	return Object{
		Obj:              ref,
		OptionalRelation: optRel,
		builder:          b,
	}
}

// Object represents a SpiceDB object reference with an optional relation.
type Object struct {
	Obj              *v1.ObjectReference
	OptionalRelation string
	builder          Builder
}

func (obj *Object) Touch() Relationship {
	return Relationship{
		Resource:  obj,
		Operation: v1.RelationshipUpdate_OPERATION_TOUCH,
		builder:   obj.builder,
	}
}

func (obj *Object) Delete() Relationship {
	return Relationship{
		Resource:  obj,
		Operation: v1.RelationshipUpdate_OPERATION_DELETE,
		builder:   obj.builder,
	}
}

func (obj *Object) Create() Relationship {
	return Relationship{
		Resource:  obj,
		Operation: v1.RelationshipUpdate_OPERATION_CREATE,
		builder:   obj.builder,
	}
}

// Relationship represents a pending relationship operation.
type Relationship struct {
	Resource  *Object
	Operation v1.RelationshipUpdate_Operation
	builder   Builder
}

// Add adds a relationship with the given relation name, subject object, and optional subject relation.
// This is the core method used by generated code to build relationships.
func (r Relationship) Add(relation string, subjectObj *v1.ObjectReference, subjectRelation string) {
	r.builder.Add(r.Operation, &v1.Relationship{
		Resource: r.Resource.Obj,
		Relation: relation,
		Subject: &v1.SubjectReference{
			Object:           subjectObj,
			OptionalRelation: subjectRelation,
		},
	})
}

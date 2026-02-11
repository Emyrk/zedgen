// Package policy code generated. DO NOT EDIT.
package policy

import (
	"fmt"

	. "github.com/Emyrk/zedgen/relbuilder"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/gochugaru/rel"
)

// SchemaBuilder is the entry point for building relationships and permission checks.
// It embeds relbuilder.Build for access to Updates() and Preconditions().
type SchemaBuilder struct {
	*Build
}

// New creates a new SchemaBuilder instance.
func New() *SchemaBuilder {
	return &SchemaBuilder{
		Build: NewBuild(),
	}
}

type ObjOther struct {
	src Object
}

func (b *SchemaBuilder) Other(id fmt.Stringer) *ObjOther {
	return &ObjOther{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "other",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjOther) Object() rel.Object {
	return obj.src.Object()
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjOther) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

type ObjPerson struct {
	src Object
}

func (b *SchemaBuilder) Person(id fmt.Stringer) *ObjPerson {
	return &ObjPerson{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "person",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjPerson) Object() rel.Object {
	return obj.src.Object()
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjPerson) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjPerson) RelationTest() string {
	return "test"
}

func (obj *ObjPerson) RelationUser() string {
	return "user"
}

func (obj *ObjPerson) PermissionRead() string {
	return "read"
}

type PersonRelates struct {
	obj *ObjPerson
	rel Relationship
}

func (obj *ObjPerson) Touch() *PersonRelates {
	return &PersonRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjPerson) Delete() *PersonRelates {
	return &PersonRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjPerson) Create() *PersonRelates {
	return &PersonRelates{obj: obj, rel: obj.src.Create()}
}

// User_User dupename.zed:8
// Relationship: person:<id>#user@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().User_User() etc.
func (obj *ObjPerson) User_User(subs ...*ObjUser) *ObjPerson {
	for _, sub := range subs {
		obj.src.Touch().Add("user", sub.src.Obj, "")
	}
	return obj
}

// User_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *PersonRelates) User_User(subs ...*ObjUser) *PersonRelates {
	for _, sub := range subs {
		r.rel.Add("user", sub.src.Obj, "")
	}
	return r
}

// User_Other dupename.zed:8
// Relationship: person:<id>#user@other:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().User_Other() etc.
func (obj *ObjPerson) User_Other(subs ...*ObjOther) *ObjPerson {
	for _, sub := range subs {
		obj.src.Touch().Add("user", sub.src.Obj, "")
	}
	return obj
}

// User_Other on Relates uses the specified operation (Touch/Create/Delete)
func (r *PersonRelates) User_Other(subs ...*ObjOther) *PersonRelates {
	for _, sub := range subs {
		r.rel.Add("user", sub.src.Obj, "")
	}
	return r
}

// Test dupename.zed:9
// Relationship: person:<id>#test@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Test() etc.
func (obj *ObjPerson) Test(subs ...*ObjUser) *ObjPerson {
	for _, sub := range subs {
		obj.src.Touch().Add("test", sub.src.Obj, "")
	}
	return obj
}

// Test on Relates uses the specified operation (Touch/Create/Delete)
func (r *PersonRelates) Test(subs ...*ObjUser) *PersonRelates {
	for _, sub := range subs {
		r.rel.Add("test", sub.src.Obj, "")
	}
	return r
}

// CanRead_User checks if the subject has read permission
// // Object: person:<id>
// Schema: permission read = user
func (obj *ObjPerson) CanRead_User(sub *ObjUser) rel.Relationship {
	r, s := obj.src.Obj, sub.src
	return rel.Relationship{
		ResourceType:     r.ObjectType,
		ResourceID:       r.ObjectId,
		ResourceRelation: "read",
		SubjectType:      s.Obj.ObjectType,
		SubjectID:        s.Obj.ObjectId,
		SubjectRelation:  s.OptionalRelation,
	}
}

// CanRead_Other checks if the subject has read permission
// // Object: person:<id>
// Schema: permission read = user
func (obj *ObjPerson) CanRead_Other(sub *ObjOther) rel.Relationship {
	r, s := obj.src.Obj, sub.src
	return rel.Relationship{
		ResourceType:     r.ObjectType,
		ResourceID:       r.ObjectId,
		ResourceRelation: "read",
		SubjectType:      s.Obj.ObjectType,
		SubjectID:        s.Obj.ObjectId,
		SubjectRelation:  s.OptionalRelation,
	}
}

type ObjUser struct {
	src Object
}

func (b *SchemaBuilder) User(id fmt.Stringer) *ObjUser {
	return &ObjUser{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "user",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjUser) Object() rel.Object {
	return obj.src.Object()
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjUser) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

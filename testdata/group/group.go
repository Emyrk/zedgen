// Package policy code generated. DO NOT EDIT.
package policy

import (
	"fmt"

	//nolint:staticcheck
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

type ObjGroup struct {
	src Object
}

func (b *SchemaBuilder) Group(id fmt.Stringer) *ObjGroup {
	return &ObjGroup{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "group",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjGroup) Object() rel.Object {
	return obj.src.Object()
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjGroup) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjGroup) RelationMember() string {
	return "member"
}

func (obj *ObjGroup) PermissionMembership() string {
	return "membership"
}

type GroupRelates struct {
	obj *ObjGroup
	rel Relationship
}

func (obj *ObjGroup) Touch() *GroupRelates {
	return &GroupRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjGroup) Delete() *GroupRelates {
	return &GroupRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjGroup) Create() *GroupRelates {
	return &GroupRelates{obj: obj, rel: obj.src.Create()}
}

// Member_User group.zed:4
// Relationship: group:<id>#member@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Member_User() etc.
func (obj *ObjGroup) Member_User(subs ...*ObjUser) *ObjGroup {
	for _, sub := range subs {
		obj.src.Touch().Add("member", sub.src.Obj, "")
	}
	return obj
}

// Member_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *GroupRelates) Member_User(subs ...*ObjUser) *GroupRelates {
	for _, sub := range subs {
		r.rel.Add("member", sub.src.Obj, "")
	}
	return r
}

// Member_Group group.zed:4
// Relationship: group:<id>#member@group:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Member_Group() etc.
func (obj *ObjGroup) Member_Group(subs ...*ObjGroup) *ObjGroup {
	for _, sub := range subs {
		obj.src.Touch().Add("member", sub.src.Obj, "member")
	}
	return obj
}

// Member_Group on Relates uses the specified operation (Touch/Create/Delete)
func (r *GroupRelates) Member_Group(subs ...*ObjGroup) *GroupRelates {
	for _, sub := range subs {
		r.rel.Add("member", sub.src.Obj, "member")
	}
	return r
}

// MemberWildcard group.zed:4
// Relationship: group:<id>#member@user:*
func (obj *ObjGroup) MemberWildcard() *ObjGroup {
	obj.src.Touch().Add("member", &v1.ObjectReference{
		ObjectType: "user",
		ObjectId:   "*",
	}, "")
	return obj
}

// MemberWildcard on Relates uses the specified operation
func (r *GroupRelates) MemberWildcard() *GroupRelates {
	r.rel.Add("member", &v1.ObjectReference{
		ObjectType: "user",
		ObjectId:   "*",
	}, "")
	return r
}

// CanMembership_User checks if the subject has membership permission
// // Object: group:<id>
// Schema: permission membership = member
func (obj *ObjGroup) CanMembership_User(sub *ObjUser) rel.Relationship {
	r, s := obj.src.Obj, sub.src
	return rel.Relationship{
		ResourceType:     r.ObjectType,
		ResourceID:       r.ObjectId,
		ResourceRelation: "membership",
		SubjectType:      s.Obj.ObjectType,
		SubjectID:        s.Obj.ObjectId,
		SubjectRelation:  s.OptionalRelation,
	}
}

// CanMembership_GroupMember checks if the subject has membership permission
// // Object: group:<id>
// Schema: permission membership = member
func (obj *ObjGroup) CanMembership_GroupMember(sub *ObjGroup) rel.Relationship {
	r, s := obj.src.Obj, sub.src
	return rel.Relationship{
		ResourceType:     r.ObjectType,
		ResourceID:       r.ObjectId,
		ResourceRelation: "membership",
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

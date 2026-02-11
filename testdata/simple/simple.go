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

type ObjResource struct {
	src Object
}

func (b *SchemaBuilder) Resource(id fmt.Stringer) *ObjResource {
	return &ObjResource{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "resource",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjResource) Object() rel.Object {
	return obj.src.Object()
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjResource) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjResource) RelationViewer() string {
	return "viewer"
}

func (obj *ObjResource) RelationWriter() string {
	return "writer"
}

func (obj *ObjResource) PermissionWrite() string {
	return "write"
}

func (obj *ObjResource) PermissionView() string {
	return "view"
}

type ResourceRelates struct {
	obj *ObjResource
	rel Relationship
}

func (obj *ObjResource) Touch() *ResourceRelates {
	return &ResourceRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjResource) Delete() *ResourceRelates {
	return &ResourceRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjResource) Create() *ResourceRelates {
	return &ResourceRelates{obj: obj, rel: obj.src.Create()}
}

// Writer simple.zed:7
// Relationship: resource:<id>#writer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Writer() etc.
func (obj *ObjResource) Writer(subs ...*ObjUser) *ObjResource {
	for _, sub := range subs {
		obj.src.Touch().Add("writer", sub.src.Obj, "")
	}
	return obj
}

// Writer on Relates uses the specified operation (Touch/Create/Delete)
func (r *ResourceRelates) Writer(subs ...*ObjUser) *ResourceRelates {
	for _, sub := range subs {
		r.rel.Add("writer", sub.src.Obj, "")
	}
	return r
}

// Viewer simple.zed:8
// Relationship: resource:<id>#viewer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Viewer() etc.
func (obj *ObjResource) Viewer(subs ...*ObjUser) *ObjResource {
	for _, sub := range subs {
		obj.src.Touch().Add("viewer", sub.src.Obj, "")
	}
	return obj
}

// Viewer on Relates uses the specified operation (Touch/Create/Delete)
func (r *ResourceRelates) Viewer(subs ...*ObjUser) *ResourceRelates {
	for _, sub := range subs {
		r.rel.Add("viewer", sub.src.Obj, "")
	}
	return r
}

// CanWrite_User checks if the subject has write permission
// // Object: resource:<id>
// Schema: permission write = writer
func (obj *ObjResource) CanWrite_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "write",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanView_User checks if the subject has view permission
// // Object: resource:<id>
// Schema: permission view = viewer + writer
func (obj *ObjResource) CanView_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "view",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
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

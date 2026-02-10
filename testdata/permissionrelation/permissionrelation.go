// Package policy code generated. DO NOT EDIT.
package policy

import (
	"fmt"

	. "github.com/Emyrk/zedgen/relbuilder"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
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

type ObjFile struct {
	src Object
}

func (b *SchemaBuilder) File(id fmt.Stringer) *ObjFile {
	return &ObjFile{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "file",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjFile) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjFile) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjFile) RelationFolder() string {
	return "folder"
}

func (obj *ObjFile) PermissionRead() string {
	return "read"
}

type FileRelates struct {
	obj *ObjFile
	rel Relationship
}

func (obj *ObjFile) Touch() *FileRelates {
	return &FileRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjFile) Delete() *FileRelates {
	return &FileRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjFile) Create() *FileRelates {
	return &FileRelates{obj: obj, rel: obj.src.Create()}
}

// Folder permissionrelation.zed:18
// Relationship: file:<id>#folder@folder:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Folder() etc.
func (obj *ObjFile) Folder(subs ...*ObjFolder) *ObjFile {
	for _, sub := range subs {
		obj.src.Touch().Add("folder", sub.src.Obj, "")
	}
	return obj
}

// Folder on Relates uses the specified operation (Touch/Create/Delete)
func (r *FileRelates) Folder(subs ...*ObjFolder) *FileRelates {
	for _, sub := range subs {
		r.rel.Add("folder", sub.src.Obj, "")
	}
	return r
}

// CanRead_Folder checks if the subject has read permission
// // Object: file:<id>
// Schema: permission read = folder->read
func (obj *ObjFile) CanRead_Folder(sub *ObjFolder) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

type ObjFolder struct {
	src Object
}

func (b *SchemaBuilder) Folder(id fmt.Stringer) *ObjFolder {
	return &ObjFolder{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "folder",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjFolder) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjFolder) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjFolder) RelationOwner() string {
	return "owner"
}

func (obj *ObjFolder) PermissionRead() string {
	return "read"
}

type FolderRelates struct {
	obj *ObjFolder
	rel Relationship
}

func (obj *ObjFolder) Touch() *FolderRelates {
	return &FolderRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjFolder) Delete() *FolderRelates {
	return &FolderRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjFolder) Create() *FolderRelates {
	return &FolderRelates{obj: obj, rel: obj.src.Create()}
}

// Owner_User permissionrelation.zed:12
// Relationship: folder:<id>#owner@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Owner_User() etc.
func (obj *ObjFolder) Owner_User(subs ...*ObjUser) *ObjFolder {
	for _, sub := range subs {
		obj.src.Touch().Add("owner", sub.src.Obj, "")
	}
	return obj
}

// Owner_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *FolderRelates) Owner_User(subs ...*ObjUser) *FolderRelates {
	for _, sub := range subs {
		r.rel.Add("owner", sub.src.Obj, "")
	}
	return r
}

// Owner_Group permissionrelation.zed:12
// Relationship: folder:<id>#owner@group:<id>#membership
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Owner_Group() etc.
func (obj *ObjFolder) Owner_Group(subs ...*ObjGroup) *ObjFolder {
	for _, sub := range subs {
		obj.src.Touch().Add("owner", sub.src.Obj, "membership")
	}
	return obj
}

// Owner_Group on Relates uses the specified operation (Touch/Create/Delete)
func (r *FolderRelates) Owner_Group(subs ...*ObjGroup) *FolderRelates {
	for _, sub := range subs {
		r.rel.Add("owner", sub.src.Obj, "membership")
	}
	return r
}

// CanRead_User checks if the subject has read permission
// // Object: folder:<id>
// Schema: permission read = owner
func (obj *ObjFolder) CanRead_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanRead_GroupMembership checks if the subject has read permission
// // Object: folder:<id>
// Schema: permission read = owner
func (obj *ObjFolder) CanRead_GroupMembership(sub *ObjGroup) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "membership",
		},
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
func (obj *ObjGroup) Object() *v1.ObjectReference {
	return obj.src.Obj
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

// Member_User permissionrelation.zed:4
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

// Member_Group permissionrelation.zed:4
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

// CanMembership_User checks if the subject has membership permission
// // Object: group:<id>
func (obj *ObjGroup) CanMembership_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "membership",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanMembership_GroupMember checks if the subject has membership permission
// // Object: group:<id>
func (obj *ObjGroup) CanMembership_GroupMember(sub *ObjGroup) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "membership",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
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
func (obj *ObjUser) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjUser) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

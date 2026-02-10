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

type ObjOrganization struct {
	src Object
}

func (b *SchemaBuilder) Organization(id fmt.Stringer) *ObjOrganization {
	return &ObjOrganization{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "organization",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjOrganization) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjOrganization) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjOrganization) RelationBilling_manager() string {
	return "billing_manager"
}

func (obj *ObjOrganization) RelationMember() string {
	return "member"
}

func (obj *ObjOrganization) RelationOwn() string {
	return "own"
}

func (obj *ObjOrganization) RelationTeam_maintainer() string {
	return "team_maintainer"
}

func (obj *ObjOrganization) PermissionCreate_repository() string {
	return "create_repository"
}

func (obj *ObjOrganization) PermissionManage_billing() string {
	return "manage_billing"
}

func (obj *ObjOrganization) PermissionUser_seat() string {
	return "user_seat"
}

func (obj *ObjOrganization) PermissionOwner() string {
	return "owner"
}

func (obj *ObjOrganization) PermissionChange_team_name() string {
	return "change_team_name"
}

type OrganizationRelates struct {
	obj *ObjOrganization
	rel Relationship
}

func (obj *ObjOrganization) Touch() *OrganizationRelates {
	return &OrganizationRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjOrganization) Delete() *OrganizationRelates {
	return &OrganizationRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjOrganization) Create() *OrganizationRelates {
	return &OrganizationRelates{obj: obj, rel: obj.src.Create()}
}

// Own github.zed:13
// Relationship: organization:<id>#own@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Own() etc.
func (obj *ObjOrganization) Own(subs ...*ObjUser) *ObjOrganization {
	for _, sub := range subs {
		obj.src.Touch().Add("own", sub.src.Obj, "")
	}
	return obj
}

// Own on Relates uses the specified operation (Touch/Create/Delete)
func (r *OrganizationRelates) Own(subs ...*ObjUser) *OrganizationRelates {
	for _, sub := range subs {
		r.rel.Add("own", sub.src.Obj, "")
	}
	return r
}

// Member github.zed:14
// Relationship: organization:<id>#member@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Member() etc.
func (obj *ObjOrganization) Member(subs ...*ObjUser) *ObjOrganization {
	for _, sub := range subs {
		obj.src.Touch().Add("member", sub.src.Obj, "")
	}
	return obj
}

// Member on Relates uses the specified operation (Touch/Create/Delete)
func (r *OrganizationRelates) Member(subs ...*ObjUser) *OrganizationRelates {
	for _, sub := range subs {
		r.rel.Add("member", sub.src.Obj, "")
	}
	return r
}

// Billing_manager github.zed:15
// Relationship: organization:<id>#billing_manager@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Billing_manager() etc.
func (obj *ObjOrganization) Billing_manager(subs ...*ObjUser) *ObjOrganization {
	for _, sub := range subs {
		obj.src.Touch().Add("billing_manager", sub.src.Obj, "")
	}
	return obj
}

// Billing_manager on Relates uses the specified operation (Touch/Create/Delete)
func (r *OrganizationRelates) Billing_manager(subs ...*ObjUser) *OrganizationRelates {
	for _, sub := range subs {
		r.rel.Add("billing_manager", sub.src.Obj, "")
	}
	return r
}

// Team_maintainer github.zed:16
// Relationship: organization:<id>#team_maintainer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Team_maintainer() etc.
func (obj *ObjOrganization) Team_maintainer(subs ...*ObjUser) *ObjOrganization {
	for _, sub := range subs {
		obj.src.Touch().Add("team_maintainer", sub.src.Obj, "")
	}
	return obj
}

// Team_maintainer on Relates uses the specified operation (Touch/Create/Delete)
func (r *OrganizationRelates) Team_maintainer(subs ...*ObjUser) *OrganizationRelates {
	for _, sub := range subs {
		r.rel.Add("team_maintainer", sub.src.Obj, "")
	}
	return r
}

// CanCreate_repository_User checks if the subject has create_repository permission
// // Object: organization:<id>
func (obj *ObjOrganization) CanCreate_repository_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_repository",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanManage_billing_User checks if the subject has manage_billing permission
// // Object: organization:<id>
func (obj *ObjOrganization) CanManage_billing_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_billing",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanUser_seat_User checks if the subject has user_seat permission
// // Object: organization:<id>
// Schema: permission user_seat = owner + member + team_maintainer
func (obj *ObjOrganization) CanUser_seat_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "user_seat",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanOwner_User checks if the subject has owner permission
// // Object: organization:<id>
// Schema: permission owner = own
func (obj *ObjOrganization) CanOwner_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "owner",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanChange_team_name_User checks if the subject has change_team_name permission
// // Object: organization:<id>
func (obj *ObjOrganization) CanChange_team_name_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "change_team_name",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

type ObjRepository struct {
	src Object
}

func (b *SchemaBuilder) Repository(id fmt.Stringer) *ObjRepository {
	return &ObjRepository{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "repository",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjRepository) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjRepository) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjRepository) RelationAdmin() string {
	return "admin"
}

func (obj *ObjRepository) RelationMaintainer() string {
	return "maintainer"
}

func (obj *ObjRepository) RelationOrganization() string {
	return "organization"
}

func (obj *ObjRepository) RelationPublic() string {
	return "public"
}

func (obj *ObjRepository) RelationReader() string {
	return "reader"
}

func (obj *ObjRepository) RelationTriager() string {
	return "triager"
}

func (obj *ObjRepository) RelationWriter() string {
	return "writer"
}

func (obj *ObjRepository) PermissionClone() string {
	return "clone"
}

func (obj *ObjRepository) PermissionPush() string {
	return "push"
}

func (obj *ObjRepository) PermissionRead() string {
	return "read"
}

func (obj *ObjRepository) PermissionDelete() string {
	return "delete"
}

func (obj *ObjRepository) PermissionCreate_issue() string {
	return "create_issue"
}

func (obj *ObjRepository) PermissionClose_issue() string {
	return "close_issue"
}

func (obj *ObjRepository) PermissionCreate_pull_request() string {
	return "create_pull_request"
}

func (obj *ObjRepository) PermissionMerge_pull_request() string {
	return "merge_pull_request"
}

func (obj *ObjRepository) PermissionClose_pull_request() string {
	return "close_pull_request"
}

func (obj *ObjRepository) PermissionManage_setting() string {
	return "manage_setting"
}

func (obj *ObjRepository) PermissionManage_sensitive_setting() string {
	return "manage_sensitive_setting"
}

type RepositoryRelates struct {
	obj *ObjRepository
	rel Relationship
}

func (obj *ObjRepository) Touch() *RepositoryRelates {
	return &RepositoryRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjRepository) Delete() *RepositoryRelates {
	return &RepositoryRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjRepository) Create() *RepositoryRelates {
	return &RepositoryRelates{obj: obj, rel: obj.src.Create()}
}

// Organization github.zed:31
// Relationship: repository:<id>#organization@organization:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Organization() etc.
func (obj *ObjRepository) Organization(subs ...*ObjOrganization) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("organization", sub.src.Obj, "")
	}
	return obj
}

// Organization on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Organization(subs ...*ObjOrganization) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("organization", sub.src.Obj, "")
	}
	return r
}

// Reader_User github.zed:34
// Relationship: repository:<id>#reader@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Reader_User() etc.
func (obj *ObjRepository) Reader_User(subs ...*ObjUser) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("reader", sub.src.Obj, "")
	}
	return obj
}

// Reader_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Reader_User(subs ...*ObjUser) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("reader", sub.src.Obj, "")
	}
	return r
}

// Reader_Team github.zed:34
// Relationship: repository:<id>#reader@team:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Reader_Team() etc.
func (obj *ObjRepository) Reader_Team(subs ...*ObjTeam) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("reader", sub.src.Obj, "member")
	}
	return obj
}

// Reader_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Reader_Team(subs ...*ObjTeam) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("reader", sub.src.Obj, "member")
	}
	return r
}

// Triager_User github.zed:35
// Relationship: repository:<id>#triager@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Triager_User() etc.
func (obj *ObjRepository) Triager_User(subs ...*ObjUser) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("triager", sub.src.Obj, "")
	}
	return obj
}

// Triager_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Triager_User(subs ...*ObjUser) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("triager", sub.src.Obj, "")
	}
	return r
}

// Triager_Team github.zed:35
// Relationship: repository:<id>#triager@team:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Triager_Team() etc.
func (obj *ObjRepository) Triager_Team(subs ...*ObjTeam) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("triager", sub.src.Obj, "member")
	}
	return obj
}

// Triager_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Triager_Team(subs ...*ObjTeam) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("triager", sub.src.Obj, "member")
	}
	return r
}

// Writer_User github.zed:36
// Relationship: repository:<id>#writer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Writer_User() etc.
func (obj *ObjRepository) Writer_User(subs ...*ObjUser) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("writer", sub.src.Obj, "")
	}
	return obj
}

// Writer_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Writer_User(subs ...*ObjUser) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("writer", sub.src.Obj, "")
	}
	return r
}

// Writer_Team github.zed:36
// Relationship: repository:<id>#writer@team:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Writer_Team() etc.
func (obj *ObjRepository) Writer_Team(subs ...*ObjTeam) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("writer", sub.src.Obj, "member")
	}
	return obj
}

// Writer_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Writer_Team(subs ...*ObjTeam) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("writer", sub.src.Obj, "member")
	}
	return r
}

// Maintainer_User github.zed:37
// Relationship: repository:<id>#maintainer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Maintainer_User() etc.
func (obj *ObjRepository) Maintainer_User(subs ...*ObjUser) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("maintainer", sub.src.Obj, "")
	}
	return obj
}

// Maintainer_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Maintainer_User(subs ...*ObjUser) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("maintainer", sub.src.Obj, "")
	}
	return r
}

// Maintainer_Team github.zed:37
// Relationship: repository:<id>#maintainer@team:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Maintainer_Team() etc.
func (obj *ObjRepository) Maintainer_Team(subs ...*ObjTeam) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("maintainer", sub.src.Obj, "member")
	}
	return obj
}

// Maintainer_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Maintainer_Team(subs ...*ObjTeam) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("maintainer", sub.src.Obj, "member")
	}
	return r
}

// Admin_User github.zed:38
// Relationship: repository:<id>#admin@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Admin_User() etc.
func (obj *ObjRepository) Admin_User(subs ...*ObjUser) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("admin", sub.src.Obj, "")
	}
	return obj
}

// Admin_User on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Admin_User(subs ...*ObjUser) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("admin", sub.src.Obj, "")
	}
	return r
}

// Admin_Team github.zed:38
// Relationship: repository:<id>#admin@team:<id>#member
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Admin_Team() etc.
func (obj *ObjRepository) Admin_Team(subs ...*ObjTeam) *ObjRepository {
	for _, sub := range subs {
		obj.src.Touch().Add("admin", sub.src.Obj, "member")
	}
	return obj
}

// Admin_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *RepositoryRelates) Admin_Team(subs ...*ObjTeam) *RepositoryRelates {
	for _, sub := range subs {
		r.rel.Add("admin", sub.src.Obj, "member")
	}
	return r
}

// PublicWildcard github.zed:39
// Relationship: repository:<id>#public@user:*
func (obj *ObjRepository) PublicWildcard() *ObjRepository {
	obj.src.Touch().Add("public", &v1.ObjectReference{
		ObjectType: "user",
		ObjectId:   "*",
	}, "")
	return obj
}

// PublicWildcard on Relates uses the specified operation
func (r *RepositoryRelates) PublicWildcard() *RepositoryRelates {
	r.rel.Add("public", &v1.ObjectReference{
		ObjectType: "user",
		ObjectId:   "*",
	}, "")
	return r
}

// CanClone_Organization checks if the subject has clone permission
// // Object: repository:<id>
func (obj *ObjRepository) CanClone_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "clone",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClone_User checks if the subject has clone permission
// // Object: repository:<id>
func (obj *ObjRepository) CanClone_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "clone",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClone_TeamMember checks if the subject has clone permission
// // Object: repository:<id>
func (obj *ObjRepository) CanClone_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "clone",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanPush_Organization checks if the subject has push permission
// // Object: repository:<id>
// Schema: permission push = writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanPush_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "push",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanPush_User checks if the subject has push permission
// // Object: repository:<id>
// Schema: permission push = writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanPush_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "push",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanPush_TeamMember checks if the subject has push permission
// // Object: repository:<id>
// Schema: permission push = writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanPush_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "push",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanRead_Organization checks if the subject has read permission
// // Object: repository:<id>
func (obj *ObjRepository) CanRead_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanRead_User checks if the subject has read permission
// // Object: repository:<id>
func (obj *ObjRepository) CanRead_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanRead_TeamMember checks if the subject has read permission
// // Object: repository:<id>
func (obj *ObjRepository) CanRead_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "read",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanDelete_Organization checks if the subject has delete permission
// // Object: repository:<id>
// Schema: permission delete = admin + organization->owner
func (obj *ObjRepository) CanDelete_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "delete",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanDelete_User checks if the subject has delete permission
// // Object: repository:<id>
// Schema: permission delete = admin + organization->owner
func (obj *ObjRepository) CanDelete_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "delete",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanDelete_TeamMember checks if the subject has delete permission
// // Object: repository:<id>
// Schema: permission delete = admin + organization->owner
func (obj *ObjRepository) CanDelete_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "delete",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanCreate_issue_Organization checks if the subject has create_issue permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_issue_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanCreate_issue_User checks if the subject has create_issue permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_issue_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanCreate_issue_TeamMember checks if the subject has create_issue permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_issue_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanClose_issue_Organization checks if the subject has close_issue permission
// // Object: repository:<id>
// Schema: permission close_issue = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_issue_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClose_issue_User checks if the subject has close_issue permission
// // Object: repository:<id>
// Schema: permission close_issue = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_issue_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClose_issue_TeamMember checks if the subject has close_issue permission
// // Object: repository:<id>
// Schema: permission close_issue = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_issue_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_issue",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanCreate_pull_request_Organization checks if the subject has create_pull_request permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_pull_request_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanCreate_pull_request_User checks if the subject has create_pull_request permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_pull_request_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanCreate_pull_request_TeamMember checks if the subject has create_pull_request permission
// // Object: repository:<id>
func (obj *ObjRepository) CanCreate_pull_request_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "create_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanMerge_pull_request_Organization checks if the subject has merge_pull_request permission
// // Object: repository:<id>
// Schema: permission merge_pull_request = maintainer + organization->owner
func (obj *ObjRepository) CanMerge_pull_request_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "merge_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanMerge_pull_request_User checks if the subject has merge_pull_request permission
// // Object: repository:<id>
// Schema: permission merge_pull_request = maintainer + organization->owner
func (obj *ObjRepository) CanMerge_pull_request_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "merge_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanMerge_pull_request_TeamMember checks if the subject has merge_pull_request permission
// // Object: repository:<id>
// Schema: permission merge_pull_request = maintainer + organization->owner
func (obj *ObjRepository) CanMerge_pull_request_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "merge_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanClose_pull_request_Organization checks if the subject has close_pull_request permission
// // Object: repository:<id>
// Schema: permission close_pull_request = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_pull_request_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClose_pull_request_User checks if the subject has close_pull_request permission
// // Object: repository:<id>
// Schema: permission close_pull_request = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_pull_request_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanClose_pull_request_TeamMember checks if the subject has close_pull_request permission
// // Object: repository:<id>
// Schema: permission close_pull_request = triager + writer + maintainer + admin + organization->owner
func (obj *ObjRepository) CanClose_pull_request_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "close_pull_request",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanManage_setting_Organization checks if the subject has manage_setting permission
// // Object: repository:<id>
func (obj *ObjRepository) CanManage_setting_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanManage_setting_User checks if the subject has manage_setting permission
// // Object: repository:<id>
func (obj *ObjRepository) CanManage_setting_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanManage_setting_TeamMember checks if the subject has manage_setting permission
// // Object: repository:<id>
func (obj *ObjRepository) CanManage_setting_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

// CanManage_sensitive_setting_Organization checks if the subject has manage_sensitive_setting permission
// // Object: repository:<id>
// Schema: permission manage_sensitive_setting = admin + organization->owner
func (obj *ObjRepository) CanManage_sensitive_setting_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_sensitive_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanManage_sensitive_setting_User checks if the subject has manage_sensitive_setting permission
// // Object: repository:<id>
// Schema: permission manage_sensitive_setting = admin + organization->owner
func (obj *ObjRepository) CanManage_sensitive_setting_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_sensitive_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanManage_sensitive_setting_TeamMember checks if the subject has manage_sensitive_setting permission
// // Object: repository:<id>
// Schema: permission manage_sensitive_setting = admin + organization->owner
func (obj *ObjRepository) CanManage_sensitive_setting_TeamMember(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "manage_sensitive_setting",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "member",
		},
	}
}

type ObjTeam struct {
	src Object
}

func (b *SchemaBuilder) Team(id fmt.Stringer) *ObjTeam {
	return &ObjTeam{
		src: b.Object(&v1.ObjectReference{
			ObjectType: "team",
			ObjectId:   id.String(),
		}, ""),
	}
}

// Object returns the underlying ObjectReference for use in SpiceDB API calls.
func (obj *ObjTeam) Object() *v1.ObjectReference {
	return obj.src.Obj
}

// AsSubject returns this object as a SubjectReference for use in checks.
func (obj *ObjTeam) AsSubject() *v1.SubjectReference {
	return &v1.SubjectReference{
		Object:           obj.src.Obj,
		OptionalRelation: obj.src.OptionalRelation,
	}
}

func (obj *ObjTeam) RelationDirect_member() string {
	return "direct_member"
}

func (obj *ObjTeam) RelationMaintainer() string {
	return "maintainer"
}

func (obj *ObjTeam) RelationParent() string {
	return "parent"
}

func (obj *ObjTeam) PermissionMember() string {
	return "member"
}

func (obj *ObjTeam) PermissionChange_team_name() string {
	return "change_team_name"
}

type TeamRelates struct {
	obj *ObjTeam
	rel Relationship
}

func (obj *ObjTeam) Touch() *TeamRelates {
	return &TeamRelates{obj: obj, rel: obj.src.Touch()}
}

func (obj *ObjTeam) Delete() *TeamRelates {
	return &TeamRelates{obj: obj, rel: obj.src.Delete()}
}

func (obj *ObjTeam) Create() *TeamRelates {
	return &TeamRelates{obj: obj, rel: obj.src.Create()}
}

// Parent_Organization github.zed:4
// Relationship: team:<id>#parent@organization:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Parent_Organization() etc.
func (obj *ObjTeam) Parent_Organization(subs ...*ObjOrganization) *ObjTeam {
	for _, sub := range subs {
		obj.src.Touch().Add("parent", sub.src.Obj, "")
	}
	return obj
}

// Parent_Organization on Relates uses the specified operation (Touch/Create/Delete)
func (r *TeamRelates) Parent_Organization(subs ...*ObjOrganization) *TeamRelates {
	for _, sub := range subs {
		r.rel.Add("parent", sub.src.Obj, "")
	}
	return r
}

// Parent_Team github.zed:4
// Relationship: team:<id>#parent@team:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Parent_Team() etc.
func (obj *ObjTeam) Parent_Team(subs ...*ObjTeam) *ObjTeam {
	for _, sub := range subs {
		obj.src.Touch().Add("parent", sub.src.Obj, "")
	}
	return obj
}

// Parent_Team on Relates uses the specified operation (Touch/Create/Delete)
func (r *TeamRelates) Parent_Team(subs ...*ObjTeam) *TeamRelates {
	for _, sub := range subs {
		r.rel.Add("parent", sub.src.Obj, "")
	}
	return r
}

// Maintainer github.zed:5
// Relationship: team:<id>#maintainer@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Maintainer() etc.
func (obj *ObjTeam) Maintainer(subs ...*ObjUser) *ObjTeam {
	for _, sub := range subs {
		obj.src.Touch().Add("maintainer", sub.src.Obj, "")
	}
	return obj
}

// Maintainer on Relates uses the specified operation (Touch/Create/Delete)
func (r *TeamRelates) Maintainer(subs ...*ObjUser) *TeamRelates {
	for _, sub := range subs {
		r.rel.Add("maintainer", sub.src.Obj, "")
	}
	return r
}

// Direct_member github.zed:6
// Relationship: team:<id>#direct_member@user:<id>
// Uses Touch operation implicitly. For Delete/Create, use obj.Delete().Direct_member() etc.
func (obj *ObjTeam) Direct_member(subs ...*ObjUser) *ObjTeam {
	for _, sub := range subs {
		obj.src.Touch().Add("direct_member", sub.src.Obj, "")
	}
	return obj
}

// Direct_member on Relates uses the specified operation (Touch/Create/Delete)
func (r *TeamRelates) Direct_member(subs ...*ObjUser) *TeamRelates {
	for _, sub := range subs {
		r.rel.Add("direct_member", sub.src.Obj, "")
	}
	return r
}

// CanMember_Organization checks if the subject has member permission
// // Object: team:<id>
// Schema: permission member = maintainer + direct_member
func (obj *ObjTeam) CanMember_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "member",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanMember_Team checks if the subject has member permission
// // Object: team:<id>
// Schema: permission member = maintainer + direct_member
func (obj *ObjTeam) CanMember_Team(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "member",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanMember_User checks if the subject has member permission
// // Object: team:<id>
// Schema: permission member = maintainer + direct_member
func (obj *ObjTeam) CanMember_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "member",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanChange_team_name_Organization checks if the subject has change_team_name permission
// // Object: team:<id>
// Schema: permission change_team_name = maintainer + parent->change_team_name
func (obj *ObjTeam) CanChange_team_name_Organization(sub *ObjOrganization) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "change_team_name",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanChange_team_name_Team checks if the subject has change_team_name permission
// // Object: team:<id>
// Schema: permission change_team_name = maintainer + parent->change_team_name
func (obj *ObjTeam) CanChange_team_name_Team(sub *ObjTeam) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "change_team_name",
		Subject: &v1.SubjectReference{
			Object:           sub.src.Obj,
			OptionalRelation: "",
		},
	}
}

// CanChange_team_name_User checks if the subject has change_team_name permission
// // Object: team:<id>
// Schema: permission change_team_name = maintainer + parent->change_team_name
func (obj *ObjTeam) CanChange_team_name_User(sub *ObjUser) *v1.CheckPermissionRequest {
	return &v1.CheckPermissionRequest{
		Resource:   obj.src.Obj,
		Permission: "change_team_name",
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

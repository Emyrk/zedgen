package relbuilder

import (
	"fmt"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/gochugaru/rel"
)

type Builder struct {
	txn rel.Txn
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) Add(op v1.RelationshipUpdate_Operation, r *v1.Relationship) *Builder {
	switch op {
	case v1.RelationshipUpdate_OPERATION_TOUCH:
		return b.Touch(r)
	case v1.RelationshipUpdate_OPERATION_CREATE:
		return b.Create(r)
	case v1.RelationshipUpdate_OPERATION_DELETE:
		return b.Delete(r)
	default:
		panic(fmt.Sprintf("unknown operation: %v", op))
	}
}

func (b *Builder) Touch(r *v1.Relationship) *Builder {
	b.txn.Touch(*rel.FromV1Proto(r))
	return b
}

func (b *Builder) Create(r *v1.Relationship) *Builder {
	b.txn.Create(*rel.FromV1Proto(r))
	return b
}

func (b *Builder) Delete(r *v1.Relationship) *Builder {
	b.txn.Delete(*rel.FromV1Proto(r))
	return b
}

func (b *Builder) User(id fmt.Stringer) *ObjUser {
	return &ObjUser{
		src: Object{
			Obj: &v1.ObjectReference{
				ObjectType: "user",
				ObjectId:   id.String(),
			},
			OptionalRelation: "",
			builder:          b,
		},
	}
}

type ObjUser struct {
	src Object
}

func (obj *ObjUser) Touch() *ObjUserRelates {
	return &ObjUserRelates{rel: obj.src.Touch()}
}

//

func (b *Builder) Resource(id fmt.Stringer) *ObjResource {
	return &ObjResource{
		src: Object{
			Obj: &v1.ObjectReference{
				ObjectType: "resource",
				ObjectId:   id.String(),
			},
			OptionalRelation: "",
			builder:          b,
		},
	}
}

type ObjUserRelates struct {
	rel Relationship
}

type ObjResource struct {
	src Object
}

func (obj *ObjResource) Touch() *ObjResourceRelates {
	return &ObjResourceRelates{rel: obj.src.Touch()}
}

type ObjResourceRelates struct {
	rel Relationship
}

func (obj *ObjResourceRelates) Writer(subs ...*ObjUser) *ObjResourceRelates {
	for _, sub := range subs {
		obj.rel.builder.Add(obj.rel.Operation, &v1.Relationship{
			Resource: sub.src.Obj,
			Relation: "writer",
			Subject: &v1.SubjectReference{
				Object: sub.src.Obj,
			},
			OptionalCaveat:    nil,
			OptionalExpiresAt: nil,
		})
	}
	return obj
}

func foo() {
	b := New()

	b.Resource(String("foo")).Touch().Writer(String("hello"))
}

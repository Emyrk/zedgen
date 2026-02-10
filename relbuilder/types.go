package relbuilder

import v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"

type String string

func (s String) String() string {
	return string(s)
}

type Object struct {
	Obj              *v1.ObjectReference
	OptionalRelation string
	builder          *Builder
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

type Relationship struct {
	Resource  *Object
	Operation v1.RelationshipUpdate_Operation
	builder   *Builder
}

package relbuilder

import (
	"fmt"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/gochugaru/rel"
)

// Builder is the interface for relationship building operations.
type Builder interface {
	Add(op v1.RelationshipUpdate_Operation, r *v1.Relationship) Builder
	Touch(r *v1.Relationship) Builder
	Create(r *v1.Relationship) Builder
	Delete(r *v1.Relationship) Builder
}

// Build is the concrete implementation of Builder that accumulates relationships.
type Build struct {
	txn rel.Txn
}

// NewBuild creates a new relationship builder.
func NewBuild() *Build {
	return &Build{}
}

// Add adds a relationship with the specified operation type.
func (b *Build) Add(op v1.RelationshipUpdate_Operation, r *v1.Relationship) Builder {
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

// Touch idempotently creates or updates a relationship.
func (b *Build) Touch(r *v1.Relationship) Builder {
	b.txn.Touch(*rel.FromV1Proto(r))
	return b
}

// Create inserts a new relationship (fails if already exists).
func (b *Build) Create(r *v1.Relationship) Builder {
	b.txn.Create(*rel.FromV1Proto(r))
	return b
}

// Delete removes a relationship.
func (b *Build) Delete(r *v1.Relationship) Builder {
	b.txn.Delete(*rel.FromV1Proto(r))
	return b
}

// Updates returns the accumulated relationship updates for use with WriteRelationships.
func (b *Build) Updates() []*v1.RelationshipUpdate {
	return b.txn.V1Updates
}

// Preconditions returns any preconditions that were set on the transaction.
func (b *Build) Preconditions() []*v1.Precondition {
	return b.txn.V1Preconds
}

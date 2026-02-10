// Package example demonstrates zedgen's type-safe SpiceDB helpers.
//
// This example uses a GitHub-like authorization schema with users, teams,
// organizations, and repositories.
package example

import (
	"context"
	"fmt"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/google/uuid"

	"github.com/Emyrk/zedgen/relbuilder"
	policy "github.com/Emyrk/zedgen/testdata/github"
)

func Example() {
	b := policy.New()

	// =========================================
	// Create typed objects from IDs
	// =========================================
	alice := b.User(uuid.New())
	bob := b.User(uuid.New())
	charlie := b.User(uuid.New())

	acmeCorp := b.Organization(relbuilder.String("acme-corp"))
	devTeam := b.Team(uuid.New())
	repo := b.Repository(relbuilder.String("acme-corp/api"))

	// =========================================
	// Build relationships (Touch is implicit)
	// =========================================

	// Organization membership
	acmeCorp.Own(alice)           // alice owns acme-corp
	acmeCorp.Member(bob, charlie) // bob and charlie are members

	// Team membership - fluent chaining
	devTeam.
		Parent_Organization(acmeCorp). // team belongs to org
		Maintainer(alice). // alice maintains the team
		Direct_member(bob) // bob is a direct member

	// Repository permissions - multiple subject types use _Suffix pattern
	repo.
		Organization(acmeCorp).
		Admin_User(alice). // direct user admin
		Writer_User(bob). // direct user writer
		Reader_Team(devTeam). // team#member as reader (subject relation)
		PublicWildcard() // user:* - public read access

	// =========================================
	// Get relationship updates for SpiceDB
	// =========================================
	updates := b.Updates()
	fmt.Printf("Built %d relationship updates\n", len(updates))

	// Use with SpiceDB client:
	// client.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{Updates: updates})

	// =========================================
	// Type-safe permission checks
	// =========================================

	// Check if alice can push - returns *v1.CheckPermissionRequest
	canPush := repo.CanPush_User(alice)
	fmt.Printf("Check: can %s %s on %s:%s?\n",
		canPush.Subject.Object.ObjectId,
		canPush.Permission,
		canPush.Resource.ObjectType,
		canPush.Resource.ObjectId,
	)

	// Check with team#member subject
	canRead := repo.CanRead_TeamMember(devTeam)
	_ = canRead // use with client.CheckPermission(ctx, canRead)

	// =========================================
	// Explicit operations (Delete/Create)
	// =========================================

	// Remove bob as writer (explicit Delete)
	repo.Delete().Writer_User(bob)

	// Create fails if relationship exists (explicit Create)
	repo.Create().Triager_User(charlie)

	// =========================================
	// Access helpers and constants
	// =========================================

	// Get underlying protobuf types
	_ = repo.Object()     // *v1.ObjectReference
	_ = alice.AsSubject() // *v1.SubjectReference

	// Relation and permission constants (prevent typos)
	_ = repo.RelationWriter() // "writer"
	_ = repo.PermissionPush() // "push"
}

// ExampleWithClient shows integration with a SpiceDB client.
func ExampleWithClient(ctx context.Context, client v1.PermissionsServiceClient) error {
	b := policy.New()

	alice := b.User(relbuilder.String("alice"))
	repo := b.Repository(relbuilder.String("acme/api"))

	// Build relationships
	repo.Writer_User(alice)

	// Write to SpiceDB
	_, err := client.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{
		Updates:               b.Updates(),
		OptionalPreconditions: b.Preconditions(),
	})
	if err != nil {
		return fmt.Errorf("write relationships: %w", err)
	}

	// Check permission
	resp, err := client.CheckPermission(ctx, repo.CanPush_User(alice))
	if err != nil {
		return fmt.Errorf("check permission: %w", err)
	}

	if resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION {
		fmt.Println("alice can push!")
	}

	return nil
}

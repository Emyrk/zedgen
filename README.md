# zedgen

Generate type-safe Go code from [SpiceDB](https://authzed.com/spicedb) schemas.

## Features

- **Type-safe relationships**: `repo.Writer_User(alice)` instead of raw tuples
- **Type-safe permission checks**: `repo.CanPush_User(alice)` returns `*v1.CheckPermissionRequest`
- **Implicit Touch**: Common case is concise, explicit `Delete()`/`Create()` when needed
- **IDE-friendly**: Autocomplete shows valid relations and permissions per type

## Installation

```bash
go install github.com/Emyrk/zedgen/cmd/zedgen@latest
```

## Usage

```bash
# Generate to stdout
zedgen -schema schema.zed -package policy

# Generate to file
zedgen -schema schema.zed -package policy -out policy/policy.go
```

## Example

Given a SpiceDB schema:

```zed
definition user {}

definition team {
    relation direct_member: user
    permission member = direct_member
}

definition repository {
    relation writer: user | team#member
    relation public: user:*
    permission push = writer
    permission read = public + push
}
```

Generate Go code and use it:

```go
package main

import (
    "context"
    "github.com/your/project/policy"
    "github.com/Emyrk/zedgen/relbuilder"
    "github.com/authzed/authzed-go/v1"
)

func main() {
    b := policy.New()

    // Create objects
    alice := b.User(relbuilder.String("alice"))
    devTeam := b.Team(relbuilder.String("dev-team"))
    repo := b.Repository(relbuilder.String("my/repo"))

    // Build relationships (Touch is implicit)
    devTeam.Direct_member(alice)
    repo.Writer_User(alice).Writer_TeamMember(devTeam)
    repo.PublicWildcard() // user:* - anyone can read

    // Get transaction for WriteRelationships
    txn := b.Txn()
    _, err := client.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{
        Updates: txn.V1Updates,
    })

    // Permission checks - type-safe, returns rel.Relationship
    check := repo.CanPush_User(alice)
    resp, err := client.CheckPermission(ctx, check.V1CheckPermissionRequest())

    // Explicit Delete when needed
    repo.Delete().Writer_User(alice)
}
```

## API Patterns

| Pattern | Example |
|---------|---------|
| Create object | `b.User(id)` → `*ObjUser` |
| Add relation (implicit Touch) | `repo.Writer(alice)` |
| Add with subject relation | `repo.Writer_TeamMember(team)` |
| Add wildcard | `repo.PublicWildcard()` |
| Explicit operation | `repo.Delete().Writer(alice)` |
| Permission check | `repo.CanPush_User(alice)` → `rel.Relationship` |
| Get transaction | `b.Txn()` → `*rel.Txn` (use `.V1Updates`, `.V1Preconds`) |
| Access constants | `repo.RelationWriter()`, `repo.PermissionPush()` |

## License

MIT

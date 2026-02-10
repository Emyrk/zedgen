package example

import (
	"github.com/Emyrk/zedgen/relbuilder"
	policy "github.com/Emyrk/zedgen/testdata/github"
	"github.com/google/uuid"
)

func main() {
	b := policy.New()

	alice := b.User(uuid.New())
	bob := b.User(uuid.New())

	admins := b.Team(uuid.New()).
		Direct_member(alice)

	//public := b.

	b.Repository(relbuilder.String("my/repo")).
		Writer_User(alice, bob).Admin_Team(admins)
}

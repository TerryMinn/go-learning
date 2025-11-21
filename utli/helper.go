package utli

// AddMember adds a new member to the members slice and returns the updated slice.
func AddMember(name string, members []string) []string {
	return append(members, name)
}

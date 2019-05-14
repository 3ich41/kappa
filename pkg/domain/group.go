package domain

type Group struct {
	CN          string
	Description string
}

type GroupRepository interface {
	GetGroupsOfUser(username string) ([]Group, error)
}

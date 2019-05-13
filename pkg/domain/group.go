package domain

type Group struct {
	Name        string
	Description string
}

type GroupRepository interface {
	GetGroupsOfUser(username string) ([]Group, error)
}

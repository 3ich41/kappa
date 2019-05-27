package usecases

import (
	"m15.io/kappa/pkg/domain"
)

// ConfInteractor is an interface defining GetConf usecase
type ConfInteractor interface {
	GetConf(username string) (*domain.Conf, error)
}

// confInteractor implements ConfInteractor interface providing GetConf usecase
type confInteractor struct {
	groupRepo domain.GroupRepository
}

// NewConfInteractor creates and initializes ConfInteractor
func NewConfInteractor(groupRepo domain.GroupRepository) ConfInteractor {
	confinteractor := new(confInteractor)
	confinteractor.groupRepo = groupRepo

	return confinteractor
}

// GetConf returns domain.Conf for given username
func (interactor *confInteractor) GetConf(username string) (*domain.Conf, error) {
	conf := new(domain.Conf)
	conf.Username = username

	groups, err := interactor.groupRepo.GetGroupsOfUser(username)
	if err != nil {
		return nil, err
	}

	buttons := groupsToButtons(groups)
	conf.Buttons = buttons

	return conf, nil
}

// groupsToButtons transforms list of Groups to list of Buttons
func groupsToButtons(groups []domain.Group) []domain.Button {
	buttons := make([]domain.Button, 0, 25)
	for _, group := range groups {
		b := domain.Button{Text: group.Description, Value: group.CN}
		buttons = append(buttons, b)
	}
	return buttons
}

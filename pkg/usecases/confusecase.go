package usecases

import (
	"fmt"

	"m15.io/kappa/pkg/domain"
)

type ConfUsecase interface {
	GetConf(username string) (*domain.Conf, error)
}

type confUsecase struct {
	groupRepo domain.GroupRepository
}

func NewConfUsecase(groupRepo domain.GroupRepository) ConfUsecase {
	uc := new(confUsecase)
	uc.groupRepo = groupRepo

	return uc
}

func (usecase *confUsecase) GetConf(username string) (*domain.Conf, error) {
	conf := new(domain.Conf)
	conf.Username = username

	groups, err := usecase.groupRepo.GetGroupsOfUser(username)
	if err != nil {
		return conf, err
	}

	buttons := groupsToButtons(groups)
	conf.Buttons = buttons
	fmt.Println(conf)

	return conf, nil
}

func groupsToButtons(groups []domain.Group) []domain.Button {
	buttons := make([]domain.Button, 0, 25)
	for _, group := range groups {
		b := domain.Button{Text: group.CN, Value: group.Description}
		buttons = append(buttons, b)
	}
	return buttons
}

package repositories

import (
	"regexp"

	"m15.io/kappa/pkg/domain"
)

type LDAPClient interface {
	GetDNsOfUserGroups(username string) ([]string, error)
}

type LdapGroupsRepository struct {
	ldapClient LDAPClient
}

func (repo *LdapGroupsRepository) GetGroupsOfUser(username string) ([]domain.Group, error) {
	groups := make([]domain.Group, 0, 25)

	groupsDNs, err := repo.ldapClient.GetDNsOfUserGroups(username)
	if err != nil {
		return groups, err
	}

	for _, dn := range groupsDNs {
		groupname := getGroupCN(dn)
		group := domain.Group{
			Name: groupname,
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func getGroupCN(groupdn string) string {
	re := regexp.MustCompile("CN=([^,]+)")
	matches := re.FindStringSubmatch(groupdn)
	return matches[1]
}

package repositories

import (
	"m15.io/kappa/pkg/domain"
)

type LDAPClient interface {
	GetDNsOfUserGroups(username string) ([]string, error)
	GetGroupByDN(groupDN string) (string, string, error)
}

type LdapGroupRepository struct {
	ldapClient LDAPClient
}

func NewLdapGroupRepository(ldapClient LDAPClient) domain.GroupRepository {
	r := new(LdapGroupRepository)
	r.ldapClient = ldapClient

	return r
}

func (repo *LdapGroupRepository) GetGroupsOfUser(username string) ([]domain.Group, error) {
	groups := make([]domain.Group, 0, 25)

	groupsDNs, err := repo.ldapClient.GetDNsOfUserGroups(username)
	if err != nil {
		return groups, err
	}

	for _, dn := range groupsDNs {
		groupCN, groupDescr, err := repo.ldapClient.GetGroupByDN(dn)
		if err != nil {
			return groups, err
		}
		group := domain.Group{
			CN:          groupCN,
			Description: groupDescr,
		}
		groups = append(groups, group)
	}
	return groups, nil
}

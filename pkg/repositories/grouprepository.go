package repositories

import (
	"m15.io/kappa/pkg/domain"
)

type LDAPClient interface {
	GetDNsOfUsersGroups(username string) ([]string, error)
	GetGroupInfoByDN(groupDN string) (string, string, error)
}

// LdapGroupRepository is an implementation of GroupRepository
type ldapGroupRepository struct {
	ldapClient LDAPClient
}

// NewLdapGroupRepository returns creates and initializes ldapGroupRepository
func NewLdapGroupRepository(ldapClient LDAPClient) domain.GroupRepository {
	r := new(ldapGroupRepository)
	r.ldapClient = ldapClient

	return r
}

// GetGroupsOfUser fetches from LDAP groups that user belongs to. Returns list of Group{CN, Description}
func (repo *ldapGroupRepository) GetGroupsOfUser(username string) ([]domain.Group, error) {
	groups := make([]domain.Group, 0, 25)

	groupsDNs, err := repo.ldapClient.GetDNsOfUsersGroups(username)
	if err != nil {
		return groups, err
	}

	for _, dn := range groupsDNs {
		groupCN, groupDescr, err := repo.ldapClient.GetGroupInfoByDN(dn)
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

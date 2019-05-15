// Package ldap provides a simple ldap client to authenticate,
// retrieve basic information and groups for a user.
package infrastructure

import (
	"crypto/tls"
	"errors"
	"fmt"

	"gopkg.in/ldap.v2"
)

// dodać pole prefix
type LDAPClient struct {
	Attributes         []string
	Base               string
	BindDN             string
	BindPassword       string
	GroupFilter        string // e.g. "(memberUid=%s)"
	Host               string
	ServerName         string
	UserFilter         string // e.g. "(uid=%s)"
	Conn               *ldap.Conn
	Port               int
	InsecureSkipVerify bool
	UseSSL             bool
	SkipTLS            bool
	ClientCertificates []tls.Certificate // Adding client certificates
}

// Connect connects to the ldap backend.
func (lc *LDAPClient) Connect() error {
	if lc.Conn == nil {
		var l *ldap.Conn
		var err error
		address := fmt.Sprintf("%s:%d", lc.Host, lc.Port)
		if !lc.UseSSL {
			l, err = ldap.Dial("tcp", address)
			if err != nil {
				return err
			}

			// Reconnect with TLS
			if !lc.SkipTLS {
				err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
				if err != nil {
					return err
				}
			}
		} else {
			config := &tls.Config{
				InsecureSkipVerify: lc.InsecureSkipVerify,
				ServerName:         lc.ServerName,
			}
			if lc.ClientCertificates != nil && len(lc.ClientCertificates) > 0 {
				config.Certificates = lc.ClientCertificates
			}
			l, err = ldap.DialTLS("tcp", address, config)
			if err != nil {
				return err
			}
		}

		lc.Conn = l
	}
	return nil
}

// Close closes the ldap backend connection.
func (lc *LDAPClient) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}

// Bind binds to LDAP server using read-only users credentials
func (lc *LDAPClient) Bind() error {
	if lc.BindDN != "" && lc.BindPassword != "" {
		err := lc.Conn.Bind(lc.BindDN, lc.BindPassword)
		return err
	}

	return errors.New("BindDN and BindPassword can not be empty")
}

// GetDNsOfUsersGroups returns list of group DNs that user belongs to
func (lc *LDAPClient) GetDNsOfUsersGroups(username string) ([]string, error) {
	err := lc.Connect()
	if err != nil {
		return nil, err
	}

	// First bind with a read only user
	err = lc.Bind()
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		lc.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.UserFilter, username),
		[]string{"memberOf"},
		nil,
	)
	searchResult, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if len(searchResult.Entries) == 0 {
		return nil, errors.New("Could not find user")
	}

	groups := searchResult.Entries[0].GetAttributeValues("memberOf")
	lc.Close()
	return groups, nil
}

// GetGroupInfoByDN returns CN and Description of a group described by groupDN
func (lc *LDAPClient) GetGroupInfoByDN(groupDN string) (string, string, error) {
	err := lc.Connect()
	if err != nil {
		return "", "", err
	}

	// First bind with a read only user
	err = lc.Bind()
	if err != nil {
		return "", "", err
	}

	searchRequest := ldap.NewSearchRequest(
		lc.Base, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=group)(distinguishedName=%v))", groupDN), // The filter to apply
		[]string{"dn", "cn", "description"},                                  // A list attributes to retrieve
		nil,
	)

	searchResult, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return "", "", err
	}

	groupCN := searchResult.Entries[0].GetAttributeValue("cn")
	groupDescr := searchResult.Entries[0].GetAttributeValue("description")

	lc.Close()
	return groupCN, groupDescr, nil
}

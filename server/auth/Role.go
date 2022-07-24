package auth

type Role struct {
	Name        string      `json:"name"`
	Authorities []Authority `json:"authority"`
}

var adminAuthorities = []Authority{AUTHORITY_BROWSE, AUTHORITY_EDIT}
var visitorAuthorities = []Authority{AUTHORITY_BROWSE}

var (
	ROLE_ADMIN   = &Role{"admin", adminAuthorities}
	ROLE_VISITOR = &Role{"visitor", visitorAuthorities}
)

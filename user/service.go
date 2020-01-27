package user

import (
	"github.com/hellyab/techreview/entities"
)

// UserService specifies application user related services
type UserService interface {
	Users() ([]entities.User, []error)
	User(id string) (*entities.User, []error)
	UserByUsername(username string) (*entities.User, []error)
	UpdateUser(user *entities.User) (*entities.User, []error)
	DeleteUser(id string) (*entities.User, []error)
	StoreUser(user *entities.User) (*entities.User, []error)
	// PhoneExists(phone string) bool
	EmailExists(email string) bool
	UserRoles(*entities.User) ([]entities.Role, []error)
}

// RoleService speifies application user role related services
type RoleService interface {
	Roles() ([]entities.Role, []error)
	Role(id string) (*entities.Role, []error)
	RoleByName(name string) (*entities.Role, []error)
	UpdateRole(role *entities.Role) (*entities.Role, []error)
	DeleteRole(id string) (*entities.Role, []error)
	StoreRole(role *entities.Role) (*entities.Role, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionID string) (*entities.Session, []error)
	StoreSession(session *entities.Session) (*entities.Session, []error)
	DeleteSession(sessionID string) (*entities.Session, []error)
}

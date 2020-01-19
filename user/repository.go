package user

import (
	"github.com/hellyab/techreview/entities"
)

// UserRepository specifies application user related database operations
type UserRepository interface {
	Users() ([]entities.User, []error)
	User(id string) (*entities.User, []error)
	UserByEmail(email string) (*entities.User, []error)
	UpdateUser(user *entities.User) (*entities.User, []error)
	DeleteUser(id string) (*entities.User, []error)
	StoreUser(user *entities.User) (*entities.User, []error)
	// PhoneExists(phone string) bool
	EmailExists(email string) bool
	// UserRoles(*entities.User) ([]entities.Role, []error)
}

// // RoleRepository speifies application user role related database operations
// type RoleRepository interface {
// 	Roles() ([]entities.Role, []error)
// 	Role(id uint) (*entities.Role, []error)
// 	RoleByName(name string) (*entities.Role, []error)
// 	UpdateRole(role *entities.Role) (*entities.Role, []error)
// 	DeleteRole(id uint) (*entities.Role, []error)
// 	StoreRole(role *entities.Role) (*entities.Role, []error)
// }

// // SessionRepository specifies logged in user session related database operations
// type SessionRepository interface {
// 	Session(sessionID string) (*entities.Session, []error)
// 	StoreSession(session *entities.Session) (*entities.Session, []error)
// 	DeleteSession(sessionID string) (*entities.Session, []error)
// }

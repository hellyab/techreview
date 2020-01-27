package repository

import (
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/user"
	"github.com/jinzhu/gorm"
)

// UserGormRepo Implements the menu.UserRepository interface
type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of UserGormRepo
func NewUserGormRepo(db *gorm.DB) user.UserRepository {
	return &UserGormRepo{conn: db}
}

// Users return all users from the database
func (userRepo *UserGormRepo) Users() ([]entities.User, []error) {
	users := []entities.User{}
	errs := userRepo.conn.Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return users, errs
}

// User retrieves a user by its id from the database
func (userRepo *UserGormRepo) User(id uint) (*entities.User, []error) {
	user := entities.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

// UserByEmail retrieves a user by its email address from the database
func (userRepo *UserGormRepo) UserByEmail(email string) (*entities.User, []error) {
	user := entities.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

// UpdateUser updates a given user in the database
func (userRepo *UserGormRepo) UpdateUser(user *entities.User) (*entities.User, []error) {
	usr := user
	errs := userRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user from the database
func (userRepo *UserGormRepo) DeleteUser(id uint) (*entities.User, []error) {
	usr, errs := userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = userRepo.conn.Delete(usr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// StoreUser stores a new user into the database
func (userRepo *UserGormRepo) StoreUser(user *entities.User) (*entities.User, []error) {
	usr := user
	errs := userRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// PhoneExists check if a given phone number is found
func (userRepo *UserGormRepo) PhoneExists(phone string) bool {
	user := entities.User{}
	errs := userRepo.conn.Find(&user, "phone=?", phone).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// EmailExists check if a given email is found
func (userRepo *UserGormRepo) EmailExists(email string) bool {
	user := entities.User{}
	errs := userRepo.conn.Find(&user, "email=?", email).GetErrors()
	if len(errs) > 0 {
		return false
	}
	return true
}

// UserRoles returns list of application roles that a given user has
func (userRepo *UserGormRepo) UserRoles(user *entities.User) ([]entities.Role, []error) {
	userRoles := []entities.Role{}
	errs := userRepo.conn.Model(user).Related(&userRoles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return userRoles, errs
}

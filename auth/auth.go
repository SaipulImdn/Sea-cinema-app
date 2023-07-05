package auth

import (
	"errors"

	"github.com/SaipulImdn/sea-cinema-app/db"
)

// AuthService provides authentication and authorization functionalities
type AuthService struct {
	db *db.Database
}

// NewAuthService creates a new instance of the authentication service
func NewAuthService(database *db.Database) *AuthService {
	return &AuthService{
		db: database,
	}
}

// RegisterUser registers a new user in the system
func (a *AuthService) RegisterUser(username, password, name string, age int) error {
	// Implement user registration logic here
	// Example code to check if a user with the given username already exists
	if a.db.UserExists(username) {
		return errors.New("username already exists")
	}

	// Example code to create a new user and add it to the database
	user := &db.User{
		Username: username,
		Password: password,
		Name:     name,
		Age:      age,
	}
	a.db.AddUser(user)

	return nil
}

// LoginUser performs user login and returns the user object if successful
func (a *AuthService) LoginUser(username, password string) (*db.User, error) {
	// Implement user login logic here
	// Example code to check if the username and password match a user in the database
	user := a.db.FindUserByUsername(username)
	if user == nil || user.Password != password {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

// LogoutUser logs out the currently logged in user
func (a *AuthService) LogoutUser() {
	// Implement user logout logic here
	// Example code to clear the logged-in user session
	a.db.ClearLoggedInUser()
}

// ... Implement additional authentication and authorization methods as needed

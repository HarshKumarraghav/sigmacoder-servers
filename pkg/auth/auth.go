package auth

// The `import` statement is used to import packages that are required for the code to run. In this
// case, the code is importing the `http` package, the `auth` package from the `sigmacoder/pkg` directory,
// and the `fiber` package from the `github.com/gofiber/fiber/v2` repository. These packages are used
// in the code to handle HTTP requests and responses, and to implement user authentication
// functionality.
import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	Email string `json:"email"`
	Password    string `json:"password"`
}

// The above type defines a user with various properties such as ID, name, password, phone number,
// email, and gender.
// @property {string} ID - A unique identifier for the user, typically stored as a string.
// @property {string} Name - The name of the user.
// @property {string} Password - The password property is a string that stores the user's password. It
// is important to ensure that this property is properly secured and encrypted to prevent unauthorized
// access to the user's account.
// @property {string} PhoneNumber - The phone number of the user.
// @property {string} ProfilePic - The property "ProfilePic" is a string that represents the URL or
// file path of the user's profile picture.
// @property {string} Email - The email address of the user.
// @property {string} Username - The username property is a string that represents the unique username
// of a user. It is used for authentication and identification purposes.
// @property {string} UserType - UserType is a property of the User struct that represents the type of
// user. It can be used to differentiate between different types of users, such as regular users,
// administrators, or moderators. The value of UserType can be set to any string that represents the
// type of user.
// @property {string} DateOfBirth - This property represents the date of birth of a user. It is stored
// as a string in the format of "YYYY-MM-DD".
// @property {string} Gender - The gender of the user. It can be a string value such as "male",
// "female", "non-binary", etc.
// @property CreatedAt - CreatedAt is a property of the User struct that represents the date and time
// when the user was created. It is of type time.Time and is formatted as "YYYY-MM-DD HH:MM:SS". This
// property can be used to track when a user was added to a system or database.

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	UserType    string    `json:"usertype"`
	DateOfBirth string    `json:"dob"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
}

// The above type defines the structure of an input user object in Go, with various fields such as
// name, password, phone number, email, and gender.
// @property {string} Name - The name of the user.
// @property {string} Password - The "Password" property is a string that represents the user's
// password. It is likely used for authentication purposes to ensure that only authorized users can
// access the system or application. It is important to ensure that passwords are securely stored and
// encrypted to prevent unauthorized access.
// @property {string} PhoneNumber - PhoneNumber is a property of the InUser struct that represents the
// phone number of a user. It is of type string and is tagged with `json:"phonenumber"` to specify its
// name in JSON serialization.
// @property {string} ProfilePic - ProfilePic is a property of the InUser struct that represents the
// URL or file path of the user's profile picture. It is of type string and is tagged with
// `json:"profilepic"` to indicate that it should be marshaled and unmarshaled as "profilepic" in JSON.
// @property {string} Email - Email is a property of the InUser struct which represents the email
// address of a user. It is of type string and is tagged with `json:"email"` for JSON serialization.
// @property {string} Username - The username property is a string that represents the unique
// identifier for a user's account. It is often used as a way for users to log in to their account and
// can be displayed publicly on their profile.
// @property {string} DateOfBirth - This property represents the date of birth of a user. It is stored
// as a string in the format "YYYY-MM-DD".
// @property {string} Gender - Gender refers to the classification of individuals based on their
// biological sex, typically male or female. It is often used as a demographic variable in various
// contexts, including social, medical, and legal. In the context of the InUser struct, it is a
// property that stores the gender of a user.
// @property {string} UserType - UserType is a property of the InUser struct that represents the type
// of user. It can be used to differentiate between different types of users, such as regular users,
// administrators, or moderators. The value of UserType can be set to any string that represents the
// type of user.
type InUser struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
	ProfilePic  string `json:"profilepic"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dob"`
	Gender      string `json:"gender"`
	UserType    string `json:"usertype"`
}

// The above type defines the structure of an output user object in Go, including various user details
// such as ID, name, email, phone number, profile picture, user type, username, date of birth, gender,
// and creation date.
// @property {string} ID - A unique identifier for the user, typically stored as a string.
// @property {string} Name - The name of the user.
// @property {string} Email - The email address of the user.
// @property {string} PhoneNumber - The phone number of the user.
// @property {string} ProfilePic - ProfilePic is a property of the OutUser struct that represents the
// URL or file path of the user's profile picture.
// @property {string} UserType - UserType is a property of the OutUser struct that represents the type
// of user. It could be a customer, admin, or any other type of user.
// @property {string} Username - The username property is a string that represents the unique username
// of a user. It is used for authentication and identification purposes.
// @property {string} DateOfBirth - DateOfBirth is a property of the OutUser struct that represents the
// date of birth of a user. It is of type string and is represented in the format "YYYY-MM-DD".
// @property {string} Gender - The gender of the user. It can be male, female, non-binary, or any other
// gender identity.
// @property CreatedAt - CreatedAt is a property of the OutUser struct that represents the date and
// time when the user was created. It is of type time.Time and is formatted as "YYYY-MM-DD HH:MM:SS".
type OutUser struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	UserType    string    `json:"user_type"`
	Username    string    `json:"username"`
	DateOfBirth string    `json:"dob"`
	Gender      string    `json:"gender"`
	CreatedAt   time.Time `json:"created_at"`
}

// The `ToUser()` function is a method of the `InUser` struct that converts an input user object of
// type `InUser` to an output user object of type `User`. It generates a new UUID for the user ID,
// hashes the user's password using the `hashPassword()` function, and sets the remaining user
// properties based on the input `InUser` object. The function returns a new `User` object with the
// generated UUID and hashed password, along with the other user properties.
func (in *InUser) ToUser() User {
	uuid := uuid.New().String()
	return User{
		ID:          uuid,
		Name:        in.Name,
		ProfilePic:  in.ProfilePic,
		PhoneNumber: in.PhoneNumber,
		Password:    hashPassword(in.Password),
		Email:       in.Email,
		UserType:    in.UserType,
		Username:    in.Username,
		DateOfBirth: in.DateOfBirth,
		Gender:      in.Gender,
		CreatedAt:   time.Now(),
	}
}

// The `func (u *User) ToOutUser() OutUser` method is a function that takes a `User` object as a
// receiver and returns an `OutUser` object. It converts a `User` object to an `OutUser` object by
// mapping the properties of the `User` object to the corresponding properties of the `OutUser` object.
// The function creates a new `OutUser` object and sets its properties to the values of the
// corresponding properties of the `User` object. The resulting `OutUser` object is then returned.
func (u *User) ToOutUser() OutUser {
	return OutUser{
		ID:          u.ID,
		Name:        u.Name,
		ProfilePic:  u.ProfilePic,
		PhoneNumber: u.PhoneNumber,
		UserType:    u.UserType,
		Email:       u.Email,
		Username:    u.Username,
		DateOfBirth: u.DateOfBirth,
		Gender:      u.Gender,
		CreatedAt:   u.CreatedAt,
	}
}

// The function takes a password string, generates a hash using bcrypt algorithm with minimum cost, and
// returns the hash as a string.
func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes)
}

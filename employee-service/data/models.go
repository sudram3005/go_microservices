package data

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = time.Second * 3

var db *sqlx.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sqlx.DB) Models {
	db = dbPool

	return Models{
		User: User{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	User User
}

// User is the structure which holds one user from the database.
type User struct {
	ID                 int       `json:"id"`
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name,omitempty"`
	LastName           string    `json:"last_name,omitempty"`
	Role               string    `json:"role"`
	PrimaryTechStack   string    `json:"primaryTechStack"`
	SecondaryTechStack string    `json:"secondaryTechStack"`
	Password           string    `json:"-"`
	Active             int       `json:"active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// GetAll returns a slice of all users, sorted by last name
func (u *User) GetAllUsers() (*[]User, error) {
	query := "select * from users order by id ASC"
	users := []User{}
	err := db.Select(&users, query)
	if err != nil {
		log.Println("Error scanning", err)
		return nil, err
	}
	return &users, nil
}

// GetOne returns one user by id
func (u *User) GetOne(id int) (*User, error) {
	query := "select * from users where id=$1"
	user := User{}
	err := db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users set
		email = $1,
		first_name = $2,
		last_name = $3,
		role = $4,
		primaryTechStack = $5,
		secondaryTechStack = $6,
		user_active = $7,
		updated_at = $8
		where id = $9
	`

	_, err := db.ExecContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Role,
		u.PrimaryTechStack,
		u.SecondaryTechStack,
		u.Active,
		time.Now(),
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into users (email, first_name, last_name, role, primaryTechStack, secondaryTechStack, password, user_active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id`

	err = db.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role,
		user.PrimaryTechStack,
		user.SecondaryTechStack,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password = $1 where id = $2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, u.ID)
	if err != nil {
		return err
	}

	return nil
}

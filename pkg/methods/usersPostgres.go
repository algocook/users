package methods

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	// Нужно для подключения тест
	_ "github.com/lib/pq"
)

// User struct
type User struct {
	ID          int64  `db:"id"`
	Username    string `db:"username"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

// PostgresClient struct
type PostgresClient struct {
	db *sqlx.DB
}

// NewPostgresClient init
func NewPostgresClient(dbname string) (*PostgresClient, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=postgres user=users password=users dbname=%s sslmode=disable", dbname))
	if err != nil {
		return nil, err
	}

	return &PostgresClient{db}, nil
}

// GetUserByID method
func (client *PostgresClient) GetUserByID(id int64) (User, error) {
	user := User{}
	err := client.db.Get(&user, fmt.Sprintf("SELECT * FROM users WHERE id=%d", id))
	if err != nil {
		return User{}, errors.New("Fucked up in reading user")
	}

	return user, nil
}

// CheckUsernameAvialability method
func (client *PostgresClient) CheckUsernameAvialability(username string) (bool, error) {
	var isNotAvailable bool
	err := client.db.Get(&isNotAvailable, fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM users WHERE username='%s')", username))
	if isNotAvailable || err != nil {
		return false, err
	}
	return true, nil
}

// InsertNewUser method
func (client *PostgresClient) InsertNewUser(username string, title string, description string) (int64, error) {
	resCheck, _ := client.CheckUsernameAvialability(username)
	if !resCheck {
		return 0, errors.New("Username is not available")
	}

	fmt.Print(resCheck)

	var lastInsertedID int64
	err := client.db.QueryRow("INSERT INTO users (username, title, description) VALUES ($1, $2, $3) RETURNING id", username, title, description).Scan(&lastInsertedID)
	if err != nil {
		return 0, errors.New("Fucked up in inserting user")
	}

	return lastInsertedID, nil
}

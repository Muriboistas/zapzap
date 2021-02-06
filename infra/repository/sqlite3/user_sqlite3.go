package sqlite3

import (
	"database/sql"

	"github.com/muriboistas/zapzap/entity"
)

// User repo
type User struct {
	db *sql.DB
}

// NewUser create new repository
func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

// Get user
func (r *User) Get(urn string) (*entity.User, error) {
	var u entity.User
	rows, err := r.db.Query(getUserQuery, urn)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.URN, &u.Language, &u.CreatedAt, &u.UpdatedAt)
	}

	return &u, nil
}

// Create user
func (r *User) Create(e *entity.User) (entity.ID, error) {
	stmt, err := r.db.Prepare(createUserQuery)
	if err != nil {
		return e.ID, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.Name, e.URN, e.Language, e.CreatedAt)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Update user
func (r *User) Update(e *entity.User) error {
	stmt, err := r.db.Prepare(updateUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.URN, e.Language, e.UpdatedAt, e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete user
func (r *User) Delete(id entity.ID) error {
	stmt, err := r.db.Prepare(deleteUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

const (
	getUserQuery = `
		SELECT id, name, urn, language, created_at, updated_at
		FROM users
		WHERE urn = ?`
	createUserQuery = `
		INSERT INTO users(id, name, urn, language, created_at)
		VALUES(?, ?, ?, ?, ?)`
	updateUserQuery = `
		UPDATE users
		SET name = ?, urn = ?, language = ?, updated_at = ?
		WHERE id = ?`
	deleteUserQuery = `
		DELETE FROM users
		WHERE id = ?`
)

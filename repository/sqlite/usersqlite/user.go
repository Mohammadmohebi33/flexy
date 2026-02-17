package usersqlite

import (
	"database/sql"
	"errors"
	"flexy/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (d *DB) RegisterUser(u entity.User) (entity.User, error) {

	query := `
        INSERT INTO users (name, email, password, is_verified, created_at, updated_at)
        VALUES (?,?, ?, 0, ?, ?)
    `

	now := time.Now()
	result, err := d.conn.Conn().Exec(query, u.Name, u.Email, u.Password, now, now)
	if err != nil {
		return entity.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	u.ID = int(id)
	u.IsVerified = false
	u.CreatedAt = now
	u.UpdatedAt = now

	return u, nil
}

func (d *DB) LoginUser(u entity.User) (entity.User, error) {

	query := `
        SELECT id, name, email, password, is_verified, created_at, updated_at
        FROM users
        WHERE email = ?
    `

	var user entity.User

	err := d.conn.Conn().QueryRow(query, u.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.New("invalid email or password")
		}
		return entity.User{}, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(u.Password),
	)
	if err != nil {
		return entity.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

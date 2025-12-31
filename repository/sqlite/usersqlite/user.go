package usersqlite

import (
	"flexy/entity"
	"time"
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

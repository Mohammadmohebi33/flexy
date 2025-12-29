package usersqlite

import "flexy/entity"

func (d *DB) RegisterUser(entity.User) (entity.User, error) {
	return entity.User{}, nil
}

package contract

import "todo_cli/entity"

type UserReadStore interface {
	Load() []entity.User
}

type UserWriteStore interface {
	Save(user entity.User)
}

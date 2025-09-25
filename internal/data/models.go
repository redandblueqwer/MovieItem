package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrConflict       = errors.New("edict conflict")
)

type Models struct {
	Movie       MovieModel
	Permissions PermissionModel
	Token       TokenModel
	User        UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movie:       MovieModel{DB: db},
		Permissions: PermissionModel{DB: db},
		Token:       TokenModel{DB: db},
		User:        UserModel{DB: db},
	}
}

package repos

import (
	"context"
	"fmt"
	"project/types"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserByUsername(context.Context, string) (*types.User, error)
	CreateUser(context.Context, *types.User) error
}

type userRepositoryImpl struct {
	dbConn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		dbConn: conn,
	}
}

const SQL_GET_USER_BY_USERNAME = `
	SELECT
		u.id,
		u.username,
		u.pass
	FROM
		"user" AS u
	WHERE u.username = $1;
`

// Get a User
func (repo *userRepositoryImpl) GetUserByUsername(c context.Context, username string) (*types.User, error) {
	row := repo.dbConn.QueryRow(c, SQL_GET_USER_BY_USERNAME, username)

	user := &types.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

const SQL_INSERT_USER = `
	insert INTO "user" (id, username, pass) VALUES ($1, $2, $3)
	RETURNING id;`

// Create a new user in DB
func (repo *userRepositoryImpl) CreateUser(c context.Context, user *types.User) error {
	var userId string
	err := repo.dbConn.QueryRow(c, SQL_INSERT_USER, user.Id, user.Username, user.Password).Scan(&userId)
	if err != nil {
		return fmt.Errorf("error during user creation: %v", err)
	}
	return nil
}

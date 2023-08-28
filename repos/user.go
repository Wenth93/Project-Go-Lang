package repos

import (
	"context"
	"fmt"

	"github.com/Wenth93/Project-Go-Lang/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUser(context.Context, string) (*types.User, error)
	CreateUser(context.Context, *types.User) error
	GetUserByUsername(context.Context, string) (*types.User, error)
}

type userRepositoryImpl struct {
	dbConn *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		dbConn: conn,
	}
}

const SQL_GET_USER = `
		select 
			u.id,
			u.username,
			u.pass
		from
			"user" as u
		where u.id = $1;`

func (repo *userRepositoryImpl) GetUser(c context.Context, userId string) (*types.User, error) {
	rows, err := repo.dbConn.Query(c, SQL_GET_USER, userId)
	if err != nil {
		return nil, fmt.Errorf("error during query to get user: %v", err)
	}

	if rows.Next() {
		user := &types.User{}
		err = rows.Scan(
			&user.Id,
			&user.Password,
			&user.Username,
		)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, nil
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

func (repo *userRepositoryImpl) GetUserByUsername(c context.Context, username string) (*types.User, error) {
	row := repo.dbConn.QueryRow(c, SQL_GET_USER_BY_USERNAME, username)

	user := &types.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

const SQL_INSERT_USER = `
	INSERT INTO "user" (id, username, pass) VALUES ($1, $2, $3)
	RETURNING id;`

func (repo *userRepositoryImpl) CreateUser(c context.Context, user *types.User) error {
	var userId string
	err := repo.dbConn.QueryRow(c, SQL_INSERT_USER, user.Id, user.Username, user.Password).Scan(&userId)
	if err != nil {
		return fmt.Errorf("error during user creation: %v", err)
	}
	return nil
}

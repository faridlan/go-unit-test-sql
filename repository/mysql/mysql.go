package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/faridlan/go-unit-test-sql/model"
	repo "github.com/faridlan/go-unit-test-sql/repository"
	_ "github.com/go-sql-driver/mysql"
)

type RepositoryImpl struct {
	DB *sql.DB
}

func NewRepository(dialect, dsn string, idleConn, maxConn int) (repo.Repository, error) {

	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &RepositoryImpl{db}, nil
}

func (repository *RepositoryImpl) Close() {
	repository.DB.Close()
}

func (repository *RepositoryImpl) FindById(id string) (*model.User, error) {
	user := new(model.User)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repository.DB.QueryRowContext(ctx, "select id, name, email, phone from users where id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *RepositoryImpl) FindAll() ([]*model.User, error) {
	users := make([]*model.User, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := repository.DB.QueryContext(ctx, "select id,name,email,phone from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := new(model.User)
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func (repository *RepositoryImpl) Create(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "insert into users (id, name, email, phone) values (?,?,?,?)"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Email, user.Phone)
	return err
}

func (repository *RepositoryImpl) Update(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "update users set name = ?, email = ?, phone = ?, where id = ?"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Phone, user.ID)
	return err
}

func (repository *RepositoryImpl) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "delete from users where id = ?"
	stmt, err := repository.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}

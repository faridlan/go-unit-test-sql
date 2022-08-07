package mysql

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/go-unit-test-sql/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var User = &model.User{
	ID:    uuid.New().String(),
	Name:  "Faridlan",
	Email: "faridlan@gmail.com",
	Phone: "087663527189",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connecntion", err)
	}

	return db, mock
}

func TestFindByIdSuccess(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "select id, name, email, phone from users where id = ?"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(User.ID, User.Name, User.Email, User.Phone)

	mock.ExpectQuery(query).WithArgs(User.ID).WillReturnRows(rows)

	user, err := repo.FindById(User.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindByIdFailed(t *testing.T) {

	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "select id, name, email, phone from users where id = ?"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"})

	mock.ExpectQuery(query).WithArgs(User.ID).WillReturnRows(rows)
	user, err := repo.FindById(User.ID)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestCreateSuccess(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "insert into users \\(id, name, email, phone\\) values \\(\\?,\\?,\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.ID, User.Name, User.Email, User.Phone).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(User)
	assert.NoError(t, err)
}

func TestCreateFailed(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "insert into users \\(id, name, email, phone\\) values \\(\\?,\\?,\\?,\\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.ID, User.Name, User.Email, User.Phone).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Create(User)
	assert.Nil(t, err)
}

func TestUpdateSuccess(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "update users set name = \\?, email = \\?, phone = \\?, where id = \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.Name, User.Email, User.Phone, User.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(User)
	assert.NoError(t, err)
}

func TestUpdateFailed(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "update users set name = \\?, email = \\?, phone = \\?, where id = \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.Name, User.Email, User.Phone, User.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Update(User)
	assert.Nil(t, err)
}

func TestDeleteSuccess(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "delete from users where id = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(User.ID)
	assert.NoError(t, err)
}

func TestDeleteFailed(t *testing.T) {
	db, mock := NewMock()
	repo := &RepositoryImpl{db}
	defer func() {
		repo.Close()
	}()

	query := "delete from users where id = ?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(User.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(User.ID)
	assert.Nil(t, err)
}

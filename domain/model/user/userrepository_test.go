package user

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
)

func Test_FindByUserName(t *testing.T) {
	userName, _ := NewUserName("userName")
	userId, _ := NewUserId("userId")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM users WHERE name = $1`)).
			WithArgs("userName").
			WillReturnRows(mock.NewRows([]string{"userId", "userName"}).AddRow("userId", "userName"))
		mock.ExpectCommit()

		got, err := userRepository.FindByUserName(userName)
		if err != nil {
			t.Error(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		want := &User{id: *userId, name: *userName}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(User{}, UserName{}, UserId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM users WHERE name = $1`)).
			WithArgs("userName").
			WillReturnRows(mock.NewRows([]string{}))
		mock.ExpectCommit()

		got, err := userRepository.FindByUserName(userName)
		if err != nil {
			t.Error(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if got != nil {
			t.Errorf("want: nil, got: %v", got)
		}
	})
}

func Test_Save(t *testing.T) {
	userName, _ := NewUserName("userName")
	userId, _ := NewUserId("userId")
	user, _ := NewUser(*userId, *userName)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO users").
			WithArgs("userId", "userName").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		got := userRepository.Save(user)
		if got != nil {
			t.Errorf("got must be nil, but %v", got)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("fail", func(t *testing.T) {
		var saveQueryRowError *SaveQueryRowError
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO users").
			WithArgs("userId", "userName").
			WillReturnError(saveQueryRowError)
		mock.ExpectRollback()

		got := userRepository.Save(user)
		if !errors.As(got, &saveQueryRowError) {
			t.Errorf("err type: %v, expect err type: %v", reflect.TypeOf(err), reflect.TypeOf(saveQueryRowError))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_FindByUserId(t *testing.T) {
	userName, _ := NewUserName("userName")
	userId, _ := NewUserId("userId")

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM users WHERE id = $1`)).
			WithArgs("userId").
			WillReturnRows(mock.NewRows([]string{"userId", "userName"}).AddRow("userId", "userName"))
		mock.ExpectCommit()

		got, err := userRepository.FindByUserId(userId)
		if err != nil {
			t.Error(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		want := &User{id: *userId, name: *userName}
		if diff := cmp.Diff(want, got, cmp.AllowUnexported(User{}, UserName{}, UserId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name FROM users WHERE id = $1`)).
			WithArgs("userId").
			WillReturnRows(mock.NewRows([]string{}))
		mock.ExpectCommit()

		got, err := userRepository.FindByUserId(userId)
		if err != nil {
			t.Error(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if got != nil {
			t.Errorf("want: nil, got: %v", got)
		}
	})
}

func Test_UpdateRepository(t *testing.T) {
	userName, _ := NewUserName("updateUserName")
	userId, _ := NewUserId("userId")

	user, _ := NewUser(*userId, *userName)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE users").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		got := userRepository.Update(user)
		if got != nil {
			t.Errorf("got must be nil, but %v", got)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("fail", func(t *testing.T) {
		var updateQueryRowError *UpdateQueryError
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE users").
			WillReturnError(updateQueryRowError)
		mock.ExpectRollback()

		got := userRepository.Update(user)
		if !errors.As(got, &updateQueryRowError) {
			t.Errorf("err type: %v, expect err type: %v", reflect.TypeOf(err), reflect.TypeOf(updateQueryRowError))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_DeleteRepository(t *testing.T) {
	userName, _ := NewUserName("updateUserName")
	userId, _ := NewUserId("userId")

	user, _ := NewUser(*userId, *userName)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%v' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userRepository, err := NewUserRepository(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM users").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		got := userRepository.Delete(user)
		if got != nil {
			t.Errorf("got must be nil, but %v", got)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("fail", func(t *testing.T) {
		var deleteQueryRowError *DeleteQueryError
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM users").
			WillReturnError(deleteQueryRowError)
		mock.ExpectRollback()

		got := userRepository.Delete(user)
		if !errors.As(got, &deleteQueryRowError) {
			t.Errorf("err type: %v, expect err type: %v", reflect.TypeOf(err), reflect.TypeOf(deleteQueryRowError))
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

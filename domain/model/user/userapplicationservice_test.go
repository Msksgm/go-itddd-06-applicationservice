package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Register(t *testing.T) {
	data := []struct {
		testname       string
		userName       string
		create         func(userName UserName) (*User, error)
		findByUserName func(name UserName) (*User, error)
		exists         func(user User) (bool, error)
		save           func(user User) error
		errMsg         string
	}{
		{
			"success",
			"userName",
			func(userName UserName) (*User, error) {
				return &User{name: UserName{value: "userName"}, id: UserId{value: "userId"}}, nil
			},
			func(name UserName) (*User, error) { return nil, nil },
			func(user User) (bool, error) { return false, nil },
			func(user User) error { return nil },
			"",
		},
	}
	userApplicationService := UserApplicationService{}
	userService := UserService{}

	for _, d := range data {
		t.Run(d.testname, func(t *testing.T) {
			userApplicationService.userRepository = &UserRepositorierStub{save: d.save}
			userService.userRepository = &UserRepositorierStub{findByUserName: d.findByUserName}
			userApplicationService.userService = userService

			err := userApplicationService.Register("userName")
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}

func Test_Get(t *testing.T) {
	data := []struct {
		testname     string
		userId       string
		findByUserId func(userId UserId) (*User, error)
		want         *UserData
	}{
		{
			"found",
			"userId",
			func(userId UserId) (*User, error) {
				return &User{name: UserName{value: "userName"}, id: UserId{value: "userId"}}, nil
			},
			&UserData{Id: "userId", Name: "userName"},
		},
		{
			"not found",
			"userId",
			func(userId UserId) (*User, error) {
				return nil, nil
			},
			nil,
		},
	}
	userApplicationService := UserApplicationService{}

	for _, d := range data {
		t.Run(d.testname, func(t *testing.T) {
			userApplicationService.userRepository = &UserRepositorierStub{findByUserId: d.findByUserId}

			got, err := userApplicationService.Get(d.userId)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported()); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
			var expectedErr *RegisterError
			if err != nil && !errors.As(err, &expectedErr) {
				t.Errorf("Expected error `%v`, got `%v`", reflect.TypeOf(err), reflect.TypeOf(expectedErr))
			}
		})
	}
}

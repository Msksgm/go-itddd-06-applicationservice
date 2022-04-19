package user

import (
	"testing"
)

type UserRepositorierStub struct {
	findByUserName func(name UserName) (*User, error)
	save           func(user User) error
}

func (urs *UserRepositorierStub) FindByUserName(name *UserName) (*User, error) {
	return urs.findByUserName(*name)
}

func (urs *UserRepositorierStub) Save(user *User) error {
	return urs.save(*user)
}

func Test_Exists(t *testing.T) {
	userId, _ := NewUserId("userId")
	userName, _ := NewUserName("userName")
	data := []struct {
		name           string
		findByUserName func(name UserName) (*User, error)
		want           bool
		user           *User
		testErrMsg     string
	}{
		{"exists", func(name UserName) (*User, error) { return &User{name: *userName, id: *userId}, nil }, true, &User{name: *userName, id: *userId}, "isExists must be true but false"},
		{"not exists", func(name UserName) (*User, error) { return nil, nil }, false, &User{name: *userName, id: *userId}, "isExists must be false but true"},
	}
	userService := UserService{}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			userService.userRepository = &UserRepositorierStub{findByUserName: d.findByUserName}
			got, err := userService.Exists(d.user)
			if err != nil {
				t.Fatal(err)
			}
			if got != d.want {
				t.Errorf(d.testErrMsg)
			}
		})
	}
}

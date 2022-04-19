package user

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserName(t *testing.T) {
	data := []struct {
		testName string
		userName string
		want     *UserName
		errMsg   string
	}{
		{"success", "username", &UserName{value: "username"}, ""},
		{"fail because value is less than 3 characters", "us", nil, "UserName is more than 3 characters."},
		{"fail because value is more than 20 characters", "usernameusernameusername", nil, "UserName is less than 20 characters."},
	}
	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			got, err := NewUserName(d.userName)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported(UserName{})); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
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

func Test_UserNameEquals(t *testing.T) {
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	data := []struct {
		testName      string
		otherUserName *UserName
		want          bool
	}{
		{"equal", &UserName{value: "username"}, true},
		{"not equal", &UserName{value: "otherusername"}, false},
	}
	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			got := userName.Equals(*d.otherUserName)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported()); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}

func Test_UserNameString(t *testing.T) {
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("UserName: [value: %s]", userName.value)
	got := fmt.Sprint(userName)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

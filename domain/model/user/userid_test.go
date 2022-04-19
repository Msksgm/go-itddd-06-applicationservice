package user

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserId(t *testing.T) {
	got, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}

	want := &UserId{value: "id"}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(UserId{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func Test_UserIdEquals(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	data := []struct {
		testName    string
		otherUserId *UserId
		want        bool
	}{
		{"equal", &UserId{value: "id"}, true},
		{"not equal", &UserId{value: "otherId"}, false},
	}
	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			got := userId.Equals(d.otherUserId)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported()); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}

func Test_UserIdString(t *testing.T) {
	userId, err := NewUserId("id")
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(userId)
	want := fmt.Sprintf("UserId [value: %s]", userId.value)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

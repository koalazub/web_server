package server

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	want := &Server{
		users: make(map[string]UserInfo),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got.users, want.users)
	}
}

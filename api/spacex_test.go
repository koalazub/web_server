package api

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	ls := New()
	want := &LaunchServer{map[string]CustomLaunchData{}}

	if !reflect.DeepEqual(ls, want) {
		t.Errorf("Got: %v | want: %v", ls, want)
	}
}

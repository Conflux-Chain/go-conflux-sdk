package utils

import "testing"

func TestGetObjJsonFieldTags(t *testing.T) {
	s := struct {
		Name string `json:"name"`
	}{}

	actual := GetObjJsonFieldTags(s)
	if actual[0] != "name" {
		t.Errorf("expect %v, actual %v", "name", actual[0])
	}
}

package gobot

import (
	"testing"
)

func TestSessionStorage_get(t *testing.T) {
	stor := NewSessionStorage()
	s1 := stor.get("aaa")
	s2 := stor.get("bbb")
	s3 := stor.get("aaa")

	s1.ValueMap["check"] = "12341234"
	s1.ValueMap["check1"] = "2341234"
	if s1 != s3 {
		t.Error("storages are different")
	}
	if s1 == s2 || s2 == s3 {
		t.Error("storages are not different")
	}
	if s1.ValueMap["check"] != s3.ValueMap["check"] {
		t.Error("storages are different")
	}
}

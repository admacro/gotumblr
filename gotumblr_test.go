package main

import (
	"testing"
)

func TestGetBlog(t *testing.T) {
	blogName := "nixiaoqing"
	bf := getBlog()
	if blogName != bf.Name {
		t.Errorf("Blog name is %q, want %q", bf.Name, blogName)
	}
}

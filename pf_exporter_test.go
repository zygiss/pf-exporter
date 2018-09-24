package main

import (
	"testing"
)

func TestCanGetStats(t *testing.T) {
	filename := "/home/zygis/tmp/pf.stats"
	got, err := getStats(filename)
	if err != nil {
		t.Fatalf("getStats('%v') err: %v", filename, err)
	}
	if len(got) == 0 {
		t.Errorf("getStats(%v) = %v", filename, got)
	}
}

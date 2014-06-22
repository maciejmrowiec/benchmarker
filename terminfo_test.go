package main

import (
	"testing"
)

func TestGetColumnCount(t *testing.T) {
	ti := TermInfoNew()
	_, err := ti.GetColumnCount()
	if err != nil {
		t.Failed()
	}
}

func TestGetLineCount(t *testing.T) {
	ti := TermInfoNew()
	_, err := ti.GetLineCount()
	if err != nil {
		t.Failed()
	}
}

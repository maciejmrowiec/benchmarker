package main

import (
	"os/exec"
	"strconv"
)

const (
	TPUT_CMD     = "tput"
	TPUT_COLUMNS = "cols"
	TPUT_LINES   = "lines"
)

type TermInfo struct {
	path string
}

func TermInfoNew() *TermInfo {
	return new(TermInfo)
}

func (t *TermInfo) SetPath(path string) {
	t.path = path
}

func (t *TermInfo) GetColumnCount() (int, error) {
	out, err := t.getAttribute(TPUT_COLUMNS)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(out)
}

func (t *TermInfo) GetLineCount() (int, error) {
	out, err := t.getAttribute(TPUT_LINES)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(out)
}

func (t *TermInfo) getAttribute(attibute string) (string, error) {

	out, err := exec.Command(t.path+TPUT_CMD, attibute).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

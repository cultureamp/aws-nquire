package command

import (
	"testing"
)

func TestMatchByRegex(t *testing.T) {
	str := "ca-base-ami-01010101"
	p := "ca-base*"

	actual := matchByRegex(p, str)
	if !actual {
		t.Fail()
	}
}

func TestNotMatchByRegex(t *testing.T) {
	str := "ca-docker-ami-01010101"
	p := "ca-base*"

	actual := matchByRegex(p, str)
	if actual {
		t.Fail()
	}

	blank := ""
	actual = matchByRegex(p, blank)
	if actual {
		t.Fail()
	}
}

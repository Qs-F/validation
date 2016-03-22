package validation

import (
	"regexp"
	"unicode/utf8"
)

func Required(s string) (r bool) {
	if s == "" {
		r = false
	} else {
		r = true
	}
	return
}

func MaxSize(s string, max int) (r bool) {
	if utf8.RuneCountInString(s) <= max {
		r = true
	} else {
		r = false
	}
	return
}

func MinSize(s string, min int) (r bool) {
	if utf8.RuneCountInString(s) >= min {
		r = true
	} else {
		r = false
	}
	return
}

func OnlyAlphabet(s string) (r bool) {
	reg, _ := regexp.Compile("^[a-zA-Z]*$")
	if reg.MatchString(s) {
		r = true
	} else {
		r = false
	}
	return
}

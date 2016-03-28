package validation

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// Basic funcs

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

func OnlySafeString(s string) (r bool) {
	reg, _ := regexp.Compile("^[a-zA-Z0-9-_.]*$")
	if reg.MatchString(s) {
		r = true
	} else {
		r = false
	}
	return
}

func RegexpMatch(s string, pattern string) (r bool) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		r = false
		return
	}
	if reg.MatchString(s) {
		r = true
	} else {
		r = false
	}
	return
}

// Advanced funcs
// future design:
// 	v := validation.Validation{name}
// 	v.Required() // if not setted Error, set default err message
// 	v.MaxSize(2, "user name must be bigger than 3 chars.")
// 	for _,m := range v.Errors() {
// 		println(m)
// 	}

type Validation struct {
	String           string
	validationErrors []error
}

func SetValue(s string) *Validation {
	return &Validation{s, nil}
}

const (
	RequiredDefaultMessage_       = "field is required."
	MaxSizeDefaultMessage_        = "field is too big."
	MinSizeDefaultMessage_        = "field is too small."
	OnlyAlphabetDefaultMessage_   = "field is required only alphabet."
	OnlySafeStringDefaultMessage_ = "field is required only a-z, A-Z, and some signs(_, - , .)."
	RegexpMatchDefaultMessage_    = "field is not fullfilled."
)

func processErrMsg(defaultMsg string, s ...string) (r string) {
	if len(s) <= 0 {
		r = defaultMsg
	} else {
		r = s[0]
	}
	return
}

func (v *Validation) Required(errMsg ...string) (ok bool) {
	ok = Required(v.String)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(RequiredDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) MaxSize(max int, errMsg ...string) (ok bool) {
	ok = MaxSize(v.String, max)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(MaxSizeDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) MinSize(min int, errMsg ...string) (ok bool) {
	ok = MinSize(v.String, min)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(MinSizeDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) OnlyAlphabet(errMsg ...string) (ok bool) {
	ok = OnlyAlphabet(v.String)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(OnlyAlphabetDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) OnlySafeString(errMsg ...string) (ok bool) {
	ok = OnlySafeString(v.String)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(OnlySafeStringDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) RegexpMatch(pattern string, errMsg ...string) (ok bool) {
	ok = RegexpMatch(v.String, pattern)
	if !ok {
		v.validationErrors = append(v.validationErrors, errors.New(processErrMsg(RegexpMatchDefaultMessage_, errMsg...)))
	}
	return
}

func (v *Validation) Errors() (r []error) {
	return v.validationErrors
}

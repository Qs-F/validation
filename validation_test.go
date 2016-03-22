package validation

import (
	"testing"
)

func TestRequired(t *testing.T) {
	if Required("abcd") {
		t.Log("Required Test(abcd): ok")
	} else {
		t.Error("Required Test(abcd): false")
	}
	if Required("") {
		t.Error("Required Test( ): false")
	} else {
		t.Log("Required Test( ): ok")
	}
}

func TestMaxSize(t *testing.T) {
	tests := []string{"abc", "123", "あいう", "   ", "ab", "12", "あい", " ab", "abcd", "ab123", "あいうえ", "𡈽𡈽𡈽𡈽"}
	for i, v := range tests {
		if i < 8 {
			if MaxSize(v, 3) {
				t.Log("MaxSize Test(" + v + "): ok")
			} else {
				t.Error("MaxSize Test(" + v + "): false")
			}
		} else {
			if MaxSize(v, 3) {
				t.Error("MaxSize Test(" + v + "): false")
			} else {
				t.Log("MaxSize Test(" + v + "): ok")
			}
		}
	}
}

func TestMinSize(t *testing.T) {
	tests := []string{"abc", "123", "あいう", "   ", " ab", "abcd", "ab123", "あいうえ", "  ", "ab", "12", "あい", "𡈽𡈽"}
	for i, v := range tests {
		if i < 8 {
			if MinSize(v, 3) {
				t.Log("MaxSize Test(" + v + "): ok")
			} else {
				t.Error("MaxSize Test(" + v + "): false")
			}
		} else {
			if MinSize(v, 3) {
				t.Error("MaxSize Test(" + v + "): false")
			} else {
				t.Log("MaxSize Test(" + v + "): ok")
			}
		}
	}
}

func TestOnlyAlphabet(t *testing.T) {
	tests := []string{"abcd", "ab1234", " ", " ab", "ab ", "あいうえお", " あいうえお", "あいうえお "}
	for i, v := range tests {
		if i != 0 {
			if OnlyAlphabet(v) {
				t.Error("OnlyAlphabet Test(" + v + "): false")
			} else {
				t.Log("OnlyAlphabet Test(" + v + "): ok")
			}
		} else {
			if OnlyAlphabet(v) {
				t.Log("OnlyAlphabet Test(abcd): ok")
			} else {
				t.Error("OnlyAlphabet Test" + v + "): false")
			}
		}
	}
}

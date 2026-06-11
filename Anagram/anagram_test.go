package anagram

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
    assert.Equal(t, true, Anagram("care", "race"))
}

func TestNotAnagram(t *testing.T) {
    assert.Equal(t, false, Anagram("cat", "rat"))
}

func TestAnagramCaseSensitivity(t *testing.T) {
    assert.Equal(t, false, Anagram("SiLenT", "lIstEn"))
}

func TestAnagramSpacesInString(t *testing.T) {
    assert.Equal(t, false, Anagram("ca r e", "r a c e"))
}
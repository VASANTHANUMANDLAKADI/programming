package main
import "testing"

func TestAnagram(t*testing.T) {
	result := Anagram("care", "race")

	if result != true {
		t.Errorf("Expected true, got false")
	}
}

func TestNotAnagram(t*testing.T){
	result := Anagram("cat", "rat")

	if result != false {
		t.Errorf("Expected false, got true")
	}
}

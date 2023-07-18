package feature

import (
	"strings"
	"testing"
)

func TestCheckIfWorktreeIsClean(t *testing.T) {
	if CheckIfWorktreeIsClean("../git_test") {
		t.Error("Not works")
	}
}

func TestFormatFeatureName(t *testing.T) {
	featureDescription := "Feature Command"
	formatedFeatureName := FormatFeatureName(featureDescription)
	expectedFeature := "feature_command"
	if formatedFeatureName != expectedFeature {
		t.Errorf("Expected %v branch name", expectedFeature)
	}
}

func TestRemoveAccents(t *testing.T) {
	helloHowAreYouInPortuguese := "Olá, como está?"
	expected := "Ola, como esta?"
	result := RemoveAccents(helloHowAreYouInPortuguese)
	if result != expected {
		t.Errorf("Expected %v", expected)
	}
}

func TestRemoveSpecialChars(t *testing.T) {
	cannotHaveQuestionMark := "will this on the 1 try?"
	expected := strings.Replace(cannotHaveQuestionMark, "?", "", 1)
	result := RemoveSpecialChars(cannotHaveQuestionMark)
	if result != expected {
		t.Error("There is still a question mark on the string")
	}
}

package feature

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"

	"github.com/AlecAivazis/survey/v2"
	"github.com/KaueSabinoSRV17/Flower/internal/catch"
	"github.com/KaueSabinoSRV17/Flower/internal/staging"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func CheckIfWorktreeIsClean(repo string) bool {
	var modifiedFiles []string
	files := staging.GetAllFilesFromStatus(repo)
	for _, file := range files {
		if file == "" {
			continue // There is always an empty string in the list. If we reached and dont continue the loop, we'll have some problems
		}
		status, fileName := staging.ExtractStatusAndFile(file)
		if staging.IsModified(status) {
			modifiedFiles = append(modifiedFiles, fileName)
		}
	}
	cleanWorktree := len(modifiedFiles) == 0
	return cleanWorktree
}

func AskToStashUncommitedChanges() bool {
	var stashChanges bool 
	err := survey.AskOne(
		&survey.Confirm{
			Message: "You have uncommited changes, so you cannot checkout to a new branch. Do you want to stash your changes?",
		},
		&stashChanges,
	)
	if err != nil {
		log.Fatal("Could not ask if user wants to stash changes")
	}
	return stashChanges
}

func AskNewFeatureName() string {
	var featureName string
	err := survey.AskOne(
		&survey.Input{
			Message: "Describe The New Feature:",
		},
		&featureName,
	)
	catch.HandleError("Could not ask new feature name", err)
	formatedName := FormatFeatureName(featureName)
	return formatedName
}

func FormatFeatureName(featureDescription string) string {
	lowerCased := strings.ToLower(featureDescription)
	spacesReplacedByUnderlines := strings.Replace(lowerCased, " ", "_", 1)
	noAccents := RemoveAccents(spacesReplacedByUnderlines)
	noSpecialChars := RemoveSpecialChars(noAccents)
	fullResult := fmt.Sprintf("feature/%v", noSpecialChars)
	return fullResult
}

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func RemoveSpecialChars(s string) string {
	anythingThatIsNotSpacesOrLowerCasedChars := regexp.MustCompile(`[^a-z0-9 ]+`)
	result := anythingThatIsNotSpacesOrLowerCasedChars.ReplaceAllString(s, "")
	return result
}

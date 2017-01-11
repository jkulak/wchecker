package main

import (
	"regexp"
	"strings"
)

/*

1. Consistent hyphenation
    - single hypen front-end vs. front-end
    - TODO: what about double hyphens? front--end?
    - TODO what about tripple and more hypends? front---end, front----end?

2. Do you always use space after puctuation marks? (but what about i.e.?)
3. Capital letter after dots (whta about i.e. )

*/

/*

- Should return list of errors and list of suggestions/warnings - like linter
- should return a list of structs with
- should support a list of checks to run
- should support different languages (string punctuation removes all Polish characters)
- should have a list of checks not dependent on a language
- should have a list of checks per language


- sholud have linting options
    - remove whitespace?
    - new line at the end?
*/

// ResultTypeError defines an error
const ResultTypeError = 0

// ResultTypeWarning defines a warning
const ResultTypeWarning = 1

// CheckResult defines type of the result and gives a text explanation
type CheckResult struct {
	resType     byte
	description string
}

// string punctuation, leaving n-dashes and spaces and newlines
func stripPunctuation(s *string) string {
	// match not whitespace and not word characters and 4 different dashes
	res := regexp.MustCompile(`[^\n0-9A-Za-z \-]`).ReplaceAllString(*s, " ")
	// remove dashes that are not inside the words
	res = regexp.MustCompile(`\s-\w|\w-\s|\s-\s|^-|-$`).ReplaceAllString(res, "")
	return res
}

func trimWhiteSpace(s *string) string {
	// `\s+` - replaces all whitespace (tabs, newlines)
	return regexp.MustCompile(` +`).ReplaceAllString(*s, " ")
}

// Replace all kinds of four dashes with the regular n-dash "-"
func normalizeDashes(s string) string {
	return regexp.MustCompile(`‒|—|―`).ReplaceAllString(s, `-`)
}

func checkHyphenationConsistency(s string) []CheckResult {
	var res []CheckResult

	// true - hyphenated, false - not hyphenatd
	usedWords := make(map[string]bool)

	s = normalizeDashes(s)

	// process line by line
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			hyphenated := regexp.MustCompile(`-`).MatchString(word)
			word := regexp.MustCompile(`-`).ReplaceAllString(word, "")
			if _, isset := usedWords[word]; !isset {
				usedWords[word] = hyphenated
			} else {
				if usedWords[word] != hyphenated {
					res = append(res, CheckResult{ResultTypeError, "Inconsistent use of hyphen in word: " + word})
				}
			}
		}
	}
	return res
}

// Check runs list of checkers on given string and returns an array of errors
func Check(s *string) []CheckResult {
	res := []CheckResult{}

	// return no errors if string has 0 length
	len := len(*s)
	if len == 0 {
		return res
	}

	// immediately return no errors if string has no whitespace
	re := regexp.MustCompile(`\s`)
	m := re.MatchString(*s)
	if !m {
		return res
	}

	stripped := stripPunctuation(s)
	trimmed := trimWhiteSpace(&stripped)

	return append(res, checkHyphenationConsistency(trimmed)...)
}

package main

import "regexp"

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

func stripPunctuation(s *string) string {
	res := ""

	return res
}

func checkHyphenationConsistency(s string) []CheckResult {
	var res []CheckResult
	return res
}

// Check runs list of checkers on given string and returns an array of errors
func Check(s *string) []CheckResult {
	res := []CheckResult{}

	len := len(*s)
	if len == 0 {
		return res
	}

	m, _ := regexp.MatchString("^\\s", *s)
	if m {
		return res
	}

	stripped := stripPunctuation(s)
	res = checkHyphenationConsistency(stripped)

	return res
}

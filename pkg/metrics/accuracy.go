// Package metrics holds plain scoring functions: input in, number out.
package metrics

import (
	"regexp"
	"strings"
	"unicode"
)

// Accuracy is the share of true values.
func Accuracy(correct []bool) float64 {
	if len(correct) == 0 {
		return 0
	}
	n := 0
	for _, c := range correct {
		if c {
			n++
		}
	}
	return float64(n) / float64(len(correct))
}

// Rate is a safe division, used for hit rate, error rate, etc.
func Rate(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total)
}

var articles = regexp.MustCompile(`\b(a|an|the)\b`)

// Normalize lowercases, drops punctuation and articles, squeezes spaces (SQuAD style).
func Normalize(s string) string {
	s = strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	return strings.Join(strings.Fields(articles.ReplaceAllString(s, " ")), " ")
}

// ExactMatch compares two answers after normalization.
func ExactMatch(prediction, reference string) bool {
	return Normalize(prediction) == Normalize(reference)
}

var number = regexp.MustCompile(`-?\d[\d,]*(\.\d+)?`)

// LastNumber returns the last number in the text without thousands separators ("" if none).
func LastNumber(s string) string {
	all := number.FindAllString(s, -1)
	if len(all) == 0 {
		return ""
	}
	n := strings.ReplaceAll(all[len(all)-1], ",", "")
	if strings.Contains(n, ".") {
		n = strings.TrimRight(strings.TrimRight(n, "0"), ".")
	}
	return n
}

var choice = regexp.MustCompile(`(?i)answer\s*(?:is)?\s*[:\-]?\s*\(?\*{0,2}([A-J])\b`)
var bareChoice = regexp.MustCompile(`\b([A-J])\b`)

// Choice extracts a multiple-choice letter: "answer is X" first, else the last standalone letter.
func Choice(s string) string {
	if m := choice.FindAllStringSubmatch(s, -1); len(m) > 0 {
		return strings.ToUpper(m[len(m)-1][1])
	}
	if m := bareChoice.FindAllStringSubmatch(s, -1); len(m) > 0 {
		return m[len(m)-1][1]
	}
	return ""
}

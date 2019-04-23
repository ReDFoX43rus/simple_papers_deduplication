package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/agnivade/levenshtein"
)

func nameToInitials(name string) string {
	name = strings.ReplaceAll(name, ".", "")
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, ",", "")

	result := ""

	for _, char := range name {
		if unicode.IsUpper(char) {
			result += string(char)
		}
	}

	return result
}

func iniSim(a, b string) bool {
	len1 := len(a)
	len2 := len(b)

	if len1 < 2 || len2 < 2 {
		return false
	}

	afirst := a[0]
	bfirst := b[0]

	alast := a[len1-1]
	blast := b[len2-1]

	var asecond, bsecond byte

	if len1 < 2 {
		asecond = a[0]
	} else {
		asecond = a[1]
	}

	if len2 < 2 {
		bsecond = b[0]
	} else {
		bsecond = b[1]
	}

	if afirst == bfirst && alast == blast ||
		afirst == bsecond && alast == bfirst ||
		afirst == bfirst && asecond == bsecond ||
		afirst == blast && asecond == bfirst {
		return true
	}

	return false
}

func nameMatch(initials1, initials2 []string, threshold float64) bool {
	len1 := len(initials1)
	len2 := len(initials2)

	counter := 0

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if initials2[j] != "" && iniSim(initials1[i], initials2[j]) {
				counter++
				initials2[j] = ""
			}
		}
	}

	return float64(counter)/math.Max(float64(len1), float64(len2)) >= threshold
}

// IsPapersMatch returns true if papers match, false otherwise
func IsPapersMatch(paper1, paper2 CiteSeerPaper, thresholdYear, thresholdEditDistancePercentage int, thresholdName float64) bool {
	yearRegexp := regexp.MustCompile("\\d{4}")

	year1str := yearRegexp.FindString(paper1.Year)
	year2str := yearRegexp.FindString(paper2.Year)

	if year1str != "" && year2str != "" {
		year1, err1 := strconv.Atoi(year1str)
		year2, err2 := strconv.Atoi(year2str)

		if err1 != nil || err2 != nil || absDiffInt(year1, year2) > thresholdYear {
			return false
		}
	}

	/* authors1 := getAuthors(paper1.Author)
	authors2 := getAuthors(paper2.Author) */

	initials1 := []string{}
	initials2 := []string{}

	for _, author := range paper1.Authors {
		initials1 = append(initials1, nameToInitials(author))
	}

	for _, author := range paper2.Authors {
		initials2 = append(initials2, nameToInitials(author))
	}

	maxDiff := minInt(len(paper1.Title), len(paper2.Title))
	if maxDiff > 10 {
		maxDiff = maxDiff * thresholdEditDistancePercentage / 100
	} else if maxDiff > thresholdEditDistancePercentage/10 && thresholdEditDistancePercentage/10 != 0 {
		maxDiff = thresholdEditDistancePercentage / 10
	}

	if nameMatch(initials1, initials2, thresholdName) &&
		levenshtein.ComputeDistance(paper1.Title, paper2.Title) <= maxDiff {
		/* if paper1.Meta == "fahlman1990a" && paper2.Meta == "fahlman1990b" {
			fmt.Println(paper1, paper2)
		} */
		return true
	}

	return false
}

func getAuthors(raw string) []string {
	if strings.Contains(raw, "., ") {
		raw := strings.ReplaceAll(raw, "and ", ",. ")
		return strings.Split(raw, ",. ")
	}

	raw = strings.ReplaceAll(raw, "and ", ", ")
	return strings.Split(raw, ", ")
}

func clearAuthors(authors []string) []string {
	len := len(authors)

	for i := 0; i < len; i++ {
		authors[i] = strings.ReplaceAll(authors[i], "and ", "")
	}

	return authors
}

func absDiffInt(a, b int) int {
	tmp := a - b

	if tmp < 0 {
		return -tmp
	}

	return tmp
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}

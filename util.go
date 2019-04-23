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

func IsPapersMatchCS(paper1, paper2 CiteSeerPaper, thresholdYear, thresholdEditDistancePercentage int, thresholdName float64) bool {
	yearRegexp := regexp.MustCompile("\\d{4}")

	year1str := yearRegexp.FindString(paper1.Year)
	year2str := yearRegexp.FindString(paper2.Year)

	year1, _ := strconv.Atoi(year1str)
	year2, _ := strconv.Atoi(year2str)

	return IsPapersMatch(paper1.Title, paper2.Title, paper1.Authors, paper2.Authors, year1, year2, thresholdYear, thresholdEditDistancePercentage, thresholdName)
}

// IsPapersMatch returns true if papers match, false otherwise
func IsPapersMatch(title1, title2 string, authors1, authors2 []string, year1, year2 int, thresholdYear, thresholdEditDistancePercentage int, thresholdName float64) bool {
	if year1 != 0 && year2 != 0 && absDiffInt(year1, year2) > thresholdYear {
		return false
	}

	initials1 := []string{}
	initials2 := []string{}

	for _, author := range authors1 {
		initials1 = append(initials1, nameToInitials(author))
	}

	for _, author := range authors2 {
		initials2 = append(initials2, nameToInitials(author))
	}

	maxDiff := minInt(len(title1), len(title2))
	if maxDiff > 10 {
		maxDiff = maxDiff * thresholdEditDistancePercentage / 100
	} else if maxDiff > thresholdEditDistancePercentage/10 && thresholdEditDistancePercentage/10 != 0 {
		maxDiff = thresholdEditDistancePercentage / 10
	}

	if nameMatch(initials1, initials2, thresholdName) &&
		levenshtein.ComputeDistance(title1, title2) <= maxDiff {
		return true
	}

	return false
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
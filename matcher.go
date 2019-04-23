package main
/*
type MatchResult struct {
	Paper1 Paper
	Paper2 Paper
}

func MatchPapers(papers []Paper) []MatchResult {
	result := []MatchResult{}

	len := len(papers)

	for i := 0; i < len; i++ {
		paper1 := papers[i]
		for j := i + 1; j < len; j++ {
			paper2 := papers[j]

			if IsPapersMatch(paper1, paper2, 1, 10, 0.8) {
				result = append(result, MatchResult{paper1, paper2})
			}
		}
	}

	return result
}
*/
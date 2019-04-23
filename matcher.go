package main

type MatchResult struct {
	Paper1 DBPaper
	Paper2 DBPaper
}

func MatchDBPapers(papers []DBPaper) []MatchResult {
	result := []MatchResult{}

	len := len(papers)

	for i := 0; i < len; i++ {
		paper1 := papers[i]
		for j := i + 1; j < len; j++ {
			paper2 := papers[j]

			if IsPapersMatch(paper1.Title.String, paper2.Title.String, paper1.Authors, paper2.Authors, int(paper1.Year.Int64), int(paper2.Year.Int64), 1, 10, 0.8) {
				result = append(result, MatchResult{paper1, paper2})
			}
		}
	}

	return result
}
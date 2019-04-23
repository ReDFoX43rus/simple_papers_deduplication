package main

type CSMatchResult struct {
	Paper1 CiteSeerPaper
	Paper2 CiteSeerPaper
}

func MatchPapers(papers []CiteSeerPaper) []CSMatchResult {
	result := []CSMatchResult{}

	len := len(papers)

	for i := 0; i < len; i++ {
		paper1 := papers[i]
		for j := i + 1; j < len; j++ {
			paper2 := papers[j]

			if IsPapersMatchCS(paper1, paper2, 1, 10, 0.8) {
				result = append(result, CSMatchResult{paper1, paper2})
			}
		}
	}

	return result
}

func MatchPapersByMeta(papers []CiteSeerPaper) []CSMatchResult {
	result := []CSMatchResult{}

	len := len(papers)

	for i := 0; i < len; i++ {
		paper1 := papers[i]
		for j := i + 1; j < len; j++ {
			paper2 := papers[j]

			if paper1.Cluster == paper2.Cluster {
				result = append(result, CSMatchResult{paper1, paper2})
			}
		}
	}

	return result
}

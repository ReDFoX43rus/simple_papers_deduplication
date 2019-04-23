package main

import (
	// "strconv"
	"fmt"
)

func removeInvalidPapers(papers []Paper) []Paper {
	validPapers := []Paper{}

	for _, paper := range papers {
		if paper.Title != "" && paper.Author != "" && paper.Year != "" {
			validPapers = append(validPapers, paper)
		}
	}

	return validPapers
}

func removeDuplicates(papers []Paper) []Paper {
	distinct := []Paper{}

	was := make(map[Paper]bool)

	for _, paper := range papers {
		if _, ok := was[paper]; !ok {
			was[paper] = true
			distinct = append(distinct, paper)
		}
	}

	return distinct
}

func main() {
	cspath := "citeseer_ie/constraintOut"
	papers, err := LoadCiteSeerPapers(cspath)
	if err != nil {
		panic(err)
	}

	result := MatchPapers(papers)
	actual := MatchPapersByMeta(papers)

	var TP, FP float32

	for _, r := range result {
		p1 := r.Paper1
		p2 := r.Paper2

		// fmt.Println(strconv.Itoa(p1.Cluster) + ":" + strconv.Itoa(p1.Reference), strconv.Itoa(p2.Cluster) + ":" + strconv.Itoa(p2.Reference))

		if p1.Cluster == p2.Cluster {
			TP++
		} else {
			FP++
		}
	}

	precision := TP / (TP + FP)
	recall := TP / float32(len(actual))
	f1 := 2 * precision * recall / (precision + recall)

	fmt.Println(precision, recall, f1)

	// fmt.Println(TP, FP, TP / (TP + FP), len(actual))

	/* path := "cora-ref/fahl-labeled"

	papers, err := LoadPapers(path)
	if err != nil {
		panic(err)
	}

	papers = removeInvalidPapers(papers)
	papers = removeDuplicates(papers)

	matched := MatchPapers(papers)

	var TP, FP float32

	for _, result := range matched {
		meta1 := result.Paper1.Meta
		meta2 := result.Paper2.Meta

		if meta1 == meta2 {
			TP++
		} else {
			FP++

			fmt.Println(meta1, meta2, levenshtein.ComputeDistance(meta1, meta2))
		}
	}

	fmt.Println(TP, FP, TP/(TP+FP)) */
}

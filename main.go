package main

import (
	"fmt"
)

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
}

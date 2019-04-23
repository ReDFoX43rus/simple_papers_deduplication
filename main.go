package main

import (
	"fmt"
)

type testResult struct {
	Precision float32
	Recall float32
}

func (this *testResult) f1() float32 {
	return 2 * this.Precision * this.Recall / (this.Precision + this.Recall)
}

func testOnDataset(path string) (testResult, error) {
	var testRes testResult

	papers, err := LoadCiteSeerPapers(path)
	if err != nil {
		return testRes, err
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

	testRes.Precision = TP / (TP + FP)
	testRes.Recall = TP / float32(len(actual))

	return testRes, nil
}

func main() {
	cspaths := []string{"constraintOut", "faceOut", "reasoningOut", "reinforcementOut"}

	var totalPrecision, totalRecall, totalF1 float32

	for _, path := range cspaths {
		result, err := testOnDataset("citeseer_ie/" + path)
		if err != nil {
			panic(err)
		}

		totalPrecision += result.Precision
		totalRecall += result.Recall
		totalF1 += result.f1()

		fmt.Println(result.Precision, result.Recall, result.f1())
	}

	fmt.Print(totalPrecision / 4, totalRecall / 4, totalF1 / 4)
}

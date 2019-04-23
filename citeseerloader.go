package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type CiteSeerPaper struct {
	Authors []string
	Title   string
	Year    string

	Reference int
	Cluster   int
	TrueID    int
}

func LoadCiteSeerPapers(filename string) ([]CiteSeerPaper, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	raw := string(bytes)

	raw = strings.ReplaceAll(raw, "\r", "")
	raw = strings.ReplaceAll(raw, "\n\n", "\n")

	split := strings.Split(raw, "\n")

	referenceNoRegExp := regexp.MustCompile("reference_no=\"\\d+\"")
	clusterNoRegExp := regexp.MustCompile("cluster_no=\"\\d+\"")
	trueIDRegExp := regexp.MustCompile("true_id=\"\\d+\"")

	authorsRegExp := regexp.MustCompile("<authors>.+</authors>")
	authorRegExp := regexp.MustCompile("<author>.+</author>")

	dateRegExp := regexp.MustCompile("<date>.+</date>")
	titleRegExp := regexp.MustCompile("<title>.+</title>")

	len := len(split)
	papers := []CiteSeerPaper{}

	for i := 1; i < len; i += 2 {
		var paper CiteSeerPaper

		meta := split[i-1]
		data := split[i]

		ref, err := extractIntAttribute(referenceNoRegExp.FindString(meta), "reference_no")
		if err == nil {
			paper.Reference = ref
		}

		cluster, err := extractIntAttribute(clusterNoRegExp.FindString(meta), "cluster_no")
		if err == nil {
			paper.Cluster = cluster
		}

		trueID, err := extractIntAttribute(trueIDRegExp.FindString(meta), "true_id")
		if err == nil {
			paper.TrueID = trueID
		}

		authors := authorsRegExp.FindString(data)
		authors = removeTag(authors, "authors")

		authorsSplit := authorRegExp.FindAllString(authors, -1)
		for _, author := range authorsSplit {
			paper.Authors = append(paper.Authors, removeTag(author, "author"))
		}

		date := dateRegExp.FindString(data)
		date = removeTag(date, "date")

		paper.Year = date

		title := titleRegExp.FindString(data)
		title = removeTag(title, "title")

		paper.Title = title

		if paper.Title != "" && paper.Reference != 0 && paper.Cluster != 0 && paper.TrueID != 0 {
			papers = append(papers, paper)
		}
	}

	return papers, nil
}

func removeTag(raw, tag string) string {
	raw = strings.ReplaceAll(raw, "<"+tag+">", "")
	raw = strings.ReplaceAll(raw, "</"+tag+">", "")

	return raw
}

func extractIntAttribute(raw, attr string) (int, error) {
	raw = strings.ReplaceAll(raw, attr+"=", "")
	raw = strings.ReplaceAll(raw, "\"", "")

	res, err := strconv.Atoi(raw)

	if err != nil {
		return 0, err
	}

	return res, nil
}

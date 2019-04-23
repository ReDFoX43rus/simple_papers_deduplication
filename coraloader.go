package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// Paper is simplified information about paper
type Paper struct {
	Author string
	Title  string
	Year   string
	Meta   string
}

// LoadPapers takes raw string from file and converts it to a Paper struct
func LoadPapers(coraFilename string) ([]Paper, error) {
	bytes, err := ioutil.ReadFile(coraFilename)
	if err != nil {
		return nil, err
	}

	rawString := string(bytes)

	split := strings.Split(rawString, "<NEWREFERENCE>")

	papers := []Paper{}
	authorExp := regexp.MustCompile("<author>(\\s+\\D+)+\\s+<\\/author>")
	titleExp := regexp.MustCompile("<title>(\\s+.+)+\\s+<\\/title>")
	yearExp := regexp.MustCompile("<date>(\\s+.+)+\\s+<\\/date>")
	metaExp := regexp.MustCompile("^.+\\s<author>")

	fahlmanExp := regexp.MustCompile("^fahlman\\d{4}\\w")

	for _, str := range split {
		replaced := strings.ReplaceAll(str, "\n", "")

		paper := Paper{}

		author := authorExp.FindString(replaced)
		title := titleExp.FindString(replaced)
		year := yearExp.FindString(replaced)
		meta := metaExp.FindString(replaced)
		meta = strings.ReplaceAll(meta, " <author>", "")
		meta = strings.ReplaceAll(meta, " ", "")
		meta = strings.ReplaceAll(meta, ".", "")
		meta = strings.ReplaceAll(meta, "\r", "")
		meta = strings.ReplaceAll(meta, "fahlman1990b7", "fahlman1990b")
		meta = fahlmanExp.FindString(meta)

		paper.Author = replaceTag(author, "author")
		paper.Title = replaceTag(title, "title")
		paper.Year = replaceTag(year, "date")
		paper.Meta = meta

		papers = append(papers, paper)
	}

	return papers, nil
}

func replaceTag(src, tag string) string {
	openTag := "<" + tag + ">"
	closeTag := "</" + tag + ">"
	src = strings.ReplaceAll(src, openTag, "")
	src = strings.ReplaceAll(src, closeTag, "")

	return src
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// FindTags finds tags in cora databse, prints it and saved to file tags.txt
func FindTags(coraFilename string) {
	bytes, err := ioutil.ReadFile(coraFilename)
	if err != nil {
		panic(err)
	}

	str := string(bytes)

	rexp := regexp.MustCompile("<\\w+?>")
	strings := rexp.FindAllString(str, -1)

	unique := make(map[string]bool)
	for _, element := range strings {
		if _, value := unique[element]; !value {
			unique[element] = true
		}
	}

	file, err := os.Create("./tags.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	for k := range unique {
		fmt.Println(k)

		file.WriteString(k + "\n")
	}
}

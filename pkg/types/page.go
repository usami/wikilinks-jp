package types

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Page struct {
	PageID json.Number
	Lines  map[int]string
}

func ParsePageID(s string) json.Number {
	base := filepath.Base(s)
	fe := strings.Split(base, ".")
	return json.Number(fe[0])
}

func LoadPageFromHTML(filepath string, p json.Number, lineIDs []int) *Page {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	wholeLines := strings.Split(string(raw), "\n")

	lines := make(map[int]string)

	for _, id := range lineIDs {
		lines[id] = wholeLines[id]
	}

	return &Page{
		PageID: p,
		Lines:  lines,
	}
}

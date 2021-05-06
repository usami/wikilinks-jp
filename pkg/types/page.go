package types

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Page struct {
	PageID string
	Lines  map[int]string
}

func ParsePageID(s string) string {
	base := filepath.Base(s)
	fe := strings.Split(base, ".")
	return fe[0]
}

func LoadPageFromHTML(filepath, p string, lineIDs []int) *Page {
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

package linker

import (
	"bufio"
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/usami/wikilinks-jp/pkg/types"
)

type Linker struct {
	Category      string
	Annotations   []*types.Annotation
	Pages         map[json.Number]*types.Page
	TitleToPageID map[string]json.Number
}

func NewLinker(c string) *Linker {
	return &Linker{
		Category:      c,
		Annotations:   make([]*types.Annotation, 0),
		Pages:         make(map[json.Number]*types.Page),
		TitleToPageID: make(map[string]json.Number),
	}
}

func (l *Linker) Load(a, d, t string) {
	log.Printf("linker[%s]: load annotaions", l.Category)
	l.loadAnnotations(a)
	log.Printf("linker[%s]: load pages", l.Category)
	l.loadPages(d)
	log.Printf("linker[%s]: load title to pageid mappings", l.Category)
	l.loadTitlePageIDs(t)
}

func (l *Linker) Run() {
	log.Printf("linker[%s]: check links", l.Category)
	l.checkLinks()
}

func (l *Linker) Output(filepath string) {
	log.Printf("linker[%s]: output analyzed results", l.Category)

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, an := range l.Annotations {
		if an.LinkPageID != "" {
			jsonStr, err := json.Marshal(an)
			if err != nil {
				log.Fatal(err)
			}
			s := string(jsonStr) + "\n"

			if _, err := f.WriteString(s); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (l *Linker) loadAnnotations(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var aa []*types.Annotation

	for s.Scan() {
		var an types.Annotation
		if err := json.Unmarshal([]byte(s.Text()), &an); err != nil {
			log.Fatal(err)
		}
		aa = append(aa, &an)
	}

	l.Annotations = aa
}

func (l *Linker) loadPages(dirpath string) {
	files := listHTMLFiles(dirpath)

	amap := make(map[json.Number][]int)

	for _, an := range l.Annotations {
		lines := amap[an.PageID]
		if lines == nil {
			lines = make([]int, 0)
		}
		lines = append(lines, an.HTMLOffset.Start.LineID)
		amap[an.PageID] = lines
	}

	for _, file := range files {
		pid := types.ParsePageID(file)
		page := types.LoadPageFromHTML(file, pid, amap[pid])
		l.Pages[page.PageID] = page
	}
}

func (l *Linker) loadTitlePageIDs(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		var t types.TitlePageIDWithRedirect
		if err := json.Unmarshal([]byte(s.Text()), &t); err != nil {
			log.Fatal(err)
		}
		l.TitleToPageID[t.Title] = t.Resolve()
	}
}

func (l *Linker) checkLinks() {
	for _, an := range l.Annotations {
		p := l.Pages[an.PageID]
		li := an.HTMLOffset.Start.LineID

		doc, err := html.Parse(strings.NewReader(p.Lines[li]))
		if err != nil {
			log.Fatal(err)
		}

		var f func(*html.Node, int) int

		f = func(n *html.Node, offset int) int {
			switch n.Type {
			case html.ElementNode:
				if n.DataAtom == atom.A {
					for _, attr := range n.Attr {
						if isLinkToEntity(attr) {
							if matchesAnnotation(n.FirstChild, an, offset) {
								title := extractTitle(attr.Val)
								if pageID, ok := l.TitleToPageID[title]; ok {
									an.LinkPageID = pageID
								}
								break
							}
						}
					}
				}
			case html.TextNode:
				offset += utf8.RuneCountInString(n.Data)
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				offset = f(c, offset)
			}

			return offset
		}
		f(doc, 0)
	}
}

func listHTMLFiles(dirpath string) []string {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	var files []string

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func isLinkToEntity(a html.Attribute) bool {
	return a.Key == "href" && (strings.HasPrefix(a.Val, "/index.php/") || strings.HasPrefix(a.Val, "/a-sumida/wiki2019_1/index.php/"))
}

func extractTitle(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Path, "/")
	return parts[len(parts)-1]
}

func matchesAnnotation(n *html.Node, a *types.Annotation, offset int) bool {
	if n == nil || n.Type != html.TextNode {
		return false
	}

	return a.TextOffset.Start.Offset == offset && a.TextOffset.Text == n.Data
}

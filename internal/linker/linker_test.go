package linker

import (
	"encoding/json"
	"sort"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/usami/wikilinks-jp/pkg/types"
)

func TestNewLinker(t *testing.T) {
	l := NewLinker("airport")

	assert.Equal(t, l.Category, "airport")
}

func TestLinker_loadAnnotations(t *testing.T) {
	l := NewLinker("airport")

	assert.Equal(t, len(l.Annotations), 0)
	l.loadAnnotations("../../testdata/annotations.json")
	assert.Equal(t, len(l.Annotations), 10)

	assert.Equal(t, l.Annotations[2].HTMLOffset.Text, "Pekoa Airfield")
	assert.Equal(t, l.Annotations[3].HTMLOffset.Text, "ペコア飛行場")
}

func TestLinker_loadPages(t *testing.T) {
	l := NewLinker("airport")
	l.loadAnnotations("../../testdata/annotations.json")

	assert.Equal(t, len(l.Pages), 0)
	l.loadPages("../../testdata")
	assert.Equal(t, len(l.Pages), 2)

	p := l.Pages[json.Number("1017261")]
	assert.Equal(t, p.PageID, json.Number("1017261"))
	assert.Equal(t, len(p.Lines), 6)
}

func TestLinker_loadTitlePageIDs(t *testing.T) {
	l := NewLinker("airport")
	l.loadTitlePageIDs("../../testdata/title2pageid.json")

	assert.Equal(t, len(l.TitleToPageID), 10)
	assert.Equal(t, l.TitleToPageID["!"], json.Number("124376"))
}

func TestLinker_checkLinks(t *testing.T) {
	l := NewLinker("airport")

	t.Run("one annotation", func(t *testing.T) {
		an := types.Annotation{
			PageID: json.Number("1017261"),
			TextOffset: types.OffsetPair{
				Start: types.Offset{
					LineID: 63,
					Offset: 63,
				},
				End: types.Offset{
					LineID: 63,
					Offset: 67,
				},
				Text: "バヌアツ",
			},
			HTMLOffset: types.OffsetPair{
				Start: types.Offset{
					LineID: 63,
				},
				Text: "バヌアツ",
			},
		}

		l.Annotations = append(l.Annotations, &an)
		l.loadPages("../../testdata/1017261.html")
		l.TitleToPageID["バヌアツ"] = json.Number("10")

		l.checkLinks()

		assert.Equal(t, string(l.Annotations[0].LinkPageID), "10")
	})

	t.Run("more annotations", func(t *testing.T) {
		l.loadAnnotations("../../testdata/annotations.json")
		l.loadPages("../../testdata")
		l.TitleToPageID["エスピリトゥサント島"] = json.Number("100")
		l.TitleToPageID["ルーガンビル"] = json.Number("1000")

		l.checkLinks()

		expected := []string{
			"",
			"",
			"",
			"",
			"",
			"",
			"10",
			"10",
			"100",
			"1000",
		}
		for i, an := range l.Annotations {
			assert.Check(t, string(an.LinkPageID) == expected[i], an.LinkPageID)
		}
	})
}

func TestListHTMLFiles(t *testing.T) {
	files := sort.StringSlice(listHTMLFiles("../../testdata"))
	files.Sort()

	assert.Equal(t, len(files), 2)
	assert.DeepEqual(t, files, sort.StringSlice{"../../testdata/1017261.html", "../../testdata/4189.html"})
}

func TestExtractTitle(t *testing.T) {
	t.Run("/index.php", func(t *testing.T) {
		s := "/index.php/%E3%83%90%E3%83%8C%E3%82%A2%E3%83%84"

		title := extractTitle(s)
		assert.Equal(t, title, "バヌアツ")
	})

	t.Run("/a-sumida/wiki2019_1/index.php", func(t *testing.T) {
		s := "/a-sumida/wiki2019_1/index.php/1965%E5%B9%B4"

		title := extractTitle(s)
		assert.Equal(t, title, "1965年")
	})
}

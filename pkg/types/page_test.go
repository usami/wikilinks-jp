package types

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPage_LoadPageFromHTML(t *testing.T) {
	p := LoadPageFromHTML("../../testdata/1017261.html", json.Number("1017261"), []int{1, 10, 14, 18})

	assert.Equal(t, p.PageID, json.Number("1017261"))
	assert.Equal(t, len(p.Lines), 4)
}

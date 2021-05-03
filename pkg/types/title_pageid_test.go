package types

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestTitlePageID(t *testing.T) {
	f, err := os.Open("../../testdata/title2pageid.json")
	assert.NilError(t, err)
	defer f.Close()

	s := bufio.NewScanner(f)

	var tps []*TitlePageIDWithRedirect

	for s.Scan() {
		var tp TitlePageIDWithRedirect
		err := json.Unmarshal([]byte(s.Text()), &tp)

		assert.NilError(t, err)
		tps = append(tps, &tp)
	}

	assert.Equal(t, len(tps), 10)

	tp := tps[0]
	assert.Equal(t, tp.PageID, json.Number("305230"))
	assert.Equal(t, tp.Title, "!")
	assert.Equal(t, tp.IsRedirect, true)
	assert.Equal(t, tp.RedirectTo.PageID, json.Number("124376"))
	assert.Equal(t, tp.Resolve(), json.Number("124376"))

	tp = tps[2]
	assert.Equal(t, tp.PageID, json.Number("617718"))
	assert.Equal(t, tp.Title, "!!!")
	assert.Equal(t, tp.IsRedirect, false)
	assert.Equal(t, tp.Resolve(), json.Number("617718"))
}

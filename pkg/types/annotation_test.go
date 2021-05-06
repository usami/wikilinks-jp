package types

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"gotest.tools/v3/assert"
)

func TestAnnotation(t *testing.T) {
	var a Annotation

	f, err := ioutil.ReadFile("../../testdata/annotation_sample.json")
	assert.NilError(t, err)

	assert.NilError(t, json.Unmarshal([]byte(f), &a))

	assert.Equal(t, a.ENETag, ENETag("1.6.5.3"))
	assert.Equal(t, a.AttributeName, AttributeName("別名"))
	assert.Equal(t, a.PageID, "1017261")
	assert.Equal(t, a.Title, "サントペコア国際空港")
	assert.Equal(t, a.LinkPageID, "")

	ho := a.HTMLOffset
	assert.Equal(t, ho.Text, "Santo-Pekoa International Airport")
	assert.Equal(t, ho.Start.LineID, 63)
	assert.Equal(t, ho.Start.Offset, 39)
	assert.Equal(t, ho.End.LineID, 63)
	assert.Equal(t, ho.End.Offset, 72)

	to := a.TextOffset
	assert.Equal(t, to.Text, "Santo-Pekoa International Airport")
	assert.Equal(t, to.Start.LineID, 63)
	assert.Equal(t, to.Start.Offset, 26)
	assert.Equal(t, to.End.LineID, 63)
	assert.Equal(t, to.End.Offset, 59)
}

package types

import "encoding/json"

type ENETag string

type AttributeName string

type Annotation struct {
	ENETag        ENETag        `json:"ENE"`
	AttributeName AttributeName `json:"attribute"`
	HTMLOffset    OffsetPair    `json:"html_offset"`
	TextOffset    OffsetPair    `json:"text_offset"`
	PageID        json.Number   `json:"page_id"`
	Title         string        `json:"title"`
	// HasLink       bool          `json:"has_link"`
	HasLink   bool   `json:"-"`
	LinkTitle string `json:"link_title"`
}

type OffsetPair struct {
	Start Offset `json:"start"`
	End   Offset `json:"end"`
	Text  string `json:"text"`
}

type Offset struct {
	LineID int `json:"line_id"`
	Offset int `json:"offset"`
}

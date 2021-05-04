package types

import "encoding/json"

type TitlePageID struct {
	PageID     json.Number `json:"page_id"`
	Title      string      `json:"title"`
	IsRedirect bool        `json:"is_redirect"`
}

type TitlePageIDWithRedirect struct {
	TitlePageID
	RedirectTo TitlePageID `json:"redirect_to"`
}

func (t *TitlePageIDWithRedirect) Resolve() json.Number {
	if t.IsRedirect {
		return t.RedirectTo.PageID
	}
	return t.PageID
}

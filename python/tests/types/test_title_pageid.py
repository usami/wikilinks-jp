import json
from dacite import from_dict
from typing import List
from src.types.title_pageid import TitlePageIDWithRedirect


def test_title_page_id():
    with open("../testdata/title2pageid.json", mode="r") as f:
        json_lines = f.readlines()

    tps: List[TitlePageIDWithRedirect] = []
    for json_line in json_lines:
        d = json.loads(json_line)
        tp = from_dict(TitlePageIDWithRedirect, d)
        tps.append(tp)

    assert len(tps) == 10

    tp = tps[0]
    assert tp.page_id == 305230
    assert tp.title == "!"
    assert tp.is_redirect
    assert tp.redirect_to.page_id == 124376
    assert tp.resolve() == 124376

    tp = tps[2]
    assert tp.page_id == 617718
    assert tp.title == "!!!"
    assert not tp.is_redirect
    assert tp.resolve() == 617718

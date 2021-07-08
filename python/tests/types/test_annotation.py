import json
from dacite import from_dict
from src.types.annotation import Annotation


def test_annotation():
    with open("../testdata/annotation_sample.json", mode="r") as f:
        d = json.load(f)

    a = from_dict(Annotation, d)

    assert a.ENE == "1.6.5.3"
    assert a.attribute == "別名"
    assert a.page_id == "1017261"
    assert a.title == "サントペコア国際空港"
    assert a.link_page_id == ""

    ho = a.html_offset
    assert ho.text == "Santo-Pekoa International Airport"
    assert ho.start.line_id == 63
    assert ho.start.offset == 39
    assert ho.end.line_id == 63
    assert ho.end.offset == 72

    to = a.text_offset
    assert to.text == "Santo-Pekoa International Airport"
    assert to.start.line_id == 63
    assert to.start.offset == 26
    assert to.end.line_id == 63
    assert to.end.offset == 59

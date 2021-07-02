from src.types.annotation import Annotation, Offset, OffsetPair
from src.linker.linker import Linker, list_html_files, extract_title


def test_new_linker():
    l = Linker("airport")

    assert l.category == "airport"


def test_load_annotations():
    l = Linker("airport")

    assert len(l.annotations) == 0
    l.load_annotations("../testdata/annotations.json")
    assert len(l.annotations) == 10

    assert l.annotations[2].html_offset.text == "Pekoa Airfield"
    assert l.annotations[3].html_offset.text == "ペコア飛行場"


def test_load_pages():
    l = Linker("airport")
    l.load_annotations("../testdata/annotations.json")

    assert len(l.pages) == 0
    l.load_pages("../testdata")
    assert len(l.pages) == 1

    p = l.pages["1017261"]
    assert p.page_id == "1017261"
    assert len(p.lines) == 6


def test_load_title_page_ids():
    l = Linker("airport")
    l.load_title_page_ids("../testdata/title2pageid.json")

    assert len(l.title_to_page_id) == 10
    assert l.title_to_page_id["!"] == 124376


def test_check_links():
    l = Linker("airport")

    an = Annotation(
        ENE="1.6.5.3",
        attribute="国",
        text_offset=OffsetPair(
            start=Offset(line_id=63, offset=63),
            end=Offset(line_id=63, offset=67),
            text="バヌアツ",
        ),
        html_offset=OffsetPair(
            start=Offset(line_id=63, offset=151),
            end=Offset(line_id=63, offset=155),
            text="バヌアツ",
        ),
        page_id="1017261",
        title="サントペコア国際空港",
    )

    l.annotations.append(an)
    l.load_pages("../testdata/")
    l.title_to_page_id["バヌアツ"] = 10

    l.check_links()

    assert l.annotations[0].link_page_id == "10"

    l.load_annotations("../testdata/annotations.json")
    l.load_pages("../testdata")
    l.title_to_page_id["エスピリトゥサント島"] = 100
    l.title_to_page_id["ルーガンビル"] = 1000

    l.check_links()

    expected = [
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
    ]

    for i, an in enumerate(l.annotations):
        assert an.link_page_id == expected[i]


def test_list_html_files():
    files = sorted(list_html_files("../testdata"))
    assert len(files) == 2
    assert files == sorted(
        ["../testdata/1017261.html", "../testdata/4189.html"])


def test_extract_title():
    s = "/index.php/%E3%83%90%E3%83%8C%E3%82%A2%E3%83%84"

    title = extract_title(s)
    assert title == "バヌアツ"

    s = "/a-sumida/wiki2019_1/index.php/1965%E5%B9%B4"

    title = extract_title(s)
    assert title == "1965年"

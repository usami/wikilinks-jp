from src.types.page import load_page_from_html


def test_load_page_from_html():
    p = load_page_from_html("../testdata/1017261.html",
                            "1017261", [1, 10, 14, 18])

    assert p.page_id == "1017261"
    assert len(p.lines) == 4

from dataclasses import dataclass
import os
from typing import Dict, List


@dataclass
class Page:
    page_id: str
    lines: Dict[int, str]


def parse_page_id(s: str) -> str:
    base = os.path.basename(s)
    fe = base.split(".")
    return fe[0]


def load_page_from_html(filepath: str, p: str, line_ids: List[int]) -> Page:
    with open(filepath, mode="r") as f:
        raw = f.read()

    whole_lines = raw.split("\n")

    lines = {}
    for line_id in line_ids:
        lines[line_id] = whole_lines[line_id]

    return Page(page_id=p, lines=lines)

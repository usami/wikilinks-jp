from dataclasses import dataclass
from dataclasses_json import dataclass_json


@dataclass
class Offset:
    line_id: int
    offset: int


@dataclass
class OffsetPair:
    start: Offset
    end: Offset
    text: str


@dataclass_json
@dataclass
class Annotation:
    ENE: str
    attribute: str
    html_offset: OffsetPair
    text_offset: OffsetPair
    page_id: str
    title: str
    link_page_id: str = ""

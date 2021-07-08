import json
import ujson
import logging
import os
from typing import List, Dict
from urllib.parse import urlparse, unquote

from bs4 import BeautifulSoup
from bs4.element import Tag, NavigableString
from dacite import from_dict

from src.types.annotation import Annotation
from src.types.page import Page, parse_page_id, load_page_from_html
from src.types.title_pageid import TitlePageIDWithRedirect

logger = logging.getLogger(__name__)
logging.basicConfig(
    format="%(asctime)s - %(levelname)s - %(name)s - %(message)s",
    datefmt="%m/%d/%Y %H:%M:%S",
    level=logging.INFO,
)


class Linker:
    def __init__(
            self, category: str, annotations: List[Annotation] = None,
            pages: Dict[str, Page] = None, title_to_page_id: Dict[str, int] = None) -> None:
        self.category: str = category
        self.annotations: List[Annotation] = [
        ] if annotations is None else annotations
        self.pages: Dict[str, Page] = {} if pages is None else pages
        self.title_to_page_id: Dict[str, int] = {
        } if title_to_page_id is None else title_to_page_id

    def load(self, a: str, d: str, t: str) -> None:
        logger.info("linker[%s]: load annotaions", self.category)
        self.load_annotations(a)
        logger.info("linker[%s]: load pages", self.category)
        self.load_pages(d)
        logger.info("linker[%s]: load title to pageid mappings", self.category)
        self.load_title_page_ids(t)

    def run(self) -> None:
        logger.info("linker[%s]: check links", self.category)
        self.check_links()

    def output(self, filepath: str) -> None:
        logger.info("linker[%s]: output analyzed results", self.category)

        with open(filepath, mode="w") as f:
            for an in self.annotations:
                if an.link_page_id != "":
                    json_str = an.to_json(ensure_ascii=False)
                    s = json_str + "\n"
                    f.write(s)

    def load_annotations(self, filepath: str) -> None:
        aa: List[Annotation] = []
        with open(filepath, mode="r") as f:
            for line in f:
                d = json.loads(line)
                an = from_dict(Annotation, d)
                aa.append(an)

        self.annotations = aa

    def load_pages(self, dirpath: str) -> None:
        files = list_html_files(dirpath)

        amap: Dict[str, List[int]] = {}

        for an in self.annotations:
            lines = amap.get(an.page_id, [])
            lines.append(an.html_offset.start.line_id)
            amap[an.page_id] = lines

        for f in files:
            pid = parse_page_id(f)
            if pid in amap:
                page = load_page_from_html(f, pid, amap[pid])
                self.pages[page.page_id] = page

    def load_title_page_ids(self, filepath: str) -> None:
        base_name = os.path.basename(filepath)
        dirname = os.path.dirname(filepath)
        dic_path = os.path.join(dirname, "dic-" + base_name)
        if os.path.exists(dic_path):
            logger.info(
                "exists dictionary file of title to page id. loading dictionary file...")
            with open(dic_path, mode="r") as f:
                self.title_to_page_id = json.load(f)
            return

        logger.info("not exists dictionary file of title to page id.")
        with open(filepath, mode="r") as f:
            for i, line in enumerate(f):
                if i % 10000 == 0:
                    logger.info("step of loading title page ids: %s", i)
                d = ujson.loads(line)
                if "redirect_to" in d:
                    if d["redirect_to"]["page_id"] is None or d["redirect_to"]["title"] is None:
                        del d["redirect_to"]
                        d["is_redirect"] = False

                t = from_dict(TitlePageIDWithRedirect, d)
                self.title_to_page_id[t.title] = t.resolve()

        with open(dic_path, mode="w") as f:
            json.dump(self.title_to_page_id, f, indent=2, ensure_ascii=False)

    def check_links(self):
        for an in self.annotations:
            p = self.pages[an.page_id]
            li = an.html_offset.start.line_id
            soup = BeautifulSoup(p.lines[li], 'html.parser')

            def f(n, offset: int) -> int:
                if type(n) == Tag:
                    if n.name == "a":
                        href = n.attrs["href"]
                        if is_link_to_entity(href):
                            if hasattr(n, "children"):
                                first_child = list(n.children)[0]
                                if matches_annotation(first_child, an, offset):
                                    title = extract_title(href)
                                    if title in self.title_to_page_id:
                                        page_id = self.title_to_page_id[title]
                                        an.link_page_id = str(page_id)
                elif type(n) == NavigableString:
                    offset += len(n.string)

                if hasattr(n, "children"):
                    for c in n.children:
                        offset = f(c, offset)

                return offset

            f(soup, 0)


def list_html_files(dirpath: str) -> List[str]:
    file_list: List[str] = []
    for cur_dir, dirs, files in os.walk(dirpath):
        for f in files:
            _, ext = os.path.splitext(f)
            if ext == '.html':
                path = os.path.join(cur_dir, f)
                file_list.append(path)
    return file_list


def is_link_to_entity(a: str) -> bool:
    return a.startswith("/index.php/") or a.startswith("/a-sumida/wiki2019_1/index.php/")


def extract_title(s: str) -> str:
    u = urlparse(s)
    path = u.path
    parts = path.split("/")
    title = unquote(parts[len(parts)-1])
    return title


def matches_annotation(n, a: Annotation, offset: int) -> bool:
    if not type(n) == NavigableString:
        return False
    return a.text_offset.start.offset == offset and a.text_offset.text == n.string

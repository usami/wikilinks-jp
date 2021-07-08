from abc import ABCMeta
from dataclasses import dataclass
from typing import Optional


@dataclass
class TitlePageID(metaclass=ABCMeta):
    page_id: int
    title: str
    is_redirect: bool

    def __new__(cls, *args, **kwargs):
        dataclass(cls)
        return super().__new__(cls)


# dataclassの継承について
# https://zenn.dev/enven/articles/8b80ff38461b4ff329aa#%E3%83%87%E3%83%BC%E3%82%BF%E3%82%AF%E3%83%A9%E3%82%B9%E3%82%92%E7%B6%99%E6%89%BF%E3%81%99%E3%82%8B
@dataclass
class TitlePageIDWithRedirect(TitlePageID):
    redirect_to: Optional[TitlePageID]

    def resolve(self) -> int:
        if self.is_redirect:
            return self.redirect_to.page_id
        return self.page_id

# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class QueryType(str, Enum):
    """
    Determines how query words are [interpreted as prefixes](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/in-depth/prefix-searching/).
    """

    """
    allowed enum values
    """
    PREFIXLAST = "prefixLast"
    PREFIXALL = "prefixAll"
    PREFIXNONE = "prefixNone"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of QueryType from a JSON string"""
        return cls(loads(json_str))

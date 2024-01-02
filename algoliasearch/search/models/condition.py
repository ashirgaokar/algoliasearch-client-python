# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Any, Dict, Optional, Self

from pydantic import BaseModel, Field, StrictBool, StrictStr

from algoliasearch.search.models.anchoring import Anchoring


class Condition(BaseModel):
    """
    Condition
    """

    pattern: Optional[StrictStr] = Field(
        default=None, description="Query pattern syntax."
    )
    anchoring: Optional[Anchoring] = None
    alternatives: Optional[StrictBool] = Field(
        default=False,
        description="Whether the pattern matches on plurals, synonyms, and typos.",
    )
    context: Optional[StrictStr] = Field(
        default=None, description="Rule context format: [A-Za-z0-9_-]+)."
    )

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of Condition from a JSON string"""
        return cls.from_dict(loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={},
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of Condition from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "pattern": obj.get("pattern"),
                "anchoring": obj.get("anchoring"),
                "alternatives": obj.get("alternatives")
                if obj.get("alternatives") is not None
                else False,
                "context": obj.get("context"),
            }
        )
        return _obj

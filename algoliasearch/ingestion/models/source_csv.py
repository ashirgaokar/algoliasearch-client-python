# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Annotated, Any, Dict, Optional, Self

from pydantic import BaseModel, Field, StrictStr

from algoliasearch.ingestion.models.mapping_type_csv import MappingTypeCSV
from algoliasearch.ingestion.models.method_type import MethodType


class SourceCSV(BaseModel):
    """
    SourceCSV
    """

    url: StrictStr = Field(description="The URL of the file.")
    unique_id_column: Optional[StrictStr] = Field(
        default=None,
        description="The name of the column that contains the unique ID, used as `objectID` in Algolia.",
        alias="uniqueIDColumn",
    )
    mapping: Optional[Dict[str, MappingTypeCSV]] = Field(
        default=None,
        description='Mapping of type for every column. For example {"myColumn": "boolean", "myOtherColumn": "json"}. ',
    )
    method: Optional[MethodType] = None
    delimiter: Optional[
        Annotated[str, Field(min_length=1, strict=True, max_length=1)]
    ] = Field(
        default=",",
        description="The character used to split the value on each line, default to a comma (\\r, \\n, 0xFFFD, and space are forbidden).",
    )

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of SourceCSV from a JSON string"""
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
        """Create an instance of SourceCSV from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "url": obj.get("url"),
                "uniqueIDColumn": obj.get("uniqueIDColumn"),
                "mapping": dict((_k, _v) for _k, _v in obj.get("mapping").items()),
                "method": obj.get("method"),
                "delimiter": obj.get("delimiter")
                if obj.get("delimiter") is not None
                else ",",
            }
        )
        return _obj

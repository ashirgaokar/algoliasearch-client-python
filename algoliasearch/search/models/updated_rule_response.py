# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Any, ClassVar, Dict, List

from pydantic import BaseModel, Field, StrictInt, StrictStr

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class UpdatedRuleResponse(BaseModel):
    """
    UpdatedRuleResponse
    """

    object_id: StrictStr = Field(
        description="Unique object identifier.", alias="objectID"
    )
    updated_at: StrictStr = Field(
        description="Timestamp of the last update in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.",
        alias="updatedAt",
    )
    task_id: StrictInt = Field(
        description="Unique identifier of a task. A successful API response means that a task was added to a queue. It might not run immediately. You can check the task's progress with the `task` operation and this `taskID`. ",
        alias="taskID",
    )
    __properties: ClassVar[List[str]] = ["objectID", "updatedAt", "taskID"]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of UpdatedRuleResponse from a JSON string"""
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
        """Create an instance of UpdatedRuleResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "objectID": obj.get("objectID"),
                "updatedAt": obj.get("updatedAt"),
                "taskID": obj.get("taskID"),
            }
        )
        return _obj

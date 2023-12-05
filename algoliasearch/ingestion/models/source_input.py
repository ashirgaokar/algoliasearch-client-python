# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import dumps
from typing import Dict, List, Optional, Union

from pydantic import BaseModel, ValidationError, field_validator
from typing_extensions import Literal

from algoliasearch.ingestion.models.source_big_commerce import SourceBigCommerce
from algoliasearch.ingestion.models.source_big_query import SourceBigQuery
from algoliasearch.ingestion.models.source_commercetools import SourceCommercetools
from algoliasearch.ingestion.models.source_csv import SourceCSV
from algoliasearch.ingestion.models.source_docker import SourceDocker
from algoliasearch.ingestion.models.source_json import SourceJSON

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

SOURCEINPUT_ONE_OF_SCHEMAS = [
    "SourceBigCommerce",
    "SourceBigQuery",
    "SourceCSV",
    "SourceCommercetools",
    "SourceDocker",
    "SourceJSON",
]


class SourceInput(BaseModel):
    """
    SourceInput
    """

    # data type: SourceCommercetools
    oneof_schema_1_validator: Optional[SourceCommercetools] = None
    # data type: SourceBigCommerce
    oneof_schema_2_validator: Optional[SourceBigCommerce] = None
    # data type: SourceJSON
    oneof_schema_3_validator: Optional[SourceJSON] = None
    # data type: SourceCSV
    oneof_schema_4_validator: Optional[SourceCSV] = None
    # data type: SourceBigQuery
    oneof_schema_5_validator: Optional[SourceBigQuery] = None
    # data type: SourceDocker
    oneof_schema_6_validator: Optional[SourceDocker] = None
    actual_instance: Optional[
        Union[
            SourceBigCommerce,
            SourceBigQuery,
            SourceCSV,
            SourceCommercetools,
            SourceDocker,
            SourceJSON,
        ]
    ] = None
    one_of_schemas: List[str] = Literal[
        "SourceBigCommerce",
        "SourceBigQuery",
        "SourceCSV",
        "SourceCommercetools",
        "SourceDocker",
        "SourceJSON",
    ]

    model_config = {"validate_assignment": True}

    def __init__(self, *args, **kwargs) -> None:
        if args:
            if len(args) > 1:
                raise ValueError(
                    "If a position argument is used, only 1 is allowed to set `actual_instance`"
                )
            if kwargs:
                raise ValueError(
                    "If a position argument is used, keyword arguments cannot be used."
                )
            super().__init__(actual_instance=args[0])
        else:
            super().__init__(**kwargs)

    @field_validator("actual_instance")
    def actual_instance_must_validate_oneof(cls, v):
        SourceInput.model_construct()
        error_messages = []
        match = 0
        # validate data type: SourceCommercetools
        if not isinstance(v, SourceCommercetools):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceCommercetools'""")
        else:
            match += 1
        # validate data type: SourceBigCommerce
        if not isinstance(v, SourceBigCommerce):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceBigCommerce'""")
        else:
            match += 1
        # validate data type: SourceJSON
        if not isinstance(v, SourceJSON):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceJSON'""")
        else:
            match += 1
        # validate data type: SourceCSV
        if not isinstance(v, SourceCSV):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceCSV'""")
        else:
            match += 1
        # validate data type: SourceBigQuery
        if not isinstance(v, SourceBigQuery):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceBigQuery'""")
        else:
            match += 1
        # validate data type: SourceDocker
        if not isinstance(v, SourceDocker):
            error_messages.append(f"""Type '{type(v)}' is not 'SourceDocker'""")
        else:
            match += 1
        if match > 1:
            # more than 1 match
            raise ValueError(
                "Multiple matches found when setting `actual_instance` in SourceInput with oneOf schemas: SourceBigCommerce, SourceBigQuery, SourceCSV, SourceCommercetools, SourceDocker, SourceJSON. Details: "
                + ", ".join(error_messages)
            )
        elif match == 0:
            # no match
            raise ValueError(
                "No match found when setting `actual_instance` in SourceInput with oneOf schemas: SourceBigCommerce, SourceBigQuery, SourceCSV, SourceCommercetools, SourceDocker, SourceJSON. Details: "
                + ", ".join(error_messages)
            )
        else:
            return v

    @classmethod
    def from_dict(cls, obj: dict) -> Self:
        return cls.from_json(dumps(obj))

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Returns the object represented by the json string"""
        instance = cls.model_construct()
        error_messages = []
        match = 0

        # deserialize data into SourceCommercetools
        try:
            instance.actual_instance = SourceCommercetools.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SourceBigCommerce
        try:
            instance.actual_instance = SourceBigCommerce.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SourceJSON
        try:
            instance.actual_instance = SourceJSON.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SourceCSV
        try:
            instance.actual_instance = SourceCSV.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SourceBigQuery
        try:
            instance.actual_instance = SourceBigQuery.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))
        # deserialize data into SourceDocker
        try:
            instance.actual_instance = SourceDocker.from_json(json_str)
            match += 1
        except (ValidationError, ValueError) as e:
            error_messages.append(str(e))

        if match > 1:
            # more than 1 match
            raise ValueError(
                "Multiple matches found when deserializing the JSON string into SourceInput with oneOf schemas: SourceBigCommerce, SourceBigQuery, SourceCSV, SourceCommercetools, SourceDocker, SourceJSON. Details: "
                + ", ".join(error_messages)
            )
        elif match == 0:
            # no match
            raise ValueError(
                "No match found when deserializing the JSON string into SourceInput with oneOf schemas: SourceBigCommerce, SourceBigQuery, SourceCSV, SourceCommercetools, SourceDocker, SourceJSON. Details: "
                + ", ".join(error_messages)
            )
        else:
            return instance

    def to_json(self) -> str:
        """Returns the JSON representation of the actual instance"""
        if self.actual_instance is None:
            return "null"

        to_json = getattr(self.actual_instance, "to_json", None)
        if callable(to_json):
            return self.actual_instance.to_json()
        else:
            return dumps(self.actual_instance)

    def to_dict(self) -> Dict:
        """Returns the dict representation of the actual instance"""
        if self.actual_instance is None:
            return None

        to_dict = getattr(self.actual_instance, "to_dict", None)
        if callable(to_dict):
            return self.actual_instance.to_dict()
        else:
            # primitive type
            return self.actual_instance

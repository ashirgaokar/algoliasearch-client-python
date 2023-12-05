# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Annotated, Any, ClassVar, Dict, List, Optional

from pydantic import BaseModel, Field, StrictInt, StrictStr

from algoliasearch.recommend.models.search_params_object import SearchParamsObject
from algoliasearch.recommend.models.trending_items_model import TrendingItemsModel

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class TrendingItemsQuery(BaseModel):
    """
    TrendingItemsQuery
    """

    facet_name: Optional[StrictStr] = Field(
        default=None, description="Facet name for trending models.", alias="facetName"
    )
    facet_value: Optional[StrictStr] = Field(
        default=None, description="Facet value for trending models.", alias="facetValue"
    )
    model: Optional[TrendingItemsModel] = None
    query_parameters: Optional[SearchParamsObject] = Field(
        default=None, alias="queryParameters"
    )
    fallback_parameters: Optional[SearchParamsObject] = Field(
        default=None, alias="fallbackParameters"
    )
    index_name: StrictStr = Field(description="Algolia index name.", alias="indexName")
    threshold: Optional[Annotated[int, Field(le=100, strict=True, ge=0)]] = Field(
        default=None,
        description="Recommendations with a confidence score lower than `threshold` won't appear in results. > **Note**: Each recommendation has a confidence score of 0 to 100. The closer the score is to 100, the more relevant the recommendations are. ",
    )
    max_recommendations: Optional[StrictInt] = Field(
        default=0,
        description="Maximum number of recommendations to retrieve. If 0, all recommendations will be returned.",
        alias="maxRecommendations",
    )
    __properties: ClassVar[List[str]] = [
        "facetName",
        "facetValue",
        "model",
        "queryParameters",
        "fallbackParameters",
        "indexName",
        "threshold",
        "maxRecommendations",
    ]

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of TrendingItemsQuery from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of
        # query_parameters
        if self.query_parameters:
            _dict["queryParameters"] = self.query_parameters.to_dict()
        # override the default output from pydantic by calling `to_dict()` of
        # fallback_parameters
        if self.fallback_parameters:
            _dict["fallbackParameters"] = self.fallback_parameters.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of TrendingItemsQuery from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "facetName": obj.get("facetName"),
                "facetValue": obj.get("facetValue"),
                "model": obj.get("model"),
                "queryParameters": SearchParamsObject.from_dict(
                    obj.get("queryParameters")
                )
                if obj.get("queryParameters") is not None
                else None,
                "fallbackParameters": SearchParamsObject.from_dict(
                    obj.get("fallbackParameters")
                )
                if obj.get("fallbackParameters") is not None
                else None,
                "indexName": obj.get("indexName"),
                "threshold": obj.get("threshold"),
                "maxRecommendations": obj.get("maxRecommendations")
                if obj.get("maxRecommendations") is not None
                else 0,
            }
        )
        return _obj

# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from re import match
from typing import Annotated, Any, Dict, List, Optional, Tuple, Union
from urllib.parse import quote

from pydantic import Field, StrictInt, StrictStr

from algoliasearch.abtesting.config import Config
from algoliasearch.abtesting.models import (
    ABTest,
    ABTestResponse,
    AddABTestsRequest,
    ListABTestsResponse,
)
from algoliasearch.http import ApiResponse, RequestOptions, Transporter, Verb

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class AbtestingClient:
    PRIMITIVE_TYPES = (float, bool, bytes, str, int)
    NATIVE_TYPES_MAPPING = {
        "int": int,
        "float": float,
        "str": str,
        "bool": bool,
        "object": object,
    }

    def app_id(self) -> str:
        return self._config.app_id

    def __init__(self, transporter: Transporter, config: Config) -> None:
        self._transporter = transporter
        self._config = config

    def create_with_config(config: Config) -> Self:
        transporter = Transporter(config)

        return AbtestingClient(transporter, config)

    def create(app_id: Optional[str] = None, api_key: Optional[str] = None) -> Self:
        return AbtestingClient.create_with_config(Config(app_id, api_key))

    async def close(self) -> None:
        return await self._transporter.close()

    def __deserialize(self, data: Optional[dict], klass: any = None) -> dict:
        """Deserializes dict, list, str into an object.

        :param data: dict, list or str.
        :param klass: class literal, or string of class name.

        :return: object.
        """
        if data is None:
            return None

        if isinstance(klass, str):
            if klass.startswith("List["):
                sub_kls = match(r"List\[(.*)]", klass).group(1)
                return [self.__deserialize(sub_data, sub_kls) for sub_data in data]

            if klass.startswith("Dict["):
                sub_kls = match(r"Dict\[([^,]*), (.*)]", klass).group(2)
                return {k: self.__deserialize(v, sub_kls) for k, v in data.items()}

            if klass in self.NATIVE_TYPES_MAPPING:
                klass = self.NATIVE_TYPES_MAPPING[klass]

        if klass in self.PRIMITIVE_TYPES:
            try:
                return klass(data)
            except UnicodeEncodeError:
                return str(data)
            except TypeError:
                return data
        elif klass == object:
            return data
        else:
            return klass.from_json(data)

    async def add_ab_tests_with_http_info(
        self,
        add_ab_tests_request: AddABTestsRequest,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Create an A/B test.

        Creates an A/B test.

        :param add_ab_tests_request: (required)
        :type add_ab_tests_request: AddABTestsRequest
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if add_ab_tests_request is None:
            raise ValueError(
                "'add_ab_tests_request' is required when calling 'add_ab_tests'"
            )

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/2/abtests"

        if add_ab_tests_request is not None:
            _body = add_ab_tests_request

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.POST,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def add_ab_tests(
        self,
        add_ab_tests_request: AddABTestsRequest,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Create an A/B test.

        Creates an A/B test.

        :param add_ab_tests_request: (required)
        :type add_ab_tests_request: AddABTestsRequest
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.add_ab_tests_with_http_info(
            add_ab_tests_request, request_options
        )

        return self.__deserialize(response.raw_data, ABTestResponse)

    async def custom_delete_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("'path' is required when calling 'custom_delete'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/1{path}".replace("{path}", path)

        if parameters is not None:
            _query_params.append(("parameters", parameters))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.DELETE,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def custom_delete(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """

        response = await self.custom_delete_with_http_info(
            path, parameters, request_options
        )

        return self.__deserialize(response.raw_data, object)

    async def custom_get_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("'path' is required when calling 'custom_get'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/1{path}".replace("{path}", path)

        if parameters is not None:
            _query_params.append(("parameters", parameters))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.GET,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def custom_get(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """

        response = await self.custom_get_with_http_info(
            path, parameters, request_options
        )

        return self.__deserialize(response.raw_data, object)

    async def custom_post_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("'path' is required when calling 'custom_post'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/1{path}".replace("{path}", path)

        if parameters is not None:
            _query_params.append(("parameters", parameters))

        if body is not None:
            _body = body

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.POST,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def custom_post(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """

        response = await self.custom_post_with_http_info(
            path, parameters, body, request_options
        )

        return self.__deserialize(response.raw_data, object)

    async def custom_put_with_http_info(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if path is None:
            raise ValueError("'path' is required when calling 'custom_put'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/1{path}".replace("{path}", path)

        if parameters is not None:
            _query_params.append(("parameters", parameters))

        if body is not None:
            _body = body

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.PUT,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def custom_put(
        self,
        path: Annotated[
            StrictStr,
            Field(
                description='Path of the endpoint, anything after "/1" must be specified.'
            ),
        ],
        parameters: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Query parameters to apply to the current query."),
        ] = None,
        body: Annotated[
            Optional[Dict[str, Any]],
            Field(description="Parameters to send with the custom request."),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> object:
        """
        Send requests to the Algolia REST API.

        This method allow you to send requests to the Algolia REST API.

        :param path: Path of the endpoint, anything after \"/1\" must be specified. (required)
        :type path: str
        :param parameters: Query parameters to apply to the current query.
        :type parameters: Dict[str, object]
        :param body: Parameters to send with the custom request.
        :type body: object
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'object' result object.
        """

        response = await self.custom_put_with_http_info(
            path, parameters, body, request_options
        )

        return self.__deserialize(response.raw_data, object)

    async def delete_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Delete an A/B test.

        Delete an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError("'id' is required when calling 'delete_ab_test'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/2/abtests/{id}".replace("{id}", quote(str(id)))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.DELETE,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def delete_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Delete an A/B test.

        Delete an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.delete_ab_test_with_http_info(id, request_options)

        return self.__deserialize(response.raw_data, ABTestResponse)

    async def get_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Get A/B test details.

        Get specific details for an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError("'id' is required when calling 'get_ab_test'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/2/abtests/{id}".replace("{id}", quote(str(id)))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.GET,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def get_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTest:
        """
        Get A/B test details.

        Get specific details for an A/B test. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTest' result object.
        """

        response = await self.get_ab_test_with_http_info(id, request_options)

        return self.__deserialize(response.raw_data, ABTest)

    async def list_ab_tests_with_http_info(
        self,
        offset: Annotated[
            Optional[StrictInt],
            Field(
                description="Position of the starting record. Used for paging. 0 is the first record."
            ),
        ] = None,
        limit: Annotated[
            Optional[StrictInt],
            Field(description="Number of records to return (page size)."),
        ] = None,
        index_prefix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices starting with this prefix."
            ),
        ] = None,
        index_suffix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices ending with this suffix."
            ),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        List all A/B tests.

        List all A/B tests.

        :param offset: Position of the starting record. Used for paging. 0 is the first record.
        :type offset: int
        :param limit: Number of records to return (page size).
        :type limit: int
        :param index_prefix: Only return A/B tests for indices starting with this prefix.
        :type index_prefix: str
        :param index_suffix: Only return A/B tests for indices ending with this suffix.
        :type index_suffix: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/2/abtests"

        if offset is not None:
            _query_params.append(("offset", offset))
        if limit is not None:
            _query_params.append(("limit", limit))
        if index_prefix is not None:
            _query_params.append(("indexPrefix", index_prefix))
        if index_suffix is not None:
            _query_params.append(("indexSuffix", index_suffix))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.GET,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def list_ab_tests(
        self,
        offset: Annotated[
            Optional[StrictInt],
            Field(
                description="Position of the starting record. Used for paging. 0 is the first record."
            ),
        ] = None,
        limit: Annotated[
            Optional[StrictInt],
            Field(description="Number of records to return (page size)."),
        ] = None,
        index_prefix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices starting with this prefix."
            ),
        ] = None,
        index_suffix: Annotated[
            Optional[StrictStr],
            Field(
                description="Only return A/B tests for indices ending with this suffix."
            ),
        ] = None,
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ListABTestsResponse:
        """
        List all A/B tests.

        List all A/B tests.

        :param offset: Position of the starting record. Used for paging. 0 is the first record.
        :type offset: int
        :param limit: Number of records to return (page size).
        :type limit: int
        :param index_prefix: Only return A/B tests for indices starting with this prefix.
        :type index_prefix: str
        :param index_suffix: Only return A/B tests for indices ending with this suffix.
        :type index_suffix: str
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ListABTestsResponse' result object.
        """

        response = await self.list_ab_tests_with_http_info(
            offset, limit, index_prefix, index_suffix, request_options
        )

        return self.__deserialize(response.raw_data, ListABTestsResponse)

    async def stop_ab_test_with_http_info(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ApiResponse[str]:
        """
        Stop an A/B test.

        If stopped, the test is over and can't be restarted. There is now only one index, receiving 100% of all search requests. The data gathered for stopped A/B tests is retained. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the raw algoliasearch 'APIResponse' object.
        """

        if id is None:
            raise ValueError("'id' is required when calling 'stop_ab_test'")

        _query_params: List[Tuple[str, str]] = []
        _body: Optional[bytes] = None
        _path = "/2/abtests/{id}/stop".replace("{id}", quote(str(id)))

        _param = self._transporter.param_serialize(
            query_params=_query_params,
            body=_body,
            request_options=request_options,
        )

        response = await self._transporter.request(
            verb=Verb.POST,
            path=_path,
            data=_param[0],
            request_options=_param[1],
        )

        response.data = response.raw_data

        return response

    async def stop_ab_test(
        self,
        id: Annotated[StrictInt, Field(description="Unique A/B test ID.")],
        request_options: Optional[Union[dict, RequestOptions]] = None,
    ) -> ABTestResponse:
        """
        Stop an A/B test.

        If stopped, the test is over and can't be restarted. There is now only one index, receiving 100% of all search requests. The data gathered for stopped A/B tests is retained. To determine the `id` for an A/B test, use the [`listABTests` operation](#tag/abtest/operation/listABTests).

        :param id: Unique A/B test ID. (required)
        :type id: int
        :param request_options: The request options to send along with the query, they will be merged with the transporter base parameters (headers, query params, timeouts, etc.). (optional)
        :return: Returns the deserialized response in a 'ABTestResponse' result object.
        """

        response = await self.stop_ab_test_with_http_info(id, request_options)

        return self.__deserialize(response.raw_data, ABTestResponse)

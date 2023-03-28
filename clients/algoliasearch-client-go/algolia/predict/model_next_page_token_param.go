// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// NextPageTokenParam struct for NextPageTokenParam
type NextPageTokenParam struct {
	// The token is used to navigate forward in the user list. To navigate from the current user list to the next page, the API generates the next page token and it sends the token in the response, beside the current user list. NOTE: This body param cannot be used with `previousPageToken` at the same time.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

type NextPageTokenParamOption func(f *NextPageTokenParam)

func WithNextPageTokenParamNextPageToken(val string) NextPageTokenParamOption {
	return func(f *NextPageTokenParam) {
		f.NextPageToken = &val
	}
}

// NewNextPageTokenParam instantiates a new NextPageTokenParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextPageTokenParam(opts ...NextPageTokenParamOption) *NextPageTokenParam {
	this := &NextPageTokenParam{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewNextPageTokenParamWithDefaults instantiates a new NextPageTokenParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextPageTokenParamWithDefaults() *NextPageTokenParam {
	this := &NextPageTokenParam{}
	return this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *NextPageTokenParam) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NextPageTokenParam) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *NextPageTokenParam) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *NextPageTokenParam) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o NextPageTokenParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.NextPageToken != nil {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

func (o NextPageTokenParam) String() string {
	out := "NextPageTokenParam {\n"
	out += fmt.Sprintf("  nextPageToken=%v\n", o.NextPageToken)
	out += "}"
	return out
}

type NullableNextPageTokenParam struct {
	value *NextPageTokenParam
	isSet bool
}

func (v NullableNextPageTokenParam) Get() *NextPageTokenParam {
	return v.value
}

func (v *NullableNextPageTokenParam) Set(val *NextPageTokenParam) {
	v.value = val
	v.isSet = true
}

func (v NullableNextPageTokenParam) IsSet() bool {
	return v.isSet
}

func (v *NullableNextPageTokenParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextPageTokenParam(val *NextPageTokenParam) *NullableNextPageTokenParam {
	return &NullableNextPageTokenParam{value: val, isSet: true}
}

func (v NullableNextPageTokenParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextPageTokenParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

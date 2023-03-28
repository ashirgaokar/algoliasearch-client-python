// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// Properties Properties for the user profile.
type Properties struct {
	// Raw user properties (key-value pairs).
	Raw map[string]interface{} `json:"raw,omitempty"`
	// Computed user properties (key-value pairs).
	Computed map[string]interface{} `json:"computed,omitempty"`
	// Custom user properties (key-value pairs).
	Custom map[string]interface{} `json:"custom,omitempty"`
}

type PropertiesOption func(f *Properties)

func WithPropertiesRaw(val map[string]interface{}) PropertiesOption {
	return func(f *Properties) {
		f.Raw = val
	}
}

func WithPropertiesComputed(val map[string]interface{}) PropertiesOption {
	return func(f *Properties) {
		f.Computed = val
	}
}

func WithPropertiesCustom(val map[string]interface{}) PropertiesOption {
	return func(f *Properties) {
		f.Custom = val
	}
}

// NewProperties instantiates a new Properties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperties(opts ...PropertiesOption) *Properties {
	this := &Properties{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewPropertiesWithDefaults instantiates a new Properties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertiesWithDefaults() *Properties {
	this := &Properties{}
	return this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *Properties) GetRaw() map[string]interface{} {
	if o == nil || o.Raw == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetRawOk() (map[string]interface{}, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *Properties) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given map[string]interface{} and assigns it to the Raw field.
func (o *Properties) SetRaw(v map[string]interface{}) {
	o.Raw = v
}

// GetComputed returns the Computed field value if set, zero value otherwise.
func (o *Properties) GetComputed() map[string]interface{} {
	if o == nil || o.Computed == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Computed
}

// GetComputedOk returns a tuple with the Computed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetComputedOk() (map[string]interface{}, bool) {
	if o == nil || o.Computed == nil {
		return nil, false
	}
	return o.Computed, true
}

// HasComputed returns a boolean if a field has been set.
func (o *Properties) HasComputed() bool {
	if o != nil && o.Computed != nil {
		return true
	}

	return false
}

// SetComputed gets a reference to the given map[string]interface{} and assigns it to the Computed field.
func (o *Properties) SetComputed(v map[string]interface{}) {
	o.Computed = v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *Properties) GetCustom() map[string]interface{} {
	if o == nil || o.Custom == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Properties) GetCustomOk() (map[string]interface{}, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *Properties) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given map[string]interface{} and assigns it to the Custom field.
func (o *Properties) SetCustom(v map[string]interface{}) {
	o.Custom = v
}

func (o Properties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Computed != nil {
		toSerialize["computed"] = o.Computed
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	return json.Marshal(toSerialize)
}

func (o Properties) String() string {
	out := "Properties {\n"
	out += fmt.Sprintf("  raw=%v\n", o.Raw)
	out += fmt.Sprintf("  computed=%v\n", o.Computed)
	out += fmt.Sprintf("  custom=%v\n", o.Custom)
	out += "}"
	return out
}

type NullableProperties struct {
	value *Properties
	isSet bool
}

func (v NullableProperties) Get() *Properties {
	return v.value
}

func (v *NullableProperties) Set(val *Properties) {
	v.value = val
	v.isSet = true
}

func (v NullableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperties(val *Properties) *NullableProperties {
	return &NullableProperties{value: val, isSet: true}
}

func (v NullableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

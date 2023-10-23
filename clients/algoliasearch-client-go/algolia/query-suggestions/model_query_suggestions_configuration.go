// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// QuerySuggestionsConfiguration Query Suggestions configuration.
type QuerySuggestionsConfiguration struct {
	// Algolia indices from which to get the popular searches for query suggestions.
	SourceIndices []SourceIndex `json:"sourceIndices"`
	Languages     *Languages    `json:"languages,omitempty"`
	// Patterns to exclude from query suggestions.
	Exclude []string `json:"exclude,omitempty"`
	// Turn on personalized query suggestions.
	EnablePersonalization *bool `json:"enablePersonalization,omitempty"`
	// Allow suggestions with special characters.
	AllowSpecialCharacters *bool `json:"allowSpecialCharacters,omitempty"`
}

type QuerySuggestionsConfigurationOption func(f *QuerySuggestionsConfiguration)

func WithQuerySuggestionsConfigurationLanguages(val Languages) QuerySuggestionsConfigurationOption {
	return func(f *QuerySuggestionsConfiguration) {
		f.Languages = &val
	}
}

func WithQuerySuggestionsConfigurationExclude(val []string) QuerySuggestionsConfigurationOption {
	return func(f *QuerySuggestionsConfiguration) {
		f.Exclude = val
	}
}

func WithQuerySuggestionsConfigurationEnablePersonalization(val bool) QuerySuggestionsConfigurationOption {
	return func(f *QuerySuggestionsConfiguration) {
		f.EnablePersonalization = &val
	}
}

func WithQuerySuggestionsConfigurationAllowSpecialCharacters(val bool) QuerySuggestionsConfigurationOption {
	return func(f *QuerySuggestionsConfiguration) {
		f.AllowSpecialCharacters = &val
	}
}

// NewQuerySuggestionsConfiguration instantiates a new QuerySuggestionsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySuggestionsConfiguration(sourceIndices []SourceIndex, opts ...QuerySuggestionsConfigurationOption) *QuerySuggestionsConfiguration {
	this := &QuerySuggestionsConfiguration{}
	this.SourceIndices = sourceIndices
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewQuerySuggestionsConfigurationWithDefaults instantiates a new QuerySuggestionsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySuggestionsConfigurationWithDefaults() *QuerySuggestionsConfiguration {
	this := &QuerySuggestionsConfiguration{}
	var enablePersonalization bool = false
	this.EnablePersonalization = &enablePersonalization
	var allowSpecialCharacters bool = false
	this.AllowSpecialCharacters = &allowSpecialCharacters
	return this
}

// GetSourceIndices returns the SourceIndices field value
func (o *QuerySuggestionsConfiguration) GetSourceIndices() []SourceIndex {
	if o == nil {
		var ret []SourceIndex
		return ret
	}

	return o.SourceIndices
}

// GetSourceIndicesOk returns a tuple with the SourceIndices field value
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfiguration) GetSourceIndicesOk() ([]SourceIndex, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIndices, true
}

// SetSourceIndices sets field value
func (o *QuerySuggestionsConfiguration) SetSourceIndices(v []SourceIndex) {
	o.SourceIndices = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *QuerySuggestionsConfiguration) GetLanguages() Languages {
	if o == nil || o.Languages == nil {
		var ret Languages
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfiguration) GetLanguagesOk() (*Languages, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *QuerySuggestionsConfiguration) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given Languages and assigns it to the Languages field.
func (o *QuerySuggestionsConfiguration) SetLanguages(v Languages) {
	o.Languages = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuerySuggestionsConfiguration) GetExclude() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuerySuggestionsConfiguration) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *QuerySuggestionsConfiguration) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *QuerySuggestionsConfiguration) SetExclude(v []string) {
	o.Exclude = v
}

// GetEnablePersonalization returns the EnablePersonalization field value if set, zero value otherwise.
func (o *QuerySuggestionsConfiguration) GetEnablePersonalization() bool {
	if o == nil || o.EnablePersonalization == nil {
		var ret bool
		return ret
	}
	return *o.EnablePersonalization
}

// GetEnablePersonalizationOk returns a tuple with the EnablePersonalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfiguration) GetEnablePersonalizationOk() (*bool, bool) {
	if o == nil || o.EnablePersonalization == nil {
		return nil, false
	}
	return o.EnablePersonalization, true
}

// HasEnablePersonalization returns a boolean if a field has been set.
func (o *QuerySuggestionsConfiguration) HasEnablePersonalization() bool {
	if o != nil && o.EnablePersonalization != nil {
		return true
	}

	return false
}

// SetEnablePersonalization gets a reference to the given bool and assigns it to the EnablePersonalization field.
func (o *QuerySuggestionsConfiguration) SetEnablePersonalization(v bool) {
	o.EnablePersonalization = &v
}

// GetAllowSpecialCharacters returns the AllowSpecialCharacters field value if set, zero value otherwise.
func (o *QuerySuggestionsConfiguration) GetAllowSpecialCharacters() bool {
	if o == nil || o.AllowSpecialCharacters == nil {
		var ret bool
		return ret
	}
	return *o.AllowSpecialCharacters
}

// GetAllowSpecialCharactersOk returns a tuple with the AllowSpecialCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySuggestionsConfiguration) GetAllowSpecialCharactersOk() (*bool, bool) {
	if o == nil || o.AllowSpecialCharacters == nil {
		return nil, false
	}
	return o.AllowSpecialCharacters, true
}

// HasAllowSpecialCharacters returns a boolean if a field has been set.
func (o *QuerySuggestionsConfiguration) HasAllowSpecialCharacters() bool {
	if o != nil && o.AllowSpecialCharacters != nil {
		return true
	}

	return false
}

// SetAllowSpecialCharacters gets a reference to the given bool and assigns it to the AllowSpecialCharacters field.
func (o *QuerySuggestionsConfiguration) SetAllowSpecialCharacters(v bool) {
	o.AllowSpecialCharacters = &v
}

func (o QuerySuggestionsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["sourceIndices"] = o.SourceIndices
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.EnablePersonalization != nil {
		toSerialize["enablePersonalization"] = o.EnablePersonalization
	}
	if o.AllowSpecialCharacters != nil {
		toSerialize["allowSpecialCharacters"] = o.AllowSpecialCharacters
	}
	return json.Marshal(toSerialize)
}

func (o QuerySuggestionsConfiguration) String() string {
	out := ""
	out += fmt.Sprintf("  sourceIndices=%v\n", o.SourceIndices)
	out += fmt.Sprintf("  languages=%v\n", o.Languages)
	out += fmt.Sprintf("  exclude=%v\n", o.Exclude)
	out += fmt.Sprintf("  enablePersonalization=%v\n", o.EnablePersonalization)
	out += fmt.Sprintf("  allowSpecialCharacters=%v\n", o.AllowSpecialCharacters)
	return fmt.Sprintf("QuerySuggestionsConfiguration {\n%s}", out)
}

type NullableQuerySuggestionsConfiguration struct {
	value *QuerySuggestionsConfiguration
	isSet bool
}

func (v NullableQuerySuggestionsConfiguration) Get() *QuerySuggestionsConfiguration {
	return v.value
}

func (v *NullableQuerySuggestionsConfiguration) Set(val *QuerySuggestionsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySuggestionsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySuggestionsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySuggestionsConfiguration(val *QuerySuggestionsConfiguration) *NullableQuerySuggestionsConfiguration {
	return &NullableQuerySuggestionsConfiguration{value: val, isSet: true}
}

func (v NullableQuerySuggestionsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySuggestionsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package insights

import (
	"encoding/json"
	"fmt"
)

// ViewedFilters Use this method to capture active filters. For example, when browsing a category page, users see content filtered on that specific category.
type ViewedFilters struct {
	// Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string    `json:"eventName"`
	EventType ViewEvent `json:"eventType"`
	// Name of the Algolia index.
	Index string `json:"index"`
	// Facet filters.  Each facet filter string must be URL-encoded, such as, `discount:10%25`.
	Filters []string `json:"filters"`
	// Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// User token for authenticated users.
	AuthenticatedUserToken *string `json:"authenticatedUserToken,omitempty"`
}

type ViewedFiltersOption func(f *ViewedFilters)

func WithViewedFiltersTimestamp(val int64) ViewedFiltersOption {
	return func(f *ViewedFilters) {
		f.Timestamp = &val
	}
}

func WithViewedFiltersAuthenticatedUserToken(val string) ViewedFiltersOption {
	return func(f *ViewedFilters) {
		f.AuthenticatedUserToken = &val
	}
}

// NewViewedFilters instantiates a new ViewedFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewedFilters(eventName string, eventType ViewEvent, index string, filters []string, userToken string, opts ...ViewedFiltersOption) *ViewedFilters {
	this := &ViewedFilters{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.Filters = filters
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewViewedFiltersWithDefaults instantiates a new ViewedFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewedFiltersWithDefaults() *ViewedFilters {
	this := &ViewedFilters{}
	return this
}

// GetEventName returns the EventName field value
func (o *ViewedFilters) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *ViewedFilters) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value
func (o *ViewedFilters) GetEventType() ViewEvent {
	if o == nil {
		var ret ViewEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetEventTypeOk() (*ViewEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ViewedFilters) SetEventType(v ViewEvent) {
	o.EventType = v
}

// GetIndex returns the Index field value
func (o *ViewedFilters) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ViewedFilters) SetIndex(v string) {
	o.Index = v
}

// GetFilters returns the Filters field value
func (o *ViewedFilters) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetFiltersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filters, true
}

// SetFilters sets field value
func (o *ViewedFilters) SetFilters(v []string) {
	o.Filters = v
}

// GetUserToken returns the UserToken field value
func (o *ViewedFilters) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *ViewedFilters) SetUserToken(v string) {
	o.UserToken = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ViewedFilters) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ViewedFilters) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ViewedFilters) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetAuthenticatedUserToken returns the AuthenticatedUserToken field value if set, zero value otherwise.
func (o *ViewedFilters) GetAuthenticatedUserToken() string {
	if o == nil || o.AuthenticatedUserToken == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatedUserToken
}

// GetAuthenticatedUserTokenOk returns a tuple with the AuthenticatedUserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewedFilters) GetAuthenticatedUserTokenOk() (*string, bool) {
	if o == nil || o.AuthenticatedUserToken == nil {
		return nil, false
	}
	return o.AuthenticatedUserToken, true
}

// HasAuthenticatedUserToken returns a boolean if a field has been set.
func (o *ViewedFilters) HasAuthenticatedUserToken() bool {
	if o != nil && o.AuthenticatedUserToken != nil {
		return true
	}

	return false
}

// SetAuthenticatedUserToken gets a reference to the given string and assigns it to the AuthenticatedUserToken field.
func (o *ViewedFilters) SetAuthenticatedUserToken(v string) {
	o.AuthenticatedUserToken = &v
}

func (o ViewedFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["filters"] = o.Filters
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.AuthenticatedUserToken != nil {
		toSerialize["authenticatedUserToken"] = o.AuthenticatedUserToken
	}
	return json.Marshal(toSerialize)
}

func (o ViewedFilters) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  filters=%v\n", o.Filters)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	out += fmt.Sprintf("  authenticatedUserToken=%v\n", o.AuthenticatedUserToken)
	return fmt.Sprintf("ViewedFilters {\n%s}", out)
}

type NullableViewedFilters struct {
	value *ViewedFilters
	isSet bool
}

func (v NullableViewedFilters) Get() *ViewedFilters {
	return v.value
}

func (v *NullableViewedFilters) Set(val *ViewedFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableViewedFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableViewedFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewedFilters(val *ViewedFilters) *NullableViewedFilters {
	return &NullableViewedFilters{value: val, isSet: true}
}

func (v NullableViewedFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewedFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

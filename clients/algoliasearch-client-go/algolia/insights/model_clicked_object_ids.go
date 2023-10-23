// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package insights

import (
	"encoding/json"
	"fmt"
)

// ClickedObjectIDs Use this event to track when users click items unrelated to a previous Algolia request. For example, if you don't use Algolia to build your category pages, use this event.  To track click events related to Algolia requests, use the \"Clicked object IDs after search\" event.
type ClickedObjectIDs struct {
	// Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
	EventName string     `json:"eventName"`
	EventType ClickEvent `json:"eventType"`
	// Name of the Algolia index.
	Index string `json:"index"`
	// List of object identifiers for items of an Algolia index.
	ObjectIDs []string `json:"objectIDs"`
	// Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
	UserToken string `json:"userToken"`
	// Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type ClickedObjectIDsOption func(f *ClickedObjectIDs)

func WithClickedObjectIDsTimestamp(val int64) ClickedObjectIDsOption {
	return func(f *ClickedObjectIDs) {
		f.Timestamp = &val
	}
}

// NewClickedObjectIDs instantiates a new ClickedObjectIDs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickedObjectIDs(eventName string, eventType ClickEvent, index string, objectIDs []string, userToken string, opts ...ClickedObjectIDsOption) *ClickedObjectIDs {
	this := &ClickedObjectIDs{}
	this.EventName = eventName
	this.EventType = eventType
	this.Index = index
	this.ObjectIDs = objectIDs
	this.UserToken = userToken
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewClickedObjectIDsWithDefaults instantiates a new ClickedObjectIDs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickedObjectIDsWithDefaults() *ClickedObjectIDs {
	this := &ClickedObjectIDs{}
	return this
}

// GetEventName returns the EventName field value
func (o *ClickedObjectIDs) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *ClickedObjectIDs) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value
func (o *ClickedObjectIDs) GetEventType() ClickEvent {
	if o == nil {
		var ret ClickEvent
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetEventTypeOk() (*ClickEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ClickedObjectIDs) SetEventType(v ClickEvent) {
	o.EventType = v
}

// GetIndex returns the Index field value
func (o *ClickedObjectIDs) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ClickedObjectIDs) SetIndex(v string) {
	o.Index = v
}

// GetObjectIDs returns the ObjectIDs field value
func (o *ClickedObjectIDs) GetObjectIDs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectIDs
}

// GetObjectIDsOk returns a tuple with the ObjectIDs field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetObjectIDsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectIDs, true
}

// SetObjectIDs sets field value
func (o *ClickedObjectIDs) SetObjectIDs(v []string) {
	o.ObjectIDs = v
}

// GetUserToken returns the UserToken field value
func (o *ClickedObjectIDs) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *ClickedObjectIDs) SetUserToken(v string) {
	o.UserToken = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClickedObjectIDs) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickedObjectIDs) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClickedObjectIDs) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ClickedObjectIDs) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o ClickedObjectIDs) MarshalJSON() ([]byte, error) {
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
		toSerialize["objectIDs"] = o.ObjectIDs
	}
	if true {
		toSerialize["userToken"] = o.UserToken
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

func (o ClickedObjectIDs) String() string {
	out := ""
	out += fmt.Sprintf("  eventName=%v\n", o.EventName)
	out += fmt.Sprintf("  eventType=%v\n", o.EventType)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  objectIDs=%v\n", o.ObjectIDs)
	out += fmt.Sprintf("  userToken=%v\n", o.UserToken)
	out += fmt.Sprintf("  timestamp=%v\n", o.Timestamp)
	return fmt.Sprintf("ClickedObjectIDs {\n%s}", out)
}

type NullableClickedObjectIDs struct {
	value *ClickedObjectIDs
	isSet bool
}

func (v NullableClickedObjectIDs) Get() *ClickedObjectIDs {
	return v.value
}

func (v *NullableClickedObjectIDs) Set(val *ClickedObjectIDs) {
	v.value = val
	v.isSet = true
}

func (v NullableClickedObjectIDs) IsSet() bool {
	return v.isSet
}

func (v *NullableClickedObjectIDs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickedObjectIDs(val *ClickedObjectIDs) *NullableClickedObjectIDs {
	return &NullableClickedObjectIDs{value: val, isSet: true}
}

func (v NullableClickedObjectIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickedObjectIDs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

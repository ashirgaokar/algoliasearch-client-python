// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// ModelMetrics struct for ModelMetrics
type ModelMetrics struct {
	Precision *float64 `json:"precision,omitempty"`
	Recall    *float64 `json:"recall,omitempty"`
	Mrr       *float64 `json:"mrr,omitempty"`
	Coverage  *float64 `json:"coverage,omitempty"`
	F1Score   *float64 `json:"f1_score,omitempty"`
	// Date of last update (ISO-8601 format).
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type ModelMetricsOption func(f *ModelMetrics)

func WithModelMetricsPrecision(val float64) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.Precision = &val
	}
}

func WithModelMetricsRecall(val float64) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.Recall = &val
	}
}

func WithModelMetricsMrr(val float64) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.Mrr = &val
	}
}

func WithModelMetricsCoverage(val float64) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.Coverage = &val
	}
}

func WithModelMetricsF1Score(val float64) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.F1Score = &val
	}
}

func WithModelMetricsUpdatedAt(val string) ModelMetricsOption {
	return func(f *ModelMetrics) {
		f.UpdatedAt = &val
	}
}

// NewModelMetrics instantiates a new ModelMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelMetrics(opts ...ModelMetricsOption) *ModelMetrics {
	this := &ModelMetrics{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewModelMetricsWithDefaults instantiates a new ModelMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelMetricsWithDefaults() *ModelMetrics {
	this := &ModelMetrics{}
	return this
}

// GetPrecision returns the Precision field value if set, zero value otherwise.
func (o *ModelMetrics) GetPrecision() float64 {
	if o == nil || o.Precision == nil {
		var ret float64
		return ret
	}
	return *o.Precision
}

// GetPrecisionOk returns a tuple with the Precision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetPrecisionOk() (*float64, bool) {
	if o == nil || o.Precision == nil {
		return nil, false
	}
	return o.Precision, true
}

// HasPrecision returns a boolean if a field has been set.
func (o *ModelMetrics) HasPrecision() bool {
	if o != nil && o.Precision != nil {
		return true
	}

	return false
}

// SetPrecision gets a reference to the given float64 and assigns it to the Precision field.
func (o *ModelMetrics) SetPrecision(v float64) {
	o.Precision = &v
}

// GetRecall returns the Recall field value if set, zero value otherwise.
func (o *ModelMetrics) GetRecall() float64 {
	if o == nil || o.Recall == nil {
		var ret float64
		return ret
	}
	return *o.Recall
}

// GetRecallOk returns a tuple with the Recall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetRecallOk() (*float64, bool) {
	if o == nil || o.Recall == nil {
		return nil, false
	}
	return o.Recall, true
}

// HasRecall returns a boolean if a field has been set.
func (o *ModelMetrics) HasRecall() bool {
	if o != nil && o.Recall != nil {
		return true
	}

	return false
}

// SetRecall gets a reference to the given float64 and assigns it to the Recall field.
func (o *ModelMetrics) SetRecall(v float64) {
	o.Recall = &v
}

// GetMrr returns the Mrr field value if set, zero value otherwise.
func (o *ModelMetrics) GetMrr() float64 {
	if o == nil || o.Mrr == nil {
		var ret float64
		return ret
	}
	return *o.Mrr
}

// GetMrrOk returns a tuple with the Mrr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetMrrOk() (*float64, bool) {
	if o == nil || o.Mrr == nil {
		return nil, false
	}
	return o.Mrr, true
}

// HasMrr returns a boolean if a field has been set.
func (o *ModelMetrics) HasMrr() bool {
	if o != nil && o.Mrr != nil {
		return true
	}

	return false
}

// SetMrr gets a reference to the given float64 and assigns it to the Mrr field.
func (o *ModelMetrics) SetMrr(v float64) {
	o.Mrr = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *ModelMetrics) GetCoverage() float64 {
	if o == nil || o.Coverage == nil {
		var ret float64
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetCoverageOk() (*float64, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *ModelMetrics) HasCoverage() bool {
	if o != nil && o.Coverage != nil {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given float64 and assigns it to the Coverage field.
func (o *ModelMetrics) SetCoverage(v float64) {
	o.Coverage = &v
}

// GetF1Score returns the F1Score field value if set, zero value otherwise.
func (o *ModelMetrics) GetF1Score() float64 {
	if o == nil || o.F1Score == nil {
		var ret float64
		return ret
	}
	return *o.F1Score
}

// GetF1ScoreOk returns a tuple with the F1Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetF1ScoreOk() (*float64, bool) {
	if o == nil || o.F1Score == nil {
		return nil, false
	}
	return o.F1Score, true
}

// HasF1Score returns a boolean if a field has been set.
func (o *ModelMetrics) HasF1Score() bool {
	if o != nil && o.F1Score != nil {
		return true
	}

	return false
}

// SetF1Score gets a reference to the given float64 and assigns it to the F1Score field.
func (o *ModelMetrics) SetF1Score(v float64) {
	o.F1Score = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelMetrics) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMetrics) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelMetrics) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelMetrics) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Precision != nil {
		toSerialize["precision"] = o.Precision
	}
	if o.Recall != nil {
		toSerialize["recall"] = o.Recall
	}
	if o.Mrr != nil {
		toSerialize["mrr"] = o.Mrr
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.F1Score != nil {
		toSerialize["f1_score"] = o.F1Score
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o ModelMetrics) String() string {
	out := "ModelMetrics {\n"
	out += fmt.Sprintf("  precision=%v\n", o.Precision)
	out += fmt.Sprintf("  recall=%v\n", o.Recall)
	out += fmt.Sprintf("  mrr=%v\n", o.Mrr)
	out += fmt.Sprintf("  coverage=%v\n", o.Coverage)
	out += fmt.Sprintf("  f1_score=%v\n", o.F1Score)
	out += fmt.Sprintf("  updatedAt=%v\n", o.UpdatedAt)
	out += "}"
	return out
}

type NullableModelMetrics struct {
	value *ModelMetrics
	isSet bool
}

func (v NullableModelMetrics) Get() *ModelMetrics {
	return v.value
}

func (v *NullableModelMetrics) Set(val *ModelMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMetrics(val *ModelMetrics) *NullableModelMetrics {
	return &NullableModelMetrics{value: val, isSet: true}
}

func (v NullableModelMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

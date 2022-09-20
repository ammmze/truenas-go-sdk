/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// CatalogCreate0 struct for CatalogCreate0
type CatalogCreate0 struct {
	Label           *string       `json:"label,omitempty"`
	Repository      *string       `json:"repository,omitempty"`
	Branch          *string       `json:"branch,omitempty"`
	PreferredTrains []interface{} `json:"preferred_trains,omitempty"`
	Force           *bool         `json:"force,omitempty"`
}

// NewCatalogCreate0 instantiates a new CatalogCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogCreate0() *CatalogCreate0 {
	this := CatalogCreate0{}
	var force bool = false
	this.Force = &force
	return &this
}

// NewCatalogCreate0WithDefaults instantiates a new CatalogCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogCreate0WithDefaults() *CatalogCreate0 {
	this := CatalogCreate0{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CatalogCreate0) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCreate0) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CatalogCreate0) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CatalogCreate0) SetLabel(v string) {
	o.Label = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CatalogCreate0) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCreate0) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CatalogCreate0) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *CatalogCreate0) SetRepository(v string) {
	o.Repository = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *CatalogCreate0) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCreate0) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *CatalogCreate0) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *CatalogCreate0) SetBranch(v string) {
	o.Branch = &v
}

// GetPreferredTrains returns the PreferredTrains field value if set, zero value otherwise.
func (o *CatalogCreate0) GetPreferredTrains() []interface{} {
	if o == nil || o.PreferredTrains == nil {
		var ret []interface{}
		return ret
	}
	return o.PreferredTrains
}

// GetPreferredTrainsOk returns a tuple with the PreferredTrains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCreate0) GetPreferredTrainsOk() ([]interface{}, bool) {
	if o == nil || o.PreferredTrains == nil {
		return nil, false
	}
	return o.PreferredTrains, true
}

// HasPreferredTrains returns a boolean if a field has been set.
func (o *CatalogCreate0) HasPreferredTrains() bool {
	if o != nil && o.PreferredTrains != nil {
		return true
	}

	return false
}

// SetPreferredTrains gets a reference to the given []interface{} and assigns it to the PreferredTrains field.
func (o *CatalogCreate0) SetPreferredTrains(v []interface{}) {
	o.PreferredTrains = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *CatalogCreate0) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogCreate0) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *CatalogCreate0) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *CatalogCreate0) SetForce(v bool) {
	o.Force = &v
}

func (o CatalogCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.PreferredTrains != nil {
		toSerialize["preferred_trains"] = o.PreferredTrains
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogCreate0 struct {
	value *CatalogCreate0
	isSet bool
}

func (v NullableCatalogCreate0) Get() *CatalogCreate0 {
	return v.value
}

func (v *NullableCatalogCreate0) Set(val *CatalogCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogCreate0(val *CatalogCreate0) *NullableCatalogCreate0 {
	return &NullableCatalogCreate0{value: val, isSet: true}
}

func (v NullableCatalogCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

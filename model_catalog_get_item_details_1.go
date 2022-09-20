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

// CatalogGetItemDetails1 struct for CatalogGetItemDetails1
type CatalogGetItemDetails1 struct {
	Cache   *bool   `json:"cache,omitempty"`
	Catalog *string `json:"catalog,omitempty"`
	Train   *string `json:"train,omitempty"`
}

// NewCatalogGetItemDetails1 instantiates a new CatalogGetItemDetails1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogGetItemDetails1() *CatalogGetItemDetails1 {
	this := CatalogGetItemDetails1{}
	return &this
}

// NewCatalogGetItemDetails1WithDefaults instantiates a new CatalogGetItemDetails1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogGetItemDetails1WithDefaults() *CatalogGetItemDetails1 {
	this := CatalogGetItemDetails1{}
	return &this
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *CatalogGetItemDetails1) GetCache() bool {
	if o == nil || o.Cache == nil {
		var ret bool
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGetItemDetails1) GetCacheOk() (*bool, bool) {
	if o == nil || o.Cache == nil {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *CatalogGetItemDetails1) HasCache() bool {
	if o != nil && o.Cache != nil {
		return true
	}

	return false
}

// SetCache gets a reference to the given bool and assigns it to the Cache field.
func (o *CatalogGetItemDetails1) SetCache(v bool) {
	o.Cache = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *CatalogGetItemDetails1) GetCatalog() string {
	if o == nil || o.Catalog == nil {
		var ret string
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGetItemDetails1) GetCatalogOk() (*string, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *CatalogGetItemDetails1) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given string and assigns it to the Catalog field.
func (o *CatalogGetItemDetails1) SetCatalog(v string) {
	o.Catalog = &v
}

// GetTrain returns the Train field value if set, zero value otherwise.
func (o *CatalogGetItemDetails1) GetTrain() string {
	if o == nil || o.Train == nil {
		var ret string
		return ret
	}
	return *o.Train
}

// GetTrainOk returns a tuple with the Train field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogGetItemDetails1) GetTrainOk() (*string, bool) {
	if o == nil || o.Train == nil {
		return nil, false
	}
	return o.Train, true
}

// HasTrain returns a boolean if a field has been set.
func (o *CatalogGetItemDetails1) HasTrain() bool {
	if o != nil && o.Train != nil {
		return true
	}

	return false
}

// SetTrain gets a reference to the given string and assigns it to the Train field.
func (o *CatalogGetItemDetails1) SetTrain(v string) {
	o.Train = &v
}

func (o CatalogGetItemDetails1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cache != nil {
		toSerialize["cache"] = o.Cache
	}
	if o.Catalog != nil {
		toSerialize["catalog"] = o.Catalog
	}
	if o.Train != nil {
		toSerialize["train"] = o.Train
	}
	return json.Marshal(toSerialize)
}

type NullableCatalogGetItemDetails1 struct {
	value *CatalogGetItemDetails1
	isSet bool
}

func (v NullableCatalogGetItemDetails1) Get() *CatalogGetItemDetails1 {
	return v.value
}

func (v *NullableCatalogGetItemDetails1) Set(val *CatalogGetItemDetails1) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogGetItemDetails1) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogGetItemDetails1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogGetItemDetails1(val *CatalogGetItemDetails1) *NullableCatalogGetItemDetails1 {
	return &NullableCatalogGetItemDetails1{value: val, isSet: true}
}

func (v NullableCatalogGetItemDetails1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogGetItemDetails1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

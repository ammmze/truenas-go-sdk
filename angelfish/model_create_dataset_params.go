/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package angelfish

import (
	"encoding/json"
)

// CreateDatasetParams struct for CreateDatasetParams
type CreateDatasetParams struct {
	Atime                *string                               `json:"atime,omitempty"`
	Aclmode              *string                               `json:"aclmode,omitempty"`
	Name                 string                                `json:"name"`
	Comments             *string                               `json:"comments,omitempty"`
	Compression          *string                               `json:"compression,omitempty"`
	Casesensitivity      *string                               `json:"casesensitivity,omitempty"`
	Copies               *int32                                `json:"copies,omitempty"`
	Deduplication        *string                               `json:"deduplication,omitempty"`
	Encryption           *bool                                 `json:"encryption,omitempty"`
	EncryptionOptions    *CreateDatasetParamsEncryptionOptions `json:"encryption_options,omitempty"`
	Exec                 *string                               `json:"exec,omitempty"`
	ForceSize            *bool                                 `json:"force_size,omitempty"`
	InheritEncryption    *bool                                 `json:"inherit_encryption,omitempty"`
	Quota                *int64                                `json:"quota,omitempty"`
	QuotaCritical        *int64                                `json:"quota_critical,omitempty"`
	QuotaWarning         *int64                                `json:"quota_warning,omitempty"`
	Volsize              *int64                                `json:"volsize,omitempty"`
	Volblocksize         *string                               `json:"volblocksize,omitempty"`
	Readonly             *string                               `json:"readonly,omitempty"`
	Recordsize           *string                               `json:"recordsize,omitempty"`
	Refquota             *int64                                `json:"refquota,omitempty"`
	RefquotaCritical     *int64                                `json:"refquota_critical,omitempty"`
	RefquotaWarning      *int64                                `json:"refquota_warning,omitempty"`
	Refreservation       *int64                                `json:"refreservation,omitempty"`
	Reservation          *int64                                `json:"reservation,omitempty"`
	ShareType            *string                               `json:"share_type,omitempty"`
	Snapdir              *string                               `json:"snapdir,omitempty"`
	Sync                 *string                               `json:"sync,omitempty"`
	Type                 *string                               `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDatasetParams CreateDatasetParams

// NewCreateDatasetParams instantiates a new CreateDatasetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatasetParams(name string) *CreateDatasetParams {
	this := CreateDatasetParams{}
	this.Name = name
	return &this
}

// NewCreateDatasetParamsWithDefaults instantiates a new CreateDatasetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatasetParamsWithDefaults() *CreateDatasetParams {
	this := CreateDatasetParams{}
	return &this
}

// GetAtime returns the Atime field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetAtime() string {
	if o == nil || isNil(o.Atime) {
		var ret string
		return ret
	}
	return *o.Atime
}

// GetAtimeOk returns a tuple with the Atime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetAtimeOk() (*string, bool) {
	if o == nil || isNil(o.Atime) {
		return nil, false
	}
	return o.Atime, true
}

// HasAtime returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasAtime() bool {
	if o != nil && !isNil(o.Atime) {
		return true
	}

	return false
}

// SetAtime gets a reference to the given string and assigns it to the Atime field.
func (o *CreateDatasetParams) SetAtime(v string) {
	o.Atime = &v
}

// GetAclmode returns the Aclmode field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetAclmode() string {
	if o == nil || isNil(o.Aclmode) {
		var ret string
		return ret
	}
	return *o.Aclmode
}

// GetAclmodeOk returns a tuple with the Aclmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetAclmodeOk() (*string, bool) {
	if o == nil || isNil(o.Aclmode) {
		return nil, false
	}
	return o.Aclmode, true
}

// HasAclmode returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasAclmode() bool {
	if o != nil && !isNil(o.Aclmode) {
		return true
	}

	return false
}

// SetAclmode gets a reference to the given string and assigns it to the Aclmode field.
func (o *CreateDatasetParams) SetAclmode(v string) {
	o.Aclmode = &v
}

// GetName returns the Name field value
func (o *CreateDatasetParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDatasetParams) SetName(v string) {
	o.Name = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetComments() string {
	if o == nil || isNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetCommentsOk() (*string, bool) {
	if o == nil || isNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasComments() bool {
	if o != nil && !isNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *CreateDatasetParams) SetComments(v string) {
	o.Comments = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetCompression() string {
	if o == nil || isNil(o.Compression) {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetCompressionOk() (*string, bool) {
	if o == nil || isNil(o.Compression) {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasCompression() bool {
	if o != nil && !isNil(o.Compression) {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *CreateDatasetParams) SetCompression(v string) {
	o.Compression = &v
}

// GetCasesensitivity returns the Casesensitivity field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetCasesensitivity() string {
	if o == nil || isNil(o.Casesensitivity) {
		var ret string
		return ret
	}
	return *o.Casesensitivity
}

// GetCasesensitivityOk returns a tuple with the Casesensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetCasesensitivityOk() (*string, bool) {
	if o == nil || isNil(o.Casesensitivity) {
		return nil, false
	}
	return o.Casesensitivity, true
}

// HasCasesensitivity returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasCasesensitivity() bool {
	if o != nil && !isNil(o.Casesensitivity) {
		return true
	}

	return false
}

// SetCasesensitivity gets a reference to the given string and assigns it to the Casesensitivity field.
func (o *CreateDatasetParams) SetCasesensitivity(v string) {
	o.Casesensitivity = &v
}

// GetCopies returns the Copies field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetCopies() int32 {
	if o == nil || isNil(o.Copies) {
		var ret int32
		return ret
	}
	return *o.Copies
}

// GetCopiesOk returns a tuple with the Copies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetCopiesOk() (*int32, bool) {
	if o == nil || isNil(o.Copies) {
		return nil, false
	}
	return o.Copies, true
}

// HasCopies returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasCopies() bool {
	if o != nil && !isNil(o.Copies) {
		return true
	}

	return false
}

// SetCopies gets a reference to the given int32 and assigns it to the Copies field.
func (o *CreateDatasetParams) SetCopies(v int32) {
	o.Copies = &v
}

// GetDeduplication returns the Deduplication field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetDeduplication() string {
	if o == nil || isNil(o.Deduplication) {
		var ret string
		return ret
	}
	return *o.Deduplication
}

// GetDeduplicationOk returns a tuple with the Deduplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetDeduplicationOk() (*string, bool) {
	if o == nil || isNil(o.Deduplication) {
		return nil, false
	}
	return o.Deduplication, true
}

// HasDeduplication returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasDeduplication() bool {
	if o != nil && !isNil(o.Deduplication) {
		return true
	}

	return false
}

// SetDeduplication gets a reference to the given string and assigns it to the Deduplication field.
func (o *CreateDatasetParams) SetDeduplication(v string) {
	o.Deduplication = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetEncryption() bool {
	if o == nil || isNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetEncryptionOk() (*bool, bool) {
	if o == nil || isNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasEncryption() bool {
	if o != nil && !isNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *CreateDatasetParams) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetEncryptionOptions returns the EncryptionOptions field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetEncryptionOptions() CreateDatasetParamsEncryptionOptions {
	if o == nil || isNil(o.EncryptionOptions) {
		var ret CreateDatasetParamsEncryptionOptions
		return ret
	}
	return *o.EncryptionOptions
}

// GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetEncryptionOptionsOk() (*CreateDatasetParamsEncryptionOptions, bool) {
	if o == nil || isNil(o.EncryptionOptions) {
		return nil, false
	}
	return o.EncryptionOptions, true
}

// HasEncryptionOptions returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasEncryptionOptions() bool {
	if o != nil && !isNil(o.EncryptionOptions) {
		return true
	}

	return false
}

// SetEncryptionOptions gets a reference to the given CreateDatasetParamsEncryptionOptions and assigns it to the EncryptionOptions field.
func (o *CreateDatasetParams) SetEncryptionOptions(v CreateDatasetParamsEncryptionOptions) {
	o.EncryptionOptions = &v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetExec() string {
	if o == nil || isNil(o.Exec) {
		var ret string
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetExecOk() (*string, bool) {
	if o == nil || isNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasExec() bool {
	if o != nil && !isNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given string and assigns it to the Exec field.
func (o *CreateDatasetParams) SetExec(v string) {
	o.Exec = &v
}

// GetForceSize returns the ForceSize field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetForceSize() bool {
	if o == nil || isNil(o.ForceSize) {
		var ret bool
		return ret
	}
	return *o.ForceSize
}

// GetForceSizeOk returns a tuple with the ForceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetForceSizeOk() (*bool, bool) {
	if o == nil || isNil(o.ForceSize) {
		return nil, false
	}
	return o.ForceSize, true
}

// HasForceSize returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasForceSize() bool {
	if o != nil && !isNil(o.ForceSize) {
		return true
	}

	return false
}

// SetForceSize gets a reference to the given bool and assigns it to the ForceSize field.
func (o *CreateDatasetParams) SetForceSize(v bool) {
	o.ForceSize = &v
}

// GetInheritEncryption returns the InheritEncryption field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetInheritEncryption() bool {
	if o == nil || isNil(o.InheritEncryption) {
		var ret bool
		return ret
	}
	return *o.InheritEncryption
}

// GetInheritEncryptionOk returns a tuple with the InheritEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetInheritEncryptionOk() (*bool, bool) {
	if o == nil || isNil(o.InheritEncryption) {
		return nil, false
	}
	return o.InheritEncryption, true
}

// HasInheritEncryption returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasInheritEncryption() bool {
	if o != nil && !isNil(o.InheritEncryption) {
		return true
	}

	return false
}

// SetInheritEncryption gets a reference to the given bool and assigns it to the InheritEncryption field.
func (o *CreateDatasetParams) SetInheritEncryption(v bool) {
	o.InheritEncryption = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetQuota() int64 {
	if o == nil || isNil(o.Quota) {
		var ret int64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetQuotaOk() (*int64, bool) {
	if o == nil || isNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasQuota() bool {
	if o != nil && !isNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int64 and assigns it to the Quota field.
func (o *CreateDatasetParams) SetQuota(v int64) {
	o.Quota = &v
}

// GetQuotaCritical returns the QuotaCritical field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetQuotaCritical() int64 {
	if o == nil || isNil(o.QuotaCritical) {
		var ret int64
		return ret
	}
	return *o.QuotaCritical
}

// GetQuotaCriticalOk returns a tuple with the QuotaCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetQuotaCriticalOk() (*int64, bool) {
	if o == nil || isNil(o.QuotaCritical) {
		return nil, false
	}
	return o.QuotaCritical, true
}

// HasQuotaCritical returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasQuotaCritical() bool {
	if o != nil && !isNil(o.QuotaCritical) {
		return true
	}

	return false
}

// SetQuotaCritical gets a reference to the given int64 and assigns it to the QuotaCritical field.
func (o *CreateDatasetParams) SetQuotaCritical(v int64) {
	o.QuotaCritical = &v
}

// GetQuotaWarning returns the QuotaWarning field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetQuotaWarning() int64 {
	if o == nil || isNil(o.QuotaWarning) {
		var ret int64
		return ret
	}
	return *o.QuotaWarning
}

// GetQuotaWarningOk returns a tuple with the QuotaWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetQuotaWarningOk() (*int64, bool) {
	if o == nil || isNil(o.QuotaWarning) {
		return nil, false
	}
	return o.QuotaWarning, true
}

// HasQuotaWarning returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasQuotaWarning() bool {
	if o != nil && !isNil(o.QuotaWarning) {
		return true
	}

	return false
}

// SetQuotaWarning gets a reference to the given int64 and assigns it to the QuotaWarning field.
func (o *CreateDatasetParams) SetQuotaWarning(v int64) {
	o.QuotaWarning = &v
}

// GetVolsize returns the Volsize field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetVolsize() int64 {
	if o == nil || isNil(o.Volsize) {
		var ret int64
		return ret
	}
	return *o.Volsize
}

// GetVolsizeOk returns a tuple with the Volsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetVolsizeOk() (*int64, bool) {
	if o == nil || isNil(o.Volsize) {
		return nil, false
	}
	return o.Volsize, true
}

// HasVolsize returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasVolsize() bool {
	if o != nil && !isNil(o.Volsize) {
		return true
	}

	return false
}

// SetVolsize gets a reference to the given int64 and assigns it to the Volsize field.
func (o *CreateDatasetParams) SetVolsize(v int64) {
	o.Volsize = &v
}

// GetVolblocksize returns the Volblocksize field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetVolblocksize() string {
	if o == nil || isNil(o.Volblocksize) {
		var ret string
		return ret
	}
	return *o.Volblocksize
}

// GetVolblocksizeOk returns a tuple with the Volblocksize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetVolblocksizeOk() (*string, bool) {
	if o == nil || isNil(o.Volblocksize) {
		return nil, false
	}
	return o.Volblocksize, true
}

// HasVolblocksize returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasVolblocksize() bool {
	if o != nil && !isNil(o.Volblocksize) {
		return true
	}

	return false
}

// SetVolblocksize gets a reference to the given string and assigns it to the Volblocksize field.
func (o *CreateDatasetParams) SetVolblocksize(v string) {
	o.Volblocksize = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetReadonly() string {
	if o == nil || isNil(o.Readonly) {
		var ret string
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetReadonlyOk() (*string, bool) {
	if o == nil || isNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasReadonly() bool {
	if o != nil && !isNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given string and assigns it to the Readonly field.
func (o *CreateDatasetParams) SetReadonly(v string) {
	o.Readonly = &v
}

// GetRecordsize returns the Recordsize field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetRecordsize() string {
	if o == nil || isNil(o.Recordsize) {
		var ret string
		return ret
	}
	return *o.Recordsize
}

// GetRecordsizeOk returns a tuple with the Recordsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetRecordsizeOk() (*string, bool) {
	if o == nil || isNil(o.Recordsize) {
		return nil, false
	}
	return o.Recordsize, true
}

// HasRecordsize returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasRecordsize() bool {
	if o != nil && !isNil(o.Recordsize) {
		return true
	}

	return false
}

// SetRecordsize gets a reference to the given string and assigns it to the Recordsize field.
func (o *CreateDatasetParams) SetRecordsize(v string) {
	o.Recordsize = &v
}

// GetRefquota returns the Refquota field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetRefquota() int64 {
	if o == nil || isNil(o.Refquota) {
		var ret int64
		return ret
	}
	return *o.Refquota
}

// GetRefquotaOk returns a tuple with the Refquota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetRefquotaOk() (*int64, bool) {
	if o == nil || isNil(o.Refquota) {
		return nil, false
	}
	return o.Refquota, true
}

// HasRefquota returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasRefquota() bool {
	if o != nil && !isNil(o.Refquota) {
		return true
	}

	return false
}

// SetRefquota gets a reference to the given int64 and assigns it to the Refquota field.
func (o *CreateDatasetParams) SetRefquota(v int64) {
	o.Refquota = &v
}

// GetRefquotaCritical returns the RefquotaCritical field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetRefquotaCritical() int64 {
	if o == nil || isNil(o.RefquotaCritical) {
		var ret int64
		return ret
	}
	return *o.RefquotaCritical
}

// GetRefquotaCriticalOk returns a tuple with the RefquotaCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetRefquotaCriticalOk() (*int64, bool) {
	if o == nil || isNil(o.RefquotaCritical) {
		return nil, false
	}
	return o.RefquotaCritical, true
}

// HasRefquotaCritical returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasRefquotaCritical() bool {
	if o != nil && !isNil(o.RefquotaCritical) {
		return true
	}

	return false
}

// SetRefquotaCritical gets a reference to the given int64 and assigns it to the RefquotaCritical field.
func (o *CreateDatasetParams) SetRefquotaCritical(v int64) {
	o.RefquotaCritical = &v
}

// GetRefquotaWarning returns the RefquotaWarning field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetRefquotaWarning() int64 {
	if o == nil || isNil(o.RefquotaWarning) {
		var ret int64
		return ret
	}
	return *o.RefquotaWarning
}

// GetRefquotaWarningOk returns a tuple with the RefquotaWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetRefquotaWarningOk() (*int64, bool) {
	if o == nil || isNil(o.RefquotaWarning) {
		return nil, false
	}
	return o.RefquotaWarning, true
}

// HasRefquotaWarning returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasRefquotaWarning() bool {
	if o != nil && !isNil(o.RefquotaWarning) {
		return true
	}

	return false
}

// SetRefquotaWarning gets a reference to the given int64 and assigns it to the RefquotaWarning field.
func (o *CreateDatasetParams) SetRefquotaWarning(v int64) {
	o.RefquotaWarning = &v
}

// GetRefreservation returns the Refreservation field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetRefreservation() int64 {
	if o == nil || isNil(o.Refreservation) {
		var ret int64
		return ret
	}
	return *o.Refreservation
}

// GetRefreservationOk returns a tuple with the Refreservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetRefreservationOk() (*int64, bool) {
	if o == nil || isNil(o.Refreservation) {
		return nil, false
	}
	return o.Refreservation, true
}

// HasRefreservation returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasRefreservation() bool {
	if o != nil && !isNil(o.Refreservation) {
		return true
	}

	return false
}

// SetRefreservation gets a reference to the given int64 and assigns it to the Refreservation field.
func (o *CreateDatasetParams) SetRefreservation(v int64) {
	o.Refreservation = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetReservation() int64 {
	if o == nil || isNil(o.Reservation) {
		var ret int64
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetReservationOk() (*int64, bool) {
	if o == nil || isNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasReservation() bool {
	if o != nil && !isNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given int64 and assigns it to the Reservation field.
func (o *CreateDatasetParams) SetReservation(v int64) {
	o.Reservation = &v
}

// GetShareType returns the ShareType field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetShareType() string {
	if o == nil || isNil(o.ShareType) {
		var ret string
		return ret
	}
	return *o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetShareTypeOk() (*string, bool) {
	if o == nil || isNil(o.ShareType) {
		return nil, false
	}
	return o.ShareType, true
}

// HasShareType returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasShareType() bool {
	if o != nil && !isNil(o.ShareType) {
		return true
	}

	return false
}

// SetShareType gets a reference to the given string and assigns it to the ShareType field.
func (o *CreateDatasetParams) SetShareType(v string) {
	o.ShareType = &v
}

// GetSnapdir returns the Snapdir field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetSnapdir() string {
	if o == nil || isNil(o.Snapdir) {
		var ret string
		return ret
	}
	return *o.Snapdir
}

// GetSnapdirOk returns a tuple with the Snapdir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetSnapdirOk() (*string, bool) {
	if o == nil || isNil(o.Snapdir) {
		return nil, false
	}
	return o.Snapdir, true
}

// HasSnapdir returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasSnapdir() bool {
	if o != nil && !isNil(o.Snapdir) {
		return true
	}

	return false
}

// SetSnapdir gets a reference to the given string and assigns it to the Snapdir field.
func (o *CreateDatasetParams) SetSnapdir(v string) {
	o.Snapdir = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetSync() string {
	if o == nil || isNil(o.Sync) {
		var ret string
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetSyncOk() (*string, bool) {
	if o == nil || isNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasSync() bool {
	if o != nil && !isNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given string and assigns it to the Sync field.
func (o *CreateDatasetParams) SetSync(v string) {
	o.Sync = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateDatasetParams) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatasetParams) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateDatasetParams) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateDatasetParams) SetType(v string) {
	o.Type = &v
}

func (o CreateDatasetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Atime) {
		toSerialize["atime"] = o.Atime
	}
	if !isNil(o.Aclmode) {
		toSerialize["aclmode"] = o.Aclmode
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !isNil(o.Compression) {
		toSerialize["compression"] = o.Compression
	}
	if !isNil(o.Casesensitivity) {
		toSerialize["casesensitivity"] = o.Casesensitivity
	}
	if !isNil(o.Copies) {
		toSerialize["copies"] = o.Copies
	}
	if !isNil(o.Deduplication) {
		toSerialize["deduplication"] = o.Deduplication
	}
	if !isNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !isNil(o.EncryptionOptions) {
		toSerialize["encryption_options"] = o.EncryptionOptions
	}
	if !isNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	if !isNil(o.ForceSize) {
		toSerialize["force_size"] = o.ForceSize
	}
	if !isNil(o.InheritEncryption) {
		toSerialize["inherit_encryption"] = o.InheritEncryption
	}
	if !isNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !isNil(o.QuotaCritical) {
		toSerialize["quota_critical"] = o.QuotaCritical
	}
	if !isNil(o.QuotaWarning) {
		toSerialize["quota_warning"] = o.QuotaWarning
	}
	if !isNil(o.Volsize) {
		toSerialize["volsize"] = o.Volsize
	}
	if !isNil(o.Volblocksize) {
		toSerialize["volblocksize"] = o.Volblocksize
	}
	if !isNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !isNil(o.Recordsize) {
		toSerialize["recordsize"] = o.Recordsize
	}
	if !isNil(o.Refquota) {
		toSerialize["refquota"] = o.Refquota
	}
	if !isNil(o.RefquotaCritical) {
		toSerialize["refquota_critical"] = o.RefquotaCritical
	}
	if !isNil(o.RefquotaWarning) {
		toSerialize["refquota_warning"] = o.RefquotaWarning
	}
	if !isNil(o.Refreservation) {
		toSerialize["refreservation"] = o.Refreservation
	}
	if !isNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	if !isNil(o.ShareType) {
		toSerialize["share_type"] = o.ShareType
	}
	if !isNil(o.Snapdir) {
		toSerialize["snapdir"] = o.Snapdir
	}
	if !isNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateDatasetParams) UnmarshalJSON(bytes []byte) (err error) {
	varCreateDatasetParams := _CreateDatasetParams{}

	if err = json.Unmarshal(bytes, &varCreateDatasetParams); err == nil {
		*o = CreateDatasetParams(varCreateDatasetParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "atime")
		delete(additionalProperties, "aclmode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "compression")
		delete(additionalProperties, "casesensitivity")
		delete(additionalProperties, "copies")
		delete(additionalProperties, "deduplication")
		delete(additionalProperties, "encryption")
		delete(additionalProperties, "encryption_options")
		delete(additionalProperties, "exec")
		delete(additionalProperties, "force_size")
		delete(additionalProperties, "inherit_encryption")
		delete(additionalProperties, "quota")
		delete(additionalProperties, "quota_critical")
		delete(additionalProperties, "quota_warning")
		delete(additionalProperties, "volsize")
		delete(additionalProperties, "volblocksize")
		delete(additionalProperties, "readonly")
		delete(additionalProperties, "recordsize")
		delete(additionalProperties, "refquota")
		delete(additionalProperties, "refquota_critical")
		delete(additionalProperties, "refquota_warning")
		delete(additionalProperties, "refreservation")
		delete(additionalProperties, "reservation")
		delete(additionalProperties, "share_type")
		delete(additionalProperties, "snapdir")
		delete(additionalProperties, "sync")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDatasetParams struct {
	value *CreateDatasetParams
	isSet bool
}

func (v NullableCreateDatasetParams) Get() *CreateDatasetParams {
	return v.value
}

func (v *NullableCreateDatasetParams) Set(val *CreateDatasetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatasetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatasetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatasetParams(val *CreateDatasetParams) *NullableCreateDatasetParams {
	return &NullableCreateDatasetParams{value: val, isSet: true}
}

func (v NullableCreateDatasetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatasetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

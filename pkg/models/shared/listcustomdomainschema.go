// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListCustomDomainSchemaOwnershipVerification struct {
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (o *ListCustomDomainSchemaOwnershipVerification) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListCustomDomainSchemaOwnershipVerification) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListCustomDomainSchemaOwnershipVerification) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type ListCustomDomainSchemaSslDcvDelegationRecords struct {
	Cname       *string `json:"cname,omitempty"`
	CnameTarget *string `json:"cname_target,omitempty"`
}

func (o *ListCustomDomainSchemaSslDcvDelegationRecords) GetCname() *string {
	if o == nil {
		return nil
	}
	return o.Cname
}

func (o *ListCustomDomainSchemaSslDcvDelegationRecords) GetCnameTarget() *string {
	if o == nil {
		return nil
	}
	return o.CnameTarget
}

type ListCustomDomainSchemaSslSettings struct {
	MinTLSVersion *string `json:"min_tls_version,omitempty"`
}

func (o *ListCustomDomainSchemaSslSettings) GetMinTLSVersion() *string {
	if o == nil {
		return nil
	}
	return o.MinTLSVersion
}

type ListCustomDomainSchemaSslValidationRecords struct {
	Status   *string `json:"status,omitempty"`
	TxtName  *string `json:"txt_name,omitempty"`
	TxtValue *string `json:"txt_value,omitempty"`
}

func (o *ListCustomDomainSchemaSslValidationRecords) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ListCustomDomainSchemaSslValidationRecords) GetTxtName() *string {
	if o == nil {
		return nil
	}
	return o.TxtName
}

func (o *ListCustomDomainSchemaSslValidationRecords) GetTxtValue() *string {
	if o == nil {
		return nil
	}
	return o.TxtValue
}

type ListCustomDomainSchemaSsl struct {
	BundleMethod         *string                                         `json:"bundle_method,omitempty"`
	CertificateAuthority *string                                         `json:"certificate_authority,omitempty"`
	DcvDelegationRecords []ListCustomDomainSchemaSslDcvDelegationRecords `json:"dcv_delegation_records,omitempty"`
	ID                   *string                                         `json:"id,omitempty"`
	Method               *string                                         `json:"method,omitempty"`
	Settings             *ListCustomDomainSchemaSslSettings              `json:"settings,omitempty"`
	Status               *string                                         `json:"status,omitempty"`
	TxtName              *string                                         `json:"txt_name,omitempty"`
	TxtValue             *string                                         `json:"txt_value,omitempty"`
	Type                 *string                                         `json:"type,omitempty"`
	ValidationRecords    []ListCustomDomainSchemaSslValidationRecords    `json:"validation_records,omitempty"`
	Wildcard             *bool                                           `json:"wildcard,omitempty"`
}

func (o *ListCustomDomainSchemaSsl) GetBundleMethod() *string {
	if o == nil {
		return nil
	}
	return o.BundleMethod
}

func (o *ListCustomDomainSchemaSsl) GetCertificateAuthority() *string {
	if o == nil {
		return nil
	}
	return o.CertificateAuthority
}

func (o *ListCustomDomainSchemaSsl) GetDcvDelegationRecords() []ListCustomDomainSchemaSslDcvDelegationRecords {
	if o == nil {
		return nil
	}
	return o.DcvDelegationRecords
}

func (o *ListCustomDomainSchemaSsl) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListCustomDomainSchemaSsl) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *ListCustomDomainSchemaSsl) GetSettings() *ListCustomDomainSchemaSslSettings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *ListCustomDomainSchemaSsl) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ListCustomDomainSchemaSsl) GetTxtName() *string {
	if o == nil {
		return nil
	}
	return o.TxtName
}

func (o *ListCustomDomainSchemaSsl) GetTxtValue() *string {
	if o == nil {
		return nil
	}
	return o.TxtValue
}

func (o *ListCustomDomainSchemaSsl) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListCustomDomainSchemaSsl) GetValidationRecords() []ListCustomDomainSchemaSslValidationRecords {
	if o == nil {
		return nil
	}
	return o.ValidationRecords
}

func (o *ListCustomDomainSchemaSsl) GetWildcard() *bool {
	if o == nil {
		return nil
	}
	return o.Wildcard
}

type ListCustomDomainSchema struct {
	CreatedAt             *string                                      `json:"created_at,omitempty"`
	Hostname              *string                                      `json:"hostname,omitempty"`
	ID                    *string                                      `json:"id,omitempty"`
	OwnershipVerification *ListCustomDomainSchemaOwnershipVerification `json:"ownership_verification,omitempty"`
	Ssl                   *ListCustomDomainSchemaSsl                   `json:"ssl,omitempty"`
	Status                *string                                      `json:"status,omitempty"`
	VerificationErrors    []string                                     `json:"verification_errors,omitempty"`
}

func (o *ListCustomDomainSchema) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ListCustomDomainSchema) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *ListCustomDomainSchema) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListCustomDomainSchema) GetOwnershipVerification() *ListCustomDomainSchemaOwnershipVerification {
	if o == nil {
		return nil
	}
	return o.OwnershipVerification
}

func (o *ListCustomDomainSchema) GetSsl() *ListCustomDomainSchemaSsl {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *ListCustomDomainSchema) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ListCustomDomainSchema) GetVerificationErrors() []string {
	if o == nil {
		return nil
	}
	return o.VerificationErrors
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListCustomDomainSchemaOwnershipVerification struct {
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ListCustomDomainSchemaSslDcvDelegationRecords struct {
	Cname       *string `json:"cname,omitempty"`
	CnameTarget *string `json:"cname_target,omitempty"`
}

type ListCustomDomainSchemaSslSettings struct {
	MinTLSVersion *string `json:"min_tls_version,omitempty"`
}

type ListCustomDomainSchemaSslValidationRecords struct {
	Status   *string `json:"status,omitempty"`
	TxtName  *string `json:"txt_name,omitempty"`
	TxtValue *string `json:"txt_value,omitempty"`
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

type ListCustomDomainSchema struct {
	CreatedAt             *string                                      `json:"created_at,omitempty"`
	Hostname              *string                                      `json:"hostname,omitempty"`
	ID                    *string                                      `json:"id,omitempty"`
	OwnershipVerification *ListCustomDomainSchemaOwnershipVerification `json:"ownership_verification,omitempty"`
	Ssl                   *ListCustomDomainSchemaSsl                   `json:"ssl,omitempty"`
	Status                *string                                      `json:"status,omitempty"`
	VerificationErrors    []string                                     `json:"verification_errors,omitempty"`
}

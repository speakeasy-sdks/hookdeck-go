// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hookdeck-go/internal/utils"
	"github.com/speakeasy-sdks/hookdeck-go/models/components"
	"net/http"
)

type UpsertIssueTriggerConfigsType string

const (
	UpsertIssueTriggerConfigsTypeIssueTriggerDeliveryConfigs       UpsertIssueTriggerConfigsType = "IssueTriggerDeliveryConfigs"
	UpsertIssueTriggerConfigsTypeIssueTriggerTransformationConfigs UpsertIssueTriggerConfigsType = "IssueTriggerTransformationConfigs"
	UpsertIssueTriggerConfigsTypeIssueTriggerBackpressureConfigs   UpsertIssueTriggerConfigsType = "IssueTriggerBackpressureConfigs"
)

type UpsertIssueTriggerConfigs struct {
	IssueTriggerDeliveryConfigs       *components.IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *components.IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *components.IssueTriggerBackpressureConfigs

	Type UpsertIssueTriggerConfigsType
}

func CreateUpsertIssueTriggerConfigsIssueTriggerDeliveryConfigs(issueTriggerDeliveryConfigs components.IssueTriggerDeliveryConfigs) UpsertIssueTriggerConfigs {
	typ := UpsertIssueTriggerConfigsTypeIssueTriggerDeliveryConfigs

	return UpsertIssueTriggerConfigs{
		IssueTriggerDeliveryConfigs: &issueTriggerDeliveryConfigs,
		Type:                        typ,
	}
}

func CreateUpsertIssueTriggerConfigsIssueTriggerTransformationConfigs(issueTriggerTransformationConfigs components.IssueTriggerTransformationConfigs) UpsertIssueTriggerConfigs {
	typ := UpsertIssueTriggerConfigsTypeIssueTriggerTransformationConfigs

	return UpsertIssueTriggerConfigs{
		IssueTriggerTransformationConfigs: &issueTriggerTransformationConfigs,
		Type:                              typ,
	}
}

func CreateUpsertIssueTriggerConfigsIssueTriggerBackpressureConfigs(issueTriggerBackpressureConfigs components.IssueTriggerBackpressureConfigs) UpsertIssueTriggerConfigs {
	typ := UpsertIssueTriggerConfigsTypeIssueTriggerBackpressureConfigs

	return UpsertIssueTriggerConfigs{
		IssueTriggerBackpressureConfigs: &issueTriggerBackpressureConfigs,
		Type:                            typ,
	}
}

func (u *UpsertIssueTriggerConfigs) UnmarshalJSON(data []byte) error {

	issueTriggerDeliveryConfigs := components.IssueTriggerDeliveryConfigs{}
	if err := utils.UnmarshalJSON(data, &issueTriggerDeliveryConfigs, "", true, true); err == nil {
		u.IssueTriggerDeliveryConfigs = &issueTriggerDeliveryConfigs
		u.Type = UpsertIssueTriggerConfigsTypeIssueTriggerDeliveryConfigs
		return nil
	}

	issueTriggerTransformationConfigs := components.IssueTriggerTransformationConfigs{}
	if err := utils.UnmarshalJSON(data, &issueTriggerTransformationConfigs, "", true, true); err == nil {
		u.IssueTriggerTransformationConfigs = &issueTriggerTransformationConfigs
		u.Type = UpsertIssueTriggerConfigsTypeIssueTriggerTransformationConfigs
		return nil
	}

	issueTriggerBackpressureConfigs := components.IssueTriggerBackpressureConfigs{}
	if err := utils.UnmarshalJSON(data, &issueTriggerBackpressureConfigs, "", true, true); err == nil {
		u.IssueTriggerBackpressureConfigs = &issueTriggerBackpressureConfigs
		u.Type = UpsertIssueTriggerConfigsTypeIssueTriggerBackpressureConfigs
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpsertIssueTriggerConfigs) MarshalJSON() ([]byte, error) {
	if u.IssueTriggerDeliveryConfigs != nil {
		return utils.MarshalJSON(u.IssueTriggerDeliveryConfigs, "", true)
	}

	if u.IssueTriggerTransformationConfigs != nil {
		return utils.MarshalJSON(u.IssueTriggerTransformationConfigs, "", true)
	}

	if u.IssueTriggerBackpressureConfigs != nil {
		return utils.MarshalJSON(u.IssueTriggerBackpressureConfigs, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpsertIssueTriggerRequestBody struct {
	// Notification channels object for the specific channel type
	Channels *components.IssueTriggerChannels `json:"channels"`
	// Configuration object for the specific issue type selected
	Configs *UpsertIssueTriggerConfigs `json:"configs,omitempty"`
	// Required unique name to use as reference when using the API
	Name string `json:"name"`
	// Issue type
	Type components.IssueType `json:"type"`
}

func (o *UpsertIssueTriggerRequestBody) GetChannels() *components.IssueTriggerChannels {
	if o == nil {
		return nil
	}
	return o.Channels
}

func (o *UpsertIssueTriggerRequestBody) GetConfigs() *UpsertIssueTriggerConfigs {
	if o == nil {
		return nil
	}
	return o.Configs
}

func (o *UpsertIssueTriggerRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpsertIssueTriggerRequestBody) GetType() components.IssueType {
	if o == nil {
		return components.IssueType("")
	}
	return o.Type
}

type UpsertIssueTriggerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A single issue trigger
	IssueTrigger *components.IssueTrigger
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpsertIssueTriggerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertIssueTriggerResponse) GetIssueTrigger() *components.IssueTrigger {
	if o == nil {
		return nil
	}
	return o.IssueTrigger
}

func (o *UpsertIssueTriggerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertIssueTriggerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
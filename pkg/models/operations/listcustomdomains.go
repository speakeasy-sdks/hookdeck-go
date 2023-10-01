// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

type ListCustomDomainsRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
}

func (o *ListCustomDomainsRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

type ListCustomDomainsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// List of custom domains
	ListCustomDomainSchema []shared.ListCustomDomainSchema
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListCustomDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListCustomDomainsResponse) GetListCustomDomainSchema() []shared.ListCustomDomainSchema {
	if o == nil {
		return nil
	}
	return o.ListCustomDomainSchema
}

func (o *ListCustomDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListCustomDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDestinationRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteDestinationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteDestination200ApplicationJSON - A single destination
type DeleteDestination200ApplicationJSON struct {
	// ID of the destination
	ID string `json:"id"`
}

func (o *DeleteDestination200ApplicationJSON) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteDestinationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// A single destination
	DeleteDestination200ApplicationJSONObject *DeleteDestination200ApplicationJSON
}

func (o *DeleteDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteDestinationResponse) GetDeleteDestination200ApplicationJSONObject() *DeleteDestination200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteDestination200ApplicationJSONObject
}

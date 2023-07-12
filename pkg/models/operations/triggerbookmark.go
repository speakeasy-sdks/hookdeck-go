// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"net/http"
)

// TriggerBookmarkRequestBodyTarget - Bookmark target
type TriggerBookmarkRequestBodyTarget string

const (
	TriggerBookmarkRequestBodyTargetHTTP TriggerBookmarkRequestBodyTarget = "http"
	TriggerBookmarkRequestBodyTargetCli  TriggerBookmarkRequestBodyTarget = "cli"
)

func (e TriggerBookmarkRequestBodyTarget) ToPointer() *TriggerBookmarkRequestBodyTarget {
	return &e
}

func (e *TriggerBookmarkRequestBodyTarget) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "cli":
		*e = TriggerBookmarkRequestBodyTarget(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TriggerBookmarkRequestBodyTarget: %v", v)
	}
}

type TriggerBookmarkRequestBody struct {
	// Bookmark target
	Target *TriggerBookmarkRequestBodyTarget `json:"target,omitempty"`
}

type TriggerBookmarkRequest struct {
	RequestBody TriggerBookmarkRequestBody `request:"mediaType=application/json"`
	ID          string                     `pathParam:"style=simple,explode=false,name=id"`
}

type TriggerBookmarkResponse struct {
	// Bad Request
	APIErrorResponse *shared.APIErrorResponse
	ContentType      string
	// Array of created events
	EventArray  []shared.Event
	StatusCode  int
	RawResponse *http.Response
}

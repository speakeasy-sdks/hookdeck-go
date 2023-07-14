// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package hookdeck

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/operations"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/models/shared"
	"github.com/speakeasy-sdks/hookdeck-go/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// rulesets - A ruleset defines a group of rules that can be used across many connections.
type rulesets struct {
	sdkConfiguration sdkConfiguration
}

func newRulesets(sdkConfig sdkConfiguration) *rulesets {
	return &rulesets{
		sdkConfiguration: sdkConfig,
	}
}

// Get - Get rulesets
func (s *rulesets) Get(ctx context.Context, request operations.GetRulesetsRequest) (*operations.GetRulesetsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/rulesets"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetRulesetsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RulesetPaginatedResult
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.RulesetPaginatedResult = out
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIErrorResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.APIErrorResponse = out
		}
	}

	return res, nil
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// IntegrationProvider - The provider name
type IntegrationProvider string

const (
	IntegrationProviderTwitter        IntegrationProvider = "twitter"
	IntegrationProviderStripe         IntegrationProvider = "stripe"
	IntegrationProviderRecharge       IntegrationProvider = "recharge"
	IntegrationProviderGithub         IntegrationProvider = "github"
	IntegrationProviderShopify        IntegrationProvider = "shopify"
	IntegrationProviderPostmark       IntegrationProvider = "postmark"
	IntegrationProviderTypeform       IntegrationProvider = "typeform"
	IntegrationProviderHmac           IntegrationProvider = "hmac"
	IntegrationProviderBasicAuth      IntegrationProvider = "basic_auth"
	IntegrationProviderAPIKey         IntegrationProvider = "api_key"
	IntegrationProviderXero           IntegrationProvider = "xero"
	IntegrationProviderSvix           IntegrationProvider = "svix"
	IntegrationProviderZoom           IntegrationProvider = "zoom"
	IntegrationProviderAkeneo         IntegrationProvider = "akeneo"
	IntegrationProviderAdyen          IntegrationProvider = "adyen"
	IntegrationProviderGitlab         IntegrationProvider = "gitlab"
	IntegrationProviderPropertyFinder IntegrationProvider = "property-finder"
	IntegrationProviderWoocommerce    IntegrationProvider = "woocommerce"
	IntegrationProviderOura           IntegrationProvider = "oura"
	IntegrationProviderCommercelayer  IntegrationProvider = "commercelayer"
	IntegrationProviderMailgun        IntegrationProvider = "mailgun"
	IntegrationProviderPipedrive      IntegrationProvider = "pipedrive"
	IntegrationProviderSendgrid       IntegrationProvider = "sendgrid"
	IntegrationProviderWorkos         IntegrationProvider = "workos"
	IntegrationProviderSynctera       IntegrationProvider = "synctera"
	IntegrationProviderAwsSns         IntegrationProvider = "aws_sns"
)

func (e IntegrationProvider) ToPointer() *IntegrationProvider {
	return &e
}

func (e *IntegrationProvider) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "twitter":
		fallthrough
	case "stripe":
		fallthrough
	case "recharge":
		fallthrough
	case "github":
		fallthrough
	case "shopify":
		fallthrough
	case "postmark":
		fallthrough
	case "typeform":
		fallthrough
	case "hmac":
		fallthrough
	case "basic_auth":
		fallthrough
	case "api_key":
		fallthrough
	case "xero":
		fallthrough
	case "svix":
		fallthrough
	case "zoom":
		fallthrough
	case "akeneo":
		fallthrough
	case "adyen":
		fallthrough
	case "gitlab":
		fallthrough
	case "property-finder":
		fallthrough
	case "woocommerce":
		fallthrough
	case "oura":
		fallthrough
	case "commercelayer":
		fallthrough
	case "mailgun":
		fallthrough
	case "pipedrive":
		fallthrough
	case "sendgrid":
		fallthrough
	case "workos":
		fallthrough
	case "synctera":
		fallthrough
	case "aws_sns":
		*e = IntegrationProvider(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IntegrationProvider: %v", v)
	}
}

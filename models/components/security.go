// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Security struct {
	BasicAuth  *SchemeBasicAuth `security:"scheme,type=http,subtype=basic"`
	BearerAuth *string          `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *Security) GetBasicAuth() *SchemeBasicAuth {
	if o == nil {
		return nil
	}
	return o.BasicAuth
}

func (o *Security) GetBearerAuth() *string {
	if o == nil {
		return nil
	}
	return o.BearerAuth
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type DeleteCustomDomainSchema struct {
	// The custom hostname ID
	ID string `json:"id"`
}

func (o *DeleteCustomDomainSchema) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
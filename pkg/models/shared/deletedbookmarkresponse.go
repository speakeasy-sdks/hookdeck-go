// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeletedBookmarkResponse struct {
	// Bookmark ID
	ID string `json:"id"`
}

func (o *DeletedBookmarkResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransformationFailedMeta struct {
	TransformationID string `json:"transformation_id"`
}

func (o *TransformationFailedMeta) GetTransformationID() string {
	if o == nil {
		return ""
	}
	return o.TransformationID
}

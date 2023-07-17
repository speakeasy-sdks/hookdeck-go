// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IssueTriggerSlackChannel - Slack channel for an issue trigger
type IssueTriggerSlackChannel struct {
	// Channel name
	ChannelName string `json:"channel_name"`
}

func (o *IssueTriggerSlackChannel) GetChannelName() string {
	if o == nil {
		return ""
	}
	return o.ChannelName
}

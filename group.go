package analytics

import "time"

var _ Message = (*Group)(nil)

type Group struct {
	// This field is exported for serialization purposes and shouldn't be set by
	// the application, its value is always overwritten by the library.
	Type string `json:"type,omitempty"`

	MessageId    string       `json:"messageId,omitempty"`
	AnonymousId  string       `json:"anonymousId,omitempty"`
	UserId       string       `json:"userId,omitempty"`
	GroupId      string       `json:"groupId"`
	Timestamp    time.Time    `json:"timestamp,omitempty"`
	Context      *Context     `json:"context,omitempty"`
	Traits       Traits       `json:"traits,omitempty"`
	Integrations Integrations `json:"integrations,omitempty"`
}

func (msg Group) Validate() error {
	if len(msg.GroupId) == 0 {
		return FieldError{
			Type:  "analytics.Group",
			Name:  "GroupId",
			Value: msg.GroupId,
		}
	}

	return nil
}

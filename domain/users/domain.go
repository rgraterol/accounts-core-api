package users

type UsersFeedMessage struct {
	ID                string  `json:"id"`
	UID               string  `json:"uid"`
	Topic             string  `json:"topic"`
	Cluster           string  `json:"cluster"`
	Consumer          string  `json:"consumer"`
	PublishTime       int64   `json:"publish_time"`
	RecipientCallback string  `json:"recipientCallback"`
	Msg               UserMsg `json:"msg"`
}

type UserMsg struct {
	ID        int64          `json:"id"`
	CountryID string         `json:"country_id"`
	Headers   UserMsgHeaders `json:"headers"`
}

type UserMsgHeaders struct {
	NewUser bool `json:"new_user,omitempty"`
}

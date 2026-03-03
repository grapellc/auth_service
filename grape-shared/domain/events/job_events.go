package events

type JobViewedEvent struct {
	JobID     uint    `json:"job_id"`
	UserID    *uint   `json:"user_id,omitempty"`
	IPAddress *string `json:"ip_address,omitempty"`
}

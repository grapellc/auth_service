package constants

type JobStatus string

const (
	JobStatusActive        JobStatus = "active"
	JobStatusInactive      JobStatus = "inactive"
	JobStatusFoundEmployee JobStatus = "found_employee"
)

type JobType string

const (
	JobTypePartTime JobType = "part_time"
)

const (
	JobIndexName = "jobs"
)

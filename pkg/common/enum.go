package common

type Role string

const (
	RoleOld       Role = "Old"
	RoleVolunteer Role = "Volunteer"
	RoleAdmin     Role = "Admin"
)

type Status string

const (
	StatusDraft     Status = "Draft"
	StatusPublished Status = "Published"
	StatusCensoring Status = "Censoring"
	StatusPassed    Status = "Passed"
)

type Action string

const (
	ActionCreate  Action = "Create"
	ActionEdit    Action = "Edit"
	ActionDelete  Action = "Delete"
	ActionRespond Action = "Respond"
)

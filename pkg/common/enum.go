package common

type Role string

const (
	RoleOld       Role = "Old"
	RoleVolunteer Role = "Volunteer"
	RoleAdmin     Role = "Admin"
)

type Status string

const (
	Draft     Status = "Draft"
	Published Status = "Published"
	Closed    Status = "Closed"
)

type Action string

const (
	Create  Action = "Create"
	Edit    Action = "Edit"
	Delete  Action = "Delete"
	Respond Action = "Respond"
)

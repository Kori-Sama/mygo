package utils

import "mygo/internal/pkg/common"

func CheckRole(role common.Role) bool {
	switch role {
	case common.RoleOld, common.RoleVolunteer, common.RoleAdmin:
		return true
	}
	return false
}

func FilterStatus(status string) (common.Status, error) {
	switch status {
	case "Draft":
		return common.StatusDraft, nil
	case "Censoring":
		return common.StatusCensoring, nil
	case "Passed":
		return common.StatusPassed, nil
	case "Rejected":
		return common.StatusRejected, nil
	default:
		return "", common.ErrorInvalidParam
	}
}

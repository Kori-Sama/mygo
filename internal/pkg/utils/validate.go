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
	case "draft":
		return common.StatusDraft, nil
	case "censoring":
		return common.StatusCensoring, nil
	case "passed":
		return common.StatusPassed, nil
	case "rejected":
		return common.StatusRejected, nil
	default:
		return "", common.ErrorInvalidParam
	}
}

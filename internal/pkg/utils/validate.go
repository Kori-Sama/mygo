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

func FilterAction(action string) (common.Action, error) {
	switch action {
	case "Create":
		return common.ActionCreate, nil
	case "Edit":
		return common.ActionEdit, nil
	case "Delete":
		return common.ActionDelete, nil
	case "Respond":
		return common.ActionRespond, nil
	case "Save":
		return common.ActionSave, nil
	case "Censor":
		return common.ActionCensor, nil
	default:
		return "", common.ErrorInvalidParam
	}
}

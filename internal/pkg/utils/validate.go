package utils

import "mygo/internal/pkg/common"

func CheckRole(role common.Role) bool {
	switch role {
	case common.RoleOld, common.RoleVolunteer, common.RoleAdmin:
		return true
	}
	return false
}

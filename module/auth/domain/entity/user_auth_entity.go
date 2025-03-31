package entity

import "asset_management_backend/common/enums"

type UserAuthEntity struct {
	ID            int   `json:"id"`
	RoleID        int   `json:"role_id"`
	PermissionIDs []int `json:"permission_ids"`
}

func (u *UserAuthEntity) HasAuthority(userAuthority string) bool {
	if userAuthority == enums.UserAuthority.Admin {
		return u.HasAdminAuthority()
	}
	if userAuthority == enums.UserAuthority.AssetManagement {
		return u.HasAssetManagementAuthority()
	}
	if userAuthority == enums.UserAuthority.BorrowManagement {
		return u.HasBorrowManagementAuthority()
	}
	if userAuthority == enums.UserAuthority.Statistical {
		return u.HasStatisticalAuthority()
	}

	if userAuthority == enums.UserAuthority.CategoryManagement {
		return u.HasCategoryManagementAuthority()
	}
	return u.HasEmployeeAuthority()
}

func (u *UserAuthEntity) HasAdminAuthority() bool {
	return u.RoleID == enums.UserRoleID.Admin
}

func (u *UserAuthEntity) hasManagerAuthority(permissionID int) bool {
	if u.HasAdminAuthority() {
		return true
	}
	if u.RoleID == enums.UserRoleID.Manager {
		for _, permID := range u.PermissionIDs {
			if permID == permissionID {
				return true
			}
		}
	}
	return false
}

func (u *UserAuthEntity) HasAssetManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.AssetManagement)

}

func (u *UserAuthEntity) HasBorrowManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.BorrowManagement)
}

func (u *UserAuthEntity) HasStatisticalAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.Statistical)
}

func (u *UserAuthEntity) HasCategoryManagementAuthority() bool {
	return u.hasManagerAuthority(enums.PermissionID.CategoryManagement)
}

func (u *UserAuthEntity) HasEmployeeAuthority() bool {
	return true
}

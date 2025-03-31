package model

import "asset_management_backend/module/auth/domain/entity"

type UserAuthModel struct {
	ID            int      `json:"id" bun:"id"`
	RoleID        int      `json:"role_id" bun:"role_id"`
	PermissionIDs []int    `json:"permission_ids" bun:"permission_ids,type:int[]"` 
}


func (u *UserAuthModel) ToEntity() *entity.UserAuthEntity {
	if u == nil {
		return nil
	}
	return &entity.UserAuthEntity{
		ID:            u.ID,
		RoleID:        u.RoleID,
		PermissionIDs: u.PermissionIDs,
	}
}

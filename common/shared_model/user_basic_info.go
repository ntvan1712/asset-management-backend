package sharedmodel

import "github.com/uptrace/bun"


type UserBasicInfo struct {
	bun.BaseModel `bun:"table:users"`

	Name string
	Code string
}
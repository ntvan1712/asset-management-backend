package sharedmodel

import "github.com/uptrace/bun"

type UserBasicInfo struct {
	bun.BaseModel `bun:"table:users"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
}

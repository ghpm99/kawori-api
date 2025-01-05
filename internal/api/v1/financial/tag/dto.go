package tag

import "kawori/api/pkg/utils"

type TagFilter struct {
	Name   string `json:"name"`
	UserId int    `json:"-"`
}

type GetTagReturn struct {
	data     []Tag
	pageInfo utils.Pagination
}

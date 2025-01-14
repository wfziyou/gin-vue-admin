package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID uint64 `json:"id" form:"id"` // 主键ID
}

// IdSearch id查询
type IdSearch struct {
	ID uint64 `json:"id" form:"id" ` //编号
}

// IdDelete id删除
type IdDelete struct {
	ID uint64 `json:"id" form:"id" ` //编号
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []uint64 `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint64 `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

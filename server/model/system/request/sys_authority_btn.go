package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuID"`
	AuthorityId uint64 `json:"authorityId"`
	Selected    []uint `json:"selected"`
}

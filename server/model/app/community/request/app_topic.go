package request

// CreateForumTopicReq 结构体
type CreateForumTopicReq struct {
	Name string `json:"name" form:"name"` //名称
}

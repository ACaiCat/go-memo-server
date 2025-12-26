package handler

// BaseResp 基础响应结构
type BaseResp struct {
	// Status 状态码
	Status int `json:"status" example:"200"`
	// Msg 消息描述
	Msg string `json:"msg" example:"success"`
}

// Pagination 分页信息
type Pagination struct {
	// Page 当前页码
	Page int `json:"page" example:"1"`
	// PerPage 每页数量
	PerPage int `json:"per_page" example:"20"`
	// Total 总记录数
	Total int `json:"total" example:"100"`
}

func NewPagination(page int, perPage int, total int) *Pagination {
	return &Pagination{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}
}

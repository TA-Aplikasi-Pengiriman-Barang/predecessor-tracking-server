package dto

import "time"

type (
	News struct {
		ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
		Title     string    `gorm:"column:title" json:"title"`
		Detail    string    `gorm:"column:detail" json:"detail"`
		CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	}

	NewsSlice []News

	// CreateNewsDto CreateNewsDto
	CreateNewsDto struct {
		Title  string `json:"title" validate:"required"`
		Detail string `json:"detail" validate:"required"`
	}

	// CreateNewsResponse CreateNewsResponse
	CreateNewsResponse struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Detail    string `json:"detail"`
		CreatedAt string `json:"createdAt"`
	}

	// GetAllNewsResponse GetAllNewsResponse
	GetAllNewsResponse struct {
		News []News `json:"news"`
	}

	// EditNewsDto EditNewsDto
	EditNewsDto struct {
		Title  string `json:"title" validate:"omitempty"`
		Detail string `json:"detail" validate:"omitempty"`
	}
)

func (n *News) ToCreateNewsResponse() CreateNewsResponse {
	return CreateNewsResponse{
		ID:        n.ID,
		Title:     n.Title,
		Detail:    n.Detail,
		CreatedAt: n.CreatedAt.String(),
	}
}

func (n *NewsSlice) ToGetAllNewsResponse() GetAllNewsResponse {
	return GetAllNewsResponse{
		News: *n,
	}
}

func (n *News) FillNewsEdit(data EditNewsDto) {
	if data.Title != "" {
		n.Title = data.Title
	}

	if data.Detail != "" {
		n.Detail = data.Detail
	}
}

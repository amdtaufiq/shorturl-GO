package dto

type CreateGeratorRequest struct {
	LongURL string `json:"long_url" binding:"required"`
}

type CreateGeratorResponse struct {
	Message string `json:"message"`
	SortURL string `json:"sort_url"`
}

type GetSortURLRequest struct {
	SortURL string `uri:"url" binding:"required"`
}

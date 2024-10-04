package dto

import "github.com/pankop/event-porto/entity"

type (
	ShortenLinkCreateRequest struct {
		Custom_Link   string `json:"custom_link" form:"custom_link" binding:"required" `
		Original_link string `json:"original_link" form:"original_link" binding:"required" `
	}

	ShortenLinkCreateResponse struct {
		Link string `json:"link" `
	}

	ShortenLinkUpdateRequest struct {
		Custom_Link   string `json:"custom_link" form:"custom_link" binding:"required" `
		Original_link string `json:"original_link" form:"original_link" binding:"required" `
	}

	ShortenLinkUpdateResponse struct {
		Link string `json:"link" `
	}

	ShortenLinkPaginationResponse struct {
		Data []ShortenLinkCreateResponse `json:"data"`
		PaginationResponse
	}

	GetAllShortenLinkRepositoryResponse struct {
		ShortenLink []entity.ShortenLink
		PaginationResponse
	}
)

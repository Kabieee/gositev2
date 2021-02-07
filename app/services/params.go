package services

type SearchParams struct {
	Page int `form:"p" binding:"gte=1"`
	Search string `form:"s"`
}
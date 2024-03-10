package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Offset  int
	Size    int
	Keyword string
}

func NewPagination(ctx *gin.Context) *Pagination {
	size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	keyword := ctx.DefaultQuery("keyword", "")
	return &Pagination{
		Size:    size,
		Offset:  offset,
		Keyword: keyword,
	}
}

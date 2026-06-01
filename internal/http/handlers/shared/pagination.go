package shared

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParsePagination 从 query 中读取默认分页参数并归一化。
func ParsePagination(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	return NormalizePagination(page, pageSize)
}

// NormalizePagination 归一化分页参数。
func NormalizePagination(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 200 {
		pageSize = 200
	}
	return page, pageSize
}

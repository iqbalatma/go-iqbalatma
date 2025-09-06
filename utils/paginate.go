package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate[T any](c *gin.Context, db *gorm.DB, out *[]T) (*PaginationMeta, error) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, _ := strconv.Atoi(pageStr)
	perPage, _ := strconv.Atoi(perPageStr)
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	var total int64
	db.Model(out).Count(&total)

	offset := (page - 1) * perPage
	fmt.Println("total", total)
	fmt.Println("offset", offset)
	if err := db.Offset(offset).Limit(perPage).Find(out).Error; err != nil {
		return nil, err
	}

	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	from := offset + 1
	to := offset + len(*out)
	if total == 0 {
		from = 0
		to = 0
	}

	fmt.Println("total", total)
	fmt.Print("page: ", page, " perPage: ", perPage, " \n")

	return &PaginationMeta{
		CurrentPage: page,
		PerPage:     perPage,
		From:        from,
		To:          to,
		LastPage:    lastPage,
		Path:        "",
		Total:       total,
	}, nil
}

package uolo_lens

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	LensHandler struct {
	}
)

func (handler *LensHandler) LatestLens(ec echo.Context) error {

	// 从redis中获取当前用户聚合推荐的内容

	// 由聚合数据取出具体的内容组装成内容列表

	return ec.JSON(http.StatusOK, nil)
}

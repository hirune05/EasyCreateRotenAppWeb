package storeStaff

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	StudentNumber int    `json:"student_number"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *storeStaffHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	// ユースケースを呼び出してログイン処理を実行
	dto, err := h.storeStaffUseCase.Login(ctx, req.StudentNumber, req.Password)
	if err != nil {
		log.Printf("ログインエラー: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	// ログイン成功時にJWTトークンを返す
	return c.JSON(http.StatusOK, LoginResponse{Token: dto.Token})
}

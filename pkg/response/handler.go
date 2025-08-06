package response

import "github.com/gofiber/fiber/v2"

type (
	APIResponse[T any] struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data,omitempty"`
		Meta    any    `json:"meta,omitempty"`
	}

	APIResponseOptions[T any] struct {
		Message string
		Data    T
		Meta    any
	}

	PaginationMeta struct {
		Page       int   `json:"page"`
		PageSize   int   `json:"page_size"`
		TotalRows  int64 `json:"total_rows"`
		TotalPages int   `json:"total_pages"`
	}
)

func MakeAPIResponse[T any](ctx *fiber.Ctx, res APIResponse[T]) error {
	return ctx.Status(res.Code).JSON(res)
}

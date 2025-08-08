package handler

import (
	"github.com/bagusyanuar/app-hr-be/internal/config"
	"github.com/bagusyanuar/app-hr-be/internal/schema"
	"github.com/bagusyanuar/app-hr-be/internal/service"
	"github.com/bagusyanuar/app-hr-be/pkg/response"
	"github.com/bagusyanuar/app-hr-be/pkg/util"
	"github.com/gofiber/fiber/v2"
)

type BranchHandler struct {
	BranchService service.BranchService
	Config        *config.AppConfig
}

func NewBranchHandler(
	branchService service.BranchService,
	cfg *config.AppConfig,
) *BranchHandler {
	return &BranchHandler{
		BranchService: branchService,
		Config:        cfg,
	}
}

func (c *BranchHandler) Create(ctx *fiber.Ctx) error {
	request := new(schema.BranchSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: err.Error(),
			Code:    fiber.StatusBadRequest,
		})
	}

	messages, err := util.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: fiber.ErrUnprocessableEntity.Error(),
			Code:    fiber.StatusUnprocessableEntity,
			Data:    messages,
		})
	}

	_, err = c.BranchService.Create(ctx.UserContext(), request)
	if err != nil {
		return response.MakeAPIResponse(ctx, response.APIResponse[any]{
			Message: err.Error(),
			Code:    fiber.StatusInternalServerError,
		})
	}
	return response.MakeAPIResponse(ctx, response.APIResponse[any]{
		Message: "successfully create new branch",
		Code:    fiber.StatusOK,
	})
}

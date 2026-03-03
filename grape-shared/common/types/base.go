package types

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type SearchResult struct {
	ID int `json:"id"`
}

type Pagination struct {
	Page          int    `query:"page"  default:"1"`
	Limit         int    `query:"limit" default:"20"`
	Query         string `query:"query" default:""`
	LocationID    uint   `query:"location_id"`
	Location      uint   `query:"location"`
	MarketTypeID  uint   `query:"market_type_id"`
	MarketTypeIDs []uint `query:"market_type_id"`
	Ids           []uint `query:"ids[]"`
	IsFavorite    bool   `query:"is_favorite"`
}

const (
	DEFAULT_PAGE   = 1
	DEFAULT_LIMIT  = 20
	MAX_LIMIT      = 100
	DEFAULT_SORTER = "created_at DESC"
)

func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		return DEFAULT_LIMIT
	}
	if p.Limit > MAX_LIMIT {
		return MAX_LIMIT
	}
	return p.Limit
}

func (p *Pagination) GetOffset() int {
	page := p.GetPage()
	limit := p.GetLimit()
	return (page - 1) * limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		return DEFAULT_PAGE
	}
	return p.Page
}

func (p *Pagination) GetLocationID() uint {
	if p.LocationID != 0 {
		return p.LocationID
	}
	return p.Location
}

func (p *Pagination) GetMarketTypeIDs() []uint {
	if len(p.MarketTypeIDs) > 0 {
		return p.MarketTypeIDs
	}
	if p.MarketTypeID != 0 {
		return []uint{p.MarketTypeID}
	}
	return nil
}

type PaginationData struct {
	Total   int  `json:"total"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
	HasMore bool `json:"has_more"`
}

func PaginationDataResp(total int, limit int, page int) PaginationData {
	return PaginationData{
		Total:   total,
		Page:    page,
		Limit:   limit,
		HasMore: total > limit*page,
	}
}

type PaginationResponse struct {
	Data           interface{} `json:"data"`
	PaginationData `json:"pagination"`
}

type Meta struct {
	Pagination *PaginationData `json:"pagination,omitempty"`
}

type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type BaseController struct{}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) SendPaginationResponse(ctx *fiber.Ctx, data interface{}, paginationData PaginationData) error {
	return ctx.Status(fiber.StatusOK).JSON(APIResponse{
		Data: data,
		Meta: &Meta{
			Pagination: &paginationData,
		}})
}

// SendResponse supports both the new signature (ctx, data, message)
// and legacy calls that pass only (ctx, data).
func (c *BaseController) SendResponse(ctx *fiber.Ctx, data interface{}, message ...string) error {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}

	// Handle PaginationResponse by value or pointer
	switch v := data.(type) {
	case PaginationResponse:
		return ctx.Status(fiber.StatusOK).JSON(APIResponse{
			Data: v.Data,
			Meta: &Meta{
				Pagination: &v.PaginationData,
			},
		})
	case *PaginationResponse:
		if v == nil {
			return ctx.Status(fiber.StatusOK).JSON(APIResponse{Data: nil, Message: msg})
		}
		return ctx.Status(fiber.StatusOK).JSON(APIResponse{
			Data: v.Data,
			Meta: &Meta{
				Pagination: &v.PaginationData,
			},
		})
	default:
		return ctx.Status(fiber.StatusOK).JSON(APIResponse{Data: data, Message: msg})
	}
}

func (c *BaseController) SendUnauthorized(ctx *fiber.Ctx, message ...string) error {
	msg := "Unauthorized"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(APIResponse{Data: nil, Message: msg})
}

// SendError supports flexible calling styles to remain backward compatible with handlers:
// - SendError(ctx, err)
// - SendError(ctx, *fiber.Error)
// - SendError(ctx, httpStatus int, err error)
// - SendError(ctx, httpStatus int, err error, code int)
func (c *BaseController) SendError(ctx *fiber.Ctx, args ...interface{}) error {
	var httpStatus = fiber.StatusBadRequest
	var err error
	var codeError int

	// No args -> unknown error
	if len(args) == 0 {
		err = fmt.Errorf("unknown error")
	} else if len(args) == 1 {
		// Single arg could be error or *fiber.Error or int(status)
		switch v := args[0].(type) {
		case error:
			err = v
			if fe, ok := v.(*fiber.Error); ok {
				httpStatus = fe.Code
			}
		case *fiber.Error:
			err = v
			httpStatus = v.Code
		case int:
			httpStatus = v
			err = errors.New(http.StatusText(v))
		default:
			err = fmt.Errorf("%v", v)
		}
	} else {
		// len >= 2
		// First arg may be httpStatus (int) or *fiber.Error
		if hs, ok := args[0].(int); ok {
			httpStatus = hs
		} else if fe, ok := args[0].(*fiber.Error); ok {
			httpStatus = fe.Code
			// if second provided and is error set err
			if len(args) >= 2 {
				if e, ok := args[1].(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%v", args[1])
				}
			} else {
				err = fe
			}
		}

		// If err not set yet, try to take second arg as error
		if err == nil && len(args) >= 2 {
			if e, ok := args[1].(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", args[1])
			}
		}

		// Optional code (third arg)
		if len(args) >= 3 {
			if ci, ok := args[2].(int); ok {
				codeError = ci
			}
		}
	}

	if err == nil {
		err = fmt.Errorf("unknown error")
	}

	// Special handling for not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(httpStatus).JSON(APIResponse{Data: nil, Message: "Уучлаарай, Олдсонгүй", Code: codeError})
	}

	// In production hide details
	if viper.GetString("app.env") == "production" {
		logrus.Error(err.Error())
		return ctx.Status(httpStatus).JSON(APIResponse{Data: nil, Message: "Алдаа гарлаа", Code: codeError})
	}

	// Development: return error message
	return ctx.Status(httpStatus).JSON(APIResponse{Data: nil, Message: err.Error(), Code: codeError})
}

// SendHTML returns an HTML response using standardized typing and status.
func (c *BaseController) SendHTML(ctx *fiber.Ctx, html string, status ...int) error {
	st := fiber.StatusOK
	if len(status) > 0 {
		st = status[0]
	}
	return ctx.Status(st).Type("html").SendString(html)
}

// LimiterReachedHandler returns a fiber.Handler suitable for limiter callbacks that
// want to use the standardized API response shape.
func (c *BaseController) LimiterReachedHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return c.SendError(ctx, fiber.StatusTooManyRequests, errors.New("Too many requests, please try again later."))
	}
}

// MiddlewareError is a small helper that middleware can call without having a controller
// instance to produce a standardized APIResponse error.
func MiddlewareError(ctx *fiber.Ctx, httpStatus int, message string) error {
	return (&BaseController{}).SendError(ctx, httpStatus, errors.New(message))
}

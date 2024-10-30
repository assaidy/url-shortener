package handlers

import (
	s "strings"
	"time"

	"github.com/assaidy/url-shortener/database"
	"github.com/assaidy/url-shortener/models"
	"github.com/assaidy/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	db *database.DBService
}

func NewURLHandler(db *database.DBService) *URLHandler {
	return &URLHandler{db: db}
}

func (h *URLHandler) HandleCreateURL(c *fiber.Ctx) error {
	req := models.URLCreateOrUpdateReq{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := utils.Validator.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	url := models.URL{
		OriginalURL: s.ToLower(s.TrimSpace(req.OriginalURL)),
		ShortCode:   utils.GenerateShortCode(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := h.db.InsertURL(&url); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(url)
}

func (h *URLHandler) HandleGetURL(c *fiber.Ctx) error {
	sc := c.Params("sc")
	if sc == "" {
		return fiber.ErrBadRequest
	}

	url, err := h.db.GetURL(sc)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if url == nil {
		return fiber.ErrNotFound
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func (h *URLHandler) HandleUpdateURL(c *fiber.Ctx) error {
	req := models.URLCreateOrUpdateReq{}
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if err := utils.Validator.Struct(req); err != nil {
		return fiber.ErrBadRequest
	}

	sc := c.Params("sc")
	if sc == "" {
		return fiber.ErrBadRequest
	}

	url, err := h.db.GetURL(sc)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if url == nil {
		return fiber.ErrNotFound
	}

	url.OriginalURL = s.ToLower(s.TrimSpace(req.OriginalURL))
	url.UpdatedAt = time.Now().UTC()

	if err := h.db.UpdateURL(url); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

func (h *URLHandler) HandleDeleteURL(c *fiber.Ctx) error {
	sc := c.Params("sc")
	if sc == "" {
		return fiber.ErrBadRequest
	}

	if ok, err := h.db.CheckIfURLExists(sc); err != nil {
		return fiber.ErrInternalServerError
	} else if !ok {
		return fiber.ErrNotFound
	}

	if err := h.db.DeleteURL(sc); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *URLHandler) HandleGetURLWithStats(c *fiber.Ctx) error {
	sc := c.Params("sc")
	if sc == "" {
		return fiber.ErrBadRequest
	}

	url, err := h.db.GetURLWithAccessCount(sc)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if url == nil {
		return fiber.ErrNotFound
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

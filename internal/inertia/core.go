package inertia

import (
	"embed"
	"encoding/json"
	"gravel/internal/env"
	"html/template"

	"github.com/gofiber/fiber/v3"
)

//go:embed index.go.html
var FS embed.FS

const (
	Header         = "X-Inertia"
	HeaderLocation = "X-Inertia-Location"
)

type inertia struct{ Inertia template.JS }

func (inertia) AppName() string { return env.Get("APP_NAME", "Application") }

func Render(c fiber.Ctx, name string, props fiber.Map) error {
	if props == nil {
		props = fiber.Map{}
	}

	page := fiber.Map{
		"component": name,
		"props":     props,
		"url":       c.OriginalURL(),
	}

	if c.Get(Header) == "true" {
		c.Set(Header, "true")
		return c.JSON(page)
	}

	pageBytes, err := json.Marshal(page)
	if err != nil {
		return err
	}

	return c.Render("index", inertia{
		Inertia: template.JS(pageBytes),
	})
}

func Location(c fiber.Ctx, url string) error {
	c.Set(HeaderLocation, url)
	return c.SendStatus(fiber.StatusConflict)
}

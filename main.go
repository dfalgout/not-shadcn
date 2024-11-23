package main

import (
	"embed"
	"net/http"

	"github.com/a-h/templ"
	"github.com/dfalgout/not-shadcn/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets/*
var assets embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Root:       "/",
		Filesystem: http.FS(assets),
	}))

	e.GET("/", func(c echo.Context) error {
		return Render(c, 200, components.Home())
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

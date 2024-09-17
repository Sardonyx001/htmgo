// Package pages THIS FILE IS GENERATED. DO NOT EDIT.
package pages

import "github.com/labstack/echo/v4"
import "github.com/maddalax/htmgo/framework/h"

func RegisterPages(f *echo.Echo) {
	f.GET("/", func(ctx echo.Context) error {
		return h.HtmlView(ctx, IndexPage(ctx))
	})
}

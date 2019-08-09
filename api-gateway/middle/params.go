package middleware

import "github.com/labstack/echo"

func QueryParamsCheck() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			language := ctx.Request().Header.Get("Accept-Language")
			if language == "" {
				language = "en-US"
			}
			ctx.Set("language", language)
			return next(ctx)
		}
	}
}

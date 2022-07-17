package echo

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*

skipper リクエストurlの除外定義

*/
func skipper(c echo.Context) bool {
	if strings.HasPrefix(c.Request().URL.Path, "/index.html") {
		return true
	}
	if strings.HasPrefix(c.Request().URL.Path, "/resources/") {
		return true
	}
	if strings.HasPrefix(c.Request().URL.Path, "/login") && c.Request().Method == "POST" {
		return true
	}
	if strings.HasPrefix(c.Request().URL.Path, "/api") {
		return true
	}
	return false
}

/*

defineRewrite リクエストurlのリライト定義
frontendでルーティング処理を行うのでサーバーサイドはindex.htmlで受ける。
*/
func defineRewrite(e *echo.Echo) {
	config := middleware.RewriteConfig{
		Skipper: skipper,
		Rules: map[string]string{
			"/*": "/index.html",
		},
	}

	e.Pre(middleware.RewriteWithConfig(config))
}

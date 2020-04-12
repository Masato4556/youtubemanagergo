package routes

import (
	"github.com/Masato4556/youtubemanagergo/web/api"
	"github.com/labstack/echo"
)

func Init(e echo.Echo) {
	g := e.Groups("/api"){
		g.Get("/popular", api.FetchPopularVideos())
	}
}
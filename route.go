package snowflakeid

import (
	"github.com/labstack/echo"
)

type routeItem struct {
	url     string
	method  int
	handler func(echo.Context) error
}

func getRouteList() []*routeItem {
	return []*routeItem{
		&routeItem{
			url:     "/get/:num",
			handler: getUniqueId,
		},
	}
}

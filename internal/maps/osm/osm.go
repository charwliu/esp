/*

Package osm encapsulates the OpenStreetMap API.
*/
package osm

import (
	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/event"
)


func S() *zap.SugaredLogger  {
	return event.S()
}

func L() *zap.Logger  {
	return event.L()
}

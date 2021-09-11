package osm

import (
	"time"

	"github.com/melihmucuk/geocache"
	"go.uber.org/zap"
)

var geoCache *geocache.Cache

func init() {
	c, err := geocache.NewCache(time.Hour, 5*time.Minute, geocache.WithIn1M)

	if err != nil {
		L().Panic("osm ", zap.Error(err))
	}

	geoCache = c
}

package hub

import (
	"sync"

	"go.vixal.xyz/esp/internal/event"
)

var log = event.Log.Sugar()

var mutex = sync.Mutex{}

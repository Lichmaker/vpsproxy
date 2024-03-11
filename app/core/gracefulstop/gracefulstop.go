package gracefulstop

import (
	"fmt"
	"sync"

	"github.com/Lichmaker/vpsproxy/pkg/logger"
)

var registerStopInfo []Info
var locker sync.Mutex

type Info struct {
	Name   string
	StopFn func()
}

func init() {
	registerStopInfo = make([]Info, 0)
}

func RegisterStop(name string, stopFunc func()) {
	locker.Lock()
	defer locker.Unlock()
	registerStopInfo = append(registerStopInfo, Info{
		Name:   name,
		StopFn: stopFunc,
	})
}

func Stop() {
	locker.Lock()
	defer locker.Unlock()
	for _, info := range registerStopInfo {
		logger.Logger.Warn(fmt.Sprintf("stopping %s ...", info.Name))
		info.StopFn()
	}
}

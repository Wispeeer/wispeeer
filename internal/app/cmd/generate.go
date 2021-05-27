package cmd

import (
	"sync/atomic"
	"time"

	"github.com/ka1i/wispeeer/internal/pkg/config"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/logeer"
)

var (
	articleTotal   *uint64 = new(uint64)
	wispeeerConfig config.Config
)

// Generate ...
func Generate() error {
	var err error
	defer utils.Timer("wispeeer ", time.Now())

	// article count
	atomic.StoreUint64(articleTotal, 0)

	wispeeerConfig, err = config.GetWispeeerConfig()
	if err != nil {
		return err
	}
	logeer.WispeeerLog("generate").Infof("Location : %v", utils.GetWorkspace())
	logeer.WispeeerLog("generate").Infof("public in: %v", wispeeerConfig.PublicDir)

	// time.Sleep(time.Second)
	return nil
}

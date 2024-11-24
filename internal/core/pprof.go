package core

import (
	"net/http"
	_ "net/http/pprof"
)

type PprofCfg struct {
	Addr string
}

func initPprof() (err error) {
	if Cfg.App.Mode == AppDevMode {
		go func() {
			http.ListenAndServe(Cfg.Pprof.Addr, nil)
		}()
	}

	return nil
}

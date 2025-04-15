package service

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Service struct {
	mux *http.ServeMux
}

func NewService() *Service {
	svc := &Service{
		mux: http.NewServeMux(),
	}
	svc.mux.Handle("/metrics", promhttp.Handler())
	svc.mux.HandleFunc("/debug/pprof/", pprof.Index)
	svc.mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	svc.mux.HandleFunc("/debug/pprof/profile", pprof.Profile) // go tool pprof profile
	svc.mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	svc.mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	go svc.Start(":6060")
	return svc
}

func (s *Service) Start(lsnAddr string) {
	err := http.ListenAndServe(lsnAddr, s.mux)
	if err != nil {
		panic(fmt.Sprintf("init start listen monitor fail: %s", err.Error()))
	}
}

package runner

import (
	"fmt"
	"go-backend-common/util"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Service interface {
	Start() error
	Stop() error
}

type ServiceRunner interface {
	Wait()
}

func RunService(s Service) ServiceRunner {
	r := newServiceRunner(s)
	r.run()
	return r
}

func newServiceRunner(s Service) *serviceRunner {
	return &serviceRunner{
		signals: make(chan os.Signal, 1),
		service: s,
	}
}

type serviceRunner struct {
	signals chan os.Signal
	service Service

	stopped int32

	wg sync.WaitGroup
}

func (r *serviceRunner) run() {
	r.wg.Add(1)
	go r.handleSignal()
	go r.handleStart()
}

func (r *serviceRunner) handleStart() {
	func() {
		defer util.Recovery()
		err := r.service.Start()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if atomic.LoadInt32(&r.stopped) == 0 {
		fmt.Println(22222)
		r.wg.Done()
	}
}

func (r *serviceRunner) handleSignal() {
	signal.Notify(r.signals, syscall.SIGPIPE, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT)
	for {
		select {
		case sig := <-r.signals:
			switch sig {
			case syscall.SIGPIPE:
			case syscall.SIGINT:
				r.signalHandler()
				os.Exit(1)
			default:
				r.signalHandler()
				r.wg.Done()
			}
		}
	}
}

func (r *serviceRunner) signalHandler() {
	go func() {
		to := 10 * time.Second
		time.Sleep(to)
		os.Exit(1)
	}()
	atomic.StoreInt32(&r.stopped, 1)
	r.service.Stop()
}

func (r *serviceRunner) Wait() {
	r.wg.Wait()
	fmt.Println(44444)

}

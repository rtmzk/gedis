package server

import (
	"context"
	"github.com/rtmzk/gedis/pkg/logger"
	"github.com/rtmzk/gedis/pkg/tcp"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	var closeChan = make(chan struct{})
	var signalChan = make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-signalChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	ListenAndServe(listener, handler, closeChan)
	return nil
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {
	var ctx = context.Background()
	var waitDown sync.WaitGroup

	go func() {
		<-closeChan
		logger.Info("shutting down")
		_ = listener.Close()
		_ = handler.Close()
	}()
	defer func() {
		_ = listener.Close()
		_ = handler.Close()
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		waitDown.Add(1)
		logger.Info("accepted new client connection")
		go func() {
			defer waitDown.Done()
			handler.Handle(ctx, conn)
		}()
	}
	waitDown.Wait()
}

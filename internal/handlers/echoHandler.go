package handlers

import (
	"bufio"
	"context"
	"github.com/rtmzk/gedis/pkg/logger"
	"github.com/rtmzk/gedis/pkg/sync/atomic"
	"github.com/rtmzk/gedis/pkg/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (client *EchoClient) Close() error {
	client.Waiting.WaitWithTimeout(10 * time.Second)
	return client.Conn.Close()
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func (handler *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		_ = conn.Close()
	}

	client := &EchoClient{Conn: conn}
	handler.activeConn.Store(client, struct{}{})
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Info("Connection close")
				handler.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}
		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

func (handler *EchoHandler) Close() error {
	logger.Info("handler shutting down")
	handler.closing.Set(true)
	handler.activeConn.Range(func(key, value any) bool {
		client := key.(*EchoClient)
		_ = client.Conn.Close()
		return true
	})

	return nil
}

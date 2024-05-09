package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// Close 监听signal并停止
func Close(handler func()) {
	ctx := make(chan os.Signal, 1)
	signal.Notify(ctx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	<-ctx
	signal.Stop(ctx)

	handler()
}

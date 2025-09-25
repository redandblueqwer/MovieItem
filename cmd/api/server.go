package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	// 创建HTTP服务端 指针
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit // 阻塞，等待quit中捕获到停止信号
		app.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx) // 拒绝新请求，处理已经接收的请求

		app.logger.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})
		app.wg.Wait() // 等待所有后台任务结束
		if err != nil {
			shutdownError <- err
		} else {
			shutdownError <- nil
		}

	}()

	// 启动服务
	app.logger.PrintInfo("starting server",
		map[string]string{
			"addr": srv.Addr,
			"env":  app.config.env,
		})

	err := srv.ListenAndServe() // 主进程持续阻塞监听请求，直到srv.Shutdown触发
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError // 接收阻塞直到，有发送到shutdownError的操作
	if err != nil {
		return err
	}
	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
	return nil
}

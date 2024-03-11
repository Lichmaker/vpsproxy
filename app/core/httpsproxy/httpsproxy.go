package httpsproxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/Lichmaker/vpsproxy/app/core/gracefulstop"
	"github.com/Lichmaker/vpsproxy/pkg/logger"

	"go.uber.org/zap"
)

func Proxy(target string, listen string, certPath, keyPath string) {
	// 创建反向代理的目标URL
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	// 创建反向代理的处理程序
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ErrorLog = log.New(
		log.Writer(),
		"REVERSE PROXY ",
		log.LstdFlags,
	)

	// 创建HTTP服务器监听特定端口
	server := &http.Server{
		Addr: listen,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 在请求头中添加自定义信息等等
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

			// 执行反向代理
			proxy.ServeHTTP(w, r)
		}),
	}
	gracefulstop.RegisterStop("http proxy server", func() {
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Logger.Warn("stop http proxy failed", zap.Error(err))
		}
	})

	if len(certPath) == 0 || len(keyPath) == 0 {
		// 启动HTTP服务器
		logger.Logger.Warn("[HTTP] Starting reverse proxy server", zap.String("port", listen))
		go func() {
			err = server.ListenAndServe()
			if err != nil {
				logger.Logger.Warn("[HTTP] reverse proxy server error", zap.Error(err))
			}
		}()
	} else {
		// 启动HTTPS服务器
		logger.Logger.Warn("[HTTPS] Starting reverse proxy server", zap.String("port", listen))
		go func() {
			err = server.ListenAndServeTLS(certPath, keyPath)
			if err != nil {
				logger.Logger.Warn("[HTTPS] reverse proxy server error", zap.Error(err))
			}
		}()
	}
}

package proxy

import (
	"fmt"
	"net"
	"sync"

	vpsproxyController "github.com/Lichmaker/vpsproxy/app/controller/vpsproxy"
	"github.com/Lichmaker/vpsproxy/app/core/gracefulstop"
	"github.com/Lichmaker/vpsproxy/app/core/grpcproxy"
	"github.com/Lichmaker/vpsproxy/app/core/httpsproxy"
	"github.com/Lichmaker/vpsproxy/pkg/config"
	"github.com/Lichmaker/vpsproxy/pkg/logger"
	"github.com/Lichmaker/vpsproxy/protobuf/pb/vpsproxy"

	"github.com/spf13/cobra"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var CmdProxy = &cobra.Command{
	Use:   "run",
	Short: "Start proxy",
	Run:   runProxy,
	Args:  cobra.NoArgs,
}

func init() {
	// 注册子命令
	CmdProxy.AddCommand()
}

func runProxy(cmd *cobra.Command, args []string) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		httpsproxy.Proxy(
			config.GetString("proxy.https.target"),
			fmt.Sprintf(":%d", config.GetInt("proxy.https.port")),
			config.GetString("proxy.https.cert"),
			config.GetString("proxy.https.key"),
		)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcproxy.Proxy(
			config.GetString("proxy.grpc.target"),
			fmt.Sprintf(":%d", config.GetInt("proxy.grpc.port")),
			config.GetString("proxy.grpc.secret"),
		)

	}()

	wg.Wait()

	// 监听本地端口
	listenAddr := fmt.Sprintf("%s:%d", config.GetString("app.host", "0.0.0.0"), config.GetInt("app.port", "10099"))
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(grpcproxy.BuildGRPCInterceptor(config.GetString("proxy.grpc.secret"))),
	)
	gracefulstop.RegisterStop("vpsproxy grpc server", s.GracefulStop)

	// 注册服务
	vpsproxy.RegisterVPSProxyServiceServer(s, &vpsproxyController.VPSProxyServer{})

	reflection.Register(s)

	logger.Logger.Warn(fmt.Sprintf("listening vpsproxy grpc %s...", listenAddr))
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}

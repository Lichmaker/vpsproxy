package grpcproxy

import (
	"context"
	"net"

	"github.com/Lichmaker/vpsproxy/app/core/gracefulstop"
	"github.com/Lichmaker/vpsproxy/pkg/logger"

	proxyPkg "github.com/mwitkow/grpc-proxy/proxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Proxy(target string, listen string, setSecret string) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	localRPCConn, err := grpc.Dial(target, dialOpts...)
	if err != nil {
		panic(err)
	}

	// 监听本地端口
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		panic(err)
	}

	s := proxyPkg.NewProxy(localRPCConn, []proxyPkg.Interceptor{BuildReverseProxyInterceptor(setSecret)})
	gracefulstop.RegisterStop("grpc proxy server", s.GracefulStop)
	logger.Logger.Warn("[GRPC] Starting reverse proxy server", zap.String("port", listen))
	go func() {
		if err := s.Serve(lis); err != nil {
			logger.Logger.Warn("[GRPC] reverse proxy server error", zap.Error(err))
		}
	}()

}

func BuildReverseProxyInterceptor(secret string) proxyPkg.Interceptor {
	return func(serverStream grpc.ServerStream) error {
		md, ok := metadata.FromIncomingContext(serverStream.Context())
		if !ok {
			return status.Error(codes.Internal, "missing metadata")
		}
		secretValue := md.Get("secret")
		if len(secretValue) == 0 {
			return status.Error(codes.Internal, "missing secret")
		}
		if secretValue[0] != secret {
			return status.Error(codes.Internal, "secret incorrect")
		}
		return nil
	}
}

func BuildGRPCInterceptor(secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "missing metadata")
		}
		secretValue := md.Get("secret")
		if len(secretValue) == 0 {
			return nil, status.Error(codes.Internal, "missing secret")
		}
		if secretValue[0] != secret {
			return nil, status.Error(codes.Internal, "secret incorrect")
		}
		resp, err = handler(ctx, req)
		return
	}
}

package stop

import (
	"context"
	"fmt"
	"time"

	"github.com/Lichmaker/vpsproxy/pkg/config"
	"github.com/Lichmaker/vpsproxy/protobuf/pb/vpsproxy"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var CmdStop = &cobra.Command{
	Use:   "stop",
	Short: "stop vpsproxy",
	Run:   runStop,
	Args:  cobra.NoArgs,
}

func init() {
	// 注册子命令
	CmdStop.AddCommand()
}

func runStop(cmd *cobra.Command, args []string) {
	url := fmt.Sprintf("127.0.0.1:%d", config.GetInt("app.port"))
	secret := config.GetString("proxy.grpc.secret")
	backoffConfig := backoff.DefaultConfig

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoffConfig,
		}),
	}
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(
		buildUnaryClientInterceptor(secret),
	))

	dial, err := grpc.Dial(
		url,
		dialOpts...,
	)
	if err != nil {
		fmt.Printf("dial failed: %s\n", err)
		return
	}

	client := vpsproxy.NewVPSProxyServiceClient(dial)
	_, err = client.Shutdown(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Printf("shutdown failed: %s\n", err)
	} else {
		fmt.Println("shutdown success")
	}
}

func buildUnaryClientInterceptor(secret string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 获取要发送给服务端的`metadata`
		md, ok := metadata.FromOutgoingContext(ctx)
		if ok && len(md.Get("secret")) == 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, "secret", secret)
		} else {
			newMd := make(metadata.MD)
			newMd.Set("secret", secret)
			ctx = metadata.NewOutgoingContext(ctx, newMd)
		}

		// 设置5秒的超时
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Invoking the remote method
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

package vpsproxy

import (
	"context"
	"fmt"
	"time"

	"github.com/Lichmaker/vpsproxy/app/core/gracefulstop"
	"github.com/Lichmaker/vpsproxy/pkg/config"
	"github.com/Lichmaker/vpsproxy/pkg/logger"
	vpsproxyPB "github.com/Lichmaker/vpsproxy/protobuf/pb/vpsproxy"

	"google.golang.org/protobuf/types/known/emptypb"
)

type VPSProxyServer struct {
	vpsproxyPB.UnimplementedVPSProxyServiceServer
}

var _ vpsproxyPB.VPSProxyServiceServer = (*VPSProxyServer)(nil)

func (v *VPSProxyServer) UpdateHttpProxy(ctx context.Context, req *vpsproxyPB.UpdateHttpProxyReq) (*emptypb.Empty, error) {
	config.Set("proxy.https.port", req.Port)
	return &emptypb.Empty{}, nil
}

func (v *VPSProxyServer) Shutdown(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	go func() {
		for i := 3; i > 0; i-- {
			logger.Logger.Warn(fmt.Sprintf("ready shutdown server in %d...", i))
			time.Sleep(time.Second)
		}
		gracefulstop.Stop()
	}()
	return &emptypb.Empty{}, nil
}

func (v *VPSProxyServer) GetHttpProxy(ctx context.Context, _ *emptypb.Empty) (*vpsproxyPB.GetHttpProxyResponse, error) {
	p := config.GetInt("proxy.https.port")
	return &vpsproxyPB.GetHttpProxyResponse{
		Port: uint32(p),
	}, nil
}

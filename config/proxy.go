package config

import "github.com/Lichmaker/vpsproxy/pkg/config"

func init() {
	config.BindEnv("proxy.https.port", "PROXY_HTTPS_PORT")
	config.BindEnv("proxy.https.cert", "PROXY_HTTPS_CERT")
	config.BindEnv("proxy.https.key", "PROXY_HTTPS_KEY")
	config.BindEnv("proxy.https.target", "PROXY_HTTPS_TARGET")
	config.BindEnv("proxy.grpc.port", "PROXY_RPC_PORT")
	config.BindEnv("proxy.grpc.secret", "PROXY_RPC_SECRET")
	config.BindEnv("proxy.grpc.target", "PROXY_RPC_TARGET")
}

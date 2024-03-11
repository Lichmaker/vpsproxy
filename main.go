package main

import (
	"os"

	"github.com/Lichmaker/vpsproxy/app/cmd/install"
	"github.com/Lichmaker/vpsproxy/app/cmd/proxy"
	"github.com/Lichmaker/vpsproxy/app/cmd/stop"
	"github.com/Lichmaker/vpsproxy/app/runtime/app"
	"github.com/Lichmaker/vpsproxy/app/runtime/flags"
	"github.com/Lichmaker/vpsproxy/bootstrap"
	"github.com/Lichmaker/vpsproxy/config"
	configPkg "github.com/Lichmaker/vpsproxy/pkg/config"

	"github.com/spf13/cobra"
)

func main() {
	app.AbsolutePath, _ = os.Getwd()
	config.Initialize()

	// 主入口，调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "vpsproxy",
		Short: "vpsproxy",
		Long:  ``,

		// 定义所有注册命令都会执行的func
		PersistentPreRun: func(command *cobra.Command, args []string) {
			configPkg.InitConfig("", "settings.yaml")

			bootstrap.SetupLogger()

		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		proxy.CmdProxy,
		install.CmdInstall,
		stop.CmdStop,
	)

	// 注册指定命令，非全局的flag
	flags.RegisterCommandFlags()

	// 注册全局参数
	flags.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

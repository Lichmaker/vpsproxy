package install

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var CmdInstall = &cobra.Command{
	Use:   "install",
	Short: "Install VPSProxy",
	Run:   runInstall,
	Args:  cobra.NoArgs,
}

func init() {
	// 注册子命令
	CmdInstall.AddCommand()
}

const serviceFileName = "vpsproxy.service"

func runInstall(cmd *cobra.Command, args []string) {
	// 获取程序自身的路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	// 获取当前用户信息
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	// 获取程序所在目录
	exeDir := filepath.Dir(exePath)

	// 生成服务文件内容
	serviceContent := fmt.Sprintf(`[Unit]
Description=VPSProxy Service
After=network.target

[Service]
WorkingDirectory=%s
ExecStart=%s run
ExecStop=%s stop
Restart=always
StartLimitInterval=0
StartLimitBurst=2
User=%s
Group=%s

[Install]
WantedBy=multi-user.target

`, exeDir, exePath, exePath, currentUser.Username, currentUser.Gid)

	// 写入服务文件
	serviceFilePath := filepath.Join("/etc/systemd/system", serviceFileName)
	err = writeToFile(serviceFilePath, serviceContent)
	if err != nil {
		fmt.Println("Error writing service file:", err)
		return
	}

	fmt.Println("Service file written successfully:", serviceFilePath)
}

func writeToFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

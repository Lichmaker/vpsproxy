package config

import (
	"fmt"

	"github.com/Lichmaker/vpsproxy/pkg/helpers"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// viper 实例
var viperInst *viper.Viper

// 把某一个key绑定到一个环境变量上，当存在环境变量时，将会优先使用
var bindEnvBuffer map[string]string

func init() {
	bindEnvBuffer = make(map[string]string)
}

// 初始化配置信息，完成环境变量和config的加载
func InitConfig(envPrefix string, appointedConfigPath string) {

	// 初始化viper
	viperInst = viper.New()
	viperInst.SetConfigType("yaml")
	viperInst.AddConfigPath(".")
	viperInst.SetEnvPrefix(envPrefix) // 设置前缀，同一环境用以区分
	viperInst.AutomaticEnv()          // 读取环境变量
	loadBindEnv()
	loadConfig(appointedConfigPath)

}

func loadBindEnv() {
	for bindKey, envName := range bindEnvBuffer {
		viperInst.BindEnv(bindKey, envName)
	}
}

func loadConfig(appointedConfigPath string) {
	if len(appointedConfigPath) > 0 {
		viperInst.SetConfigFile(appointedConfigPath)
	}

	if err := viperInst.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Get(path string, defaultValues ...interface{}) string {
	return GetString(path, defaultValues...)
}

// 新增配置项
func BindEnv(keyName string, envName string) {
	bindEnvBuffer[keyName] = envName
}

func Set(name string, value interface{}) {
	viperInst.Set(name, value)
	WriteConfig()
}

func WriteConfig() {
	if err := viperInst.WriteConfig(); err != nil {
		fmt.Println("write config file failed! " + err.Error())
	}
}

func internalGet(path string, defaultValues ...interface{}) interface{} {
	// config 或者环境变量都不存在，使用default
	if !viperInst.IsSet(path) || helpers.Empty(viperInst.Get(path)) {
		if len(defaultValues) > 0 {
			return defaultValues[0]
		}
		return nil
	}
	return viperInst.Get(path)
}

// get 的 string 版本
func GetString(path string, defaultValues ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValues...))
}

// get 的 int 版本
func GetInt(path string, defaultValues ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValues...))
}

func GetFloat64(path string, defaultValues ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValues...))
}

func GetInt64(path string, defaultValues ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValues...))
}

func GetUint(path string, defaultValues ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValues...))
}

func GetBool(path string, defaultValues ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValues...))
}

func GetStringMapString(path string, defaultValues ...interface{}) map[string]string {
	return viperInst.GetStringMapString(path)
}

package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"reflect"
	"runtime"
	"runtime/debug"
)

type setting struct {
	// 只能定义string类型的属性, 否则会读取不到! 如果需要int等其他类型, 需在使用处进行类型转换
	DEBUG         bool
	JwtSecretKey  string `yaml:"JWT_SECRET_KEY"` // jwt 密钥
	MysqlUrl      string `yaml:"MYSQL_URL"`      // mysql 数据库连接地址
	SqliteUrl     string `yaml:"SQLITE_URL"`     // sqlite 数据库连接地址
	ObsDomain     string `yaml:"OBS_DOMAIN"`
	ObsEndPoint   string `yaml:"OBS_ENDPOINT"`
	ObsBucketName string `yaml:"OBS_BUCKET_NAME"`
	ObsAk         string `yaml:"OBS_AK"`
	ObsSk         string `yaml:"OBS_SK"`
}

// readEnvConfig 通过反射读取系统环境变量
// reference: https://geektutu.com/post/hpg-reflect.html
func readEnvConfig() *setting {
	config := setting{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		key := f.Tag.Get("yaml")
		if env, exist := os.LookupEnv(key); exist {
			value.FieldByName(f.Name).Set(reflect.ValueOf(env))
		} else {
			fmt.Println("环境变量:", key, "不存在")
		}
	}
	return &config
}

type localSetting struct {
	Data setting `yaml:"data"`
}

func readYamlConfig(env string) (*setting, error) {
	// 1. 获取当前文件的路径
	_, currFileName, _, _ := runtime.Caller(0)               // 获取当前函数的路径信息
	localProjectRootPath := path.Dir(path.Dir(currFileName)) // 获取当前项目根路径
	configPath := fmt.Sprintf("%s/config/config_%s.yaml", localProjectRootPath, env)
	// 读取本地配置文件
	data, err := os.ReadFile(configPath) // 读取指定路径下的文件内容
	if err != nil {
		return nil, errors.Wrap(err, "read local config file error") // 如果读取文件出错，则返回nil和错误信息
	}

	// 2. 解析文件
	var y localSetting
	err = yaml.Unmarshal(data, &y) // 将文件内容解析为localSetting结构体类型
	return &y.Data, nil            // 返回解析后的结构体指针和nil错误
}

var Setting = new(setting)

// 初始化项目配置
func init() {
	defer func() { // 初始化配置未知异常捕获
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s\n", r, s)
		}
	}()

	ENV := os.Getenv("ENV")
	if ENV == "local" || ENV == "dev" || ENV == "prod" {
		localConf, _ := readYamlConfig(ENV)
		localConf.DEBUG = true
		Setting = localConf
	} else {
		Setting.DEBUG = false
		Setting = readEnvConfig()
	}
}

package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

//go:embed config.toml
var ConfigFS string

type Config struct {
	Version  string   `toml:"version" env:"version"`
	Contract Contract `toml:"contract" env:"contract"`
	Endpoint Endpoint `toml:"endpoint" env:"endpoint"`
	Account  Account  `toml:"account" env:"account"`
	DataFin  DataFin  `toml:"datafin" env:"datafin"`
	MySQL    MySQL    `toml:"mysql" env:"mysql"`
	Pulsar   Pulsar   `toml:"pulsar" env:"pulsar"`
	Redis    Redis    `toml:"redis" env:"redis"`
	Minio    Minio    `toml:"minio" env:"minio"`
}

type Contract struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type Endpoint struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type Account struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
}

type DataFin struct {
	Domain   string `toml:"domain" env:"domain"`
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogFile  string `toml:"log-file" env:"log_file"`
	DataDir  string `toml:"data-dir" env:"data_dir"`
}

type MySQL struct {
	Domain   string `toml:"domain" env:"domain"`
	Port     int    `toml:"port" env:"port"`
	User     string `toml:"user" env:"user"`
	Password string `toml:"password" env:"password"`
	Database string `toml:"database" env:"database"`
}

type Redis struct {
	Address  string `toml:"address" env:"address"`
	Password string `toml:"password" env:"password"`
}

type Pulsar struct {
	Domain              string `toml:"domain" env:"domain"`
	Port                int    `toml:"port" env:"port"`
	OperationTimeout    uint64 `toml:"operation-timeout" env:"operation_timeout"`
	ConnectionTimeout   uint64 `toml:"connection-timeout" env:"connection_timeout"`
	TopicSyncTask       string `toml:"topic-sync-task" env:"topic_sync_task"`
	TopicTransformImage string `toml:"topic-transform-image" env:"topic_transform_image"`
}

type Minio struct {
	Address          string `toml:"address" env:"address"`
	AccessKey        string `toml:"access-key" env:"access_key"`
	SecretKey        string `toml:"secret-key" env:"secret_key"`
	Region           string `toml:"region" env:"region"`
	TokenImageBucket string `toml:"token-image-bucket" env:"token_image_bucket"`
}

// set default config
var config = &Config{
	Contract: Contract{
		HTTPPort: 30110,
		GrpcPort: 30111,
	},
}

type envMatcher struct {
	envMap map[string]string
}

func DetectEnv(co *Config) (err error) {
	e := &envMatcher{}
	e.envMap = make(map[string]string)
	ct := reflect.TypeOf(co)
	e.detectEnv(ct, "", "")
	_, err = toml.Decode(e.toToml(), co)
	return err
}

// read environment var
func (e *envMatcher) detectEnv(t reflect.Type, preffix, _preffix string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		it := t.Field(i)
		envKey := fmt.Sprintf("%v%v", preffix, it.Tag.Get("env"))
		_envKey := fmt.Sprintf("%v%v", _preffix, it.Tag.Get("toml"))
		if it.Type.Kind() != reflect.Struct {
			if envValue, ok := os.LookupEnv(envKey); ok {
				if it.Type.Kind() == reflect.String {
					e.envMap[_envKey] = fmt.Sprintf("\"%v\"", envValue)
				} else {
					e.envMap[_envKey] = envValue
				}
			}
			continue
		}
		envKey = fmt.Sprintf("%v%v_", preffix, it.Tag.Get("env"))
		_envKey = fmt.Sprintf("%v%v.", _preffix, it.Tag.Get("toml"))
		e.detectEnv(it.Type, envKey, _envKey)
	}
}

func (e *envMatcher) toToml() string {
	var b bytes.Buffer

	for v := range e.envMap {
		b.WriteString(fmt.Sprintf("%v=%v\n", v, e.envMap[v]))
	}

	return b.String()
}

func init() {
	md, err := toml.Decode(ConfigFS, config)
	if err != nil {
		panic(fmt.Sprintf("failed to parse config file, %v", err))
	}
	if len(md.Undecoded()) > 0 {
		fmt.Printf("cannot parse [%v] to config\n", md.Undecoded())
	}
	err = DetectEnv(config)
	if err != nil {
		fmt.Printf("environment variable parse failed, %v", err)
	}
}

func GetConfig() *Config {
	return config
}

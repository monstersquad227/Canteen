package configuration

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Configurations struct {
	ServerPort     string `ini:"server_port"`
	MysqlAddress   string `ini:"mysql_address"`
	MysqlUsername  string `ini:"mysql_username"`
	MysqlPassword  string `ini:"mysql_password"`
	MysqlPort      string `ini:"mysql_port"`
	MysqlDatabases string `ini:"mysql_databases"`
	MysqlCharset   string `ini:"mysql_charset"`
	JwtSecretKey   string `ini:"jwt_secret_key"`
	ExpireTime     int    `ini:"expire_time"`
	EncryptionKey  string `ini:"encryption_key"`
}

var Configs = Configurations{}

func LoadConfig(name string) {
	cfg, err1 := ini.Load(name)
	if err1 != nil {
		fmt.Println("Load Config File Failed: ", err1)
	}
	if err2 := cfg.MapTo(&Configs); err2 != nil {
		fmt.Println("Analyze Failed: ", err2)
	}
}

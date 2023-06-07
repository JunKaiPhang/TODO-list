package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type App struct {
	Name      string
	PrefixUrl string
}

var AppSetting = &App{}

type Server struct {
	RunMode  string
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	User     string
	Password string
	Host     string
	Db       string
}

var DatabaseSetting = &Database{}

type Cors struct {
	AllowOrigin      []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
}

var CorsSetting = &Cors{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("cors", CorsSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

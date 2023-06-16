package setting

import (
	"log"
	"time"

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

type Jwt struct {
	JwtSecret            string
	TokenDuration        time.Duration
	RandOAuthStateString string
}

var JwtSetting = &Jwt{}

type Fb struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var FbSetting = &Fb{}

type Gmail struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var GmailSetting = &Gmail{}

type Github struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var GithubSetting = &Github{}

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
	mapTo("jwt", JwtSetting)
	mapTo("fb", FbSetting)
	mapTo("gmail", GmailSetting)
	mapTo("github", GithubSetting)

	JwtSetting.TokenDuration = JwtSetting.TokenDuration * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

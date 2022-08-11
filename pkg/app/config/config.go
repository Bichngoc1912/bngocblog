package config

const (
	SERVER_NAME = "bngocblog"
)

type SessionServer struct {
	Type 					string 			`json:"serverConfig"`
	StorageServer RedisServer `json:"store"`
}

type RedisServer struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}

type MemcachedServer struct {
	List      []string `json:"list"`
	ThreadNum int      `json:"thread_num"`
}

type HttpServerSetting struct {
	Addr     string `json:"addr"`
	Network  string `json:"network"`
	Compress bool   `json:"compress"`
}

type ServerSetting struct {
	ServerName 				string 						`json:"name"`
	SessionServer			SessionServer 		`json:"session_service"`
	DebugMode         bool              `json:"debug"`
	ProductionEnv     bool              `json:"production"`
	ServerURI         string            `json:"server_uri"`
	LogPath           string            `json:"log_path"`
	HttpServerSetting HttpServerSetting `json:"http_service"`
	MemcachedServer   MemcachedServer   `json:"memcached_service"`
}
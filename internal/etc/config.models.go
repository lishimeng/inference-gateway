package etc

// Configuration 全局环境变量配置。（按需添加）
type Configuration struct {
	Web        web        `toml:"web"`
	Db         db         `toml:"db"`
	Dl         dl         `toml:"dl"`
	Mqtt       mqtt       `toml:"mqtt"`
	Redis      redis      `toml:"redis"`
	Token      token      `toml:"token"`
	FileSystem fileSystem `toml:"fileSystem"`
}
type fileSystem struct {
	Root   string `toml:"root"`
	Domain string `toml:"domain"`
}
type web struct {
	Listen    string `toml:"listen"`
	AppLogDir string `toml:"applogdir"`
}

type db struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Ssl      string `toml:"ssl"`
}

type dl struct {
	Host string `toml:"host"`
	Key  string `toml:"key"`
}

type mqtt struct {
	Broker   string `toml:"broker"`
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type redis struct {
	Enable   bool   `toml:"enable"`
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	// other
}

type token struct {
	Enable bool   `toml:"enable"`
	Issuer string `toml:"issuer"`
	Key    string `toml:"key"`
	Ttl    int    `toml:"ttl"`
}

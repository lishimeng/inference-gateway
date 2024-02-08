package etc

// Configuration 全局环境变量配置。（按需添加）
type Configuration struct {
	Web        web        `toml:"web"`
	Db         db         `toml:"db"`
	FileSystem fileSystem `toml:"fileSystem"`
	Geo        geo        `toml:"geo"`
	Wx         wx         `toml:"wx"`
}
type wx struct {
	Appid  string `toml:"appid"`
	Secret string `toml:"secret"`
}
type geo struct {
	Key string `toml:"key"`
}
type fileSystem struct {
	Root   string `toml:"root"`
	Domain string `toml:"domain"`
}
type web struct {
	Listen string `toml:"listen"`
}

type db struct {
	Host     string `toml:"host"`
	Database string `toml:"database"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Ssl      string `toml:"ssl"`
}

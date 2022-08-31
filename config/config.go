package config

const (
	JwtKey     = "RedRock"
	JwtTimeout = 7 * 24 * 60 * 60
)

type DatabaseConfiguration struct {
	DefaultDbName    string
	DefaultRoot      string
	DefaultPassword  string
	DefaultIpAndPort string
	DefaultCharset   string
}

type EmailConfiguration struct {
	EmailAuthSender   string
	EmailAuthAccount  string
	EmailAuthPassword string
}

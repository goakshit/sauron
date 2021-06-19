package types

// PostgresConf - Postgres config
type PostgresConf struct {
	User         string
	Passwd       string
	DatabaseName string
	Host         string
	Port         int
	SSLMode      string
	Timezone     string
}

// Config - Store application config data
type Config struct {
	AppPort  int
	AppEnv   string
	Postgres PostgresConf
}

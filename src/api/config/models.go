package config

type Environment struct {
	Environment                   string `env:"Environment,default=local"`
	MongodbAddrs                  string `env:"MONGO_DATABASE_ADDRS,default=mongodb+srv://adminApp:vWQuBeGLwDtr6B3I@yourfinances.slmnk.mongodb.net/?retryWrites=true&w=majority"`
	MongodbUser                   string `env:"MONGO_DATABASE_USERNAME ,default=adminApp"`
	MongodbDatabase               string `env:"MONGO_DATABASE_NAME,default=your-finances-auth"`
	MongodbPassword               string `env:"MONGO_DATABASE_PASSWORD,default=vWQuBeGLwDtr6B3I"`
	MongodbMaxPoolSize            uint64 `env:"MONGO_MAX_POOL_SIZE,default=100"`
	MongodbMaxConnIdleTine        int    `env:"MONGO_MAX_CONN_IDLE_TIME,default=2"`
	LogLevel                      string `env:"LOG_LEVEL,default=warn"`
	RefreshTokenExpireTimeMinutes int    `env:"REFRESH_EXPIRE,default=15"`
	JwtExpireTimeMinutesClient    int    `env:"JWT_EXPIRE_CLIENT,default=3"`
	JwtExpireTimeMinutes          int    `env:"JWT_EXPIRE,default=15"`
	JwtKey                        string `env:"JWT_KEY,default=654d30eae2f0496295a2e161e644b31e-06e01c70dd88480c8e07e5e89c1668da-89ddd1301a624dcda4f2da8abc1190f8"`
}

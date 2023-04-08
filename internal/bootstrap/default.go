package bootstrap

import "github.com/ayanami-desu/iropic-go/internal/utils"

func DefaultConfig() *Config {
	return &Config{
		IsDev:      "no",
		Username:   "admin",
		Password:   utils.RandomString(8),
		JwtSecret:  utils.RandomString(16),
		PageMaxNum: 16,
		PathPrefix: "/iropic-go/Pictures/",
		Port:       "5245",
		DBConfig: DBConfig{
			User:     "go",
			Password: "123456",
			DBName:   "go",
		},
	}
}

package cnf

import "encoding/json"

type Config struct {
	Timezone      string
	HttpHost      string
	HttpPort      int
	MySqlHost     string
	MySqlPort     int
	MySqlUsername string
	MySqlPassword string
	MySqlDatabase string
	IndentJson    bool
	Version       bool
}

func (config *Config) String() string {
	data, _ := json.Marshal(config)
	return string(data)
}

const (
	Protobuf string = "protobuf"
	Json            = "json"
)

package cnf

import "encoding/json"

type Config struct {
	Timezone      string
	HttpPort      int
	MySqlHost     string
	MySqlPort     int
	MySqlUsername string
	MySqlPassword string
	MySqlDatabase string
	ResponseType  string
	Indent        bool
}

func (config *Config) String() string {
	data, _ := json.Marshal(config)
	return string(data)
}

const (
	Protobuf string = "protobuf"
	Json            = "json"
)

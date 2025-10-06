package dtos

type AuthResult struct {
	UserInfo   AuthUserInfo   `json:"user_info"`
	ServerInfo AuthServerInfo `json:"server_info"`
}

type AuthUserInfo struct {
	Username             string   `json:"username"`
	Password             string   `json:"password"`
	Message              string   `json:"message"`
	Auth                 int      `json:"auth"`
	Status               string   `json:"status"`
	ExpirationDate       string   `json:"exp_date"`
	IsTrial              string   `json:"is_trial"`
	ActiveCons           string   `json:"active_cons"`
	CreatedAt            string   `json:"created_at"`
	MaxConnections       string   `json:"max_connections"`
	AllowedOutputFormats []string `json:"allowed_output_formats"`
}

type AuthServerInfo struct {
	Url            string `json:"url"`
	Port           string `json:"port"`
	HttpsPort      string `json:"https_port"`
	ServerProtocol string `json:"server_protocol"`
	RtmpPort       string `json:"rtmp_port"`
	Timezone       string `json:"timezone"`
	TimestampNow   int    `json:"timestamp_now"`
	TimeNow        string `json:"time_now"`
	Process        bool   `json:"process"`
}

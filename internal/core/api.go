package core

type ApiCfg struct {
	Addr string
}

type AppCfg struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

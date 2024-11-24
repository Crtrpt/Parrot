package core

type Crontab struct {
	Spec string `json:"spec"`
	Name string `json:"name"`
	Argv *any   `json:"argv"`
}

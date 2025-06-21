package request

import "encoding/json"

type MenuResp struct {
	ID       int64       `json:"key"`
	Name     string      `json:"label"`
	Route    string      `json:"route,omitempty"`
	Icon     string      `json:"icon,omitempty"`
	Children []*MenuResp `json:"children,omitempty"`
}

type MenuForm struct {
	Name string `json:"name" form:"name"`
}

type JobResp struct {
	ID       int64       `json:"key"`
	Name     string      `json:"label"`
	Desc     string      `json:"desc"`
	Route    string      `json:"route,omitempty"`
	Icon     string      `json:"icon,omitempty"`
	Children []*MenuResp `json:"children,omitempty"`
}

type JobForm struct {
	Name string `json:"name" form:"name"`
}

type FormResp struct {
	ID     int64           `json:"id"`
	Name   string          `json:"name"`
	Schema json.RawMessage `json:"schema"`
}

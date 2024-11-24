package request

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

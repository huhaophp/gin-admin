package model

type PageInfo struct {
	Data []*interface{}         `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

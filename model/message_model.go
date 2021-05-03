package model

type Message struct {
	Id string `json:"id"`
	ruleId string `json:"rule_id"`
	extra string `json:"extra"`
}

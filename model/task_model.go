package model

type TaskModel struct {
	TaskId   string `json:"task_id"`
	RuleId   string `json:"rule_id"`
	Operator string `json:"operator"`
	ExecuteState int64 `json:"execute_state"`
	OutputName string `json:"output_name"`
	OutputUrl string `json:"output_url"`
	OutputState int64 `json:"output_state"`
}

func TaskTableName() string{
	return "task"
}

func (TaskModel) TableName() string{
	return TaskTableName()
}
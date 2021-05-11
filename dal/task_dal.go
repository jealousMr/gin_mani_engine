package dal

import (
	"context"
	"gin_mani_engine/model"
	pb_mani "gin_mani_engine/pb"
	logx "github.com/amoghe/distillog"
)

func AddTask(ctx context.Context, task *model.TaskModel) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return err
	}
	if err := db.Create(task).Error; err != nil {
		logx.Errorf("dal AddTask error:%v", err)
		return err
	}
	return nil
}

func GetTaskByCondition(ctx context.Context, taskId, ruleId string, executeState pb_mani.ExecuteState) ([]*model.TaskModel, error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return nil, err
	}
	tasks := make([]*model.TaskModel, 0)
	db = db.Table(model.TaskTableName())
	if taskId != "" {
		db = db.Where("task_id = ?", taskId)
	}
	if ruleId != "" {
		db = db.Where("rule_id = ?", ruleId)
	}
	if executeState != pb_mani.ExecuteState_execute_unknow {
		db = db.Where("execute_state = ?", executeState)
	}
	if err := db.Find(&tasks).Order("update_at DESC").Error; err != nil {
		logx.Errorf("dal GetTaskByCondition error:%v", err)
		return nil, err
	}
	return tasks, nil

}

func UpdateTask(ctx context.Context, taskId string, updateFileds map[string]interface{}) error {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return err
	}
	if err = db.Table(model.TaskTableName()).Where("task_id = ?", taskId).Updates(updateFileds).Error; err != nil {
		logx.Errorf("dal UpdateTask error:%v", err)
		return err
	}
	return nil
}

func GetTaskByRuleIds(ctx context.Context, ids []string) ([]*model.TaskModel, error) {
	db, err := GetDBProxy()
	if err != nil {
		logx.Errorf("db get error:%v", err)
		return nil, err
	}
	tasks := make([]*model.TaskModel, 0)
	if err := db.Table(model.TaskTableName()).Where("rule_id in (?)", ids).Find(&tasks).Error; err != nil {
		logx.Errorf("dal GetTaskByRuleIds error:%v", err)
		return nil, err
	}
	return tasks, nil
}

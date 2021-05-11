package logic

import (
	"context"
	"gin_mani_engine/dal"
	"gin_mani_engine/model"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	logx "github.com/amoghe/distillog"
)

func AddTask(ctx context.Context, task *pb_mani.Task) (string, error) {
	taskModel := &model.TaskModel{
		TaskId:       util.GenUID(),
		RuleId:       task.RuleId,
		Operator:     task.Operator,
		ExecuteState: int64(task.ExecuteState),
		OutputName:   task.OutputName,
		OutputState:  int64(task.OutputState),
		OutputUrl:    task.OutputUrl,
	}
	if err := dal.AddTask(ctx, taskModel); err != nil {
		logx.Errorf("logic AddTask error:%v", err)
		return "", err
	}
	return taskModel.TaskId, nil
}

func GetTaskByCondition(ctx context.Context, taskId, ruleId string, executeState pb_mani.ExecuteState) ([]*pb_mani.Task, error) {
	tasks, err := dal.GetTaskByCondition(ctx, taskId, ruleId, executeState)
	if err != nil {
		logx.Errorf("logic GetTaskByCondition error")
		return nil, err
	}
	taskList := make([]*pb_mani.Task, 0)
	for _, t := range tasks {
		taskList = append(taskList, &pb_mani.Task{
			TaskId:       t.TaskId,
			RuleId:       t.RuleId,
			Operator:     t.Operator,
			ExecuteState: pb_mani.ExecuteState(t.ExecuteState),
			OutputName:   t.OutputName,
			OutputUrl:    t.OutputUrl,
			OutputState:  pb_mani.FileState(t.OutputState),
		})
	}
	return taskList, nil
}

func UpdateTask(ctx context.Context, req *pb_mani.UpdateTaskReq) error {
	updateFileds := make(map[string]interface{}, 0)
	if req.Task.RuleId != "" {
		updateFileds["rule_id"] = req.Task.RuleId
	}
	if req.Task.Operator != "" {
		updateFileds["operator"] = req.Task.Operator
	}
	if req.Task.ExecuteState != pb_mani.ExecuteState_execute_unknow {
		updateFileds["execute_state"] = req.Task.ExecuteState
	}
	if req.Task.OutputName != "" {
		updateFileds["output_name"] = req.Task.OutputName
	}
	if req.Task.OutputUrl != "" {
		updateFileds["output_url"] = req.Task.OutputUrl
	}
	if req.Task.OutputState != pb_mani.FileState_unknown_file_state {
		updateFileds["output_state"] = req.Task.OutputState
	}
	if err := dal.UpdateTask(ctx, req.Task.TaskId, updateFileds); err != nil {
		logx.Errorf("logic UpdateTask error")
		return err
	}
	return nil
}

func GetTaskByRuleIds(ctx context.Context, ids []string) (map[string]*pb_mani.Task, error) {
	taskMap := make(map[string]*pb_mani.Task, 0)
	tasks, err := dal.GetTaskByRuleIds(ctx, ids)
	if err != nil {
		logx.Errorf("logic GetTaskByRuleIds error")
		return nil, err
	}
	for _, t := range tasks {
		taskMap[t.RuleId] = &pb_mani.Task{
			TaskId: t.TaskId,
			RuleId: t.RuleId,
			Operator: t.Operator,
			ExecuteState: pb_mani.ExecuteState(t.ExecuteState),
			OutputName: t.OutputName,
			OutputUrl: t.OutputUrl,
			OutputState: pb_mani.FileState(t.OutputState),
		}
	}
	return taskMap, nil
}

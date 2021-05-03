package handler

import (
	"context"
	"errors"
	"gin_mani_engine/logic"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	logx "github.com/amoghe/distillog"
)

func checkCreateTask(req *pb_mani.CreateTaskReq) error {
	if req.Task == nil || req.Task.RuleId == "" {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func CreateTask(ctx context.Context, req *pb_mani.CreateTaskReq) (resp *pb_mani.CreateTaskResp, err error) {
	resp = &pb_mani.CreateTaskResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkCreateTask(req); err != nil {
		logx.Errorf("CreateTask checkCreateTask error:%v", err)
		return
	}
	taskId, err := logic.AddTask(ctx, req.Task)
	if err != nil {
		logx.Errorf("CreateTask error:%v", err)
		return
	}
	resp.TaskId = taskId
	return
}

func checkQueryTaskByConditionParam(req *pb_mani.QueryTaskByConditionReq) error {
	if req.TaskId == "" && req.RuleId == "" && req.ExecuteState == pb_mani.ExecuteState_execute_unknow {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func QueryTaskByCondition(ctx context.Context, req *pb_mani.QueryTaskByConditionReq) (resp *pb_mani.QueryTaskByConditionResp, err error) {
	resp = &pb_mani.QueryTaskByConditionResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkQueryTaskByConditionParam(req); err != nil {
		logx.Errorf("QueryTaskByCondition checkQueryTaskByConditionParam error:%v", err)
		return
	}
	tasks, err := logic.GetTaskByCondition(ctx, req.TaskId, req.RuleId, req.ExecuteState)
	if err != nil {
		logx.Errorf("QueryTaskByCondition error:%v", err)
		return
	}
	resp.TaskList = tasks
	return

}

func checkUpdateTask(req *pb_mani.UpdateTaskReq) error {
	if req.Task == nil || req.Task.TaskId == "" {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func UpdateTask(ctx context.Context, req *pb_mani.UpdateTaskReq) (resp *pb_mani.UpdateTaskResp, err error) {
	resp = &pb_mani.UpdateTaskResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkUpdateTask(req); err != nil {
		logx.Errorf("checkUpdateTask error:%v", err)
		return
	}
	if err = logic.UpdateTask(ctx, req); err != nil {
		logx.Errorf("UpdateTask error:%v", err)
		return
	}
	return
}

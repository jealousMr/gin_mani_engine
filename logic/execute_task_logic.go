package logic

import (
	"context"
	"errors"
	"fmt"
	"gin_mani_engine/clients"
	"gin_mani_engine/dal"
	"gin_mani_engine/model"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	logx "github.com/amoghe/distillog"
)

func ExecuteTask(ctx context.Context) error {
	tasks, err := dal.GetTaskByCondition(ctx, "", "", pb_mani.ExecuteState_execute_running)
	if err != nil {
		logx.Errorf("ExecuteTask GetTaskByCondition error:%v", err)
		return err
	}
	if len(tasks) > 0 {
		logx.Warningln("current has task running...")
		return nil
	}
	msg, err := dal.GetMessage()
	if err != nil {
		logx.Errorf("GetMessage error:%v", err)
		return nil
	}
	if msg == nil {
		logx.Warningln("no msg current")
		return nil
	}
	if msg.TaskId == "" || msg.RuleId == "" {
		logx.Errorf("error msg:%v", msg)
		return errors.New("invalid message")
	}
	logx.Infof("current execute task:%s,extra:%s", msg.TaskId, msg.Extra)

	updateState := make(map[string]interface{}, 0)
	updateState["execute_state"] = pb_mani.ExecuteState_execute_running
	err = dal.UpdateTask(ctx, msg.TaskId, updateState)
	if err != nil {
		logx.Errorf("ExecuteTask UpdateTask error:%v", err)
		return err
	}

	center_rpc, err := clients.GetCenterClient()
	if err != nil {
		logx.Errorf("ExecuteTask GetCenterClient error:%v", err)
		return nil
	}
	rule, err := getRule(ctx, center_rpc, msg.RuleId)
	if err != nil {
		logx.Errorf("ExecuteTask getRule error:%v", err)
		return err
	}

	imageName := fmt.Sprintf("%s@%s", msg.Id[1:5], rule.RuleConfig.SourceName)
	var oName, oUrl string
	
	var taskList []*model.TaskModel
	taskList, err = dal.GetTaskByCondition(ctx, msg.TaskId, msg.RuleId, pb_mani.ExecuteState_execute_running)
	if err != nil {
		logx.Errorf("ExecuteTask GetTaskByCondition error:%v", err)
		return err
	}
	if len(taskList) == 0 {
		logx.Errorf("no task error:%v", err)
		return err
	}
	oName, oUrl, err = util.ExecuteTask(ctx, imageName, rule.RuleConfig.SourceUrl, rule.RuleConfig.DescText, rule.RuleType)

	updateFileds := make(map[string]interface{}, 0)
	updateFileds["operator"] = msg.Operator
	updateFileds["output_name"] = oName
	updateFileds["output_url"] = oUrl
	updateFileds["output_state"] = pb_mani.FileState_file_valid
	if err != nil {
		logx.Errorf("ExecuteTask error:%v", err)
		updateFileds["execute_state"] = pb_mani.ExecuteState_execute_failed
	} else {
		updateFileds["execute_state"] = pb_mani.ExecuteState_execute_success
	}
	if err = dal.UpdateTask(ctx, msg.TaskId, updateFileds); err != nil {
		logx.Errorf("ExecuteTask UpdateTask state:%v", err)
		return err
	}
	return nil
}

func getRule(ctx context.Context, center_rpc pb_mani.GinCenterServiceClient, ruleId string) (*pb_mani.Rule, error) {
	ruleResp, err := center_rpc.GetRuleByCondition(ctx, &pb_mani.GetRuleByConditionReq{
		RuleId: ruleId,
	})
	if err != nil {
		logx.Errorf("getRule GetRuleByCondition error:%v", err)
		return nil, err
	}
	if _, ok := ruleResp.Rules[ruleId]; !ok {
		logx.Errorf("getRule no rule")
		return nil, errors.New("no rule")
	}
	rule := ruleResp.Rules[ruleId]
	if rule.RuleState == pb_mani.RuleState_rule_state_invalid {
		logx.Errorf("getRule rule invalid")
		return nil, errors.New("invalid rule")
	}
	if rule.RuleConfig == nil || rule.RuleConfig.SourceState == pb_mani.FileState_file_invalid {
		logx.Errorf("getRule rule source file invalid")
		return nil, errors.New("file invalid")
	}
	return rule, nil
}

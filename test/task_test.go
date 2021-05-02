package test

import (
	"context"
	"fmt"
	"gin_mani_engine/handler"
	pb_mani "gin_mani_engine/pb"
	"testing"
)

func TestAddTask(t *testing.T) {
	req := &pb_mani.CreateTaskReq{
		Task: &pb_mani.Task{
			RuleId:       "rrr",
			Operator:     "dsadcas",
			ExecuteState: pb_mani.ExecuteState_execute_wait,
			OutputState:  pb_mani.FileState_file_invalid,
		},
	}
	resp, err := handler.CreateTask(context.Background(), req)
	fmt.Println(resp, err)
}

func TestGetTask(t *testing.T) {
	req := &pb_mani.QueryTaskByConditionReq{
		RuleId: "rrr",
	}
	resp, err := handler.QueryTaskByCondition(context.Background(), req)
	fmt.Println(resp, err)
}

func TestUpdateTask(t *testing.T) {
	req := &pb_mani.UpdateTaskReq{
		Task: &pb_mani.Task{
			TaskId: "042c23347588acdd070ff56dbcc5c4d6",
			RuleId: "uuu",
		},
	}
	resp, err := handler.UpdateTask(context.Background(), req)
	fmt.Println(resp, err)
}

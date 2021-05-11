package main

import (
	"context"
	"gin_mani_engine/conf"
	"gin_mani_engine/handler"
	"gin_mani_engine/logic"
	pb_mani "gin_mani_engine/pb"
	logx "github.com/amoghe/distillog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type S struct {
}

func (s S) GetTaskByRules(ctx context.Context, req *pb_mani.GetTaskByRulesReq) (*pb_mani.GetTaskByRulesResp, error) {
	logx.Infoln("GetTaskByRules req: %v", req)
	resp, err := handler.GetTaskByRules(ctx, req)
	logx.Infoln("GetTaskByRules resp: %v", resp)
	return resp, err
}

func (s S) FileUriToServer(ctx context.Context, req *pb_mani.FileUriToServerReq) (*pb_mani.FileUriToServerResp, error) {
	logx.Infoln("FileUriToServer req: %v", req)
	resp, err := handler.FileUriToServer(ctx, req)
	logx.Infoln("FileUriToServer resp: %v", resp)
	return resp, err
}

func (s S) FileUriToCrm(ctx context.Context, req *pb_mani.FileUriToCrmReq) (*pb_mani.FileUriToCrmResp, error) {
	logx.Infoln("FileUriToCrm req: %v", req)
	resp, err := handler.FileUriToCrm(ctx, req)
	logx.Infoln("FileUriToCrm resp: %v", resp)
	return resp, err
}

func (s S) CreateTask(ctx context.Context, req *pb_mani.CreateTaskReq) (*pb_mani.CreateTaskResp, error) {
	logx.Infoln("CreateTask req: %v", req)
	resp, err := handler.CreateTask(ctx, req)
	logx.Infoln("CreateTask resp: %v", resp)
	return resp, err
}

func (s S) QueryTaskByCondition(ctx context.Context, req *pb_mani.QueryTaskByConditionReq) (*pb_mani.QueryTaskByConditionResp, error) {
	logx.Infoln("QueryTaskByCondition req: %v", req)
	resp, err := handler.QueryTaskByCondition(ctx, req)
	logx.Infoln("QueryTaskByCondition resp: %v", resp)
	return resp, err
}

func (s S) UpdateTask(ctx context.Context, req *pb_mani.UpdateTaskReq) (*pb_mani.UpdateTaskResp, error) {
	logx.Infoln("UpdateTask req: %v", req)
	resp, err := handler.UpdateTask(ctx, req)
	logx.Infoln("UpdateTask resp: %v", resp)
	return resp, err
}

func main() {
	cf := conf.GetConfig()
	logx.Infof("start mani engine server")

	logx.Infof("run mani engine executor success...")
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _=range ticker.C {
			err := logic.ExecuteTask(context.Background())
			if err != nil{
				logx.Errorf("ExecuteTask error:%v",err)
			}
		}
	}()

	lis, err := net.Listen("tcp", cf.Server.Port)
	if err != nil {
		log.Fatal("failed to listen")
	}
	server := grpc.NewServer()
	pb_mani.RegisterGinEngineServiceServer(server, &S{})
	reflection.Register(server)
	logx.Infof("run mani engine server success...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

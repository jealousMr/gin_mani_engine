package handler

import (
	"context"
	"errors"
	"gin_mani_engine/logic"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	logx "github.com/amoghe/distillog"
)

func checkFileUriToCrm(req *pb_mani.FileUriToCrmReq) error {
	if len(req.CrmList) == 0 || req.FileAction == pb_mani.FileAction_unknown_file_action {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func FileUriToCrm(ctx context.Context, req *pb_mani.FileUriToCrmReq) (resp *pb_mani.FileUriToCrmResp, err error) {
	resp = &pb_mani.FileUriToCrmResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkFileUriToCrm(req); err != nil {
		logx.Errorf("checkFileUriToCrm error:%v", err)
		return
	}
	urls, err := logic.FileUriToCrm(ctx, req.CrmList, req.FileAction)
	if err != nil {
		logx.Errorf("FileUriToCrm error:%v", err)
		return
	}
	resp.TagUrlMap = urls
	return

}

func checkFileUriToServerParam(req *pb_mani.FileUriToServerReq) error {
	if req.FileAction == pb_mani.FileAction_unknown_file_action || req.Files == nil {
		return errors.New(util.MsgParamError)
	}
	return nil
}

func FileUriToServer(ctx context.Context, req *pb_mani.FileUriToServerReq) (resp *pb_mani.FileUriToServerResp, err error) {
	resp = &pb_mani.FileUriToServerResp{}
	defer func() {
		resp.BaseResp = util.BuildBaseResp(err, "")
	}()
	if err = checkFileUriToServerParam(req); err != nil {
		logx.Errorf("checkFileUriToServerParam error:%v", err)
		return
	}
	fileName, url, err := logic.FileUriToServer(ctx, req)
	if err != nil {
		logx.Errorf("FileUriToServer error:%v", err)
		return
	}
	resp.FileName = fileName
	resp.SaveUrl = url
	return
}

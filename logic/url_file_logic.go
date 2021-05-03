package logic

import (
	"context"
	"errors"
	"fmt"
	"gin_mani_engine/conf"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	logx "github.com/amoghe/distillog"
	"strings"
)

func FileUriToCrm(ctx context.Context, req *pb_mani.FileUriToCrmReq) (string, error) {
	configs := conf.GetConfig()
	var fileName string
	if req.FileName == "" {
		fileName = splitFileName(req.SaveUrl)
	} else {
		fileName = req.FileName
	}
	switch req.FileAction {
	case pb_mani.FileAction_default_all_action:
		return fmt.Sprintf("%s/%s", configs.Router.DefaultAllActionFile, fileName), nil
	case pb_mani.FileAction_open_all_action:
		return fmt.Sprintf("%s/%s", configs.Router.OpenAllActionFile, fileName), nil
	case pb_mani.FileAction_default_image_action:
		return fmt.Sprintf("%s/%s", configs.Router.DefaultImageActionFile, fileName), nil

	}
	return "", errors.New("file type error")
}

func FileUriToServer(ctx context.Context, req *pb_mani.FileUriToServerReq) (fileName, url string, err error) {
	if req.FileName == "" {
		fileName = util.GenUID()[1:6]+".png" // 默认png
	}else{
		fileName = req.FileName
	}
	configs := conf.GetConfig()
	switch req.FileAction {
	case pb_mani.FileAction_default_all_action:
		url = fmt.Sprint("%s/%s",configs.Router.DefaultAllActionFile,fileName)
		err = util.SaveFile(url,req.Files)
		if err != nil{
			logx.Errorf("FileUriToServer save FileAction_default_all_action error:%v",err)
			return "", "", err
		}
		return
	case pb_mani.FileAction_open_all_action:
		url = fmt.Sprint("%s/%s",configs.Router.OpenAllActionFile,fileName)
		err = util.SaveFile(url,req.Files)
		if err != nil{
			logx.Errorf("FileUriToServer save FileAction_open_all_action error:%v",err)
			return "", "", err
		}
		return
	case pb_mani.FileAction_default_image_action:
		url = fmt.Sprint("%s/%s",configs.Router.DefaultImageActionFile,fileName)
		err = util.SaveFile(url,req.Files)
		if err != nil{
			logx.Errorf("FileUriToServer save FileAction_default_image_action error:%v",err)
			return "", "", err
		}
		return
	}
	return "", "", errors.New("none file type")
}

func splitFileName(url string) string {
	route := strings.Split(url, "/")
	return route[len(route)-1]
}

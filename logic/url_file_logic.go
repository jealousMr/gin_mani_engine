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

const Http = "http://localhost:8080"

const (
	DA = "default_all"
	DI = "default_image"
	OA = "open_all"
)

func FileUriToCrm(ctx context.Context, crms []*pb_mani.CrmUrl, action pb_mani.FileAction) (map[string]string, error) {
	var router = fmt.Sprintf("%s/static", Http)
	switch action {
	case pb_mani.FileAction_default_all_action:
		router = fmt.Sprintf("%s/%s",router,DA)
		break
	case pb_mani.FileAction_default_image_action:
		router = fmt.Sprintf("%s/%s",router,DI)
		break
	case pb_mani.FileAction_open_all_action:
		router = fmt.Sprintf("%s/%s",router,OA)
		break
	}
	tagToUrl := make(map[string]string, 0)
	for _, crm := range crms {
		var name = crm.Name
		if crm.CrmType == pb_mani.CrmType_user_crm {
			name = splitFileNameFaceUser(crm.Url)
		}
		if name == "" {
			name = splitFileName(crm.Url)
		}
		tagToUrl[crm.Tag] = fmt.Sprintf("%s/%s", router, name)
	}
	return tagToUrl, nil

}

func FileUriToServer(ctx context.Context, req *pb_mani.FileUriToServerReq) (fileName, url string, err error) {
	if req.FileName == "" {
		fileName = util.GenUID()[1:6] + ".png" // 默认png
	} else {
		fileName = req.FileName
	}
	configs := conf.GetConfig()
	switch req.FileAction {
	case pb_mani.FileAction_default_all_action:
		url = fmt.Sprint("%s/%s", configs.Router.DefaultAllActionFile, fileName)
		err = util.SaveFile(url, req.Files)
		if err != nil {
			logx.Errorf("FileUriToServer save FileAction_default_all_action error:%v", err)
			return "", "", err
		}
		return
	case pb_mani.FileAction_open_all_action:
		url = fmt.Sprint("%s/%s", configs.Router.OpenAllActionFile, fileName)
		err = util.SaveFile(url, req.Files)
		if err != nil {
			logx.Errorf("FileUriToServer save FileAction_open_all_action error:%v", err)
			return "", "", err
		}
		return
	case pb_mani.FileAction_default_image_action:
		url = fmt.Sprint("%s/%s", configs.Router.DefaultImageActionFile, fileName)
		err = util.SaveFile(url, req.Files)
		if err != nil {
			logx.Errorf("FileUriToServer save FileAction_default_image_action error:%v", err)
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

func splitFileNameFaceUser(url string) string {
	route := strings.Split(url, "/")
	dir := route[len(route)-2]
	if dir == DA || dir == DI || dir == OA{
		return route[len(route)-1]
	}
	return fmt.Sprintf("%s/%s", route[len(route)-2], route[len(route)-1])
}

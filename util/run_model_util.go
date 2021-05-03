package util

import (
	"context"
	"fmt"
	"gin_mani_engine/conf"
	pb_mani "gin_mani_engine/pb"
	logx "github.com/amoghe/distillog"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	OutputName       = "1_sf_0_SF.png"
	CommandSHFileUrl = "/home/xyl/src/gin_mani_engine/util/task.sh"

	// py env set
	Env        = "source ~/.bash_profile && source activate gan_env"
	ModelRoute = "cd /home/xyl/src/ManiGAN/code"
	RunCommon  = "python testRun.py"
)

func ExecuteTask(ctx context.Context, imageName, imageUrl, desc string, action pb_mani.RuleType) (oName, oUrl string, err error) {
	descFileUrl := genRandomDesc(action, int64(len(desc)))
	err = SaveFile(descFileUrl, []byte(desc))
	if err != nil {
		logx.Errorf("save desc to file error")
		return "", "", err
	}
	if err := genCommand(imageUrl, descFileUrl, imageName); err != nil {
		logx.Errorf("ExecuteTask genCommand error:%v", err)
		return "", "", err
	}
	shellCommand := fmt.Sprintf("sh %s", CommandSHFileUrl)
	_, err = exec.Command("/bin/bash", "-c", shellCommand).Output()
	if err != nil {
		logx.Errorf("Command running error")
		return "", "", err
	}
	cf := conf.GetConfig()
	switch action {
	case pb_mani.RuleType_default_all:
		oUrl = fmt.Sprintf("%s/%s", cf.Router.DefaultAllActionFile, OutputName)
		break
	case pb_mani.RuleType_default_image:
		oUrl = fmt.Sprintf("%s/%s", cf.Router.DefaultImageActionFile, OutputName)
		break
	case pb_mani.RuleType_open_all:
		oUrl = fmt.Sprintf("%s/%s", cf.Router.OpenAllActionFile, OutputName)
		break
	}
	return OutputName, oUrl, nil
}

func genCommand(imageUrl, descUrl string, imageName string) error {
	param := fmt.Sprintf("%s --source_image_url %s --source_text_url %s --out_name %s", RunCommon, imageUrl, descUrl, imageName)
	command := fmt.Sprintf("%s&&\n%s&&\n%s", Env, ModelRoute, param)
	f, err := os.OpenFile(CommandSHFileUrl, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(command), n)
		defer f.Close()
	}
	return err
}

func genRandomDesc(action pb_mani.RuleType, seed int64) (descFileUrl string) {
	n := rand.Int63n(seed)
	cf := conf.GetConfig()
	descFileName := fmt.Sprintf("%d@%d@%s.txt", time.Now().Unix(), n, "desc_text")
	switch action {
	case pb_mani.RuleType_default_all:
		descFileUrl = fmt.Sprintf("%s/%s", cf.Router.DefaultAllActionFile, descFileName)
		break
	case pb_mani.RuleType_open_all:
		descFileUrl = fmt.Sprintf("%s/%s", cf.Router.OpenAllActionFile, descFileName)
		break
	case pb_mani.RuleType_default_image:
		descFileUrl = fmt.Sprintf("%s/%s", cf.Router.DefaultImageActionFile, descFileName)
		break
	}
	return descFileUrl
}

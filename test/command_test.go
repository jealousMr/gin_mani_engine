package test

import (
	"context"
	"fmt"
	pb_mani "gin_mani_engine/pb"
	"gin_mani_engine/util"
	"testing"
)

func TestEx(t *testing.T) {
	oName, ourl, err := util.ExecuteTask(context.Background(), "image_test.jpg", "/Users/mjea/go/src/gin_mani_engine/test/image_test.jpg", "a red head", pb_mani.RuleType_default_all)
	fmt.Println(oName, ourl, err)
}

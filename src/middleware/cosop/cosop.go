package cosop

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/tencentyun/cos-go-sdk-v5"
)

const (
	DefalutDisplayCount = 20
)

func CosGet(srcname, dstname string) bool {
	AccessKey := cfg.Get_Local("cosaccesskey")
	AccessKeyID := cfg.Get_Local("cosaccesskeyid")
	EndpointPub := cfg.Get_Local("cosendpointpub")
	u, _ := url.Parse(EndpointPub)
	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  AccessKeyID,
			SecretKey: AccessKey,
		},
	})
	if _, err := os.Stat(dstname); err == nil {
		logger.Error("Not download because file exists: ", dstname)
		return false
	}
	rsp, err := client.Object.Head(context.Background(), srcname, nil)

	if err != nil {
		if rsp.StatusCode == 404 {
			logger.Error("file not found on cos...")
		} else {
			logger.Error("get object meta data failed ,err:", err.Error())
		}
		return false
	}
	rsp, err = client.Object.GetToFile(context.Background(), srcname, dstname, nil)
	if err != nil {
		if rsp.StatusCode == 404 {
			logger.Error("download error,file not found,error:", err)
		} else {
			logger.Error("download error,", err)
		}
		return false
	}
	return true
}

package qiniu

// 存储相关功能的引入包只有这两个，后面不再赘述
import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"github.com/qiniu/go-sdk/v7/storage"
	"gvb_server/global"
	rand2 "math/rand"
	"strconv"
	"time"
)

func GetToken() string {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := global.Config.QiNiu.Bucket

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getCfg() storage.Config {
	cfg := storage.Config{}
	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(global.Config.QiNiu.Zone))
	cfg.Zone = &zone
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	return cfg
}

// 上传图片 文件数组，前缀
func UploadImage(data []byte, ImageName string, prefix string) (filepath string, err error) {
	if global.Config.QiNiu.Enable == false {
		return "", errors.New("没有启用七牛云")
	}
	c := global.Config.QiNiu
	if c.AccessKey == "" || c.SecretKey == "" {
		return "", errors.New("请配置key")
	}
	if float64(len(data))/float64(1024*1024) > c.Size {
		return "", errors.New("文件超过设置大小")
	}
	uptoken := GetToken()
	cfg := getCfg()

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	//获取当前时间
	rand := rand2.Int31n(10001)
	now := time.Now().Format("20120304150603")
	key := fmt.Sprintf("%s/%s__%s__%s", prefix, now, strconv.Itoa(int(rand)), ImageName)
	err = formUploader.Put(context.Background(), &ret, uptoken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {

		return "", err
	}
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.CDN, ret.Key), nil
	//fmt.Println(ret.Key, ret.Hash)
}

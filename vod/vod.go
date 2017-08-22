package vod

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"encoding/json"

	"github.com/Fengxq2014/aliyun-signature/signature"
	"github.com/google/go-querystring/query"
	"github.com/goroom/rand"
)

// AliyunVod 公共参数
type AliyunVod struct {
	Format           string //返回值的类型，支持JSON与XML
	Version          string //API版本号，为日期形式：YYYY-MM-DD，本版本对应为2017-03-21
	AccessKeyID      string `url:"AccessKeyId"` //阿里云颁发给用户的访问服务所用的密钥ID
	SignatureMethod  string //签名方式，目前支持HMAC-SHA1
	Timestamp        string //请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ例如，2017-3-29T12:00:00Z(为北京时间2017年3月29日的20点0分0秒
	SignatureVersion string //签名算法版本，目前版本是1.0
	SignatureNonce   string //唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值
}

// NewAliyunVod 初始化一个新的vod client
func NewAliyunVod(accessKeyID string) (*AliyunVod, error) {
	var a AliyunVod
	a.Format = "JSON"
	a.Version = "2017-03-21"
	a.AccessKeyID = accessKeyID
	a.SignatureMethod = "HMAC-SHA1"
	a.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	a.SignatureVersion = "1.0"
	a.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	return &a, nil
}

// PlayAuthResposeEntity PlayAuth返回
type PlayAuthResposeEntity struct {
	RequestID string `json:"RequestId"`
	VideoMeta videoDetail
	PlayAuth  string
}

type videoDetail struct {
	VideoID  string `json:"VideoId"`
	Title    string
	Duration float32
	CoverURL string
	Status   string
}

// GetVideoPlayAuth 获取视频播放凭证
func (avod *AliyunVod) GetVideoPlayAuth(videoID, accessSecret string) (result PlayAuthResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `url:"VideoId"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoPlayAuth", VideoID: videoID}
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, accessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b,&result)
	if err != nil {
		return
	}
	if result.PlayAuth == ""{
		err = fmt.Errorf("%v",string(b))
		return
	}
	return
}

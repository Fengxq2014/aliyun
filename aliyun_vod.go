package aliyun_vod

import (
	"fmt"
	"time"

	"github.com/goroom/rand"
)

// AliyunVod 公共参数
type AliyunVod struct {
	Format           string //返回值的类型，支持JSON与XML
	Version          string //API版本号，为日期形式：YYYY-MM-DD，本版本对应为2017-03-21
	AccessKeyID      string `json:"AccessKeyId"` //阿里云颁发给用户的访问服务所用的密钥ID
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

// GetVideoPlayAuth 获取视频播放凭证
func (avod *AliyunVod) GetVideoPlayAuth(videoID string) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `json:"VideoId"`
	}
	type videoDetail struct{
		VideoID string `json:"VideoId"`
		Title string
		Duration float32
		CoverURL string
		Status string
	}
	type resposeEntity struct{
		RequestID string `json:"RequestId"`
		VideoMeta videoDetail
		PlayAuth string
	}
	
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoPlayAuth", VideoID: videoID}
	fmt.Println(req)
}

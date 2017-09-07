package vod

import (
	"time"

	"github.com/Fengxq2014/aliyun-signature/signature"
	"github.com/Fengxq2014/aliyun/util"
	"github.com/google/go-querystring/query"
	"github.com/goroom/rand"
)

// NewAliyunVod 初始化一个新的vod client
func NewAliyunVod(accessKeyID, accessSecret string) *AliyunVod {
	var a AliyunVod
	a.Format = "JSON"
	a.Version = "2017-03-21"
	a.AccessKeyID = accessKeyID
	a.AccessSecret = accessSecret
	a.SignatureMethod = "HMAC-SHA1"
	a.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	a.SignatureVersion = "1.0"
	return &a
}

// GetVideoPlayAuth 获取视频播放凭证
func (avod *AliyunVod) GetVideoPlayAuth(videoID string) (result PlayAuthResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `url:"VideoId"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoPlayAuth", VideoID: videoID}
	req.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "PlayAuth", nil)
	return
}

// CreateUploadVideo 获取视频上传地址和凭证
func (avod *AliyunVod) CreateUploadVideo(title, fileName, fileSize, description, coverURL, tags string, cateID int64) (result CreateUploadVideoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action      string
		Title       string
		FileName    string
		FileSize    string `url:",omitempty"`
		Description string `url:",omitempty"`
		CoverURL    string `url:",omitempty"`
		CateID      int64  `url:"CateId,omitempty"`
		Tags        string `url:",omitempty"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "CreateUploadVideo", Title: title, FileName: fileName, FileSize: fileSize, Description: description, CoverURL: coverURL, CateID: cateID, Tags: tags}
	req.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = util.GetRespOrError(url, &result, "UploadAuth", nil)
	return
}

// RefreshUploadVideo 刷新视频上传凭证
func (avod *AliyunVod) RefreshUploadVideo(videoID string) (result CreateUploadVideoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `url:"VideoId"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "RefreshUploadVideo", VideoID: videoID}
	req.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = util.GetRespOrError(url, &result, "UploadAuth", nil)
	return
}

// CreateUploadImage 获取图片上传地址和凭证
func (avod *AliyunVod) CreateUploadImage(imageType ImageType, imageExt ImageExt) (result CreateUploadImageResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action    string
		ImageType string
		ImageExt  string `url:",omitempty"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "CreateUploadImage", ImageType: imageType.String(), ImageExt: imageExt.String()}
	req.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = util.GetRespOrError(url, &result, "UploadAuth", nil)
	return
}

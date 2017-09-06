package vod

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"encoding/json"

	"reflect"

	"github.com/Fengxq2014/aliyun-signature/signature"
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
	a.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
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
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = getRespOrError(url, &result, "PlayAuth")
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
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = getRespOrError(url, &result, "UploadAuth")
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
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = getRespOrError(url, &result, "UploadAuth")
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
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = getRespOrError(url, &result, "UploadAuth")
	return
}

func getRespOrError(url string, result interface{}, validName string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return
	}
	v := reflect.ValueOf(result)

	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	validField := v.FieldByName(validName)
	if isEmptyValue(validField) {
		err = fmt.Errorf("%v", string(b))
		return
	}
	return
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	if v.Type() == reflect.TypeOf(time.Time{}) {
		return v.Interface().(time.Time).IsZero()
	}

	return false
}

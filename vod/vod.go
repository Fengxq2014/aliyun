package vod

import (
	"time"

	"github.com/Fengxq2014/aliyun-signature/signature"
	"github.com/Fengxq2014/aliyun/util"
	"github.com/google/go-querystring/query"
	rd "github.com/goroom/rand"
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

// GetPlayInfo 获取视频播放地址
// videoID 视频ID
// formats 视频流格式，多个用逗号分隔，支持格式mp4,m3u8，默认获取所有格式的流,非必填参数，可传""
// authTimeout 播放鉴权过期时间，默认为1800秒，支持设置最小值为1800秒,非必填参数，可传""
func (avod *AliyunVod) GetPlayInfo(videoID, formats, authTimeout string) (result GetPlayInfoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action      string
		VideoID     string `url:"VideoId"`
		Formats     string `url:",omitempty"`
		AuthTimeout string `url:",omitempty"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetPlayInfo", VideoID: videoID, Formats: formats, AuthTimeout: authTimeout}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "VideoBase", nil)
	return
}

// GetVideoPlayAuth 获取视频播放凭证
// videoID 视频ID
func (avod *AliyunVod) GetVideoPlayAuth(videoID string) (result PlayAuthResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `url:"VideoId"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoPlayAuth", VideoID: videoID}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "PlayAuth", nil)
	return
}

// GetVideoInfo 获取视频信息
// videoID 视频ID
func (avod *AliyunVod) GetVideoInfo(videoID string) (result GetVideoInfoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action  string
		VideoID string `url:"VideoId"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoInfo", VideoID: videoID}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "Video", nil)
	return
}

// GetVideoList 获取视频信息列表
// 所有参数均为非必填参数
// status 视频状态，默认获取所有视频，多个可以用逗号分隔，如：Uploading,Normal，取值包括：Uploading(上传中)，UploadFail(上传失败)，UploadSucc(上传完 成)，Transcoding(转码中)，TranscodeFail(转码失败)，Blocked(屏蔽)，Normal(正常)
// startTime CreationTime（创建时间）的开始时间，为开区间(大于开始时间)。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ 例如，2017-01-11T12:00:00Z（为北京时间2017年1月11日20点0分0秒）
// endTime CreationTime的结束时间，为闭区间(小于等于结束时间)。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ 例如，2017-01-11T12:00:00Z（为北京时间2017年1月11日20点0分0秒）
// sortBy 结果排序，范围：CreationTime:Desc、CreationTime:Asc，默认为CreationTime:Desc（即按创建时间倒序）
// cateID 视频分类ID
// pageNo 页号，默认1
// pageSize 可选，默认10，最大不超过100
func (avod *AliyunVod) GetVideoList(status, startTime, endTime, sortBy string, cateID, pageNo, pageSize int) (result GetVideoListResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action    string
		Status    string `url:",omitempty"`
		StartTime string `url:",omitempty"`
		EndTime   string `url:",omitempty"`
		SortBy    string `url:",omitempty"`
		CateID    int    `url:"CateId,omitempty"`
		PageNo    int    `url:",omitempty"`
		PageSize  int    `url:",omitempty"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "GetVideoList", Status: status, StartTime: startTime, EndTime: endTime, SortBy: sortBy, CateID: cateID, PageNo: pageNo, PageSize: pageSize}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "RequestID", nil)
	return
}

// UpdateVideoInfo 修改视频信息
// videoId 视频ID
// title 视频标题，长度不超过128个字节，UTF8编码,非必填
// description 视频描述，长度不超过1024个字节，UTF8编码,非必填
// coverURL 视频封面URL地址,非必填
// tags 视频标签，单个标签不超过32字节，最多不超过16个标签。多个用逗号分隔，UTF8编码,非必填
// cateID 视频分类ID,非必填
func (avod *AliyunVod) UpdateVideoInfo(videoID, title, description, coverURL, tags string, cateID int) (result UpdateVideoInfoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action      string
		VideoID     string `url:"VideoId"`
		Title       string `url:",omitempty"`
		Description string `url:",omitempty"`
		CoverURL    string `url:",omitempty"`
		CateID      int    `url:"CateId,omitempty"`
		Tags        string `url:",omitempty"`
	}
	req := requestEntity{AliyunVod: *avod, Action: "UpdateVideoInfo", VideoID: videoID, Title: title, Description: description, CoverURL: coverURL, Tags: tags, CateID: cateID}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "RequestID", nil)
	return
}

// DeleteVideo 删除视频
// videoIds 视频ID列表，多个用逗号分隔，最多支持10个
func (avod *AliyunVod) DeleteVideo(videoIds string) (result DeleteVideoResposeEntity, err error) {
	type requestEntity struct {
		AliyunVod
		Action   string
		VideoIds string
	}
	req := requestEntity{AliyunVod: *avod, Action: "DeleteVideo", VideoIds: videoIds}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "RequestID", nil)
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
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
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
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
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
	req := requestEntity{AliyunVod: *avod, Action: "CreateUploadImage", ImageType: string(imageType), ImageExt: string(imageExt)}
	rand := rd.GetRand()
	req.SignatureNonce = rand.String(16, rd.RST_NUMBER|rd.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, avod.AccessSecret, "http://vod.cn-shanghai.aliyuncs.com")

	err = util.GetRespOrError(url, &result, "UploadAuth", nil)
	return
}

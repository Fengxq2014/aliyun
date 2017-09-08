# aliyun
阿里云SDK

## 阿里云视频点播SDK

### 获取视频播放地址
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.GetPlayInfo(videoID,"","")
```

### 获取视频播放凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.GetVideoPlayAuth(videoID)
```

### 获取视频信息
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp, err := vod.GetVideoInfo(videoID)
```

### 获取视频信息列表
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp, err := vod.GetVideoList("", "", "", "", 0, 0, 0)
```

### 修改视频信息
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp, err := vod.UpdateVideoInfo(videoID, "", "describ", "", "", 0)
```

### 删除视频
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp, err := vod.DeleteVideo(videoIDS)
```
videoIDS:视频ID列表，多个用逗号分隔，最多支持10个

### 获取视频上传地址和凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.CreateUploadVideo(title, fileName, fileSize, description, coverURL, tags , cateID)
```
fileSize,description,coverURL,tags,cateID 非必填参数，可传""

### 刷新视频上传凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.RefreshUploadVideo(videoID)
```

### 获取图片上传地址和凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.CreateUploadImage(imageType, imageExt)
```
imageType:Cover/Watermark
imageExt:Png/Jpg/Jpeg

### 上传文件
请使用官方sdk[Alibaba Cloud OSS SDK for Go](https://github.com/aliyun/aliyun-oss-go-sdk)

## 阿里云短信SDK

### 短信发送
```golang
sms := NewAliyunSms("testkey", "testaccess")
resp, err := sms.SendSms("1*********", "短信签名", "短信模板ID", `{"number":"123"}`, "", "")
```

### 短信查询
```golang
sms := NewAliyunSms("testkey", "testaccess")
resp, err := sms.QuerySendDetails("1*********", "", "20170907", 10, 1)
```
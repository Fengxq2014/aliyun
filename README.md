# aliyun_vod
阿里云SDK

## 阿里云视频SDK

### 获取视频播放凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.GetVideoPlayAuth(videoID)
```

### 获取视频上传地址和凭证
```golang
vod:= NewAliyunVod("testkey","testaccess")
resp,err:=vod.CreateUploadVideo(title, fileName, fileSize, description, coverURL, tags , cateID)
```
fileSize,description,coverURL,tags,cateID 非必需参数，可传""

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
# aliyun
阿里云SDK

## 安装
```
go get github.com/Fengxq2014/aliyun/...
```

## 阿里云视频点播SDK

### 获取视频播放地址
```golang
vod:= NewAliyunVod("testkey","testaccess")
// videoID 视频ID
// formats 视频流格式，多个用逗号分隔，支持格式mp4,m3u8，默认获取所有格式的流,非必填参数，可传""
// authTimeout 播放鉴权过期时间，默认为1800秒，支持设置最小值为1800秒,非必填参数，可传""
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
// 所有参数均为非必填参数
// status 视频状态，默认获取所有视频，多个可以用逗号分隔，如：Uploading,Normal，取值包括：Uploading(上传中)，UploadFail(上传失败)，UploadSucc(上传完 成)，Transcoding(转码中)，TranscodeFail(转码失败)，Blocked(屏蔽)，Normal(正常)
// startTime CreationTime（创建时间）的开始时间，为开区间(大于开始时间)。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ 例如，2017-01-11T12:00:00Z（为北京时间2017年1月11日20点0分0秒）
// endTime CreationTime的结束时间，为闭区间(小于等于结束时间)。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ 例如，2017-01-11T12:00:00Z（为北京时间2017年1月11日20点0分0秒）
// sortBy 结果排序，范围：CreationTime:Desc、CreationTime:Asc，默认为CreationTime:Desc（即按创建时间倒序）
// cateID 视频分类ID
// pageNo 页号，默认1
// pageSize 可选，默认10，最大不超过100
resp, err := vod.GetVideoList("", "", "", "", 0, 0, 0)
```

### 修改视频信息
```golang
vod:= NewAliyunVod("testkey","testaccess")
// videoId 视频ID
// title 视频标题，长度不超过128个字节，UTF8编码,非必填
// description 视频描述，长度不超过1024个字节，UTF8编码,非必填
// coverURL 视频封面URL地址,非必填
// tags 视频标签，单个标签不超过32字节，最多不超过16个标签。多个用逗号分隔，UTF8编码,非必填
// cateID 视频分类ID,非必填
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
// phoneNumbers 短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
// signName 短信签名
// templateCode 短信模板ID
// templateParam 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败,非必填参数，可传""
// smsUpExtendCode 上行短信扩展码,无特殊需要此字段的用户请忽略此字段,非必填参数，可传""
// outID 外部流水扩展字段,非必填参数，可传""
resp, err := sms.SendSms("1*********", "短信签名", "短信模板ID", `{"number":"123"}`, "", "")
```

### 短信查询
```golang
sms := NewAliyunSms("testkey", "testaccess")
// phoneNumber 短信接收号码,如果需要查询国际短信,号码前需要带上对应国家的区号
// bizID 发送流水号,从调用发送接口返回值中获取
// sendDate 短信发送日期格式yyyyMMdd,支持最近30天记录查询
// pageSize 页大小Max=50
// currentPage 当前页码
resp, err := sms.QuerySendDetails("1*********", "", "20170907", 10, 1)
```
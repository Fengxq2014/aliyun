package sms

import (
	"time"

	"github.com/Fengxq2014/aliyun-signature/signature"
	"github.com/Fengxq2014/aliyun/util"
	"github.com/google/go-querystring/query"
	"github.com/goroom/rand"
)

// NewAliyunSms 初始化一个新的sms client
func NewAliyunSms(accessKeyID, accessSecret string) *AliyunSms {
	var a AliyunSms
	a.Format = "JSON"
	a.Version = "2017-05-25"
	a.AccessKeyID = accessKeyID
	a.AccessSecret = accessSecret
	a.SignatureMethod = "HMAC-SHA1"
	a.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	a.SignatureVersion = "1.0"
	return &a
}

// SendSms 短信发送
// phoneNumbers 短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式
// signName 短信签名
// templateCode 短信模板ID
// templateParam 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,否则会导致JSON在服务端解析失败,非必填参数，可传""
// smsUpExtendCode 上行短信扩展码,无特殊需要此字段的用户请忽略此字段,非必填参数，可传""
// outID 外部流水扩展字段,非必填参数，可传""
func (asms *AliyunSms) SendSms(phoneNumbers, signName, templateCode, templateParam, smsUpExtendCode, outID string) (result SendSmsResposeEntity, err error) {
	type requestEntity struct {
		AliyunSms
		Action          string
		RegionID        string `url:"RegionId"`
		PhoneNumbers    string
		SignName        string
		TemplateCode    string
		TemplateParam   string `url:",omitempty"`
		SmsUpExtendCode string `url:"smsUpExtendCode,omitempty"`
		OutID           string `url:"OutId,omitempty"`
	}
	req := requestEntity{AliyunSms: *asms, Action: "SendSms", RegionID: "cn-hangzhou", PhoneNumbers: phoneNumbers, SignName: signName, TemplateCode: templateCode, TemplateParam: templateParam, SmsUpExtendCode: smsUpExtendCode, OutID: outID}
	req.SignatureNonce = rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, asms.AccessSecret, "http://dysmsapi.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "Code", "OK")
	return
}

// QuerySendDetails 短信查询
// phoneNumber 短信接收号码,如果需要查询国际短信,号码前需要带上对应国家的区号
// bizID 发送流水号,从调用发送接口返回值中获取
// sendDate 短信发送日期格式yyyyMMdd,支持最近30天记录查询
// pageSize 页大小Max=50
// currentPage 当前页码
func (asms *AliyunSms) QuerySendDetails(phoneNumber,bizID,sendDate string,pageSize,currentPage int)(result QuerySendDetailsResposeEntity,err error) {
	type reequestionEntity struct {
		AliyunSms
		Action          string
		PhoneNumber string
		BizID       string `url:"BizId,omitempty"`
		SendDate    string
		PageSize    int
		CurrentPage int
	}
	req:=reequestionEntity{AliyunSms:*asms,Action:"QuerySendDetails", PhoneNumber:phoneNumber,BizID:bizID,SendDate:sendDate,PageSize:pageSize,CurrentPage:currentPage}
	req.SignatureNonce=rand.String(16, rand.RST_NUMBER|rand.RST_LOWER)
	v, _ := query.Values(req)
	url := signature.ComposeURL(v, asms.AccessSecret, "http://dysmsapi.aliyuncs.com")
	err = util.GetRespOrError(url, &result, "Code", "OK")
	return
}

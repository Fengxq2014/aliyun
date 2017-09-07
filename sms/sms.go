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

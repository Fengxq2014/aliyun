package sms

import (
	"testing"
)

var sms = NewAliyunSms("testkey", "testaccess")

func TestSendSms(t *testing.T) {
	resp, err := sms.SendSms("1*********", "短信签名", "短信模板ID", `{"number":"123"}`, "", "")
	if err != nil {
		t.Errorf("SendSms error:%v,resp:%v", err, resp)
	}
}

func TestDysmsapi(t *testing.T) {
	resp, err := sms.QuerySendDetails("1*********", "", "20170907", 10, 1)
	if err != nil {
		t.Errorf("SendSms error:%v,resp:%v", err, resp)
	}
}

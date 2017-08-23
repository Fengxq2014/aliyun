package vod

import(
	"testing"
)

func TestGetVideoPlayAuth(t *testing.T){
	vod,_:= NewAliyunVod("testkey","testaccess")
	resp,err:=vod.GetVideoPlayAuth("93ab850b4f6f44eab54b6e91d24d81d4")
	if err != nil{
		t.Errorf("resp:%v,err:%v",resp,err)
	}
}
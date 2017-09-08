package vod

import (
	"testing"
)

var vod = NewAliyunVod("testkey", "testaccess")

func TestGetVideoPlayAuth(t *testing.T) {
	resp, err := vod.GetVideoPlayAuth("93ab850b4f6f44eab54b6e91d24d81d4")
	if err != nil {
		t.Errorf("resp:%v,err:%v", resp, err)
	}
}

func TestGetPlayInfo(t *testing.T) {
	resp, err := vod.GetPlayInfo("93ab850b4f6f44eab54b6e91d24d81d4", "", "")
	if err != nil {
		t.Errorf("resp:%v,err:%v", resp, err)
	}
}

func TestGetVideoInfo(t *testing.T) {
	resp, err := vod.GetVideoInfo("93ab850b4f6f44eab54b6e91d24d81d4")
	if err != nil {
		t.Errorf("resp:%v,err:%v", resp, err)
	}
}

func TestGetVideoList(t *testing.T) {
	resp, err := vod.GetVideoList("", "", "", "", 0, 0, 0)
	if err != nil {
		t.Errorf("resp:%v,err:%v", resp, err)
	}
}

func TestUpdateVideoInfo(t *testing.T) {
	resp, err := vod.UpdateVideoInfo("93ab850b4f6f44eab54b6e91d24d81d4", "", "describ", "", "", 0)
	if err != nil {
		t.Errorf("resp:%v,err:%v", resp, err)
	}
}

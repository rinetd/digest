// Copyright 2013 M-Lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The digest package provides an implementation of http.RoundTripper that takes
// care of HTTP Digest Authentication (http://www.ietf.org/rfc/rfc2617.txt).
// This only implements the MD5 and "auth" portions of the RFC, but that covers
// the majority of avalible server side implementations including apache web
// server.
//https://en.wikipedia.org/wiki/Digest_access_authentication
package digest

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

const (
	host           = "http://admin:huawei777@192.168.0.108"
	getSystemInfo  = "http://192.168.0.108/cgi-bin/magicBox.cgi?action=getSystemInfo"
	GetBasicConfig = "http://192.168.0.108/cgi-bin/configManager.cgi?action=getConfig"
	snapshot       = "http://admin:huawei777@192.168.0.108/cgi-bin/snapshot.cgi?channel=0"
	getCaps        = "/cgi-bin/devVideoDetect.cgi?action=getCaps&channel=0"
	nas            = "/cgi-bin/configManager.cgi?action=getConfig&name=NAS"

	GetStorageDeviceCollect = "/cgi-bin/storageDevice.cgi?action=factory.getCollect"
	// 10.2.2 getDeviceAllInfo
	getDeviceAllInfo = "/cgi-bin/storageDevice.cgi?action=getDeviceAllInfo"
	// 10.4.1 GetWorkDirectoryCollect
	GetWorkDirectoryCollect = "/cgi-bin/workDirectory.cgi?action=factory.getCollect"

	// 9.System
	// 9.1.1
	GetGeneralConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=General"
	// 9.1.2 SetGeneralConfig /cgi-bin/configManager.cgi?action=setConfig&
	GetLanguageCaps = "/cgi-bin/magicBox.cgi?action=getLanguageCaps"

	GetUserGlobalConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=UserGlobal"
	GetDeviceType       = "/cgi-bin/magicBox.cgi?action=getDeviceType"
	reboot              = "/cgi-bin/magicBox.cgi?action=reboot"
	shutdown            = "/cgi-bin/magicBox.cgi?action=shutdown"

	GetVideo            = "/cgi-bin/mjpg/video.cgi?channel=0&subtype=0"
	GetStream           = "/cgi-bin/realmonitor.cgi?action=getStream&channel=1&subtype=0"
	GetVideoColorConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=VideoColor"
	GetVideoInputCaps   = "/cgi-bin/devVideoInput.cgi?action=getCaps&channel=0"
	// 4.4.3 GetVideoEncodeConfig
	GetVideoEncodeConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=Encode"

	SetVideoEncodeConfig = "/cgi-bin/configManager.cgi?action=setConfig&head.Video.FPS=16"

	//
	GetEventHandler = "/cgi-bin/configManager.cgi?action=getConfig&name=MotionDetect"
	SetEventHandler = "/cgi-bin/configManager.cgi?action=setConfig&MotionDetect[0].EventHandler.AlarmOutChannels[0]=8"
	// 6.2.1
	GetAlarmConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=Alarm"

	// 6.2.5 GetInSlots
	GetInSlots  = "/cgi-bin/alarm.cgi?action=getInSlots"
	GetOutSlots = "/cgi-bin/alarm.cgi?action=getOutSlots"
	//6.2.7 GetInState
	GetInState = "/cgi-bin/alarm.cgi?action=getInState"
	//6.2.8 GetOutState
	GetOutState = "/cgi-bin/alarm.cgi?action=getOutState"

	GetChannelInState  = "/cgi-bin/alarm.cgi?action=getInState&channel=0"
	GetChannelOutState = "/cgi-bin/alarm.cgi?action=getOutState&channel=1"
	//6.3.1 GetMotionDetectConfig
	GetMotionDetectConfig = "/cgi-bin/configManager.cgi?action=getConfig&name=MotionDetect"

	//6.9 GetEventIndexes [VideoMotion : VideoLoss : VideoBlind : AlarmLocal]
	GetEventIndexes = "/cgi-bin/eventManager.cgi?action=getEventIndexes&code=VideoMotion"
)

func TestG(t *testing.T) {
	// resp, err := Get("http://admin:huawei777@192.168.0.108/cgi-bin/magicBox.cgi?action=getSystemInfo")
	// resp, err := Get(host + SetEventHandler)
	resp, err := Get(host + GetEventIndexes)
	if err != nil {
		// return nil, err
	}
	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll((resp.Body))
	log.Println(string(body))
}

// func TestGet1(t *testing.T) {
// 	ts := NewTransport("admin", "huawei777")
// 	c, err := ts.Client()
// 	if err != nil {
// 		return
// 	}

// 	resp, err := c.Get(GetBasicConfig)
// 	if err != nil {
// 		return
// 	}
// 	bb, _ := ioutil.ReadAll(resp.Body)
// 	log.Printf("%s", bb)
// 	return
// }

// func TestGet(t *testing.T) {
// 	ts := NewTransport("admin", "huawei777")
// 	req, err := http.NewRequest("GET", "http://192.168.0.108/cgi-bin/magicBox.cgi?action=getSystemInfo", nil)
// 	if err != nil {
// 		// return err
// 	}
// 	resp, err := ts.RoundTrip(req)
// 	if err != nil {
// 		// return err
// 	}

// 	// resp, _ := ts.Do(req)
// 	log.Println(resp.Header)
// 	// var buf []byte
// 	// resp.Body.Read(buf)
// 	// 获取请求报文的内容长度
// 	len := resp.ContentLength

// 	// 新建一个字节切片，长度与请求报文的内容长度相同
// 	body := make([]byte, len)
// 	resp.Body.Read(body)

// 	log.Printf("+%s", body)
// }

var cred = &credentials{
	Username:   "Mufasa",
	Realm:      "testrealm@host.com",
	Nonce:      "dcd98b7102dd2f0e8b11d0f600bfb0c093",
	DigestURI:  "/dir/index.html",
	Algorithm:  "MD5",
	Opaque:     "5ccc069c403ebaf9f0171e9517f40e41",
	MessageQop: "auth",
	method:     "GET",
	password:   "Circle Of Life",
}

var cnonce = "0a4f113b"

func TestH(t *testing.T) {
	r1 := h("Mufasa:testrealm@host.com:Circle Of Life")
	if r1 != "939e7578ed9e3c518a452acee763bce9" {
		t.Fail()
	}

	r2 := h("GET:/dir/index.html")
	if r2 != "39aff3a2bab6126f332b942af96d3366" {
		t.Fail()
	}

	r3 := h(fmt.Sprintf("%s:dcd98b7102dd2f0e8b11d0f600bfb0c093:00000001:0a4f113b:auth:%s", r1, r2))
	if r3 != "6629fae49393a05397450978507c4ef1" {
		t.Fail()
	}
}

func TestKd(t *testing.T) {
	r1 := kd("939e7578ed9e3c518a452acee763bce9",
		"dcd98b7102dd2f0e8b11d0f600bfb0c093:00000001:0a4f113b:auth:39aff3a2bab6126f332b942af96d3366")
	if r1 != "6629fae49393a05397450978507c4ef1" {
		t.Fail()
	}
}

func TestHa1(t *testing.T) {
	r1 := cred.ha1()
	if r1 != "939e7578ed9e3c518a452acee763bce9" {
		t.Fail()
	}
}

func TestHa2(t *testing.T) {
	r1 := cred.ha2()
	if r1 != "39aff3a2bab6126f332b942af96d3366" {
		t.Fail()
	}
}

func TestResp(t *testing.T) {
	r1, err := cred.resp(cnonce)
	if err != nil {
		t.Fail()
	}
	if r1 != "6629fae49393a05397450978507c4ef1" {
		t.Fail()
	}
}

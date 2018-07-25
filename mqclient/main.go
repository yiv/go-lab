package main

import (
	"encoding/json"
	"git.spotmau.cn/cloud/bastet/protocol"
	"git.spotmau.cn/cloud/bastet/utils/idutil"
	"git.spotmau.cn/cloud/bastet/utils/stringutil"
	"git.spotmau.cn/cloud/data-analyzer/common"
	"time"
)

var (
	mdbDeviceID  = "97d7fab1593bd9f7f50a4b6556823cd4"
	mdbProductID = 2603
)

//TestMdbCapImgEvent 测试抓拍照片事件
func main() {
	// fmt.Println("*1", time.Now())
	var capInfos [1]interface{}
	capInfosArr0 := make(map[string]interface{})
	capInfosArr0["time"] = time.Now().Unix()
	capInfosArr0["url"] = "test_capImg.jpg"
	capInfosArr0["local"] = capInfosArr0["url"]
	capInfos[0] = capInfosArr0

	infos := make(map[string]interface{})
	infos["type"] = 1
	infos["capinfos"] = capInfos

	doMdbEvent("capt_img_ntf", infos)
}

func doMdbEvent(action string, infos map[string]interface{}) {
	ctime := time.Now().Unix()

	eventValue := make(map[string]interface{})
	eventValue["thread"] = stringutil.UniqueStrings(16, 1)[0]
	eventValue["pdt_id"] = mdbProductID
	eventValue["ctime"] = int(ctime)
	eventValue["level"] = 0
	eventValue["version"] = 1
	eventValue["infos"] = infos

	eventsArr0 := make(map[string]interface{})
	eventsArr0["event_Name"] = action
	eventsArr0["event_value"] = eventValue

	payload := make(map[string]interface{})
	var events [1]map[string]interface{}
	events[0] = eventsArr0

	payload["events"] = events
	payload["seri_no"] = 1
	extendMap := make(map[string]interface{})
	extendMap["device_ip"] = "10.10.0.70"
	extendMap["bind_users"] = "32754"
	extendMsg, _ := json.Marshal(extendMap)
	msgBody, _ := json.Marshal(payload)
	bizMsg := &protocol.Message{
		MqHeader: &protocol.MqHeader{
			ProtoType: protocol.PROTOCOL_COAP,
			ProtoVer:  protocol.PROTOCOL_VERSION_3,
			Ctime:     int(time.Now().Unix()),
			Extends:   string(extendMsg[:]),
			FromTo:    protocol.DEV2SYS,
		},
		MsgHeader: &protocol.MsgHeader{
			MsgId:       idutil.GetAutoIncID().AtomicInc(),
			MsgType:     protocol.MESSAGE_NON,
			ServiceName: "/dev/report/event",
			From:        mdbDeviceID,
			FromType:    protocol.OBJECT_DEVICE,
			To:          protocol.SYS_OBJ_ID_EVENT_RECR,
			ToType:      protocol.OBJECT_SYSTEM,
		},
		MsgBody: msgBody,
	}
	mqMsg, _ := protocol.Marshal(bizMsg)
	queue := common.GetQueueName("event")
	mqClient := common.GetIntraMqClient()
	for {
		mqClient.EnqueueMQ(queue, mqMsg)
		time.Sleep(time.Millisecond * 500)
	}
}

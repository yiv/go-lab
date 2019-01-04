package service

import (
	"fmt"
	"hunter/GameLogic/dao"
	"hunter/GameLogic/entity"
	"hunter/GameLogic/httpclient"
	"hunter/GameLogic/util"
	"hunter/lib/logger"
	"hunter/lib/parsebson"
	"hunter/lib/random"
	"sort"
	"strconv"
	"time"

	"labix.org/v2/mgo/bson"
)

const WT01_RANDLE_BASE = 100000000

func room_WT01sceneinit(room *entity.Room, scene *SceneWT01) {
	//初始化游戏
	room.SceneInfoId = room.TickTime.Unix() + int64(room.Id)*10000000000

	scene.StartTime = time.Now().Unix()
	scene.State = E_GUESS_DIGIT_STATE_BEGIN
	scene.PreUpdate_time = time.Now().Unix()

	//从数据库加载配置表
	WT01_LoadBetxInfoFromDb()
	WT01_LoadRtpInfoFromDb()
	WT01_LoadDrillRtpInfoFromDb()
	WT01_LoadWinTreasureInfoFromDb()

	//betx := WT01_GetLevelStoneLinkBetx(1,5,4)
	//logger.Notice("betx:",betx)

	// BetTest()

	logger.Notice("SceneWT01.game begin .sceneId:", room.SceneInfoId)

}

type SceneWT01 struct {
	StartTime      int64  //游戏开始时间
	State          int    //游戏状态
	GameResult     string //游戏结果
	PreUpdate_time int64  //上次刷新时间

}

func (scene *SceneWT01) String() string {
	str := fmt.Sprintf("State:%d, StartTime:%d ,GameResult:%s ",
		scene.State, scene.StartTime, scene.GameResult)
	return str
}

func (scene *SceneWT01) CreateGameResult(room *entity.Room) {

}

// 游戏开始
func (scene *SceneWT01) Start(room *entity.Room, sceneNext *entity.SceneNext) {
	//load config from db
	room_WT01sceneinit(room, scene)

	//test(room,scene)
}

func (scene *SceneWT01) Stop(room *entity.Room) (sceneNext *entity.SceneNext) {
	logger.Notice("SceneWT01.Stop.Room:", room)

	sceneNext = new(entity.SceneNext)

	sceneNext.SceneType = util.SCENE_TYPE_WT01

	sceneNext.Data = map[string]interface{}{
		"SceneWT01Id": room.SceneInfoId,
	}

	return sceneNext
}

func (scene *SceneWT01) Update(room *entity.Room) {
	if room.RoomState == entity.ROOM_STATE_CLOSE_WAITING {
		roomTick_tryStop(room)
		return
	}

	//60S钟检查更新，保存一次RoomTreasurbag数据
	nowTime := time.Now().Unix()
	if nowTime-room.PreSaveTime >= 60 {
		room.PreSaveTime = nowTime
		Room_SaveRoomTreasuryBag(room)
		//CheckUserLastHitTime(room,time.Now().Unix())
	}

}

func (scene *SceneWT01) GetSelf() interface{} {
	return scene
}

func (scene *SceneWT01) GetInfo(room *entity.Room, userId uint32) map[string]interface{} {

	pkgMap := bson.M{
		"op": util.OP_ROOM_GAME_INFO_GET_BACK,
	}

	pkgMap["cmd"] = util.CMD_MULTI

	pkgMap["GameId"] = util.WT01
	pkgMap["state"] = scene.State
	pkgMap["to"] = []uint32{userId}

	pkgMap["sceneInfo"] = WT01_GetSceneInfo(room, scene, userId)

	return pkgMap
}

func (scene *SceneWT01) GetUserInfo(room *entity.Room, userId uint32) map[string]interface{} {
	info := bson.M{}

	return info
}

func (scene *SceneWT01) GetExInfo(room *entity.Room, param bson.M) map[string]interface{} {
	info := bson.M{}

	return info
}

//入桌检查
func (scene *SceneWT01) EnterCheck(room *entity.Room, user *entity.UserInfo, pkgMap bson.M) int {

	return 0
}

func (scene *SceneWT01) Enter(room *entity.Room, user *entity.UserInfo, pkgMap bson.M) {
	logger.Notice("SceneWT01.Enter.room:", room, "user:", user)

	//加载用户的场景数据
	userExInfo := dao.WT01_User_fetchById(user.Id)
	if userExInfo == nil {
		userExInfo = LoadUserSceneInfoFromDb(user.Id, room.GameId)

	}
	dao.WT01_User_setUser(userExInfo)
	logger.Notice("SceneWT01.Enter...userExInfo:", userExInfo)

	sceneInfoMap := scene.GetInfo(room, user.Id)
	util.Pkg2Client(room.Ctx, sceneInfoMap)

	//下发jackpot数据
	treasureRate := 0.0
	if jpRoomConfig.Treasure_bag_Info.All_Counter > 0 {
		treasureRate = float64(room.TreasuryBag.AllCounter) / float64(jpRoomConfig.Treasure_bag_Info.All_Counter)
	}
	jackpotbb_notify_jpbb2user(room.Ctx, room.GameId, []uint32{user.Id}, jpbbRoomprize.JpBBs, treasureRate)

}

func (scene *SceneWT01) Leave(room *entity.Room, user *entity.UserInfo) {
	if !user.Disconnect { //退出离开
		//保存场景数据
		userEx := dao.WT01_User_fetchById(user.Id)
		if nil != userEx {
			dao.WT01_User_remove(userEx)
		} else {
			logger.Error("service.Leave...Room_removeUser.not find userEx info.userId:", user.Id)
		}
		logger.Notice("SceneWT01.Leave.room:", room, "user:", user)
	}

}

func (scene *SceneWT01) Pause(room *entity.Room) {
}

func (scene *SceneWT01) Restore(room *entity.Room) {
}

func (scene *SceneWT01) DealJetton(color int, jetton float64) {
}

func WT01_BetInfo_SliceMap2SliceStruct(BetSliceMap []map[string]interface{}) (betSlice []entity.WT01_Bet_Info_Struct, ok bool) {
	for _, v := range BetSliceMap {
		var item entity.WT01_Bet_Info_Struct

		BetJetton, _ := parsebson.ToFloat64(v, "BetJetton")
		item.BetJetton = BetJetton

		betSlice = append(betSlice, item)
	}

	ok = true
	return
}

func GetWT01PayOut(room *entity.Room, scene *SceneWT01, m_nResult map[int]int, m_nBetxResult int, m_nIndex int, betSlice []entity.BB_Bet_Info_Struct) (fPayOut float64) {

	return fPayOut
}

// 下注
func (scene *SceneWT01) Bet(room *entity.Room, userId uint32, sliceMap []map[string]interface{}, dataMap map[string]interface{}) (ret int) {
	var state int

	retPkg := bson.M{
		"op": util.OP_WT01_BET_BACK,
	}

	user := User_getByIdTest(userId, dataMap)
	if user == nil {
		logger.Error("SceneWT01.Scene_Bet. user:", user.Id, " is nil")
		return entity.E_WINTREASURE_USER_NIL
	}

	userEx := dao.WT01_User_fetchByIdTest(userId, dataMap)
	if nil == userEx {
		logger.Error("SceneWT01.Scene_Bet. user:", user.Id, " userExinfo is nil")
		return entity.E_WINTREASURE_USER_NIL
	}

	if userEx.Level > 3 {
		logger.Error("SceneWT01.Scene_Bet. user:", user.Id, " level:", userEx.Level)
		return entity.E_WINTREASURE_NOBET
	}

	betSlice, _ := WT01_BetInfo_SliceMap2SliceStruct(sliceMap)

	betCash := 0.0
	for _, v := range betSlice { // 玩家下注
		betCash += v.BetJetton
	}

	if util.FloatCompare(betCash, 0.0) == 0 {
		logger.Error("SceneWT01.Scene_Bet. user:", user.Id, " betCash :", betCash)
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2ClientTest(room.Ctx, retPkg, dataMap)
		return entity.E_WINTREASURE_PARAM
	}

	if util.FloatCompare(user.Cash+userEx.Point, betCash) == -1 {
		logger.Error("SceneWT01.Scene_Bet. user:", user.Id, " cash less:", user.Cash, "bet jetton:", betCash)
		retPkg["ret"] = entity.E_WINTREASURE_CASH_LESS
		util.Pkg2ClientTest(room.Ctx, retPkg, dataMap)
		return entity.E_WINTREASURE_CASH_LESS
	}

	// 跑数据
	logger.Notice("START---------------------------")
	//下注处理，扣除玩家的下注金额
	for index := 1; index < 2; index++ {
		userEx.Level = index
		for i := 0; i < 200000000; i++ {
			WT01_CreateWinTreasureResult(room, userEx, betSlice)
		}
	}
	logger.Notice("END---------------------------")

	//jackpot
	nowTimeStamp := time.Now().Unix()
	{
		if util.FloatCompare(betCash, 100.0) >= 0 { //100以上的才下注
			bb := int(betCash) / entity.Default_JP_FACTOR
			handle_jpbbTest(room.Ctx, nowTimeStamp, user, room, bb, util.SCENE_TYPE_WT01, dataMap)

			add_treasury_bag(userId, room, bb)
			if check_upload_treasury_bag(room) {
				handle_treasurybbTest(room.Ctx, userId, nowTimeStamp, room, util.SCENE_TYPE_WT01, dataMap)
			}
		}
	}
	payout, drill, betRetSlice := WT01_CreateWinTreasureResult(room, userEx, betSlice)

	// 生成订单号
	billNo := Bill_makeBillNo2("WT01", time.Now().Unix(), userId, 0)

	// 下注信息发送到http，在计算出结果后发送到钱包
	if util.FloatCompare(userEx.Point, betCash) == -1 {
		cash := betCash - userEx.Point
		betItems := make([]httpclient.BetItemInfo, 0)

		logger.Notice("[dongDebug] SessionToken", user.SessionToken)
		betInfo := httpclient.BetItemInfo{
			Type:            "bonus",
			Remarks:         "",
			TransactionId:   billNo,
			TransactionTime: time.Now().Format("2006-01-02 15:04:05"),
			BetAmount:       float64(int(cash)),
			Currency:        user.Currency,
			ResultAmount:    0.0,
			RoundId:         billNo,
			RoundClosed:     true,
		}
		betItems = append(betItems, betInfo)

		sendData := httpclient.BetSendInfo{
			Request:              "bet-and-result",
			GameId:               user.GameId,
			AccountId:            user.AccountId,
			SessionToken:         user.SessionToken,
			ValidateSessionToken: true,
			Items:                betItems,
			Mobile:               true,
			Ip:                   "",
		}

		count := 0
		for count < 2 {
			err := httpclient.BetAndResult(sendData)
			if err == httpclient.RET_ERROR {
				logger.Error("Send Http Error :", sendData, count)
				state = 1
				count += 1
			} else {
				state = 0
				break
			}
		}
		if 1 == state {
			logger.Error("SceneWT01.Scene_Bet error", user.Id, " billNo:", billNo)
			retPkg["ret"] = entity.E_SENDHTTP_FAILED
			util.Pkg2ClientTest(room.Ctx, retPkg, dataMap)
			return entity.E_SENDHTTP_FAILED
		}
	}

	user.ClearOldSettleInfo()

	if util.FloatCompare(userEx.Point, betCash) == -1 {
		user.AddCash(userEx.Point - betCash)
		userEx.Point = 0.0
	} else {
		userEx.Point = userEx.Point - betCash
	}

	//保存场景数据
	userEx.OperTime = time.Now().Unix()
	userEx.Point = userEx.Point + payout //派彩到积分
	userEx.Drill = userEx.Drill - drill  //减小砖头

	if userEx.Drill == 0 {
		userEx.Level = userEx.Level + 1
		userEx.Drill = 15
	}

	retPkg["ret"] = 0
	retPkg["uid"] = user.Id
	retPkg["left"] = user.Cash
	retPkg["payout"] = payout
	retPkg["result"] = betRetSlice

	user.BetJetton = betCash
	user.PayOut = payout

	retPkg["billNo"] = billNo

	retPkg["Level"] = userEx.Level
	retPkg["Point"] = userEx.Point
	retPkg["Drill"] = userEx.Drill

	retPkg["to"] = []uint32{user.Id}
	retPkg["cmd"] = util.CMD_MULTI
	util.Pkg2ClientTest(room.Ctx, retPkg, dataMap)

	if userEx.Level > 3 {
		//go to 龙珠夺宝,积分小于1，直接开始第1关
		if util.FloatCompare(userEx.Point, 1.0) == -1 {
			userEx.OperTime = time.Now().Unix()
			userEx.Point = 0
			userEx.Level = 1
			userEx.Drill = 15
		}
	}
	//保存场景数据
	// SaveUserSceneInfo2Db(room.Ctx, userId, room.GameId, userEx)
	SaveUserSceneInfo2DbTest(room.Ctx, userId, room.GameId, userEx, dataMap)

	//保存下注数据
	misc_record_state_scene_bill(room, user, time.Now().Unix(), billNo, "", state, dataMap)

	//清0，退出时有判断，用户是否正在下注，下了注，不让退出
	user.BetJetton = 0

	//logger.Notice("SceneWT01.Scene_Bet sucess:. user:", user.Id, "profit:", profit, "cash:", user.Cash/*, "retPkg:", retPkg*/)
	logger.Notice("[dongDebug], betEnd")
	return 0

}

// 玩家下注处理
func WT01_CreateWinTreasureResult(room *entity.Room, userEx *entity.WT01_UserInfo, betSlice []entity.WT01_Bet_Info_Struct) (float64, int, bson.M) {

	drill := WT01_Calc_Drill(userEx)
	retResult := bson.M{}

	// t1 := time.Now()

	var matrixData4 [4][4]int
	var matrixData5 [5][5]int
	var matrixData6 [6][6]int

	//生成创建参数
	linkCount := calaLinkCount(userEx.Level)              // 生成消的次数
	createParam := calaLinkParam(userEx.Level, linkCount) // 生成数组

	//createParam entity.LinkDataParamSlice
	//{1 4 0.2} {2 5 0.5} {3 6 2} {2 2 0}
	/*var createParam entity.LinkDataParamSlice
	createParam = append(createParam,entity.LinkDataParam{StoneType:1,Num:4,Betx:0.2})
	createParam = append(createParam,entity.LinkDataParam{StoneType:2,Num:5,Betx:0.5})
	createParam = append(createParam,entity.LinkDataParam{StoneType:3,Num:6,Betx:2})
	createParam = append(createParam,entity.LinkDataParam{StoneType:2,Num:2,Betx:0})*/

	if linkCount != len(createParam)-1 {
		logger.Error("calaLinkParam.error.")
	}

	bFirstTry := true

Loop:
	info := ""
	payOut := 0.0
	AllBetx := 0.0
	var matrixData44 [][4][4]int
	var matrixData55 [][5][5]int
	var matrixData66 [][6][6]int
	var linkData []map[int]int
	var inputMapData map[int]int
	ItemList := []bson.M{}
	var paneType int

	if userEx.Level == 1 {
		paneType = 4
	} else if userEx.Level == 2 {
		paneType = 5
	} else {
		paneType = 6
	}
	if paneType == 4 {
		matrixData4 = InitMatrixData4(paneType)
		inputMapData = ConverValidateMatrixData2MapData4(paneType, matrixData4)
	} else if paneType == 5 {
		matrixData5 = InitMatrixData5(paneType)
		inputMapData = ConverValidateMatrixData2MapData5(paneType, matrixData5)
	} else if paneType == 6 {
		matrixData6 = InitMatrixData6(paneType)
		inputMapData = ConverValidateMatrixData2MapData6(paneType, matrixData6)
	}

	if !bFirstTry { //失败后，排次序再试
		sort.Sort(createParam)
	}

	if bFirstTry {
		logger.Notice("uid:", userEx.Id, "calc linkCount:", linkCount, "createParamLen:", len(createParam)-1, "createParam:", createParam)
	}

	needCheckLink := false
	paramLen := len(createParam)
	for i := 0; i < len(createParam); i++ {
		param := createParam[i]

		Item := bson.M{}

		//inputMapData为本次可以生成的基础数据
		mapData1, ok := createLinkData(paneType, param.StoneType, param.Num, inputMapData)
		//logger.Notice("inputMapData:",inputMapData)

		if !ok {
			if bFirstTry {
				bFirstTry = false
				goto Loop
			} else {
				if i != paramLen-1 {
					logger.Error("service.WT01_CreateWinTreasureResult.fail.createParam:", createParam, "i:", i)
					continue
				}
			}
		}

		//填充1消的数据
		for k, v := range mapData1 {
			inputMapData[k] = v
		}

		//如果连消>2 ,随机一个第一板有2次消除的情况
		randomNum := random.Random_GetRandom(10)
		if i == 0 && linkCount >= 2 && createParam[i].Num+createParam[i+1].Num < paneType*paneType-paneType &&
			createParam[i].StoneType != createParam[i+1].StoneType && randomNum < 9 {
			i = i + 1
			param := createParam[i]
			mapData2, ok2 := createLinkData(paneType, param.StoneType, param.Num, inputMapData)

			if !ok2 {
				if bFirstTry {
					bFirstTry = false
					goto Loop
				} else {
					if i != paramLen-1 {
						logger.Error("service.WT01_CreateWinTreasureResult.fail.createParam:", createParam, "i:", i)
						continue
					}
				}
			}

			//填充1消的数据
			for k, v := range mapData2 {
				inputMapData[k] = v
				mapData1[k] = v
			}

			AllBetx = AllBetx + param.Betx
		}

		linkData = append(linkData, mapData1)
		mapBreakData, _ := createBreakLink(paneType, inputMapData)
		//矩阵数据
		if paneType == 4 {
			matrixData := convermap2Vec4(paneType, mapBreakData)
			Item["matrixData"] = matrixData
			matrixData44 = append(matrixData44, matrixData)
			//logger.Notice("i:",i,"matrixData:",matrixData)
		} else if paneType == 5 {
			matrixData := convermap2Vec5(paneType, mapBreakData)
			Item["matrixData"] = matrixData
			matrixData55 = append(matrixData55, matrixData)
		} else if paneType == 6 {
			matrixData := convermap2Vec6(paneType, mapBreakData)
			Item["matrixData"] = matrixData
			matrixData66 = append(matrixData66, matrixData)
		}

		//为下次消除准备数据
		//生成第二消的原始数据
		input2LinkData := createInputDataFrom1Link(paneType, mapBreakData, mapData1)
		inputMapData = input2LinkData

		//logger.Notice("needCheckLink:",needCheckLink,"input2LinkData:",input2LinkData)
		//检查生成的数据是否有新的满足条件的数据
		if needCheckLink {
			for k, v := range input2LinkData {
				retData := GetAllLinkPointByPos(k, paneType, v, input2LinkData)
				retDataLen := len(retData)
				if retDataLen >= paneType && param.StoneType != v && param.Num != retDataLen {
					//matrixDataCheck := convermap2Vec4(paneType, input2LinkData)
					//logger.Error("create matrixData fail.param.StoneType:",param.StoneType,"param.Num:",param.Num,"retDataLen:",retDataLen,"retData:",retData,"matrixData:",matrixDataCheck)
					goto Loop
				}
			}
		}
		ItemList = append(ItemList, Item)
		AllBetx = AllBetx + param.Betx
		info += fmt.Sprintf("%d:%d;", param.StoneType, param.Num)

		needCheckLink = true
	}
	//计算下注获利
	allJetton := 0.0
	for _, v := range betSlice {
		allJetton += v.BetJetton
	}
	payOut = allJetton * AllBetx

	retResult["paneType"] = paneType
	retResult["ItemList"] = ItemList

	if paneType == 4 {
		graph := createGraph44(matrixData44, linkData, drill)
		retResult["graph"] = graph
	} else if paneType == 5 {
		graph := createGraph55(matrixData55, linkData, drill)
		retResult["graph"] = graph
	} else if paneType == 6 {
		graph := createGraph66(matrixData66, linkData, drill)
		retResult["graph"] = graph
	}

	logger.Notice(userEx.OperTime, "Type:0, 1, ", userEx.Level, ", ", allJetton, ", ", payOut, ", ", info)
	// logger.Notice("uid:", userEx.Id, "createLink spend time:", time.Now().Sub(t1), "payOut:", payOut, "retResult:", retResult)

	return payOut, drill, retResult

}

func GetArrIndexByPos(nPos int, paneType int) (x int, y int) {
	nMod := nPos % paneType
	if nMod == 0 {
		x = nPos/paneType - 1
		y = paneType - 1
	} else {
		x = nPos / paneType
		y = nPos - paneType*x - 1
	}

	return x, y
}

func GetPosByArrIndex(x int, y int, paneType int) (nPos int) {
	nPos = paneType*x + y + 1
	return nPos
}

//检查点是否都相连
func checkAllDataLink(paneType int, stoneType int, inputData map[int]int, num int) bool {
	var checkData map[int]int
	checkData = make(map[int]int)

	for k, _ := range inputData {
		checkData[k] = stoneType
	}

	var startPos int
	for k, _ := range checkData {
		startPos = k
		break
	}

	retData := GetAllLinkPointByPos(startPos, paneType, stoneType, checkData)

	//logger.Notice("startPos:",startPos,"GetAllLinkPointByPos.len:",len(retData),"retData:",retData)

	//期望的点，就是成功
	if len(retData) == num {
		return true
	}

	//logger.Notice("startPos:",startPos,"GetAllLinkPointByPos.len:",len(checkData),"inputData:",checkData)
	return false
}

//得到与一个已知点相连的所有点
func GetAllLinkPointByPos(pos int, paneType int, stoneType int, inputData map[int]int) (retDataPos map[int]int) {
	retDataPos = make(map[int]int)
	var allDataPos []int

	allDataPos = append(allDataPos, pos)

	//logger.Notice("getalllinkpos:",pos,"inputData:",inputData)

	//广度搜索法
	for i := 0; i < len(allDataPos); i++ {
		validateData := GetPosRoundValidatePos(allDataPos[i], paneType)
		//logger.Notice("v:",allDataPos[i],"allDataPos:",allDataPos)

		//取得与这个点相连的数据
		for PosVal, _ := range validateData {
			_, ok := retDataPos[PosVal]

			//判断不在已存的点里面，且是当前要找的值
			if !ok && (inputData[PosVal] == stoneType) { //不在已生成的点里面
				retDataPos[PosVal] = stoneType
				allDataPos = append(allDataPos, PosVal)

				//logger.Notice("aaaaaaaaallDataPos:",allDataPos)
			}
		}
	}

	//logger.Notice("getalllinkpos:",retDataPos)
	return retDataPos
}

func GetPosRoundValidatePos(pos int, paneType int) (retDataPos map[int]int) {
	retDataPos = make(map[int]int)

	x, y := GetArrIndexByPos(pos, paneType)

	tempX := x - 1
	if tempX >= 0 {
		//retDataPos = append(retDataPos, GetPosByArrIndex(tempX, y))
		tempPos := GetPosByArrIndex(tempX, y, paneType)
		retDataPos[tempPos] = tempPos
	}

	tempX = x + 1
	if tempX < paneType {
		//retDataPos = append(retDataPos, GetPosByArrIndex(tempX, y))
		tempPos := GetPosByArrIndex(tempX, y, paneType)
		retDataPos[tempPos] = tempPos
	}

	tempY := y - 1
	if tempY >= 0 {
		//retDataPos = append(retDataPos, GetPosByArrIndex(x, tempY))
		tempPos := GetPosByArrIndex(x, tempY, paneType)
		retDataPos[tempPos] = tempPos
	}

	tempY = y + 1
	if tempY < paneType {
		//retDataPos = append(retDataPos, GetPosByArrIndex(x, tempY))
		tempPos := GetPosByArrIndex(x, tempY, paneType)
		retDataPos[tempPos] = tempPos
	}

	return retDataPos

}

//转换有效的矩阵数据到map
func ConverValidateMatrixData2MapData4(paneType int, matrixData [4][4]int) (retMapData map[int]int) {
	retMapData = make(map[int]int)
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			if matrixData[i][j] == 0 { //表示为有效可以填充的数据
				PosVal := GetPosByArrIndex(i, j, paneType)
				retMapData[PosVal] = 0
			}
		}
	}
	return retMapData
}

//转换有效的矩阵数据到map
func ConverValidateMatrixData2MapData5(paneType int, matrixData [5][5]int) (retMapData map[int]int) {
	retMapData = make(map[int]int)
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			if matrixData[i][j] == 0 { //表示为有效可以填充的数据
				PosVal := GetPosByArrIndex(i, j, paneType)
				retMapData[PosVal] = 0
			}
		}
	}
	return retMapData
}

//转换有效的矩阵数据到map
func ConverValidateMatrixData2MapData6(paneType int, matrixData [6][6]int) (retMapData map[int]int) {
	retMapData = make(map[int]int)
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			if matrixData[i][j] == 0 { //表示为有效可以填充的数据
				PosVal := GetPosByArrIndex(i, j, paneType)
				retMapData[PosVal] = 0
			}
		}
	}
	return retMapData
}

//生成初始的数据
func InitMatrixData4(paneType int) (matrixData [4][4]int) {
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			matrixData[i][j] = 0
		}
	}
	return
}

func InitMatrixData5(paneType int) (matrixData [5][5]int) {
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			matrixData[i][j] = 0
		}
	}
	return
}

func InitMatrixData6(paneType int) (matrixData [6][6]int) {
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			matrixData[i][j] = 0
		}
	}
	return
}

func createLinkData(paneType int, stoneType int, num int, inputMapData map[int]int) (mapData map[int]int, ok bool) {

	if (paneType-num) < paneType && len(inputMapData) == paneType {
		//如果求的相连位置数量>=13(假定是4*4的情况,5*5,6*6,类似求解)，那么，剩下的肯定是3个，则一定不会打断相连，先求出3个点，把剩下的input，返回就是要的相连的数据
		minMapData, ok := createLink(paneType, stoneType, paneType-num, inputMapData)
		if ok {
			for k, _ := range minMapData {
				delete(inputMapData, k)
			}

			return inputMapData, true
		}
	} else {
		return createLink(paneType, stoneType, num, inputMapData)
	}

	return mapData, false
}

//乌龟算法
func createLink(paneType int, stoneType int, num int, inputMapData map[int]int) (mapData map[int]int, ok bool) {
	var failData map[int]int
	failData = make(map[int]int)

	if num == 0 {
		return failData, true
	}
	LoopCount := 0
JLoop:

	mapData = make(map[int]int)
	LoopCount++

	//1.起始位置，然后通过这个位置生成指定数量的数据
	nStartPos := 0
	for k, v := range inputMapData {
		_, failFlag := failData[k]
		if v == 0 && !failFlag {
			nStartPos = k
			failData[nStartPos] = nStartPos

			break
		}

	}

	//没有一个合适的点来生成，则失败
	if nStartPos == 0 {
		failData = make(map[int]int)
		logger.Error("create fail, loop:")
		return failData, false
	}
	i := 1
	mapData[nStartPos] = i //用作生成顺序，方便对超过的数据进行删除

	retDataMap := GetAllLinkPointByPos(nStartPos, paneType, stoneType, inputMapData)
	for k, _ := range retDataMap {
		mapData[k] = stoneType * -1 //区分是上一局生成的点
	}

	//2.继续生成剩余的数据
	count := 0
	for i <= num {
		/*从已生成的数据里面，随机一个，生成这个数据，周围位置的任意一个数据为相连的格子，需要满足条件，
		这个新生成的位置，是一个合法的位置，同时也是没有在已生成的列表里面的位置*/

		//1.随机一个任意位置为起始点
		posVal := 0
		for k, _ := range mapData {
			posVal = k
			break
		}

		//2.得到这个位置周围合法的位置
		validateData := GetPosRoundValidatePos(posVal, paneType)

		//3.随机生成一个点，要求不在已生成的数据里面,并且是一个合法(not -1)的数据
		findValidatePos := false
		for randPosVal, _ := range validateData {
			_, ok := mapData[randPosVal]
			if !ok && (inputMapData[randPosVal] == 0) {
				mapData[randPosVal] = i + 1 //用作生成顺序，方便对超过的数据进行删除
				findValidatePos = true

				retDataMap := GetAllLinkPointByPos(randPosVal, paneType, stoneType, inputMapData)

				for k, _ := range retDataMap {
					mapData[k] = stoneType * -1 //区分是上一局生成的点
				}

				break
			}

		}

		//避免死循环
		count++
		if count > 100 {
			break
		}

		if !findValidatePos {
			continue
		}

		//生成完成，跳出
		if len(mapData) >= num {
			break
		}

		//生成了一个合法数据就减一个数量
		i++
	}

	needDelteNum := len(mapData) - num
	for i := 1; i <= needDelteNum; i++ {
		for k, v := range mapData {
			if v < 0 {
				continue //表示是上一次的点，不能删除哦
			}
			//检查是否允许删除，就是删除这个点
			delete(mapData, k)
			canDelete := checkAllDataLink(paneType, stoneType, mapData, num)
			if !canDelete {
				mapData[k] = v
			}

			if len(mapData) == num {
				for k, _ := range mapData {
					mapData[k] = stoneType
				}
				return mapData, true
			}
		}
	}

	//循环次数
	if LoopCount < paneType && len(mapData) != num {
		goto JLoop
	}

	//检查生成数量是否一致
	if len(mapData) != num {
		for k, _ := range mapData {
			mapData[k] = stoneType
		}
		return mapData, false
	}

	for k, _ := range mapData {
		mapData[k] = stoneType
	}

	return mapData, true
}

func createInputDataFrom1Link(paneType int, mapBreakData map[int]int, mapData1 map[int]int) map[int]int {
	if paneType == 4 {
		return createInputDataFrom1Link4(paneType, mapBreakData, mapData1)
	} else if paneType == 5 {
		return createInputDataFrom1Link5(paneType, mapBreakData, mapData1)
	} else {
		return createInputDataFrom1Link6(paneType, mapBreakData, mapData1)
	}

}

//生成2消的数据，通过1消
func createInputDataFrom1Link4(paneType int, mapBreakData map[int]int, mapData1 map[int]int) map[int]int {
	var data [4][4]int

	//1.原始数据
	for k, v := range mapBreakData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//2.把上一次相连的数据置0
	for k, _ := range mapData1 {
		mapBreakData[k] = 0
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = 0
	}

	//3.从最后一行开始把为0的数据往下消除
	for col := 0; col <= paneType-1; col++ { //例
		for row := paneType - 1; row >= 0; row-- {
			if data[row][col] == 0 {
				if row-1 == -1 {
					data[row][col] = -1
				} else {
					data[row][col] = data[row-1][col]

					for k := row - 1; k > 0; k-- {
						data[k][col] = data[k-1][col]
					}

					data[0][col] = -1

					if data[row][col] == 0 {
						row++
					}
				}
			}
		}
	}

	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			pos := GetPosByArrIndex(i, j, paneType)

			if data[i][j] > 0 {
				mapBreakData[pos] = data[i][j]
			} else {
				mapBreakData[pos] = 0
			}
		}
	}

	return mapBreakData
}

//生成2消的数据，通过1消
func createInputDataFrom1Link5(paneType int, mapBreakData map[int]int, mapData1 map[int]int) map[int]int {
	var data [5][5]int

	//1.原始数据
	for k, v := range mapBreakData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//2.把上一次相连的数据置0
	for k, _ := range mapData1 {
		mapBreakData[k] = 0
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = 0
	}

	//3.从最后一行开始把为0的数据往下消除
	for col := 0; col <= paneType-1; col++ { //例
		for row := paneType - 1; row >= 0; row-- {
			if data[row][col] == 0 {
				if row-1 == -1 {
					data[row][col] = -1
				} else {
					data[row][col] = data[row-1][col]

					for k := row - 1; k > 0; k-- {
						data[k][col] = data[k-1][col]
					}

					data[0][col] = -1

					if data[row][col] == 0 {
						row++
					}
				}
			}
		}
	}

	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			pos := GetPosByArrIndex(i, j, paneType)

			if data[i][j] > 0 {
				mapBreakData[pos] = data[i][j]
			} else {
				mapBreakData[pos] = 0
			}
		}
	}

	return mapBreakData
}

//生成2消的数据，通过1消
func createInputDataFrom1Link6(paneType int, mapBreakData map[int]int, mapData1 map[int]int) map[int]int {
	var data [6][6]int

	//1.原始数据
	for k, v := range mapBreakData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//2.把上一次相连的数据置0
	for k, _ := range mapData1 {
		mapBreakData[k] = 0
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = 0
	}

	//3.从最后一行开始把为0的数据往下消除
	for col := 0; col <= paneType-1; col++ { //例
		for row := paneType - 1; row >= 0; row-- {
			if data[row][col] == 0 {
				if row-1 == -1 {
					data[row][col] = -1
				} else {
					data[row][col] = data[row-1][col]

					for k := row - 1; k > 0; k-- {
						data[k][col] = data[k-1][col]
					}

					data[0][col] = -1

					if data[row][col] == 0 {
						row++
					}
				}
			}
		}
	}

	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			pos := GetPosByArrIndex(i, j, paneType)

			if data[i][j] > 0 {
				mapBreakData[pos] = data[i][j]
			} else {
				mapBreakData[pos] = 0
			}
		}
	}

	return mapBreakData
}

func convermap2Vec4(paneType int, mapData map[int]int) (data [4][4]int) {
	//转换数据格式输出
	for k, v := range mapData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//输出数据
	/*for i := 0; i < paneType; i++ {
		logger.Notice(strconv.Itoa(data[i][0]) + " " + strconv.Itoa(data[i][1]) + " " + strconv.Itoa(data[i][2]) + " " + strconv.Itoa(data[i][3]))
	}*/

	return
}

func convermap2Vec5(paneType int, mapData map[int]int) (data [5][5]int) {
	//转换数据格式输出
	for k, v := range mapData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//输出数据
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			if data[i][j] > 0 {
				fmt.Print(strconv.Itoa(data[i][j]) + " ")
			} else {

				fmt.Print(strconv.Itoa(data[i][j]) + " ")
			}
		}
		fmt.Println()
	}

	return
}

func convermap2Vec6(paneType int, mapData map[int]int) (data [6][6]int) {
	//转换数据格式输出
	for k, v := range mapData {
		x, y := GetArrIndexByPos(k, paneType)
		data[x][y] = v
	}

	//输出数据
	for i := 0; i < paneType; i++ {
		for j := 0; j < paneType; j++ {
			if data[i][j] > 0 {
				fmt.Print(strconv.Itoa(data[i][j]) + " ")
			} else {

				fmt.Print(strconv.Itoa(data[i][j]) + " ")
			}
		}
		fmt.Println()
	}

	return
}

//生成不相连的数据
func createBreakLink(paneType int, inputMapData map[int]int) (map[int]int, bool) {
	var fillMapData map[int]int
	fillMapData = make(map[int]int)

	//logger.Notice("inputMapData:",inputMapData)

	for k, v := range inputMapData {
		if v != 0 {
			continue
		}

		for i := 1; i <= 5; i++ {
			fillMapData[i] = i //填充可以使用的数据
		}

		//logger.Notice("fillMapData:",fillMapData)
		validateData := GetPosRoundValidatePos(k, paneType)

		//logger.Notice("k:",k,"validateData:",validateData)

		//判断四周填的数据是什么
		for randPosVal, _ := range validateData {
			if inputMapData[randPosVal] != 0 {
				delete(fillMapData, inputMapData[randPosVal])
			}
		}

		randVal := random.Random_GetRandom(len(fillMapData))

		//随机填一个值
		i := 0
		for j, _ := range fillMapData {
			if i == randVal {
				inputMapData[k] = j //填充这个值
				break
			}
			i++
		}
	}

	//logger.Notice("inputMapData:",inputMapData)
	return inputMapData, true

}

func createGraph44(matrixData [][4][4]int, link []map[int]int, drill int) (graph [][]int) {
	paneType := 4
	var colSlice [4][]int

	for k, v := range matrixData {
		curMatrix := v

		//处理第一个数据
		if k == 0 {
			for col := 0; col < paneType; col++ {
				for row := paneType - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
			//logger.Notice("colSlice:",colSlice)

		} else {
			// to do
			preMatrix := matrixData[k-1]
			mapData1 := link[k-1]

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1)
			//set preMatrix zero
			for k2, _ := range mapData1 {
				x, y := GetArrIndexByPos(k2, paneType)
				preMatrix[x][y] = 0

			}

			//check col count , move data
			var iColCount [4]int
			for iCol := 0; iCol < paneType; iCol++ {
				for iRow := paneType - 1; iRow >= 0; iRow-- {
					if preMatrix[iRow][iCol] == 0 {
						iColCount[iCol] = iColCount[iCol] + 1
					}
				}
			}

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1,"iColCount",iColCount)
			//copy data to sice
			for col := 0; col < paneType; col++ {
				for row := iColCount[col] - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
		}
	}

	drillColPos := 0
	drillPos := 0
	if drill > 0 {
		drillColPos = random.Random_GetRandom(paneType)
		drillPos = random.Random_GetRandom(len(colSlice[drillColPos]) - 1)
	}

	for col := 0; col < paneType; col++ {
		var graphColSlice []int
		for k, v := range colSlice[col] {
			if col == drillColPos && k == drillPos && drill > 0 {
				graphColSlice = append(graphColSlice, 11) //扔一个钻头进来
			}
			graphColSlice = append(graphColSlice, v)
		}

		graph = append(graph, graphColSlice)
	}

	//logger.Notice("graph:",graph)
	return graph
}

func createGraph55(matrixData [][5][5]int, link []map[int]int, drill int) (graph [][]int) {
	paneType := 5
	var colSlice [5][]int

	for k, v := range matrixData {
		curMatrix := v

		//处理第一个数据
		if k == 0 {
			for col := 0; col < paneType; col++ {
				for row := paneType - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
			//logger.Notice("colSlice:",colSlice)

		} else {
			// to do
			preMatrix := matrixData[k-1]
			mapData1 := link[k-1]

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1)
			//set preMatrix zero
			for k2, _ := range mapData1 {
				x, y := GetArrIndexByPos(k2, paneType)
				preMatrix[x][y] = 0

			}

			//check col count , move data
			var iColCount [5]int
			for iCol := 0; iCol < paneType; iCol++ {
				for iRow := paneType - 1; iRow >= 0; iRow-- {
					if preMatrix[iRow][iCol] == 0 {
						iColCount[iCol] = iColCount[iCol] + 1
					}
				}
			}

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1,"iColCount",iColCount)
			//copy data to sice
			for col := 0; col < paneType; col++ {
				for row := iColCount[col] - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
		}
	}

	drillColPos := 0
	drillPos := 0
	if drill > 0 {
		drillColPos = random.Random_GetRandom(paneType)
		drillPos = random.Random_GetRandom(len(colSlice[drillColPos]) - 1)
	}

	for col := 0; col < paneType; col++ {
		var graphColSlice []int
		for k, v := range colSlice[col] {
			if col == drillColPos && k == drillPos && drill > 0 {
				graphColSlice = append(graphColSlice, 11) //扔一个钻头进来
			}
			graphColSlice = append(graphColSlice, v)
		}

		graph = append(graph, graphColSlice)
	}

	//logger.Notice("graph:",graph)
	return graph
}

func createGraph66(matrixData [][6][6]int, link []map[int]int, drill int) (graph [][]int) {
	paneType := 6
	var colSlice [6][]int

	for k, v := range matrixData {
		curMatrix := v

		//处理第一个数据
		if k == 0 {
			for col := 0; col < paneType; col++ {
				for row := paneType - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
			//logger.Notice("colSlice:",colSlice)

		} else {
			// to do
			preMatrix := matrixData[k-1]
			mapData1 := link[k-1]

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1)
			//set preMatrix zero
			for k2, _ := range mapData1 {
				x, y := GetArrIndexByPos(k2, paneType)
				preMatrix[x][y] = 0

			}

			//check col count , move data
			var iColCount [6]int
			for iCol := 0; iCol < paneType; iCol++ {
				for iRow := paneType - 1; iRow >= 0; iRow-- {
					if preMatrix[iRow][iCol] == 0 {
						iColCount[iCol] = iColCount[iCol] + 1
					}
				}
			}

			//logger.Notice("preMatrix:",preMatrix,"mapData1:",mapData1,"iColCount",iColCount)
			//copy data to sice
			for col := 0; col < paneType; col++ {
				for row := iColCount[col] - 1; row >= 0; row-- {
					colSlice[col] = append(colSlice[col], curMatrix[row][col])
				}
			}
		}
	}

	drillColPos := 0
	drillPos := 0
	if drill > 0 {
		drillColPos = random.Random_GetRandom(paneType)
		drillPos = random.Random_GetRandom(len(colSlice[drillColPos]) - 1)
	}

	for col := 0; col < paneType; col++ {
		var graphColSlice []int
		for k, v := range colSlice[col] {
			if col == drillColPos && k == drillPos && drill > 0 {
				graphColSlice = append(graphColSlice, 11) //扔一个钻头进来
			}
			graphColSlice = append(graphColSlice, v)
		}

		graph = append(graph, graphColSlice)
	}

	//logger.Notice("graph:",graph)
	return graph
}

func WT01_User_user2Map(user *entity.WT01_UserInfo) bson.M {
	entity := bson.M{
		"Id":       user.Id,
		"GameId":   user.GameId,
		"Point":    user.Point,
		"OperTime": user.OperTime,
		"Level":    user.Level,
		"Drill":    user.Drill,
	}
	return entity
}

func LoadUserSceneInfoFromDb(id uint32, GameId string) *entity.WT01_UserInfo {
	var userExInfo *entity.WT01_UserInfo
	userExInfo = new(entity.WT01_UserInfo)

	ok := DB_getStruct(0, dao.User_ex_tblName, bson.M{"Id": id, "GameId": GameId}, &userExInfo)
	if !ok || userExInfo.OperTime < (time.Now().Unix()-30*24*60*60) { //超过30天的数据，视为无效
		userExInfo.Id = id
		userExInfo.GameId = GameId
		userExInfo.OperTime = time.Now().Unix()
		userExInfo.Level = 1
		userExInfo.Point = 0
		userExInfo.Drill = 15
	}

	//logger.Notice("servie.LoadUserSceneInfoFromDb.userExInfo:", userExInfo)
	return userExInfo
}

func SaveUserSceneInfo2Db(ctx int, id uint32, GameId string, userExInfo *entity.WT01_UserInfo) {
	userExInfoMap := WT01_User_user2Map(userExInfo)
	DB_Upsert(ctx, dao.User_ex_tblName, bson.M{"Id": id, "GameId": GameId}, userExInfoMap)

	logger.Notice("servie.SaveUserSceneInfo2Db.userExInfoMap:", userExInfoMap)

}

func SaveUserSceneInfo2DbTest(ctx int, id uint32, GameId string, userExInfo *entity.WT01_UserInfo, dataMap map[string]interface{}) {
	userExInfoMap := WT01_User_user2Map(userExInfo)
	DB_Upsert(ctx, dao.User_ex_tblName, bson.M{"Id": id, "GameId": GameId}, userExInfoMap)

	logger.Notice("servie.SaveUserSceneInfo2Db.userExInfoMap:", userExInfoMap)

}

func WT01_GetSceneInfo(room *entity.Room, scene *SceneWT01, userId uint32) (sceneInfo bson.M) {
	sceneInfo = bson.M{}

	userInfoEx := dao.WT01_User_fetchById(userId)
	if userInfoEx == nil {
		logger.Error("service.WT01_GetSceneInfo. userId:", userId, " is nil")
		return nil
	}

	sceneInfo["Point"] = userInfoEx.Point
	sceneInfo["Level"] = userInfoEx.Level
	sceneInfo["Drill"] = userInfoEx.Drill

	sceneInfo["BetxInfo"] = WT01_GetBetxInfo(room, scene)

	logger.Notice("service.WT01_GetSceneInfo.userId:", userId, "sceneInfo:", sceneInfo)

	return sceneInfo
}

func WT01_GetBetxInfo(room *entity.Room, scene *SceneWT01) (betxInfo bson.M) {
	betxInfo = bson.M{}

	betxInfo["Level1"] = wt01_betxInfo.Level1
	betxInfo["Level2"] = wt01_betxInfo.Level2
	betxInfo["Level3"] = wt01_betxInfo.Level3

	return betxInfo
}

var wt01_betxInfo entity.WT01_BetxInfo

func WT01_LoadBetxInfoFromDb() {
	ok := DB_getStruct(0, "config", bson.M{"group": "WT01"}, &wt01_betxInfo)
	if !ok {
		logger.Error("service.WT01_LoadBetxInfoFromDb.fail.")
	}

	logger.Notice("servie.WT01_LoadBetxInfoFromDb.betxInfo:", wt01_betxInfo)

}

//取得相连宝石的倍率
func WT01_GetLevelStoneLinkBetx(Level int, stoneType int, linkNum int) float64 {
	if Level < 1 && Level > 3 {
		return 0
	}

	minLinkNum := 4
	if Level == 2 {
		minLinkNum = 5
	} else if Level == 3 {
		minLinkNum = 6
	}
	if linkNum < minLinkNum {
		return 0
	}

	var stoneBetxInfo []entity.WT01_StoneBetxInfo
	if Level == 1 {
		stoneBetxInfo = wt01_betxInfo.Level1
	} else if Level == 2 {
		stoneBetxInfo = wt01_betxInfo.Level2
	} else if Level == 3 {
		stoneBetxInfo = wt01_betxInfo.Level3
	}

	for _, v := range stoneBetxInfo {
		if v.StoneType == stoneType {
			for k, betX := range v.Betx {
				if linkNum-1 == k {
					return betX
				}
			}
		}

	}

	logger.Error("servie.GetLevelStoneLinkBetx.fail.Level:", Level, "stoneType:", stoneType, "linkNum:", linkNum)
	return 0

}

var wt01_rtpInfo entity.WT01_RtpInfo

func calcSumRtp(rtp [][]float64) (rtpsum float64) {
	for _, v1 := range rtp {
		for _, v2 := range v1 {
			rtpsum += v2
		}
	}

	return rtpsum
}
func WT01_LoadRtpInfoFromDb() {
	ok := DB_getStruct(0, "config", bson.M{"group": "WT01_Rtp"}, &wt01_rtpInfo)
	if !ok {
		logger.Error("service.WT01_LoadRtpInfoFromDb.fail.")
	}

	//calc link data
	//level1
	wt01_rtpInfo.Level1.Link1Sum = calcSumRtp(wt01_rtpInfo.Level1.Link1)
	wt01_rtpInfo.Level1.Link2Sum = calcSumRtp(wt01_rtpInfo.Level1.Link2)
	wt01_rtpInfo.Level1.Link3Sum = calcSumRtp(wt01_rtpInfo.Level1.Link3)
	wt01_rtpInfo.Level1.Link4Sum = calcSumRtp(wt01_rtpInfo.Level1.Link4)
	wt01_rtpInfo.Level1.Link5Sum = calcSumRtp(wt01_rtpInfo.Level1.Link5)
	wt01_rtpInfo.Level1.Link6Sum = calcSumRtp(wt01_rtpInfo.Level1.Link6)
	wt01_rtpInfo.Level1.Link7Sum = calcSumRtp(wt01_rtpInfo.Level1.Link7)
	wt01_rtpInfo.Level1.Link8Sum = calcSumRtp(wt01_rtpInfo.Level1.Link8)

	leve1lRtp := wt01_rtpInfo.Level1.Link1Sum + wt01_rtpInfo.Level1.Link2Sum + wt01_rtpInfo.Level1.Link3Sum +
		wt01_rtpInfo.Level1.Link4Sum + wt01_rtpInfo.Level1.Link5Sum + wt01_rtpInfo.Level1.Link6Sum +
		wt01_rtpInfo.Level1.Link7Sum + wt01_rtpInfo.Level1.Link8Sum

	//level2
	wt01_rtpInfo.Level2.Link1Sum = calcSumRtp(wt01_rtpInfo.Level2.Link1)
	wt01_rtpInfo.Level2.Link2Sum = calcSumRtp(wt01_rtpInfo.Level2.Link2)
	wt01_rtpInfo.Level2.Link3Sum = calcSumRtp(wt01_rtpInfo.Level2.Link3)
	wt01_rtpInfo.Level2.Link4Sum = calcSumRtp(wt01_rtpInfo.Level2.Link4)
	wt01_rtpInfo.Level2.Link5Sum = calcSumRtp(wt01_rtpInfo.Level2.Link5)
	wt01_rtpInfo.Level2.Link6Sum = calcSumRtp(wt01_rtpInfo.Level2.Link6)
	wt01_rtpInfo.Level2.Link7Sum = calcSumRtp(wt01_rtpInfo.Level2.Link7)
	wt01_rtpInfo.Level2.Link8Sum = calcSumRtp(wt01_rtpInfo.Level2.Link8)

	leve2lRtp := wt01_rtpInfo.Level2.Link1Sum + wt01_rtpInfo.Level2.Link2Sum + wt01_rtpInfo.Level2.Link3Sum +
		wt01_rtpInfo.Level2.Link4Sum + wt01_rtpInfo.Level2.Link5Sum + wt01_rtpInfo.Level2.Link6Sum +
		wt01_rtpInfo.Level2.Link7Sum + wt01_rtpInfo.Level2.Link8Sum

	//level3
	wt01_rtpInfo.Level3.Link1Sum = calcSumRtp(wt01_rtpInfo.Level3.Link1)
	wt01_rtpInfo.Level3.Link2Sum = calcSumRtp(wt01_rtpInfo.Level3.Link2)
	wt01_rtpInfo.Level3.Link3Sum = calcSumRtp(wt01_rtpInfo.Level3.Link3)
	wt01_rtpInfo.Level3.Link4Sum = calcSumRtp(wt01_rtpInfo.Level3.Link4)
	wt01_rtpInfo.Level3.Link5Sum = calcSumRtp(wt01_rtpInfo.Level3.Link5)
	wt01_rtpInfo.Level3.Link6Sum = calcSumRtp(wt01_rtpInfo.Level3.Link6)
	wt01_rtpInfo.Level3.Link7Sum = calcSumRtp(wt01_rtpInfo.Level3.Link7)
	wt01_rtpInfo.Level3.Link8Sum = calcSumRtp(wt01_rtpInfo.Level3.Link8)

	leve3lRtp := wt01_rtpInfo.Level3.Link1Sum + wt01_rtpInfo.Level3.Link2Sum + wt01_rtpInfo.Level3.Link3Sum +
		wt01_rtpInfo.Level3.Link4Sum + wt01_rtpInfo.Level3.Link5Sum + wt01_rtpInfo.Level3.Link6Sum +
		wt01_rtpInfo.Level3.Link7Sum + wt01_rtpInfo.Level3.Link8Sum

	logger.Notice("service.wt01_rtp.level1:", leve1lRtp, "level2:", leve2lRtp, "level3:", leve3lRtp)
	logger.Notice("servie.WT01_LoadRtpInfoFromDb.wt01_rtpInfo:", wt01_rtpInfo)
}

//随机一个连消的次数
func calaLinkCount(level int) (linkCount int) {
	randnum := float64(random.Random_GetRandom(WT01_RANDLE_BASE))
	var LevelRtp entity.WT01_Rtp

	if level == 1 {
		LevelRtp = wt01_rtpInfo.Level1
	} else if level == 2 {
		LevelRtp = wt01_rtpInfo.Level2
	} else {
		LevelRtp = wt01_rtpInfo.Level3
	}

	/*Link8SumRange := LevelRtp.Link8Sum
	Link7SumRange := Link8SumRange + LevelRtp.Link7Sum
	Link6SumRange := Link7SumRange + LevelRtp.Link6Sum
	Link5SumRange := Link6SumRange + LevelRtp.Link5Sum*/
	Link5SumRange := LevelRtp.Link5Sum
	Link4SumRange := Link5SumRange + LevelRtp.Link4Sum
	Link3SumRange := Link4SumRange + LevelRtp.Link3Sum
	Link2SumRange := Link3SumRange + LevelRtp.Link2Sum
	Link1SumRange := Link2SumRange + LevelRtp.Link1Sum

	/*if randnum <= (Link8SumRange * WT01_RANDLE_BASE) {
		return 8
	} else if randnum <= (Link7SumRange * WT01_RANDLE_BASE) {
		return 7
	} else if randnum <= (Link6SumRange * WT01_RANDLE_BASE) {
		return 6
	} else */
	//logger.Notice("Link5SumRange:",Link5SumRange,"Link4SumRange:",Link4SumRange,"Link3SumRange:",Link3SumRange,"Link2SumRange:",Link2SumRange,"Link1SumRange:",Link1SumRange)
	//logger.Notice("randnum:",randnum,"kkk:",Link5SumRange * WT01_RANDLE_BASE)
	if randnum <= (Link5SumRange * WT01_RANDLE_BASE) {
		return 5
	} else if randnum <= (Link4SumRange * WT01_RANDLE_BASE) {
		return 4
	} else if randnum <= (Link3SumRange * WT01_RANDLE_BASE) {
		return 3
	} else if randnum <= (Link2SumRange * WT01_RANDLE_BASE) {
		return 2
	} else if randnum <= (Link1SumRange * WT01_RANDLE_BASE) {
		return 1
	} else {
		return 0 //表示没中
	}
}

//随机一个连消的次数
func calaLinkParam(level int, linkCount int) (createParam entity.LinkDataParamSlice) {
	var Item entity.LinkDataParam

	paneType := 0
	LinkSum := 0.0
	var rtp entity.WT01_Rtp
	var LinkParam [][]float64 //连消参数

	if level == 1 {
		rtp = wt01_rtpInfo.Level1
		paneType = 4
	} else if level == 2 {
		rtp = wt01_rtpInfo.Level2
		paneType = 5
	} else {
		rtp = wt01_rtpInfo.Level3
		paneType = 6
	}

	if linkCount == 1 {
		LinkSum = rtp.Link1Sum
		LinkParam = rtp.Link1
	} else if linkCount == 2 {
		LinkSum = rtp.Link2Sum
		LinkParam = rtp.Link2
	} else if linkCount == 3 {
		LinkSum = rtp.Link3Sum
		LinkParam = rtp.Link3
	} else if linkCount == 4 {
		LinkSum = rtp.Link4Sum
		LinkParam = rtp.Link4
	} else if linkCount == 5 {
		LinkSum = rtp.Link5Sum
		LinkParam = rtp.Link5
	} else if linkCount == 6 {
		LinkSum = rtp.Link6Sum
		LinkParam = rtp.Link6
	} else if linkCount == 7 {
		LinkSum = rtp.Link7Sum
		LinkParam = rtp.Link7
	} else if linkCount == 8 {
		LinkSum = rtp.Link8Sum
		LinkParam = rtp.Link8
	}

	//有几次连消，就取几次数据
	for i := 0; i < linkCount; i++ {

		randomNum := float64(random.Random_GetRandom(int(LinkSum * WT01_RANDLE_BASE)))
		tempRtpSum := 0.0

		bFind := false
		for stoneType := 4; stoneType >= 0; stoneType-- { // stoneType = 宝石类型
			for linkNum := 10; linkNum >= 0; linkNum-- { // linkNum = 相连个数
				tempRtpSum = tempRtpSum + LinkParam[stoneType][linkNum]

				//logger.Notice("randomNum:",randomNum,"tempRtpSum:",tempRtpSum * WT01_RANDLE_BASE)
				if randomNum <= tempRtpSum*WT01_RANDLE_BASE {
					Item.StoneType = stoneType + 1
					Item.Num = paneType + linkNum
					Item.Betx = WT01_GetLevelStoneLinkBetx(level, Item.StoneType, Item.Num)
					createParam = append(createParam, Item)
					bFind = true

					//logger.Notice("find")

					break
				}
			}

			if bFind {
				break
			}
		}
	}

	Item.StoneType = random.Random_GetRandom(5) + 1
	if linkCount == 1 {
		Item.Num = 3
	} else {
		Item.Num = 2
	}
	Item.Betx = 0.0
	createParam = append(createParam, Item)

	return createParam
}

var wt01_drill_rtp_info entity.WT01_Drill_Rtp

func WT01_LoadDrillRtpInfoFromDb() {
	ok := DB_getStruct(0, "config", bson.M{"group": "WT01_Drill_Rtp"}, &wt01_drill_rtp_info)
	if !ok {
		logger.Error("service.WT01_LoadDrillRtpInfoFromDb.fail.")
	}
	logger.Notice("servcie.WT01_LoadDrillRtpInfoFromDb...wt01_drill_rtp_info:", wt01_drill_rtp_info)
}

func WT01_Calc_Drill(userEx *entity.WT01_UserInfo) int {
	//return 0

	var drillRtpSlice []float64

	if userEx.Level == 1 {
		drillRtpSlice = wt01_drill_rtp_info.Level1
	} else if userEx.Level == 2 {
		drillRtpSlice = wt01_drill_rtp_info.Level2
	} else {
		drillRtpSlice = wt01_drill_rtp_info.Level3
	}

	rtp := 0.0
	for k, v := range drillRtpSlice {
		if k == (userEx.Drill-15)*-1 {
			rtp = v
		}
	}

	randomNum := float64(random.Random_GetRandom(100))

	//logger.Notice("Level:",userEx.Level,"DrillRtp:",rtp)
	if randomNum < rtp*100 {
		return 1
	}

	return 0
}

var wt01_win_treasure_data entity.WT01_Win_Treasure

func WT01_LoadWinTreasureInfoFromDb() {
	ok := DB_getStruct(0, "config", bson.M{"group": "WT01_Win_Treasure"}, &wt01_win_treasure_data)
	if !ok {
		logger.Error("service.WT01_LoadWinTreasureInfoFromDb.fail.")
	}
	logger.Notice("servcie.WT01_LoadWinTreasureInfoFromDb...wt01_win_treasure_data:", wt01_win_treasure_data)
}

// 第四关龙珠夺宝函数
func WT01_WinTreasure(ctx int, userId uint32, dataMap map[string]interface{}) {
	state := 0

	retPkg := bson.M{
		"op": util.OP_WT01_WINTREASURE_BACK,
	}
	retPkg["to"] = []uint32{userId}
	retPkg["cmd"] = util.CMD_MULTI

	fastitemsptr := FreqType_GetPtr()

	var user *entity.UserInfo
	user = fastitemsptr.UserMap[userId]
	if nil == user {
		logger.Error("servie.WT01_Bet...user is nil id:", userId)
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2ClientTest(ctx, retPkg, dataMap)
		return
	}

	userEx := dao.WT01_User_fetchByIdTest(userId, dataMap)
	if nil == userEx {
		logger.Error("SceneWT01.WT01_WinTreasure. user:", user.Id, " userExinfo is nil")
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2ClientTest(ctx, retPkg, dataMap)
		return
	}

	room := fastitemsptr.RoomMap[user.SceneId]
	if nil == room {
		logger.Error("servie.WT01_WinTreasure...room is nil , id:", userId, "roomId:", user.SceneId)
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2ClientTest(ctx, retPkg, dataMap)
		return
	}

	if userEx.Level != 4 {
		logger.Error("SceneWT01.WT01_WinTreasure. user:", user.Id, " not aloow wintreasure")
		retPkg["ret"] = entity.E_WINTREASURE_NOWT
		util.Pkg2ClientTest(ctx, retPkg, dataMap)
		return
	}

	//判断用户的积分在哪个范围
	var Data entity.WT01_WIN_TREASURE_LEVEL
	bFind := false
	for _, v := range wt01_win_treasure_data.Data {
		if util.FloatCompare(userEx.Point, v.PointMin) != -1 && util.FloatCompare(userEx.Point, v.PointMax) != 1 {
			Data = v
			bFind = true
			break
		}
	}

	if !bFind {
		logger.Error("SceneWT01.WT01_WinTreasure. user:", user.Id, "userPoint:", userEx.Point, " param is error.wt01_win_treasure_data:", wt01_win_treasure_data)

		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2ClientTest(ctx, retPkg, dataMap)
		return
	}

	userInfo := User_getByIdTest(userId, dataMap)
	if userInfo == nil {
		logger.Error("SceneWT01.WT01_WinTreasure. user:", userId, " is nil")
		return
	}

	//中奖奖金= 本场游戏所得金币/ 基准点数 × 基准奖级奖金
	randnum := random.Random_GetRandom(5)
	payout := userEx.Point / Data.BaseBonus * Data.Area[randnum]

	// --------------------------发送最后一关数据
	betItems := make([]httpclient.BetItemInfo, 0)

	nowTime := strconv.Itoa(int(time.Now().Unix()))
	userIdStr := strconv.Itoa(int(userId))
	transactionTime := nowTime + userIdStr

	logger.Notice("[dongDebug] SessionToken", userInfo.SessionToken)
	betInfo := httpclient.BetItemInfo{
		Type:            "bonus",
		Remarks:         "Win Treasure",
		TransactionId:   transactionTime,
		TransactionTime: time.Now().Format("2006-01-02 15:04:05"),
		BetAmount:       0.0,
		Currency:        userInfo.Currency,
		ResultAmount:    float64(int(payout)),
		RoundId:         transactionTime,
		RoundClosed:     true,
	}
	betItems = append(betItems, betInfo)

	sendData := httpclient.BetSendInfo{
		Request:              "bet-and-result",
		GameId:               userInfo.GameId,
		AccountId:            userInfo.AccountId,
		SessionToken:         userInfo.SessionToken,
		ValidateSessionToken: true,
		Items:                betItems,
		Mobile:               true,
		Ip:                   "",
	}

	count := 0
	for count < 5 {
		err := httpclient.BetAndResult(sendData)
		if err == httpclient.RET_ERROR {
			logger.Error("Send Http Error :", sendData)
			state = 1
			count += 1
		} else {
			state = 0
			break
		}
	}
	//------------------------------------------------

	user.BetJetton = 0
	user.PayOut = float64(int(payout))
	user.AddCash(float64(int(payout)))

	billNo := Bill_makeBillNo2("WT01", time.Now().Unix(), user.Id, 0)
	misc_record_state_scene_bill(room, user, time.Now().Unix(), billNo, "", state, dataMap)

	userEx.OperTime = time.Now().Unix()
	userEx.Point = 0
	userEx.Level = 1
	userEx.Drill = 15

	//保存场景数据
	SaveUserSceneInfo2DbTest(room.Ctx, userId, room.GameId, userEx, dataMap)

	retPkg["ret"] = entity.E_WINTREASURE_OK
	retPkg["payout"] = payout
	retPkg["area"] = randnum
	retPkg["uid"] = user.Id
	retPkg["left"] = user.Cash

	util.Pkg2ClientTest(room.Ctx, retPkg, dataMap)
	logger.Notice("servcie.WT01_WinTreasure...retPkg:", retPkg)
}

func WT01_Abort(ctx int, userId uint32) {
	retPkg := bson.M{
		"op": util.OP_WT01_ABORT_BACK,
	}
	retPkg["to"] = []uint32{userId}
	retPkg["cmd"] = util.CMD_MULTI

	fastitemsptr := FreqType_GetPtr()

	var user *entity.UserInfo
	user = fastitemsptr.UserMap[userId]
	if nil == user {
		logger.Error("servie.WT01_Abort...user is nil id:", userId)
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2Client(ctx, retPkg)
		return
	}

	userEx := dao.WT01_User_fetchById(userId)
	if nil == userEx {
		logger.Error("SceneWT01.WT01_Abort. user:", user.Id, " userExinfo is nil")
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2Client(ctx, retPkg)
		return
	}

	room := fastitemsptr.RoomMap[user.SceneId]
	if nil == room {
		logger.Error("servie.WT01_Abort...room is nil , id:", userId, "roomId:", user.SceneId)
		retPkg["ret"] = entity.E_WINTREASURE_PARAM
		util.Pkg2Client(ctx, retPkg)
		return
	}

	userEx.OperTime = time.Now().Unix()
	userEx.Point = 0
	userEx.Level = 1
	userEx.Drill = 15

	//保存场景数据
	SaveUserSceneInfo2Db(room.Ctx, userId, room.GameId, userEx)

	retPkg["ret"] = entity.E_WINTREASURE_OK

	util.Pkg2Client(ctx, retPkg)
	logger.Notice("servcie.WT01_Abort...retPkg:", retPkg)
}

// 获取玩家的战绩
func WT01_GetUsetBetInfo(ctx int, userId uint32, page uint32, size uint32) {
	var billscene []entity.BetInfo

	retPkg := bson.M{
		"op": util.OP_WT01_USER_BET_INFO_BACK,
	}
	retPkg["to"] = []uint32{userId}
	retPkg["cmd"] = util.CMD_MULTI

	resultSlice := make([]entity.BetInfo, 0)
	skip := (page - 1) * size

	endTime := uint32(time.Now().Unix())
	startTime := uint32(time.Now().Unix()) - 60*60*24*2

	filter := []bson.M{
		{"$match": bson.M{"bt": 1, "GameId": "WT01", "Time": bson.M{"$gte": startTime, "$lt": endTime}, "uid": userId}},
		{"$group": bson.M{"_id": "$billNo", "time": bson.M{"$first": "$Time"}, "bet": bson.M{"$sum": "$betjetton"}, "win": bson.M{"$sum": "$payout"}, "sum": bson.M{"$sum": "$md"}}},
		{"$sort": bson.M{"time": -1}},
		{"$skip": skip},
		{"$limit": size},
	}

	logger.Notice("[dongDebug] ", filter)

	ok := AggregateLimit(0, util.BILL_TBL_GAME, filter, &billscene)

	if !ok {
		logger.Error("service.WT01_GetUsetBetInfo...Data error!")
		retPkg["ret"] = entity.E_SENDHTTP_FAILED
		util.Pkg2Client(ctx, retPkg)
		return
	} else {
		for _, v := range billscene {
			var item entity.BetInfo
			item.Time = v.Time
			item.Id = v.Id
			item.Bet = v.Bet
			item.Win = v.Win
			item.Sum = v.Sum
			resultSlice = append(resultSlice, item)
		}
	}

	logger.Notice("servie.WT01_GetUsetBetInfo.resultSlice:", resultSlice)
	retPkg["ret"] = entity.E_WINTREASURE_OK
	retPkg["result"] = resultSlice
	util.Pkg2Client(ctx, retPkg)
}

// 获取玩家的jackPot
func WT01_JackPot(ctx int, userId uint32, page uint32, size uint32) {
	var billscene []entity.JackPotInfo
	type userJackPotInfo struct {
		Time int64
		Nick string
		Rank uint32
		Md   float64
	}
	retPkg := bson.M{
		"op": util.OP_WT01_USER_JACKPOT_BACK,
	}
	retPkg["to"] = []uint32{userId}
	retPkg["cmd"] = util.CMD_MULTI

	resultSlice := make([]userJackPotInfo, 0)

	endTime := uint32(time.Now().Unix())
	startTime := uint32(time.Now().Unix()) - 60*60*24*7

	skip := (page - 1) * size
	filter := []bson.M{
		{"$match": bson.M{"bt": 2, "GameId": "WT01", "Time": bson.M{"$gte": startTime, "$lt": endTime}}},
		{"$project": bson.M{"time": "$Time", "nick": "$nickName", "rank": "$ll", "md": 1}},
		{"$sort": bson.M{"time": -1}},
		{"$skip": skip},
		{"$limit": size},
	}
	logger.Notice("[dongDebug] ", filter)

	ok := AggregateLimit(0, util.BILL_TBL_GAME, filter, &billscene)

	if !ok {
		logger.Error("service.WT01_JackPot...Data error!")
		retPkg["ret"] = entity.E_SENDHTTP_FAILED
		util.Pkg2Client(ctx, retPkg)
		return
	} else {
		for _, v := range billscene {
			var item userJackPotInfo
			item.Time = v.Time
			item.Rank = v.Rank
			item.Md = v.Md
			item.Nick = v.Nick
			resultSlice = append(resultSlice, item)
		}
	}

	logger.Notice("servie.WT01_JackPot.resultSlice:", resultSlice)
	retPkg["ret"] = entity.E_WINTREASURE_OK
	retPkg["result"] = resultSlice
	util.Pkg2Client(ctx, retPkg)
}

func BetTest() int {

	logger.Notice("test begin")
	betCash := 100.0
	var betSlice []entity.WT01_Bet_Info_Struct
	betSlice = append(betSlice, entity.WT01_Bet_Info_Struct{BetJetton: betCash})

	userEx := new(entity.WT01_UserInfo)

	// 跑数据
	logger.Notice("START---------------------------")
	//下注处理，扣除玩家的下注金额
	for le := 1; le < 4; le++ {
		userEx.Level = le
		for i := 0; i < 100000000; i++ {
			userEx.OperTime = int64(i)
			WT01_CreateWinTreasureResultTest(userEx, betSlice)
		}
	}

	logger.Notice("END---------------------------")

	return 0

}

func WT01_CreateWinTreasureResultTest(userEx *entity.WT01_UserInfo, betSlice []entity.WT01_Bet_Info_Struct) {

	// t1 := time.Now()

	var matrixData4 [4][4]int
	var matrixData5 [5][5]int
	var matrixData6 [6][6]int

	//生成创建参数
	linkCount := calaLinkCount(userEx.Level)              // 生成消的次数
	createParam := calaLinkParam(userEx.Level, linkCount) // 生成数组

	//createParam entity.LinkDataParamSlice
	//{1 4 0.2} {2 5 0.5} {3 6 2} {2 2 0}
	/*var createParam entity.LinkDataParamSlice
	createParam = append(createParam,entity.LinkDataParam{StoneType:1,Num:4,Betx:0.2})
	createParam = append(createParam,entity.LinkDataParam{StoneType:2,Num:5,Betx:0.5})
	createParam = append(createParam,entity.LinkDataParam{StoneType:3,Num:6,Betx:2})
	createParam = append(createParam,entity.LinkDataParam{StoneType:2,Num:2,Betx:0})*/

	if linkCount != len(createParam)-1 {
		logger.Error("calaLinkParam.error.")
	}

	bFirstTry := true

Loop:
	info := ""
	payOut := 0.0
	AllBetx := 0.0
	var linkData []map[int]int
	var inputMapData map[int]int
	var paneType int

	if userEx.Level == 1 {
		paneType = 4
	} else if userEx.Level == 2 {
		paneType = 5
	} else {
		paneType = 6
	}
	if paneType == 4 {
		matrixData4 = InitMatrixData4(paneType)
		inputMapData = ConverValidateMatrixData2MapData4(paneType, matrixData4)
	} else if paneType == 5 {
		matrixData5 = InitMatrixData5(paneType)
		inputMapData = ConverValidateMatrixData2MapData5(paneType, matrixData5)
	} else if paneType == 6 {
		matrixData6 = InitMatrixData6(paneType)
		inputMapData = ConverValidateMatrixData2MapData6(paneType, matrixData6)
	}

	if !bFirstTry { //失败后，排次序再试
		sort.Sort(createParam)
	}

	if bFirstTry {
		// logger.Notice("uid:", userEx.Id, "calc linkCount:", linkCount, "createParamLen:", len(createParam)-1, "createParam:", createParam)
	}

	needCheckLink := false
	paramLen := len(createParam)
	for i := 0; i < len(createParam); i++ {
		param := createParam[i]

		//inputMapData为本次可以生成的基础数据
		mapData1, ok := createLinkData(paneType, param.StoneType, param.Num, inputMapData)
		//logger.Notice("inputMapData:",inputMapData)

		if !ok {
			if bFirstTry {
				bFirstTry = false
				goto Loop
			} else {
				if i != paramLen-1 {
					logger.Error("service.WT01_CreateWinTreasureResult.fail.createParam:", createParam, "i:", i)
					continue
				}
			}
		}

		//填充1消的数据
		for k, v := range mapData1 {
			inputMapData[k] = v
		}

		//如果连消>2 ,随机一个第一板有2次消除的情况
		randomNum := random.Random_GetRandom(10)
		if i == 0 && linkCount >= 2 && createParam[i].Num+createParam[i+1].Num < paneType*paneType-paneType &&
			createParam[i].StoneType != createParam[i+1].StoneType && randomNum < 9 {
			i = i + 1
			param := createParam[i]
			mapData2, ok2 := createLinkData(paneType, param.StoneType, param.Num, inputMapData)

			if !ok2 {
				if bFirstTry {
					bFirstTry = false
					goto Loop
				} else {
					if i != paramLen-1 {
						logger.Error("service.WT01_CreateWinTreasureResult.fail.createParam:", createParam, "i:", i)
						continue
					}
				}
			}

			//填充1消的数据
			for k, v := range mapData2 {
				inputMapData[k] = v
				mapData1[k] = v
			}

			AllBetx = AllBetx + param.Betx
		}

		linkData = append(linkData, mapData1)
		mapBreakData, _ := createBreakLink(paneType, inputMapData)

		//为下次消除准备数据
		//生成第二消的原始数据
		input2LinkData := createInputDataFrom1Link(paneType, mapBreakData, mapData1)
		inputMapData = input2LinkData

		//logger.Notice("needCheckLink:",needCheckLink,"input2LinkData:",input2LinkData)
		//检查生成的数据是否有新的满足条件的数据
		if needCheckLink {
			for k, v := range input2LinkData {
				retData := GetAllLinkPointByPos(k, paneType, v, input2LinkData)
				retDataLen := len(retData)
				if retDataLen >= paneType && param.StoneType != v && param.Num != retDataLen {
					//matrixDataCheck := convermap2Vec4(paneType, input2LinkData)
					//logger.Error("create matrixData fail.param.StoneType:",param.StoneType,"param.Num:",param.Num,"retDataLen:",retDataLen,"retData:",retData,"matrixData:",matrixDataCheck)
					goto Loop
				}
			}
		}
		AllBetx = AllBetx + param.Betx
		info += fmt.Sprintf("%d:%d;", param.StoneType, param.Num)

		needCheckLink = true
	}
	//计算下注获利
	allJetton := 0.0
	for _, v := range betSlice {
		allJetton += v.BetJetton
	}
	payOut = allJetton * AllBetx

	logger.Notice(userEx.OperTime, "Type:0, 1, ", userEx.Level, ", ", allJetton, ", ", payOut, ", ", info)

}

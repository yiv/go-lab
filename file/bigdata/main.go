package main

import (
	"fmt"
	. "git.cr56.cn/zygroup/zyplan/royalserver/pkg/royal"
	"github.com/go-kit/kit/log"
	"math/rand"
	"os"
	"time"
)

func main() {

	var (
		KingCount        int64
		QueenCount       int64
		TieCount         int64
		MinisterCount    int64
		KingBetCount     int64
		QueenBetCount    int64
		TieBetCount      int64
		MinisterBetCount int64
		AllBetCount      int64
		AllSettleCount   int64
		AllLoseCount     int64
	)

	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	_, err := os.Stat("test.csv")
	if os.IsNotExist(err) {
		fd, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			panic(err.Error())
		}
		line := fmt.Sprintf("\xEF\xBB\xBF")
		line += fmt.Sprintf("id,总下注,总派奖,龙下注,凤下注,和下注,三公下注,龙派奖,凤派奖,和派奖,三公派奖,龙牌型,龙牌,凤牌型,凤牌,比牌结果\n")
		if _, err := fd.Write([]byte(line)); err != nil {
			panic(err.Error())
		}
		fd.Close()
	}

	start := time.Now()
	fd, err := os.OpenFile("test.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 100000000; i++ {
		line, betPool, winPool, allBet, allSettle, allLose := Run(logger, i)
		AllBetCount += allBet
		AllSettleCount += allSettle
		AllLoseCount += allLose
		for k := range winPool {
			switch k {
			case BetKing:
				KingCount++
			case BetQueen:
				QueenCount++
			case BetTie:
				TieCount++
			case BetMinister:
				MinisterCount++
			}
		}
		for k, v := range betPool {
			switch BetType(k) {
			case BetKing:
				KingBetCount += v
			case BetQueen:
				QueenBetCount += v
			case BetTie:
				TieBetCount += v
			case BetMinister:
				MinisterBetCount += v
			}
		}
		if _, err := fd.Write([]byte(line)); err != nil {
			panic(err.Error())
		}
	}
	fd.Close()

	fmt.Println("KingCount", KingCount, "QueenCount", QueenCount, "TieCount", TieCount, "MinisterCount", MinisterCount, "AllBetCount", AllBetCount, "AllSettleCount", AllSettleCount, "AllLoseCount", AllLoseCount)
	fmt.Println("KingBetCount", KingBetCount, "QueenBetCount", QueenBetCount, "TieBetCount", TieBetCount, "MinisterBetCount", MinisterBetCount)
	fmt.Println("time used ", time.Now().Sub(start))

}

func Run(logger Logger, id int) (line string, betPool BetPool, winPool map[BetType]bool, allBet, allSettle, allLose int64) {
	var (
		ret          CompareResult
		err          error
		winner       BetType
		playerSettle int64
		playerLose   int64
	)
	winPool = make(map[BetType]bool)
	betPool = BetPool{}
	settlePool := BetPool{}

	dealer := NewCardDealer(logger)
	hands := dealer.Licensing()
	kingCards := hands[0]
	queenCards := hands[1]
	kingType := kingCards.Type()
	queenType := queenCards.Type()

	for i := 0; i < 5; i++ {
		poolProb := rand.Int31n(4)
		betProb := rand.Int31n(5)
		betCoin := RobotBetCoinMap[betProb]
		betPool[BetType(poolProb)] += betCoin
	}

	//for i := 0; i < 5; i++ {
	//	poolProb := 0
	//	betProb := 0
	//	betCoin := RobotBetCoinMap[betProb]
	//	betPool[BetType(poolProb)] += betCoin
	//}
	//判断三公
	switch kingType {
	case CardTypeSeniorMinister, CardTypeJuniorMinister, CardTypeHybridMinister:
		winPool[BetMinister] = true
	}
	switch queenType {
	case CardTypeSeniorMinister, CardTypeJuniorMinister, CardTypeHybridMinister:
		winPool[BetMinister] = true
	}
	//判断龙凤输赢
	if ret, err = kingCards.Compare(queenCards); err != nil {
		fmt.Println("err", err.Error())
		panic(err)
	}

	switch ret {
	case GreaterThan:
		//龙胜
		winPool[BetKing] = true
		winner = BetKing
	case LessThan:
		//凤胜
		winPool[BetQueen] = true
		winner = BetQueen
	case EqualTo:
		//和
		winPool[BetTie] = true
		winner = BetTie
	}

	for k, v := range betPool {
		pool := BetType(k)
		if winPool[pool] {
			var settle int64
			switch pool {
			case BetTie:
				settle += v * 123 * (100 - 5) / 100
			case BetMinister:
				settle += int64(float64(v)*41.6) * (100 - 5) / 100
			default:
				settle += v * (100 - 5) / 100
			}
			playerSettle += settle
			settlePool[pool] += settle
		} else {
			if winner == BetTie {
				switch pool {
				case BetKing, BetQueen:
				default:
					playerLose += v
				}
			} else {
				playerLose += v
			}
		}
	}

	allBet = betPool.CountBet()
	allSettle = playerSettle
	allLose = playerLose
	line = fmt.Sprintf("%v, %v, %v, %v,%v, %v, %v, %v,%v, %v, %v, %v, %v, %v, %v, %v\n", id, betPool.CountBet(), playerSettle, betPool[BetKing], betPool[BetQueen], betPool[BetTie], betPool[BetMinister], settlePool[BetKing], settlePool[BetQueen], settlePool[BetTie], settlePool[BetMinister], kingType, kingCards.PrintHex(), queenType, queenCards.PrintHex(), ret)
	return
}

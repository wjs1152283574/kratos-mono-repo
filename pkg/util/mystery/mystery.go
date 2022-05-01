package mystery

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	wr "github.com/mroth/weightedrand"
)

type Prize struct {
	PlayerId int64
	Weight   int
}

// 权重随机抽奖
func RandomDraw(prizes []*Prize) int64 {
	//权重累加求和
	var weightSum int
	for _, v := range prizes {
		weightSum += v.Weight
	}

	//生成一个权重随机数，介于0-weightSum之间
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(weightSum)

	//权重数组重组并排序
	randomNumTmp := &Prize{PlayerId: -1, Weight: randomNum}
	concatWeightArr := make([]*Prize, 0)
	aa, _ := json.Marshal(prizes)
	_ = json.Unmarshal(aa, &concatWeightArr)
	concatWeightArr = append(concatWeightArr, randomNumTmp) //将随机数加入权重数组

	//将包含随机数的新权重数组按从小到大（升序）排序
	sort.Slice(concatWeightArr, func(i, j int) bool {
		return concatWeightArr[i].Weight < concatWeightArr[j].Weight
	})
	sort.Slice(prizes, func(i, j int) bool {
		return prizes[i].Weight < prizes[j].Weight
	})

	//索引权重随机数的数组下标
	var randomNumIndex = -1 //索引随机数在新权重数组中的位置
	for p, v := range concatWeightArr {
		if v.Weight == randomNum {
			randomNumIndex = p
		}
	}
	randomNumIndexTmp := math.Min(float64(randomNumIndex), float64(len(prizes)-1)) //权重随机数的下标不得超过奖项数组的长度-1，重新计算随机数在奖项数组中的索引位置
	randomNumIndex = int(randomNumIndexTmp)

	//取出对应奖项
	res := prizes[randomNumIndex] //从奖项数组中取出本次抽奖结果
	fmt.Println("本次抽奖结果：", res.PlayerId)
	return res.PlayerId
}

// GetPrice 开启盲盒，返回中奖结果map,可一次性多开
func GetPrice(boxNum, avalable int64, prize []*Prize) map[int64]int64 {
	var prizeRes = make(map[int64]int64, avalable) // 抽奖结果
	for i := 0; i < int(boxNum); i++ {
		pid := RandomDraw(prize)  // 传入奖池数据，返回抽中的PID,但是奖池数据不会变动
		for i, v := range prize { // 每次减相应的数量，保证权重有变化
			if v.PlayerId == pid {
				if v.Weight < 1 {
					prize = append(prize[:i], prize[i+1:]...)
				}
				v.Weight -= 1
			}
		}
		// 写入中奖结果
		_, ok := prizeRes[pid]
		if ok {
			prizeRes[pid] += 1
		} else {
			prizeRes[pid] = 1
		}
	}
	return prizeRes
}

// ******** 方案2 (推荐) ******

// GetRandPrize 随机获奖 输入盲盒数量，输出中奖ID:数量
func GetRandPrize(boxNum int64, prize []Prize) map[int64]int64 {
	p := make(map[int64]int64, boxNum)

	for i := 0; i < int(boxNum); i++ {
		// 解析将池
		var wrs []wr.Choice
		for _, v := range prize {
			wrs = append(wrs, wr.Choice{
				Item:   v.PlayerId,
				Weight: uint(v.Weight),
			})
		}
		result := RandPick(wrs)
		p[result] += 1
		// 更新将池
		for i, v := range prize {
			if v.PlayerId == result {
				prize[i].Weight -= 1
			}
		}
	}

	return p
}

// RandPick 传入将池，返回中奖item
func RandPick(wrs []wr.Choice) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	chooser, _ := wr.NewChooser(wrs...)
	return chooser.Pick().(int64)
}

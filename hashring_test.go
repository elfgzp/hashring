package hashring

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/toolkits/pkg/logger"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func TestHashRing_NodeLoadBalance(t *testing.T) {
	hashRing, err := NewHashRing(150)
	if err != nil {
		t.Error(err)
	}
	nodeNum := 10
	dataAmount := 10000
	avg := float64(dataAmount) / float64(nodeNum)

	nodeMap := make(map[string]int, 0)

	for i := 1; i <= nodeNum; i++ {
		nodeMap[fmt.Sprintf("192.168.1.1%d", i)] = 1
	}
	hashRing.AddNodes(nodeMap)
	counter := make(map[string]int, 0)
	for i := 1; i <= 10000; i++ {
		key := RandStringRunes(20)
		nodeName, _ := hashRing.NodeLoadBalance(key)
		counter[nodeName]++
	}
	logger.Infof("counter %+v\n", counter)

	val := 0.0
	for nodeName := range nodeMap {
		count := counter[nodeName]
		pow := math.Pow(float64(count)-avg, 2)
		val += pow
	}
	val = val / float64(nodeNum)
	val = math.Sqrt(val)
	logger.Infof("Standard deviation %f \n", val)
}

package hashring

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"

	"github.com/toolkits/pkg/logger"
)

// Node 节点
type Node struct {
	Name   string
	Weight int
}

// VirtualNode 虚拟节点
type VirtualNode struct {
	Node
	VirtualName string
	HashValue   uint32
}

// virtualNodeArray 节点列表
type virtualNodeArray []*VirtualNode

// Len Len
func (vn virtualNodeArray) Len() int {
	return len(vn)
}

// Less Less
func (vn virtualNodeArray) Less(i, j int) bool {
	return vn[i].HashValue < vn[j].HashValue
}

// Swap Swap
func (vn virtualNodeArray) Swap(i, j int) {
	vn[i], vn[j] = vn[j], vn[i]
}

// Sort Sort
func (vn virtualNodeArray) Sort() {
	sort.Sort(vn)
}

// HashRing 哈希环
type HashRing struct {
	VirtualNodeNum int
	RealNodeMap    map[string]*Node
	virtualNodeArr virtualNodeArray
}

func (h *HashRing) hash(key string) uint32 {
	m := md5.New()
	m.Write([]byte(key))
	keyMD5 := m.Sum(nil)
	hashValue := binary.BigEndian.Uint32(keyMD5)
	return hashValue
}

func (h *HashRing) generate() {
	h.virtualNodeArr = make(virtualNodeArray, 0)

	for _, node := range h.RealNodeMap {
		virtualNodeNum := node.Weight * h.VirtualNodeNum
		for i := 1; i <= virtualNodeNum; i++ {
			virtualNode := &VirtualNode{}
			virtualNode.Name = node.Name
			virtualNode.Weight = node.Weight
			virtualNode.VirtualName = fmt.Sprintf("%s#%d", node.Name, i)
			virtualNode.HashValue = h.hash(virtualNode.VirtualName)
			h.virtualNodeArr = append(h.virtualNodeArr, virtualNode)
		}
	}

	h.virtualNodeArr.Sort()
}

// NewHashRing 新建哈希环
func NewHashRing(virtualNodeNum int) (*HashRing, error) {

	if virtualNodeNum == 0 {
		virtualNodeNum = 150
	}

	hashRing := &HashRing{
		VirtualNodeNum: virtualNodeNum,
		RealNodeMap:    make(map[string]*Node, 0),
	}

	return hashRing, nil
}

// AddNode 增加节点
func (h *HashRing) AddNode(nodeName string, nodeWeight int) {
	node := &Node{
		Name:   nodeName,
		Weight: nodeWeight,
	}
	h.RealNodeMap[nodeName] = node
	h.generate()
}

// RemoveNode 删除节点
func (h *HashRing) RemoveNode(nodeName string) {
	delete(h.RealNodeMap, nodeName)
	h.generate()
}

// NodeLoadBalance 节点负载均衡
func (h *HashRing) NodeLoadBalance(key string) (string, error) {
	if len(h.virtualNodeArr) == 0 {
		return "", errors.New("No available node. ")
	}

	hashValue := h.hash(key)

	i := sort.Search(h.virtualNodeArr.Len(), func(i int) bool {
		nodeHashValue := h.virtualNodeArr[i].HashValue
		return nodeHashValue >= hashValue
	})

	if i == h.virtualNodeArr.Len() {
		i = 0
	}

	virtualNode := h.virtualNodeArr[i]

	logger.Infof("key %s hashValue %d nodeHashValue %d virtualNodeName %s nodeName %s \n", key, hashValue, virtualNode.HashValue, virtualNode.VirtualName, virtualNode.Name)

	return virtualNode.Name, nil
}

// AddNodes 批量增加节点
func (h *HashRing) AddNodes(nodeWeightMap map[string]int) {
	for nodeName, nodeWeight := range nodeWeightMap {
		h.RealNodeMap[nodeName] = &Node{
			Name:   nodeName,
			Weight: nodeWeight,
		}
	}
	h.generate()
}

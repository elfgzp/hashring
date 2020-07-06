package hashring

import (
	"errors"
	"fmt"
	"hash/adler32"
	"sort"
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

func (h *HashRing) generate() {
	var totalWeight int
	h.virtualNodeArr = make(virtualNodeArray, 0)

	for _, node := range h.RealNodeMap {
		totalWeight += node.Weight
	}

	if totalWeight == 0 {
		return
	}

	adl := adler32.New()
	for _, node := range h.RealNodeMap {
		virtualNodeNum := node.Weight * h.VirtualNodeNum / totalWeight
		for i := 1; i <= virtualNodeNum; i++ {
			virtualNode := &VirtualNode{}
			virtualNode.Name = node.Name
			virtualNode.Weight = node.Weight
			virtualNode.VirtualName = fmt.Sprintf("%s#%d", node.Name, i)
			adl.Write([]byte(virtualNode.VirtualName))
			virtualNode.HashValue = adl.Sum32()
			adl.Reset()
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

	adl := adler32.New()
	adl.Write([]byte(key))
	hashValue := adl.Sum32()

	i := sort.Search(h.virtualNodeArr.Len(), func(i int) bool {
		return h.virtualNodeArr[i].HashValue >= hashValue
	})

	if i == h.virtualNodeArr.Len() {
		i = 0
	}

	virtualNode := h.virtualNodeArr[i]

	return virtualNode.Name, nil
}

// AddNodes 批量增加节点
func (h *HashRing) AddNodes(nodeWeightMap map[string]int) {
	for nodeName, nodeWeight := range nodeWeightMap {
		h.AddNode(nodeName, nodeWeight)
	}
}

package hashring

// HashRing 哈希环
type HashRing struct{}

// NewHashRing 新建哈希环
func NewHashRing() (*HashRing, error) {
	hashRing := &HashRing{}

	return hashRing, nil
}

// AddNode 增加节点
func (h *HashRing) AddNode(nodeName string, nodeWeight int64) error {
	return nil
}

// RemoveNode 删除节点
func (h *HashRing) RemoveNode(nodeName string) error {
	return nil
}

// GetNode 获取节点
func (h *HashRing) GetNode(key string) string {
	return "nodeName"
}

// AddNodes 批量增加节点
func (h *HashRing) AddNodes(nodeWeightMap map[string]int64) error {
	var err error

	for nodeName, nodeWeight := range nodeWeightMap {
		err = h.AddNode(nodeName, nodeWeight)
		if err != nil {
			break
		}
	}

	if err != nil {
		for nodeName := range nodeWeightMap {
			err = h.RemoveNode(nodeName)
			if err != nil {
				break
			}
		}
	}

	return err
}

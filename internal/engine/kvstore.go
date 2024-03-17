package engine

import "encoding/json"

type KVPair struct {
	Key   string
	Value string
}

func NewKVPair() *KVPair {
	return &KVPair{}
}

func (p *KVPair) ToBytes() []byte {
	bs, _ := json.Marshal(p)
	return bs
}

func (p *KVPair) FromBytes(bs []byte) error {
	return json.Unmarshal(bs, p)
}

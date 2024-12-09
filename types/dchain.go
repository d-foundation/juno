package types

import "encoding/json"

const VP_TYPE = "/dchain.tx.v1.MsgVerifiablePresentation"
const VP_MSG_INDEX = 0

type VPStandardMessage struct {
	Index int
	Type  string          `json:"@type"`
	Bytes json.RawMessage `json:"bytes"`
}

func NewVPStandardMessage(bytes []byte) *VPStandardMessage {
	return &VPStandardMessage{
		Index: VP_MSG_INDEX,
		Type:  VP_TYPE,
		Bytes: bytes,
	}
}

func (msg *VPStandardMessage) GetType() string {
	return msg.Type
}

func (msg *VPStandardMessage) GetIndex() int {
	return msg.Index
}

func (msg *VPStandardMessage) GetBytes() json.RawMessage {
	return msg.Bytes
}

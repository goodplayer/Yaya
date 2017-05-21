package global

import "github.com/goodplayer/yaya/idgen"

var _idgen idgen.Idgen

func Init(nodeId uint32) {
	_idgen = idgen.NewIdGen(nodeId)
}

func NextId() string {
	return idgen.ConvertToString(_idgen.Next())
}

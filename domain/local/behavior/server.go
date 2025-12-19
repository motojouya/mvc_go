package behavior

import (
	"github.com/motojouya/ddd_go/domain/local/core"
)

type ServerGetter interface {
	GetServer() (*core.Server, error)
}

type ServerGet struct {}

func NewServerGet() *ServerGet {
	return &ServerGet{}
}

var serverConf *core.Server

func (getter *ServerGet) GetServer() (*core.Server, error) {
	if serverConf == nil {
		var serverConfObj, err = GetEnv[core.Server]()
		if err != nil {
			return nil, err
		}

		serverConf = &serverConfObj
	}

	return serverConf, nil
}

package behavior

import (
	"github.com/motojouya/ddd_go/domain/database/core"
	localBehavior "github.com/motojouya/ddd_go/domain/local/behavior"
)

type DatabaseGetter interface {
	GetDatabase() (*core.ORP, error)
}

type DatabaseGet struct {}

func NewDatabaseGet() *DatabaseGet {
	return &DatabaseGet{}
}

var dbAccess *core.DBAccess

func (getter DatabaseGet) GetDatabase() (*database.ORP, error) {
	// access 情報はcacheするが、connectionはcacheしない
	if dbAccess == nil {
		var dbAccessData, err = localBehavior.GetEnv[*core.DBAccess]()
		if err != nil {
			return nil, err
		}
		dbAccess = &dbAccessData
	}

	var connection, err = dbAccess.CreateConnection()
	if err != nil {
		return nil, err
	}

	var db = database.CreateDatabase(connection)
	if err != nil {
		return db, err
	}

	return db, nil
}

package core

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	basic "github.com/motojouya/ddd_go/domain/basic/core"
)

// FIXME Prepare関数いる？

type Transactional interface {
	Begin() error
	Commit() error
	Rollback() error
}

type TransactionalDatabase interface {
	Transactional
	basic.Closable
}

type ORPer interface {
	TransactionalDatabase
	gorp.SqlExecutor
	Query
}

func CreateDatabase(connection *sql.DB) *ORP {
	var dbMap = &gorp.DbMap{Db: connection, Dialect: gorp.PostgresDialect{}}
	registerTable(dbMap)

	return &ORP{
		SqlExecutor: dbMap,
		dbMap:       nil,
	}
}

func registerTable(dbMap *gorp.DbMap) {
	// ここにテーブル登録を追加していく
}

// dbMapはtransactionを開始した際に、退避するためのフィールドなので、transactionが開始されていない場合はnil。
type ORP struct {
	gorp.SqlExecutor
	dbMap *gorp.DbMap
}

func (orp *ORP) Close() error {
	var insideTransaction = false

	if orp.dbMap != nil {
		insideTransaction = true
	}

	var dbMap, ok = orp.SqlExecutor.(*gorp.DbMap)
	if !ok {
		var err = orp.Rollback()
		if err != nil {
			return basic.CreateInsideTransactionError("transaction is not closed yet. and cannot closed transaction and connection.")
		}

		// rollbackしているので、`gorp.DbMap`になっているはず。失敗しているならいずれにしろcloseできないので、↑のreturnでerrorが返る。
		dbMap, ok = orp.SqlExecutor.(*gorp.DbMap)
		if !ok {
			return basic.CreateInsideTransactionError("transaction is not closed yet. and cannot closed transaction and connection.")
		}
		insideTransaction = true
	}

	var err = dbMap.Db.Close()
	if err != nil {
		return err
	}

	// closeは基本的に強制的に行うが、transactionが開いていた場合は、関数としてはエラーとする。
	if insideTransaction {
		return basic.CreateExitTransactionError("transaction is not closed yet. but closed transaction and connection already.")
	}

	return nil
}

func (orp *ORP) Begin() error {
	var dbMap, ok = orp.SqlExecutor.(*gorp.DbMap)
	if !ok || orp.dbMap != nil {
		return basic.CreateInsideTransactionError("transaction is already started")
	}

	var transaction, err = dbMap.Begin()
	if err != nil {
		return err
	}

	orp.SqlExecutor = transaction
	orp.dbMap = dbMap

	return nil
}

func (orp *ORP) Commit() error {
	var transaction, ok = orp.SqlExecutor.(*gorp.Transaction)
	if !ok || orp.dbMap == nil {
		return basic.CreateOutsideTransactionError("transaction is not started on Commit")
	}

	var err = transaction.Commit()
	if err != nil {
		return err
	}

	orp.SqlExecutor = orp.dbMap
	orp.dbMap = nil

	return nil
}

func (orp *ORP) Rollback() error {
	var transaction, ok = orp.SqlExecutor.(*gorp.Transaction)
	if !ok || orp.dbMap == nil {
		return basic.CreateOutsideTransactionError("transaction is not started on Rollback")
	}

	var err = transaction.Rollback()
	if err != nil {
		return err
	}

	orp.SqlExecutor = orp.dbMap
	orp.dbMap = nil

	return nil
}

func (orp *ORP) checkTransaction() error {
	var _, ok = orp.SqlExecutor.(*gorp.Transaction)
	if !ok || orp.dbMap == nil {
		return basic.CreateOutsideTransactionError("transaction is not started")
	}

	return nil
}

func (orp *ORP) InsideTransaction() bool {
	return orp.dbMap != nil
}

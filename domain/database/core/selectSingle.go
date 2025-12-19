package core

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/go-gorp/gorp"
	"github.com/motojouya/ddd_go/domain/basic/core"
)

var Dialect = goqu.Dialect("postgres")

/*
 * gorpのSelectOneは、レコードが見つからない場合にエラーになっちゃうがnilで返したいのでSelectSingleを作成
 */
func SelectSingle[R any](executer gorp.SqlExecutor, table string, keys map[string]string, query string, args ...interface{}) (*R, error) {
	var record []R
	var _, err = executer.Select(&record, query, args...)
	if err != nil {
		return nil, err
	}

	if len(record) == 0 {
		return nil, nil
	}

	if len(record) > 1 {
		return nil, core.NewDuplicateError(table, keys, "Duplicate record found")
	}

	return &record[0], nil
}

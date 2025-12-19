package controller

import (
	"github.com/motojouya/ddd_go/domain/database/core"
)

func RollbackWithError(database core.Transactional, err error) error {
	if rollbackErr := database.Rollback(); rollbackErr != nil {
		return rollbackErr
	}
	return err
}

// control処理の頭から最後までトランザクションとする場合に有効な関数。例えば、DBアクセスもするし、APIアクセスもして、トランザクションの粒度を操作したい場合は、control処理内でbegin/commitすべき
// FIXME templateなのでないが、操作者としての人格(Userとか)も引数にとれるようにすべき
func Transact[C core.TransactionalDatabase, E any, R any](callback func(C, E) (R, error)) func(C, E) (R, error) {
	return func(control C, entry E) (R, error) {

		var zero R
		if err := control.Begin(); err != nil {
			return zero, err
		}

		result, err := callback(control, entry)

		if err != nil {
			if rollbackErr := control.Rollback(); rollbackErr != nil {
				return zero, rollbackErr
			}
		} else {
			if commitErr := control.Commit(); commitErr != nil {
				return zero, commitErr
			}
		}

		return result, err
	}
}

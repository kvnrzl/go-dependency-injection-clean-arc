package helper

import "database/sql"

func TxCommitOrRollback(tx *sql.Tx, err error) {
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}

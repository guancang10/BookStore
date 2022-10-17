package helper

import (
	"database/sql"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrorTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		CheckError(errRollback)
		panic(err)
	} else {
		err := tx.Commit()
		CheckError(err)
	}
}

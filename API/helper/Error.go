package helper

import (
	"database/sql"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}

func CheckErrorTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		CheckError(errRollback)
		log.Fatalf("Error: %s", err)
	} else {
		err := tx.Commit()
		CheckError(err)
	}
}

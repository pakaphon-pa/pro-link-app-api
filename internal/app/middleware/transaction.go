package mdw

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBTransactionMdw(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("beginning database transaction")
		txHandle := db.Begin()

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if isStatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("commiting transaction")
			if err := txHandle.Commit(); err != nil {
				log.Print("transaction error: ", err)
			} else {
				log.Print("Rolling back transaction due to status code: ", c.Writer.Status())
				txHandle.Rollback()
			}

		}

	}
}

func isStatusInList(status int, statusList []int) bool {
	for _, s := range statusList {
		if status == s {
			return true
		}
	}

	return false
}

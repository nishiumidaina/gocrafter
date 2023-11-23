//go:generate mockgen -source=./transaction.go -destination=./transaction_mock.go -package=user
package user

import (
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Begin(shardKey int) (tx *gorm.DB, err error)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
}

package master

import (
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/config/database"
	repository "github.com/game-core/gocrafter/domain/repository/master"
)

type transactionDao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func NewTransactionDao(conn *database.SqlHandler) repository.TransactionRepository {
	return &transactionDao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
	}
}

func (d *transactionDao) Begin() (*gorm.DB, error) {
	tx := d.Write.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *transactionDao) Commit(tx *gorm.DB) error {
	tx.Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *transactionDao) Rollback(tx *gorm.DB) error {
	tx.Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

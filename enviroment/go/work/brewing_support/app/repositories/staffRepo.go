package repositories

import (
	"errors"
	"brewing_support/app/domain"

	"github.com/jmoiron/sqlx"
)

// SelectByAccountID select staaff data by accountID from database
func (s *StaffRepo) SelectByAccountID(id string) (staff *domain.Staff, err error) {
	if id == "" {
		return nil, errors.New("accountID must be required")
	}

	err = s.database.WithDbContext(func(db *sqlx.DB) error {

		queryStr := "select * " +
			"from staff s " +
			"where s.del != true and s.account_id = ?" +
			"limit 1"

		// クエリをDBドライバに併せて再構築
		queryStr = db.Rebind(queryStr)

		// データ取得処理
		err = db.Get(&staff, queryStr, id)

		return err
	})

	return
}

// NewStaffRepo constructor
func NewStaffRepo() *StaffRepo {
	return &StaffRepo{database: createDB()}
}

// StaffRepo repository struct
type StaffRepo struct {
	database db
}

package postgresql

import (
	"fmt"
	"brewing_support/app/utils/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jmoiron/sqlx"
	//sqldriver import
	_ "github.com/lib/pq"
)

// Postgresql DB接続のための構造体
type Postgresql struct {
	url      string
	port     string
	user     string
	password string
	dbname   string
}

// WithDbContext DB処理のためのラッパー関数
func (postgresql *Postgresql) WithDbContext(fn func(db *sqlx.DB) error) error {
	d, err := postgresql.Open()
	defer d.Close()

	// db openでエラー
	if err != nil {
		return err
	}

	// transaction開始
	var tx *sqlx.Tx
	tx, err = d.Beginx()
	if err != nil {
		return err
	}

	// 処理を実行
	err = fn(d)
	// 処理中にエラーだったらロールバック
	if err != nil {
		tx.Rollback()
		return err
	}

	// コミット
	tx.Commit()

	return err
}

// Open 接続情報は設定ファイルから読み込み
func (postgresql *Postgresql) Open() (*sqlx.DB, error) {
	dataSourceName := "host=" + postgresql.url +
		" port=" + postgresql.port +
		" user=" + postgresql.user +
		" password=" + postgresql.password +
		" dbname=" + postgresql.dbname +
		" sslmode=disable"
	return sqlx.Open("postgres", dataSourceName)

}

// Migrate DBスキーマ設定
func Migrate() {

	dataSourceName := config.DbUser() +
		":" + config.DbPassword() +
		"@" + config.DbUrl() +
		":" + config.DbPort() +
		"/" + config.DbName() +
		"?sslmode=disable"

	m, err := migrate.New(
		"file://./resources/migrations",
		"postgres://"+dataSourceName)

	if err != nil {
		fmt.Println(err)
	}
	m.Up()
}

// NewPostgresql コンストラクタ
func NewPostgresql() *Postgresql {
	p := &Postgresql{
		url:      config.DbUrl(),
		port:     config.DbPort(),
		user:     config.DbUser(),
		password: config.DbPassword(),
		dbname:   config.DbName(),
	}

	return p
}

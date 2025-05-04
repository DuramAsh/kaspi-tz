package store

import (
	_ "database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	// _ "github.com/sijms/go-ora/v2"
)

// postgres://username:password@localhost:5432/dbname?sslmode=disable&search_path=public
// oracle://username:password@:0/?connstr=(description=(address=(protocol=tcp)(host=localhost)(port=1521))(connect_data=(server=dedicated)(sid=dbname)))&persist security info=true&ssl=enable&ssl verify=false

func NewSQL(dataSourceName string) (db *sqlx.DB, err error) {
	if !strings.Contains(dataSourceName, "://") {
		return nil, errors.New("store: undefined data source name " + dataSourceName)
	}
	driverName := strings.ToLower(strings.Split(dataSourceName, "://")[0])

	for _ = range 5 {
		db, err = sqlx.Connect(driverName, dataSourceName)
		if err == nil {
			var pingErr error
			if pingErr = db.Ping(); pingErr == nil {
				break
			}

			err = pingErr
		}

		time.Sleep(time.Second * 2)
	}

	if err != nil {
		return
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return
}

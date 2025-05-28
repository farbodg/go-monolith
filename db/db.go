package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BaseDataProvider struct {
	DB *bun.DB
}

func ConnectToDB(ctx context.Context, conf Config) (BaseDAL, error) {
	var pgDB *bun.DB

	for numTries := uint16(0); numTries < conf.RetryNumTimes; numTries++ {
		sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.PostgreSQLConnectionString)))
		pgDB = bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())

		if err := pgDB.PingContext(ctx); err == nil {
			return &BaseDataProvider{DB: pgDB}, nil
		} else {
			fmt.Printf("got error pinging database: %s\n", err)
		}

		time.Sleep(conf.RetrySleepTime)
	}

	return nil, fmt.Errorf("could not connect to database after %d retries", conf.RetryNumTimes)
}

func (d *BaseDataProvider) GetDB() *bun.DB {
	return d.DB
}

func (d *BaseDataProvider) Close() error {
	return d.DB.Close()
}

func (d *BaseDataProvider) DoInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.DB.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(ctx)
	})
}

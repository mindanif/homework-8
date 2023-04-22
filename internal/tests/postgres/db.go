//go:build integration
// +build integration

package postgres

import (
	"context"
	"fmt"
	"homework-5/internal/pkg/db"
	"homework-5/internal/tests/config"
	"strings"
	"sync"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
)

type TBD struct {
	sync.Mutex
	DB *db.Database
}

func NewFormEnv() *TBD {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(err)
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, port, user, password, dbname)
	db, err := db.NewDB(context.Background(), psqlConn)
	if err != nil {
		panic(err)
	}
	return &TBD{DB: db}
}

func (d *TBD) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	d.Lock()
	d.Truncate(ctx)
}

func (d *TBD) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}
func (d *TBD) Truncate(ctx context.Context) {
	var tables []string

	err := d.DB.Select(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_name != 'goose_db_version'")
	if err != nil {
		panic(err)
	}
	if len(tables) == 0 {
		panic("run migration plz")
	}
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := d.DB.Exec(ctx, q); err != nil {
		panic(err)
	}
}

//go:build integration
// +build integration

package tests

import "homework-5/internal/tests/postgres"

var (
	Db *postgres.TBD
)

func init() {
	Db = postgres.NewFormEnv()
}

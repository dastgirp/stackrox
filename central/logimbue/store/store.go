package store

import (
	bolt "github.com/etcd-io/bbolt"
	"github.com/stackrox/rox/pkg/bolthelper"
)

var logsBucket = []byte("logs")

// Store provides storage functionality for alerts.
//go:generate mockgen-wrapper Store
type Store interface {
	GetLogs() ([]string, error)
	GetLogsRange() (start int64, end int64, err error)
	AddLog(log string) error
	RemoveLogs(from, to int64) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) Store {
	bolthelper.RegisterBucketOrPanic(db, logsBucket)
	return &storeImpl{
		DB: db,
	}
}

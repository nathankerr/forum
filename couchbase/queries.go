package couchbase

import "github.com/couchbase/gocb"

func Get(key string, valPtr interface{}) error {
	_, err := db.bucket.Get(key, valPtr)
	return err
}

func GetAll(list []string, valPtr interface{}) ([]gocb.BulkOp, error) {
	var items []gocb.BulkOp
	for i := 0; i < len(list); i++ {
		items = append(items, &gocb.GetOp{Key: list[i], Value: valPtr})
	}
	err := db.bucket.Do(items)
	return items, err
}

func Insert(key string, value interface{}, expiry uint32) error {
	_, err := db.bucket.Insert(key, value, expiry)
	return err
}

func Upsert(key string, value interface{}, expiry uint32) error {
	_, err := db.bucket.Upsert(key, value, expiry)
	return err
}

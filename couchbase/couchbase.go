package couchbase

import "github.com/couchbase/gocb"

type database struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

var db database

func Connect(url string) error {
	cluster, err := gocb.Connect(url)
	db = database{cluster: cluster}
	return err
}

func OpenBucket(bucket string, password string) error {
	var err error
	db.bucket, err = db.cluster.OpenBucket(bucket, password)
	return err
}

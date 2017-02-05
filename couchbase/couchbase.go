package couchbase

import "github.com/couchbase/gocb"

type database struct {
	cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	Err     error
}

var DB database

func Connect(url *string) {
	cluster, err := gocb.Connect(*url)
	DB = database{
		cluster: cluster,
		Err:     err,
	}
}

func OpenBucket(bucket *string, password *string) {
	DB.Bucket, DB.Err = DB.cluster.OpenBucket(*bucket, *password)
}

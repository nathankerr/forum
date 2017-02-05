package couchbase

import "github.com/couchbase/gocb"

type couchbase struct {
	cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	Err     error
}

var Couchbase couchbase

func Connect(url *string) {
	cluster, err := gocb.Connect(*url)
	Couchbase = couchbase{
		cluster: cluster,
		Err:     err,
	}
}

func OpenBucket(bucket *string, password *string) {
	Couchbase.Bucket, Couchbase.Err = Couchbase.cluster.OpenBucket(*bucket, *password)
}

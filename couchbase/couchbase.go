package couchbase

import "github.com/couchbase/gocb"

type Couchbase struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	Err     error
}

func Connect(url string) *Couchbase {
	cluster, err := gocb.Connect(url)
	return &Couchbase{
		Cluster: cluster,
		Err:     err,
	}
}

func (cb *Couchbase) OpenBucket(bucket string, password string) {
	cb.Bucket, cb.Err = cb.Cluster.OpenBucket(bucket, password)
}

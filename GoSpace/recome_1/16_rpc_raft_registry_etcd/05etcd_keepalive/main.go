package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	// connet etcd, create cli
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// set 续期 5s
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// set k-v to etcd
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// auto 续期
	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}

	for {
		data := <-ch
		fmt.Println(data)
	}

}

// Grant creates a new lease.
// Grant(ctx context.Context, ttl int64) (*LeaseGrantResponse, error)

// WithLease attaches a lease ID to a key in 'Put' request.
// func WithLease(leaseID LeaseID) OpOption {
// 	return func(op *Op) { op.leaseID = leaseID }
// }

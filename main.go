package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	cli.Put(context.Background(), "ab", "8888888")
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cancel()
	defer cli.Close()
	for {
		rch := cli.Watch(context.Background(), "ab")
		for wresp := range rch {
			fmt.Println(wresp)
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}

}

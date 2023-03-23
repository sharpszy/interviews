package main

import (
	"fmt"
	"testing"
	"time"
)

func invoke_proc_every_secs() {
	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			go func() {
				defer func() {
					if ex := recover(); ex != nil {
						fmt.Println(ex)
					}
				}()
				proc()
			}()
		}
	}
}

func proc() {
	panic("ok")
}

func Test_invoke_proc_every_secs(t *testing.T) {
	go invoke_proc_every_secs()
	select {}
}

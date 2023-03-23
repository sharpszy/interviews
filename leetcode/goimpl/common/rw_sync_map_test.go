package common

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test_RWMap(t *testing.T) {
	m := NewRWSyncMap[string, int]()
	var (
		wc  = 100000
		rc  = 50000
		dc  = 100
		wg1 sync.WaitGroup
		wg2 sync.WaitGroup
		wg3 sync.WaitGroup
	)
	for i := 0; i < wc; i++ {
		wg1.Add(1)
		// write map
		go func(i int) {
			defer wg1.Done()
			k, v := fmt.Sprintf("k%v", i), i
			m.Put(k, v)
			if i%100 == 0 {
				fmt.Printf("Put-{%v:%v}\n", k, v)
			}
		}(i)

		// read map
		if i < rc {
			wg2.Add(1)
			go func(i int) {
				defer wg2.Done()
				k := fmt.Sprintf("k%v", i)
				v := m.Get(k)
				if i%100 == 0 {
					fmt.Printf("Get-{%v:%v}\n", k, v)
				}
			}(i)
		}

		// delete map
		if i < dc {
			wg3.Add(1)
			go func() {
				defer wg3.Done()
				// time.Sleep(time.Microsecond)
				r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
				k := fmt.Sprintf("k%v", r.Intn(100))
				v := m.Delete(k)
				fmt.Printf("Delete-{%v:%v}\n", k, v)
			}()
		}
	}
	wg1.Wait()
	wg2.Wait()
	wg3.Wait()
	// assert.Equal(t, wc-dc, m.Len())
	fmt.Println(m.Len())
}

func BenchmarkPut(b *testing.B) {
	m := NewRWSyncMap[string, int]()
	var (
		wc  = 100000
		rc  = 5000000
		dc  = 10000
		wg1 sync.WaitGroup
		wg2 sync.WaitGroup
		wg3 sync.WaitGroup
	)
	for i := 0; i < wc; i++ {
		wg1.Add(1)
		// write map
		go func(i int) {
			defer wg1.Done()
			k, v := fmt.Sprintf("k%v", i), i
			m.Put(k, v)
		}(i)

		// read map
		if i < rc {
			wg2.Add(1)
			go func(i int) {
				defer wg2.Done()
				k := fmt.Sprintf("k%v", i)
				m.Get(k)
			}(i)
		}

		// delete map
		if i < dc {
			wg3.Add(1)
			go func() {
				defer wg3.Done()
				r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
				k := fmt.Sprintf("k%v", r.Intn(100))
				m.Delete(k)
			}()
		}
	}
	wg1.Wait()
	wg2.Wait()
	wg3.Wait()
}

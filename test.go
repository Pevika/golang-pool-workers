//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package main

import (
    "./pool"
    "log"
    "time"
)

func main () {
    _pool := pool.Create(8, func (v interface{}) {
        log.Println("Working")
        id, _ := v.(int)
        log.Println(id)
        time.Sleep(10 * time.Millisecond)
    })
    for i := 0 ; i < 15000 ; i++ {
        _pool.AddJob(i)
    }
    _pool.Start()
}
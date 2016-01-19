//
// @Author: Geoffrey Bauduin <bauduin.geo@gmail.com>
//

package pool

import (
    "sync"
)

type Pool struct {
    nbrRoutines     int
    handler         func (interface{})
    jobs            []interface{}
    jobMutex        *sync.Mutex
}

// Creates a new pool
func Create (routines int, handler func (interface{})) *Pool {
    pool := &Pool{
        nbrRoutines: routines,
        handler: handler,
        jobMutex: &sync.Mutex{},
    }
    return pool
}

func (this *Pool) AddJob (j interface{}) {
    this.jobs = append(this.jobs, j)
}

func (this *Pool) launchRoutine () {
    for {
        this.jobMutex.Lock()
        var job interface{}
        chosen := false
        if len(this.jobs) > 0 {
            job, this.jobs = this.jobs[0], this.jobs[1:]
            chosen = true
        }            
        this.jobMutex.Unlock()
        if chosen == false {
            break
        }
        this.handler(job)
    }
}

func (this *Pool) Start () error {
    wg := sync.WaitGroup{}
    for i := 0 ; i < this.nbrRoutines ; i++ {
        wg.Add(1)
        go func () {
            defer wg.Add(-1)
            this.launchRoutine()
        }()
    }
    wg.Wait()
    return nil
}
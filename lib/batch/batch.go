package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	sem := make(chan int, pool)
	ch := make(chan user)

	for i := 0; i < int(n); i++ {
		go func(i int64) {
			sem <- 1
			ch <- getOne(i)
			<-sem
		}(int64(i))
	}

	for i := 0; i < int(n); i++ {
		user := <-ch
		res = append(res, user)
	}

	return
}

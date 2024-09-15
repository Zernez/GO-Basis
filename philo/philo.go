package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Chops struct {
	sync.Mutex
}

type Philo struct {
	leftCh, rightCh *Chops
	number          int
	meals           int
}

func eat(philo *Philo, chans chan int, wg *sync.WaitGroup) {

	fmt.Println("Starting to eat #: ", philo.number)

	philo.rightCh.Lock()

	philo.leftCh.Lock()

	chans <- philo.meals + 1

	fmt.Println("Finishing eating #: ", philo.number)

	philo.rightCh.Unlock()

	philo.leftCh.Unlock()

	wg.Done()

}

func main() {

	Sticks := make([]*Chops, 6)

	for i := 0; i < 6; i++ {

		Sticks[i] = new(Chops)
	}

	Philos := make([]*Philo, 6)

	for i := 1; i < 6; i++ {

		Philos[i] = &Philo{Sticks[i], Sticks[(i+1)%6], i, 0}
	}

	var wg sync.WaitGroup

	var chans = []chan int{make(chan int), make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)}

	rand.Seed(time.Now().UnixNano())

	rand1 := rand.Intn(5) + 1

	rand2 := rand.Intn(5) + 1

	for rand2 == rand1 {

		rand2 = rand.Intn(5) + 1
	}

	waitList := 2

	for {

		wg.Add(1)

		go eat(Philos[rand1], chans[rand1], &wg)

		a := <-chans[rand1]

		Philos[rand1].meals = a

		wg.Wait()

		if waitList > 1 {

			wg.Add(1)

			go eat(Philos[rand2], chans[rand2], &wg)

			b := <-chans[rand2]

			Philos[rand2].meals = b

			wg.Wait()
		}

		if Philos[1].meals > 2 && Philos[2].meals > 2 && Philos[3].meals > 2 && Philos[4].meals > 2 && Philos[5].meals > 2 {

			break
		}

		waitList = 0

		order := false

		for waitList == 0 {

			if waitList == 0 && !order {

				rand.Seed(time.Now().UnixNano())

				rand1 = rand.Intn(5) + 1

				if Philos[rand1].meals < 3 {

					waitList++

					order = true
				}
			}

			rand2 := rand.Intn(5) + 1

			for rand2 == rand1 && order {

				rand.Seed(time.Now().UnixNano())

				rand2 = rand.Intn(5) + 1

				if Philos[rand2].meals < 3 {

					waitList++
				}
			}
		}

	}

	fmt.Printf("Eat complete")
}

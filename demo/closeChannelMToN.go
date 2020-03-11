package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumReceivers = 100
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	toStop := make(chan string, 1)

	var stoppedBy string
	// moderator
	go func() {
		stoppedBy = <-toStop // part of the trick used to notify the moderator
		// to close the additional signal channel.
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(MaxRandomNumber)
				if value == 0 {
					select {
					case toStop <- "sender#" + id:
					default:
					}
					return
				}

				select {
				case <-stopCh:
					return
				default:
				}

				select {
				case <-stopCh:
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}
	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			defer wgReceivers.Done()

			select {
			case <-stopCh:
				return
			default:
			}

			select {
			case <-stopCh:
				return
			case value := <-dataCh:
				if value == MaxRandomNumber-1 {
					select {
					case toStop <- "receiver#" + id:
					default:
					}
					return
				}
				log.Println(value)
			}
		}(strconv.Itoa(i))
	}
	wgReceivers.Wait()
}

package selects

import "net/http"

func Racer(a, b string) (winner string) {
	select {
		case <-ping(a):
			return a
		case <-ping(b):
			return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func ()  {
		http.Get(url)
		close(ch)
	}()
	return ch
}
// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// func Racer(a, b string) (winner string) {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)

// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }
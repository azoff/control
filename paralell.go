package control

type ParallelChan chan<- error
type ParallelFunc func(done ParallelChan)

func Parallel(tasks ...ParallelFunc) (err error) {
	total := len(tasks)
	buffer := make(chan error, total)
	for _, task := range(tasks) {
		go task(buffer)
	}
	for i := 0; i < total; i++ {
		if err = <-buffer; err != nil {
			return
		}
	}
	return
}

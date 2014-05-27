package control

type ParallelChan chan<- error
type ParallelFunc func(ParallelChan)

func SerialToParallelFunc(serialFunc SerialFunc) ParallelFunc {
	return func(done ParallelChan) {
		done <- serialFunc()
	}
}

func ParallelSlice(tasks []ParallelFunc) (err error) {
	total := len(tasks)
	buffer := make(chan error, total)
	for _, task := range tasks {
		go task(buffer)
	}
	for i := 0; i < total; i++ {
		if err = <-buffer; err != nil {
			return
		}
	}
	return
}

func Parallel(tasks ...ParallelFunc) error {
	return ParallelSlice(tasks)
}

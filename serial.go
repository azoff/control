package control

type SerialFunc func() error

func Serial(chain ...SerialFunc) (err error) {
	for _, link := range chain {
		if err = link(); err != nil {
			return
		}
	}
	return
}

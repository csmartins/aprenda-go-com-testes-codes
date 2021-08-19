package iteractions

func Loop(c string, times int) (repetitions string){
	for i := 0; i < times; i++ {
		repetitions += c
	}
	return
}
package sol


func Max( f int32, args... int32) int32 {
	for _ , v := range args{
		if f < v {
			f = v
		}
	}
	return f
}

func Min( f int32, args... int32) int32 {
	for _ , v := range args{
		if f > v {
			f = v
		}
	}
	return f
}
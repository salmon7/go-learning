package main

func main() {
	var xsize, ysize = 100, 100
	picture := make([][]uint64, ysize)
	for i := range picture {
		picture[i] = make([]uint64, xsize)
	}

	var xsize1, ysize1 = 200, 200
	picture1 := make([][]uint64, ysize1)
	row1 := make([]uint64, xsize1*ysize1)
	for i := range picture1 {
		picture1[i] = row1[:xsize1]
		row1 = row1[xsize1:]
	}

}

/*
二维slice的初始化，两种方式
*/

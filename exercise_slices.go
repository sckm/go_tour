package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy, dy)
    for i:=0;i<len(img);i++{
        row := make([]uint8, dx)
        for j:=0;j<len(row);j++{
            row[j] = colorFunc(j, i)
        }
        img[i] = row
    }
	return img
}

func colorFunc(x, y int) uint8 {
    return uint8(x^y)
}

func main() {
	pic.Show(Pic)
}

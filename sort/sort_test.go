package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	Convey("TestSort", t, func() {
		Convey("TestBubbleSort", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			BubbleSort(list)
			log.Println("bubble:", list)
		})

		Convey("TestSelectSort", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			SelectSort(list)
			log.Println("select:", list)
		})

		Convey("TestSelectGoodSort", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			SelectGoodSort(list)
			log.Println("good select:", list)
		})

		Convey("TestInsertSort", func() {
			list := []int{10, 88, 2, 199, 1, 199, -19, 99}
			InsertSort(list)
			log.Println("insert:", list)
		})

		Convey("TestShellSort", func() {
			list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
			ShellSort(list)
			log.Println("shell:", list)
		})
	})
}

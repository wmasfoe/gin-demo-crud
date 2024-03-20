package pkg

import "fmt"

func BubbleSorted(arr []int) []int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("BubbleSorted Error!")
		}
	}()

	for i := 1; i < len(arr); i++ {

		for j := 1; j < len(arr) - i + 1; j++ {
			prevItem := &arr[j-1]
			currentItem := &arr[j]

			if *prevItem > *currentItem {
				temp := *prevItem
				*prevItem = *currentItem
				*currentItem = temp
			}
		}
	}

	return arr
}

func Fbn(n int) []int {
	var sli []int

	for i := 0; i <= n; i++ {
		if i == 0 {
			sli = append(sli, 1)
			continue
		}
		if i == 1 {
			sli = append(sli, 1)
			continue
		}

		sli = append(sli, sli[i - 1] + sli[i - 2])
	}

	return sli
}

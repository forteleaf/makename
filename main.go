package main

import "fmt"

const (
	StrStroke = 1
	EndStroke = 40
)

// 3개의 각각의 합을 구하기
func sum(a, b, c int) {
	// arr := []int{}
	result1 := a + 1     // 천격
	result2 := a + b     // 인격
	result3 := b + c     // 지격
	result4 := c + 1     // 외격
	result5 := a + b + c // 총격

	if result5 > 81 {
		result5 = result5 - 80
	}
	func() {
		fmt.Println("[", a, b, c, "]")
		fmt.Println("천격", result1)
		fmt.Println("인격", result2)
		fmt.Println("지격", result3)
		fmt.Println("외격", result4)
		fmt.Println("총격", result5)
	}()

	if BadorGood(result1) {
		// fmt.Println("천격통과")
		if BadorGood(result5) {
			fmt.Println("총격통과")
			fmt.Println("[", a, b, c, "]")
			if BadorGood(result3) {
				fmt.Println("지격통과")
				if BadorGood(result4) {
					fmt.Println("외격통과")
					if BadorGood(result2) {
						fmt.Println("중격통과")
					}
				}
			}
		}
	}

}

func BadorGood(num int) bool {

	// goodArr := []int{}
	var numPerfect = []int{13, 16, 21, 23, 31, 41}
	// var numVerygood = []int{1, 3, 5, 6, 11, 15, 18, 24, 32, 35, 37, 39}
	// var numGood = []int{7, 8, 17, 25, 29, 33, 38, 45, 47, 48, 52, 57, 58, 61, 63, 65, 67, 68, 71, 73, 75, 77, 81}
	// var numBad = []int{2, 4, 12, 14, 27, 28, 30, 53, 59, 60, 62, 69, 72}
	// var numHell = []int{9, 10, 19, 20, 22, 26, 34, 36, 40, 42, 43, 44, 46, 50, 54, 56, 64, 66, 70, 74, 78, 80}

	for _, v := range numPerfect {
		if num == v {
			// fmt.Println("insert num is", num)
			// goodArr = append(goodArr, num)
			return true
		}
	}
	return false
}

// 첫획을 넣으면 나머지 좋은 획수 구하기
func InputNum() {
	var n1, n2, n3 int
	fmt.Print("Enter number:")
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)
	sum(n1, n2, n3)
}
func main() {
	var wordA = 12 // 황
	// reader := bufio.NewReader(os.Stdin)
	for i := StrStroke; i <= EndStroke; i++ {
		for j := StrStroke; j <= EndStroke; j++ {
			sum(wordA, i, j)
		}
	}

}

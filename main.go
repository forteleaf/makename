package main

import "fmt"

const (
	StrStroke = 1  // StrStroke is start stroke
	EndStroke = 25 // EndStroke is lastest stroke
)

// 3개의 각각의 합을 구하기
func sum(a, b, c int) {
	// arr := []int{}
	result1 := a + 1     // 천격
	result2 := a + b     // 인격, 형격
	result3 := b + c     // 지격, 원격
	result4 := c + 1     // 외격
	result5 := a + b + c // 총격, 정격
	result6 := a + c     // 이격

	if result5 > 81 {
		result5 = result5 - 80
	}

	func() {
		_ = result6 // nothing
		if GoodOrNot(result1) {
			if GoodOrNot(result5) {
				if GoodOrNot(result3) {
					if GoodOrNot(result4) {
						if GoodOrNot(result2) {
							fmt.Print("[", a, b, c, "]")
						}
					}
				}
			}
		}
	}()

	// 	(1) 수리격
	// 성과 이름을 구성하는 글자들의 획수가 길한수인가, 흉한 수인가를 구분하여 길한수로 작명하는 방법입니다.
	// 각 수가 지닌 기운에 근거하니 참고하시길 바랍니다.
	// 원격(元格) : 이름첫자와 둘째자 합 (초년운세) --- 명격,지격등으로도 부름
	// 형격(亨格) : 성과 이름 첫자의 합 (청년운세) --- 주격,인격등으로도 부름
	// 이격(利格) : 성과 이름 끝자의 합 (중년운세) --- 외격으로도 부름
	// 정격(貞格) : 성과 이름첫자, 둘째자 모두의 합 (말년운세) --- 총격으로도 부름
	// func() {
	// 	if GoodOrNot(result2) {
	// 		if GoodOrNot(result3) {
	// 			if GoodOrNot(result6) {
	// 				if GoodOrNot(result5) {
	// 					fmt.Print("[", a, b, c, "]")
	// 				}
	// 			}
	// 		}
	// 	}
	// }()
}

//  GoodOrNot is check good number or bad number
func GoodOrNot(num int) bool {
	//  http://www.finename.co.kr/gnu/bbs/board.php?bo_table=m21&wr_id=28 자료를 통해서 정리
	var ok bool = false
	var arr = []int{}
	var numPerfect = []int{13, 16, 21, 23, 31, 32, 33, 41}          // 최상운수
	var numVerygood = []int{1, 3, 5, 6, 11, 15, 18, 24, 35, 37, 39} // 상운수
	// var numGood = []int{7, 8, 17, 25, 29, 38, 45, 47, 48, 52, 57, 58, 61, 63, 65, 67, 68, 71, 73, 75, 77, 81} // 양운수
	// var numBad = []int{2, 4, 12, 14, 27, 28, 30, 53, 59, 60, 62, 69, 72} // 흉운수
	// var numHell = []int{9, 10, 19, 20, 22, 26, 34, 36, 40, 42, 43, 44, 46, 50, 54, 56, 64, 66, 70, 74, 78, 80} // 흉흉운수
	arr = append(numPerfect, numVerygood...)
	// arr = append(arr, numGood...)
	// arr = numPerfect

	// compare number which good or not
	for _, v := range arr {
		if num == v {
			ok = true
		}
	}
	return ok
}

// InputNum is via Scanner input
func InputNum() {
	var n1, n2, n3 int
	fmt.Print("Enter number:")
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)
	sum(n1, n2, n3)
}

// 첫획을 넣으면 나머지 좋은 획수 구하기
func main() {
	// var wordA = 12 // 황
	// reader := bufio.NewReader(os.Stdin)
	for i := StrStroke; i <= EndStroke; i++ {
		for j := StrStroke; j <= EndStroke; j++ {
			for k := StrStroke; k <= EndStroke; k++ {
				sum(i, j, k)
			}
		}

		fmt.Println()
	}

}

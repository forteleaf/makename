package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	// StrStroke is start stroke
	StrStroke = 2
	// EndStroke is lastest stroke
	EndStroke = 25
)

// StrokeName is name
type StrokeName struct {
	a, b, c int
}

// HanjaDataType is 성명한자에 사용되는 것들
type HanjaDataType struct {
	//  "한자","획오행","자원오행","획수","뜻","부수","발음"
	// ["加","土","水",5,"더할,있을","力","가"]
	Hanja         string
	WuxingStroke  string
	WuxingMean    string
	StrokeCount   string
	Mean          string
	Radical       string
	Pronunciation string
}

// 3개의 각각의 합을 구하기
func sum(a, b, c int) {
	// arr := []int[]
	result1 := a + 1     // 천격
	result2 := a + b     // 인격, 형격
	result3 := b + c     // 지격, 원격
	result4 := 1 + c     // 외격, 이격
	result5 := a + b + c // 총격, 정격
	if result5 >= 81 {
		result5 = result5 - 80
	}
	// fmt.Println("천격", result1, "인격", result2, "지격", result3, "외격", result4, "총격", result5, "")
	func() {
		if GoodOrNot(result1) {
			if GoodOrNot(result5) {
				if GoodOrNot(result3) {
					if GoodOrNot(result4) {
						if GoodOrNot(result2) {
							fmt.Print("{", a, b, c, "}")
						}
					}
				}
			}
		}
	}()

	// 	(1) 수리격 (수리오행)
	// 성과 이름을 구성하는 글자들의 획수가 길한수인가, 흉한 수인가를 구분하여 길한수로 작명하는 방법입니다.
	// 각 수가 지닌 기운에 근거하니 참고하시길 바랍니다.
	// 원격(元格) : 이름첫자와 둘째자 합 (초년운세) --- 명격,지격등으로도 부름
	// 형격(亨格) : 성과 이름 첫자의 합 (청년운세) --- 주격,인격등으로도 부름
	// 이격(利格) : 성과 이름 끝자의 합 (중년운세) --- 외격으로도 부름
	// 정격(貞格) : 성과 이름첫자, 둘째자 모두의 합 (말년운세) --- 총격으로도 부름
	// func() { // 수리오행
	// 	if GoodOrNot(result2) {
	// 		if GoodOrNot(result3) {
	// 			if GoodOrNot(result5) {
	// 				if GoodOrNot(result5) {
	// 					fmt.Print("{", a, b, c, "}")
	// 				}
	// 			}
	// 		}
	// 	}
	// }()
}

// ShowKe print 천격 인격, 지격 외격, 총격
func (stroke *StrokeName) ShowKe() {
	result1 := stroke.a + 1                   // 천격
	result2 := stroke.a + stroke.b            // 인격, 형격
	result3 := stroke.b + stroke.c            // 지격, 원격
	result4 := 1 + stroke.c                   // 외격, 이격
	result5 := stroke.a + stroke.b + stroke.c // 총격, 정격
	if result5 >= 81 {
		result5 = result5 - 80
	}
	fmt.Println(
		"천격", result1, GoodOrNot(result1),
		"\n인격", result2, GoodOrNot(result2),
		"\n지격", result3, GoodOrNot(result3),
		"\n외격", result4, GoodOrNot(result4),
		"\n총격", result5, GoodOrNot(result5),
	)
}
func (stroke *StrokeName) CheckKe() {

}

// GoodOrNot check waht insert number
func GoodOrNot(num int) bool {
	ok := false
	var arr = []int{}
	var numPerfect = []int{13, 16, 21, 23, 31, 32, 33, 41}                                                    // 최상운수
	var numVerygood = []int{1, 3, 5, 6, 11, 15, 18, 24, 35, 37, 39}                                           // 상운수
	var numGood = []int{7, 8, 17, 25, 29, 38, 45, 47, 48, 52, 57, 58, 61, 63, 65, 67, 68, 71, 73, 75, 77, 81} // 양운수
	// var numBad = []int{2, 4, 12, 14, 27, 28, 30, 53, 59, 60, 62, 69, 72} // 흉운수
	// var numHell = []int{9, 10, 19, 20, 22, 26, 34, 36, 40, 42, 43, 44, 46, 50, 54, 56, 64, 66, 70, 74, 78, 80} // 흉흉운수
	// jonghabGood := []int{1, 3, 5, 6, 7, 8, 11, 13, 15, 16, 17, 18, 21, 23, 24, 25, 29, 31, 32, 33, 35, 37, 38, 39, 41, 45, 47, 48, 52, 57, 58, 61, 63, 65, 67, 68, 71, 73, 75, 77, 81}
	// jonghabBad := []int{2, 4, 9, 10, 12, 14, 19, 22, 26, 27, 28, 30, 34, 36, 40, 42, 43, 44, 46, 49, 50, 51, 53, 54, 55, 56, 59, 60, 62, 64, 66, 69, 70, 72, 74, 76, 78, 79, 80}
	///////---------------------------////////

	arr = append(numPerfect, numVerygood...)
	arr = append(arr, numGood...)
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

func johabname() {
	names := []string{}
	var tmp string
	// ㅅㅈㅊ , ㅁㅂㅍ
	s := []string{"서", "선", "성", "세", "소", "수"}

	// j := []string{"자", "작", "잔", "잖", "잗", "잘", "잚", "잠", "잡", "잣", "잣", "장", "잦", "재", "잭", "잰", "잴", "잼", "잽", "잿", "쟀", "쟁", "쟈", "쟉", "쟌", "쟎", "쟐", "쟘", "쟝", "쟤", "쟨", "쟬", "저", "적", "전", "절", "젊", "점", "접", "젓", "정", "젖", "제", "젝", "젠", "젤", "젬", "젭", "젯", "젱", "져", "젼", "졀", "졈", "졉", "졌", "졍", "졔", "조", "족", "존", "졸", "졺", "좀", "좁", "좃", "종", "좆", "좇", "좋", "좌", "좍", "좔", "좝", "좟", "좡", "좨", "좼", "좽", "죄", "죈", "죌", "죔", "죕", "죗", "죙", "죠", "죡", "죤", "죵", "주", "죽", "준", "줄", "줅", "줆", "줌", "줍", "줏", "중", "줘", "줬", "줴", "쥐", "쥑", "쥔", "쥘", "쥠", "쥡", "쥣", "쥬", "쥰", "쥴", "쥼", "즈", "즉", "즌", "즐", "즘", "즙", "즛", "증", "지", "직", "진", "짇", "질", "짊", "짐", "집", "짓", "징", "짖", "짙", "짚"}
	// ch := []string{"차", "착", "찬", "찮", "찰", "참", "찹", "찻", "찼", "창", "찾", "채", "책", "챈", "챌", "챔", "챕", "챗", "챘", "챙", "챠", "챤", "챦", "챨", "챰", "챵", "처", "척", "천", "철", "첨", "첩", "첫", "첫", "청", "체", "첵", "첸", "첼", "쳄", "쳅", "쳇", "쳉", "쳐", "쳔", "쳣", "체", "촁", "초", "촉", "촌", "촐", "촘", "촙", "촛", "총", "촤", "촨", "촬", "촹", "최", "쵠", "쵤", "쵬", "쵭", "쵯", "쵱", "쵸", "춈", "추", "축", "춘", "출", "춤", "춥", "춧", "충", "취", "췄", "췌", "췐", "취", "췬", "췰", "췸", "췹", "췻", "췽", "츄", "츈", "츌", "츔", "츙", "츠", "측", "츤", "츨", "츰", "츱", "츳", "층", "치", "칙", "친", "칟", "칠", "칡", "침", "칩", "칫", "칭"}
	// m := []string{"마", "막", "만", "맏", "말", "매", "머", "먹", "먼", "메", "멕", "멘", "며", "멱", "면", "명", "모", "목", "몬", "뫼", "묘", "무", "묵", "문", "물", "뭄", "뭇", "뭉", "므", "믄", "미", "믹", "민", "밀", "밈", "밍"}
	m := []string{"마", "만", "매", "머", "메", "며", "멱", "면", "명", "모", "목", "몬", "무", "문", "미", "믹", "민", "밀"}
	b := []string{"보", "본", "부", "북", "분", "불", "비", "빈"}

	// p := []string{"파", "팍", "팍", "판", "팔", "팖", "팜", "팝", "팟", "팠", "팡", "팥", "패", "팩", "팬", "팰", "팸", "팹", "팻", "팼", "팽", "퍄", "퍅", "퍼", "퍽", "펀", "펄", "펌", "펍", "펏", "펐", "펑", "페", "펙", "펜", "펠", "펨", "펩", "펭", "펴", "편", "펼", "폄", "폅", "폈", "평", "폐", "폘", "폡", "폣", "포", "폭", "폰", "폴", "폼", "폽", "폿", "퐁", "퐈", "퐝", "푀", "푄", "표", "푠", "푤", "푭", "푯", "푸", "푹", "푼", "푿", "풀", "풂", "품", "풉", "풋", "풍", "풔", "풩", "퓌", "퓐", "퓔", "퓜", "퓟", "퓨", "퓬", "퓰", "퓸", "퓻", "퓽", "프", "픈", "플", "픔", "픕", "픗", "피", "픽", "핀", "필", "핌", "핍", "핏", "핑"}

	for _, v := range s {
		for _, w := range b {
			tmp = v + w
			names = append(names, tmp)
		}
		for _, w := range m {
			tmp = v + w
			names = append(names, tmp)
		}
	}
	fmt.Println(names)

}

func testMarshall() {
	file, err := ioutil.ReadFile("./data/peoplehanja.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var hanjaData HanjaDataType
	err = json.Unmarshal(file, &hanjaData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", hanjaData)
}
func main() {
	t0 := time.Now()
	// testMarshall()
	// reader := bufio.NewReader(os.Stdin)
	// for i := StrStroke; i <= EndStroke; i++ {
	for j := StrStroke; j <= EndStroke; j++ {
		// fmt.Printf("\n %d획 -----------------------------\n", j)
		for k := StrStroke; k <= EndStroke; k++ {
			sum(12, j, k)
		}
		fmt.Printf("\n")
		// }
	}
	var sample1 StrokeName
	sample1.a = 12
	sample1.b = 21
	sample1.c = 14

	// 황수민綏
	// 綏 13
	// 敏 11

	// 황은호
	// 誾 15
	// 顥 21
	sample1.ShowKe()
	fmt.Println("duration", time.Now().Sub(t0))
}

package iteration

func Repeat(character string, count int) string {
	// 変数を宣言するとその型のゼロ値に初期化される
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

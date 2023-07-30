package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("普通のテスト", func(t *testing.T) {

		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("回数が指定されてループが回る場合", func(t *testing.T) {

		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

package color

import (
	"fmt"
	"testing"
)

func Benchmark100char1Tag(b *testing.B) {
	s := "This is a one [red]hundred[/red] [100] char text. It isn't that long really but that is how it is...."

	for n := 0; n < b.N; n++ {
		_ = Do(s)
	}
}

func Benchmark100char3Tags(b *testing.B) {
	s := "This is a one [red]hundred[/red] [100] char [bold]text[/bold]. It isn't that [bggreen]long[/bggreen]."
	fmt.Println(Do(s))
	for n := 0; n < b.N; n++ {
		_ = Do(s)
	}
}

func Benchmark100charNoTags(b *testing.B) {
	s := "This is a one hundred [100] char text. It isn't that long really but that is what we want it to  be."

	for n := 0; n < b.N; n++ {
		_ = Do(s)
	}
}

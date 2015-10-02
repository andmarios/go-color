package color_test

import (
	"fmt"

	"github.com/andmarios/color"
)

func ExampleDo() {
	s := "Here are: a [red]coloured word[/red], [bold]bold word[/bold] and a [bggreen][underline]green background and underlined word[/underline][/bggreen]."
	fmt.Println(color.Do(s))
}

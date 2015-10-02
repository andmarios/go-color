// Package color offers a very simple “color tag to bash color code” conversion for strings.
//
// E.g:
// Hello, this is [red]coloured[/red] and this is [bold]emphasis[/bold]
// Will become:
// Hello, this is \033[31mcoloured\033[39m and this is \033[1memphasis\033[21;22m
//
// color will not check if you are on linux or windows, if you write to a terminal
// that supports colors, etc. It will only convert your string.
//
package color

import "strings"

// Do converts tags founds in s to bash escape sequences and codes.
func Do(s string) string {
	for k, v := range Colours {
		s = strings.Replace(s, k, v, -1)
	}

	return s
}

// A Colours is the map of available tags.
var Colours = map[string]string{
	"[black]":        "\033[30m",
	"[/black]":       "\033[39m",
	"[red]":          "\033[31m",
	"[/red]":         "\033[39m",
	"[green]":        "\033[32m",
	"[/green]":       "\033[39m",
	"[yellow]":       "\033[33m",
	"[/yellow]":      "\033[39m",
	"[blue]":         "\033[34m",
	"[/blue]":        "\033[39m",
	"[magenta]":      "\033[35m",
	"[/magenta]":     "\033[39m",
	"[cyan]":         "\033[36m",
	"[/cyan]":        "\033[39m",
	"[lightgray]":    "\033[37m",
	"[/lightgray]":   "\033[39m",
	"[default]":      "\033[39m",
	"[/default]":     "\033[39m",
	"[bgblack]":      "\033[40m",
	"[/bgblack]":     "\033[49m",
	"[bgred]":        "\033[41m",
	"[/bgred]":       "\033[49m",
	"[bggreen]":      "\033[42m",
	"[/bggreen]":     "\033[49m",
	"[bgyellow]":     "\033[43m",
	"[/bgyellow]":    "\033[49m",
	"[bgblue]":       "\033[44m",
	"[/bgblue]":      "\033[49m",
	"[bgmagenta]":    "\033[45m",
	"[/bgmagenta]":   "\033[49m",
	"[bgcyan]":       "\033[46m",
	"[/bgcyan]":      "\033[49m",
	"[bglightgray]":  "\033[47m",
	"[/bglightgray]": "\033[49m",
	"[bgdefault]":    "\033[49m",
	"[/bgdefault]":   "\033[49m",
	"[bold]":         "\033[1m",
	"[/bold]":        "\033[21;22m",
	"[italic]":       "\033[3m",
	"[/italic]":      "\033[23m",
	"[underline]":    "\033[4m",
	"[/underline]":   "\033[24m",
	"[blink]":        "\033[5m",
	"[/blink]":       "\033[25m",
	"[reverse]":      "\033[7m",
	"[/reverse]":     "\033[27m",
	"[reset]":        "\033[0m",
}

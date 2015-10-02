# color

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg)](http://godoc.org/github.com/andmarios/go-color)

Color is a tiny go package to convert a string with color tags to a string with
bash color escape sequences.

Given this:

    Hello, this is [red]coloured[/red] and this is [bold]emphasis[/bold]

It will produce this:

    Hello, this is \033[31mcoloured\033[39m and this is \033[1memphasis\033[21;22m

Which if your print to a linux terminal will show as expected —_coloured_ in red,
_emphasis_ in bold.

Color will not check if you are on linux or windows, if you write to a terminal
which supports colors, etc. It will only convert your string. The rest is up to you.

Supported tags are:

- For colors: black, red, green, yellow, blue, magenta, cyan, lightgray, default.
- For backgrounds: bgblack, bgred, bggreen, bgyellow, bgblue, bgmagenta, bgcyan, bglightgray, bgdefault.
- For attributes: bold, italic, underline, blink, reverse, reset.

`default` and `bgdefault` are special cases which apply the default colour, thus no need
to close them. `reset` is a special tag that resets everything: default colours, no attributes.

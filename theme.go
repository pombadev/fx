package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mazznoer/colorgrad"
	"github.com/muesli/termenv"
)

type theme struct {
	Cursor    color
	Syntax    color
	Preview   color
	StatusBar color
	Search    color
	Key       color
	String    color
	Null      color
	Boolean   color
	Number    color
}

type color func(s []byte) []byte

func init() {
	themeId, ok := os.LookupEnv("FX_THEME")
	if !ok {
		themeId = "1"
	}
	currentTheme, ok = themes[themeId]
	if !ok {
		currentTheme = themes["1"]
	}
	if termenv.ColorProfile() == termenv.Ascii {
		currentTheme = themes["0"]
	}

	colon = currentTheme.Syntax([]byte{':', ' '})
	colonPreview = currentTheme.Preview([]byte{':'})
	comma = currentTheme.Syntax([]byte{','})
	empty = currentTheme.Preview([]byte{'~'})
	dot3 = currentTheme.Preview([]byte("…"))
	closeCurlyBracket = currentTheme.Syntax([]byte{'}'})
	closeSquareBracket = currentTheme.Syntax([]byte{']'})
}

var (
	currentTheme     theme
	defaultCursor    = toColor(lipgloss.NewStyle().Reverse(true).Render)
	defaultPreview   = toColor(lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render)
	defaultStatusBar = toColor(lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("0")).Render)
	defaultSearch    = toColor(lipgloss.NewStyle().Background(lipgloss.Color("11")).Foreground(lipgloss.Color("16")).Render)
	defaultNull      = fg("243")
)

var (
	colon              []byte
	colonPreview       []byte
	comma              []byte
	empty              []byte
	dot3               []byte
	closeCurlyBracket  []byte
	closeSquareBracket []byte
)

var themes = map[string]theme{
	"0": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   noColor,
		StatusBar: noColor,
		Search:    defaultSearch,
		Key:       noColor,
		String:    noColor,
		Null:      noColor,
		Boolean:   noColor,
		Number:    noColor,
	},
	"1": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       boldFg("4"),
		String:    fg("2"),
		Null:      defaultNull,
		Boolean:   fg("5"),
		Number:    fg("6"),
	},
	"2": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("2"),
		String:    fg("4"),
		Null:      defaultNull,
		Boolean:   fg("5"),
		Number:    fg("6"),
	},
	"3": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("13"),
		String:    fg("11"),
		Null:      defaultNull,
		Boolean:   fg("1"),
		Number:    fg("14"),
	},
	"4": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("#00F5D4"),
		String:    fg("#00BBF9"),
		Null:      defaultNull,
		Boolean:   fg("#F15BB5"),
		Number:    fg("#9B5DE5"),
	},
	"5": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("#faf0ca"),
		String:    fg("#f4d35e"),
		Null:      defaultNull,
		Boolean:   fg("#ee964b"),
		Number:    fg("#ee964b"),
	},
	"6": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("#4D96FF"),
		String:    fg("#6BCB77"),
		Null:      defaultNull,
		Boolean:   fg("#FF6B6B"),
		Number:    fg("#FFD93D"),
	},
	"7": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       boldFg("42"),
		String:    boldFg("213"),
		Null:      defaultNull,
		Boolean:   boldFg("201"),
		Number:    boldFg("201"),
	},
	"8": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       gradient("rgb(125,110,221)", "rgb(90%,45%,97%)", "hsl(229,79%,85%)"),
		String:    fg("195"),
		Null:      defaultNull,
		Boolean:   fg("195"),
		Number:    fg("195"),
	},
	"9": {
		Cursor:    defaultCursor,
		Syntax:    noColor,
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       gradient("rgb(123,216,96)", "rgb(255,255,255)"),
		String:    gradient("rgb(255,255,255)", "rgb(123,216,96)"),
		Null:      defaultNull,
		Boolean:   noColor,
		Number:    noColor,
	},
	"🔵": {
		Cursor: toColor(lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("33")).
			Render),
		Syntax:    boldFg("33"),
		Preview:   defaultPreview,
		StatusBar: defaultStatusBar,
		Search:    defaultSearch,
		Key:       fg("33"),
		String:    noColor,
		Null:      noColor,
		Boolean:   noColor,
		Number:    noColor,
	},
}

func noColor(s []byte) []byte {
	return s
}

func toColor(f func(s ...string) string) color {
	return func(s []byte) []byte {
		return []byte(f(string(s)))
	}
}

func fg(color string) color {
	return toColor(lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render)
}

func boldFg(color string) color {
	return toColor(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(color)).Render)
}

func gradient(colors ...string) color {
	grad, _ := colorgrad.NewGradient().HtmlColors(colors...).Build()
	return toColor(func(s ...string) string {
		runes := []rune(s[0])
		colors := grad.ColorfulColors(uint(len(runes)))
		var out strings.Builder
		for i, r := range runes {
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[i].Hex()))
			out.WriteString(style.Render(string(r)))
		}
		return out.String()
	})
}

func themeTester() {
	title := lipgloss.NewStyle().Bold(true)
	themeNames := make([]string, 0, len(themes))
	for name := range themes {
		themeNames = append(themeNames, name)
	}
	sort.Strings(themeNames)
	for _, name := range themeNames {
		theme := themes[name]
		comma := string(theme.Syntax([]byte{','}))
		colon := string(theme.Syntax([]byte{':'}))

		fmt.Println(title.Render(fmt.Sprintf("Theme %q", name)))
		fmt.Println(string(theme.Syntax([]byte("{"))))

		fmt.Printf("  %v%v %v%v\n",
			string(theme.Key([]byte("\"string\""))),
			colon,
			string(theme.String([]byte("\"Fox jumps over the lazy dog\""))),
			comma)

		fmt.Printf("  %v%v %v%v\n",
			string(theme.Key([]byte("\"number\""))),
			colon,
			string(theme.Number([]byte("1234567890"))),
			comma)

		fmt.Printf("  %v%v %v%v\n",
			string(theme.Key([]byte("\"boolean\""))),
			colon,
			string(theme.Boolean([]byte("true"))),
			comma)
		fmt.Printf("  %v%v %v%v\n",
			string(theme.Key([]byte("\"null\""))),
			colon,
			string(theme.Null([]byte("null"))),
			comma)
		fmt.Println(string(theme.Syntax([]byte("}"))))
		println()
	}
}

func valueStyle(b []byte, selected, chunk bool) color {
	if selected {
		return currentTheme.Cursor
	} else if chunk {
		return currentTheme.String
	} else {
		switch b[0] {
		case '"':
			return currentTheme.String
		case 't', 'f':
			return currentTheme.Boolean
		case 'n':
			return currentTheme.Null
		case '{', '[', '}', ']':
			return currentTheme.Syntax
		default:
			if isDigit(b[0]) || b[0] == '-' {
				return currentTheme.Number
			}
			return noColor
		}
	}
}

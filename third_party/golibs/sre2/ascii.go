package sre2

// This file describes ASCII character ranges as per "Perl character classes"
// and "ASCII character classes" on the RE2 syntax page, found here:
// http://code.google.com/p/re2/wiki/Syntax
//
// Notably, Perl character classes should match the following ASCII classes:
//    '\d' => [:digit:]
//    '\s' => [:whitespace:]*
//    '\w' => [:word:]
//
// NB: [:whitespace:] is [:space:] with the addition of the runes "\f\r".

import "unicode"

var _alnum = []unicode.Range{
	unicode.Range{'0', '9', 1},
	unicode.Range{'A', 'Z', 1},
	unicode.Range{'a', 'z', 1},
}

var _alpha = []unicode.Range{
	unicode.Range{'A', 'Z', 1},
	unicode.Range{'a', 'z', 1},
}

var _ascii = []unicode.Range{
	unicode.Range{0x00, 0x7f, 1},
}

var _blank = []unicode.Range{
	unicode.Range{'\t', '\t', 1},
	unicode.Range{' ', ' ', 1},
}

var _cntrl = []unicode.Range{
	unicode.Range{0x00, 0x1f, 1},
	unicode.Range{0x7f, 0x7f, 1},
}

var _digit = []unicode.Range{
	unicode.Range{'0', '9', 1},
}

var _graph = []unicode.Range{
	unicode.Range{'!', '~', 1},
}

var _lower = []unicode.Range{
	unicode.Range{'a', 'z', 1},
}

var _print = []unicode.Range{
	unicode.Range{' ', '~', 1},
}

var _punct = []unicode.Range{
	unicode.Range{'!', '/', 1},
	unicode.Range{':', '@', 1},
	unicode.Range{'[', '`', 1},
	unicode.Range{'{', '~', 1},
}

var _space = []unicode.Range{
	unicode.Range{'\t', '\r', 1},
	unicode.Range{' ', ' ', 1},
}

var _upper = []unicode.Range{
	unicode.Range{'A', 'Z', 1},
}

var _word = []unicode.Range{
	unicode.Range{'0', '9', 1},
	unicode.Range{'A', 'Z', 1},
	unicode.Range{'a', 'z', 1},
}

// [:whitespace:] matches Perl's definition of '\s'.
var _whitespace = []unicode.Range{
	unicode.Range{'\t', '\n', 1},
	unicode.Range{'\f', '\r', 1},
	unicode.Range{' ', ' ', 1},
}

var _xdigit = []unicode.Range{
	unicode.Range{'0', '9', 1},
	unicode.Range{'A', 'F', 1},
	unicode.Range{'a', 'f', 1},
}

var ASCII = map[string][]unicode.Range{
	"alnum":      _alnum,
	"alpha":      _alpha,
	"ascii":      _ascii,
	"blank":      _blank,
	"cntrl":      _cntrl,
	"digit":      _digit,
	"graph":      _graph,
	"lower":      _lower,
	"print":      _print,
	"punct":      _punct,
	"space":      _space,
	"upper":      _upper,
	"word":       _word,
	"whitespace": _whitespace,
	"xdigit":     _xdigit,
}

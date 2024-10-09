package tslatex

// #cgo CFLAGS: -std=c11 -fPIC
// #include "./parser.c"
// #include "./scanner.c"
import "C"
import "unsafe"

// Get the tree-sitter Language for this grammar.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_latex())
}

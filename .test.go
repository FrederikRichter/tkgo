import (
	"context"
	"fmt"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
)

parser := sitter.NewParser()
parser.SetLanguage(javascript.GetLanguage())

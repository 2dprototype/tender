package parser

import (
	"strings"
	"math/big"

	"github.com/2dprototype/tender/token"
)

// Expr represents an expression node in the AST.
type Expr interface {
	Node
	exprNode()
}

// ArrayLit represents an array literal.
type ArrayLit struct {
	Elements []Expr
	LBrack   Pos
	RBrack   Pos
}

func (e *ArrayLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ArrayLit) Pos() Pos {
	return e.LBrack
}

// End returns the position of first character immediately after the node.
func (e *ArrayLit) End() Pos {
	return e.RBrack + 1
}

func (e *ArrayLit) String() string {
	var elements []string
	for _, m := range e.Elements {
		elements = append(elements, m.String())
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

// BadExpr represents a bad expression.
type BadExpr struct {
	From Pos
	To   Pos
}

func (e *BadExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BadExpr) Pos() Pos {
	return e.From
}

// End returns the position of first character immediately after the node.
func (e *BadExpr) End() Pos {
	return e.To
}

func (e *BadExpr) String() string {
	return "<bad expression>"
}

// BinaryExpr represents a binary operator expression.
type BinaryExpr struct {
	LHS      Expr
	RHS      Expr
	Token    token.Token
	TokenPos Pos
}

func (e *BinaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BinaryExpr) Pos() Pos {
	return e.LHS.Pos()
}

// End returns the position of first character immediately after the node.
func (e *BinaryExpr) End() Pos {
	return e.RHS.End()
}

func (e *BinaryExpr) String() string {
	return "(" + e.LHS.String() + " " + e.Token.String() +
		" " + e.RHS.String() + ")"
}

// BoolLit represents a boolean literal.
type BoolLit struct {
	Value    bool
	ValuePos Pos
	Literal  string
}

func (e *BoolLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BoolLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BoolLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BoolLit) String() string {
	return e.Literal
}

// CallExpr represents a function call expression.
type CallExpr struct {
	Func     Expr
	LParen   Pos
	Args     []Expr
	Ellipsis Pos
	RParen   Pos
	Optional bool
}

func (e *CallExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CallExpr) Pos() Pos {
	return e.Func.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CallExpr) End() Pos {
	return e.RParen + 1
}

func (e *CallExpr) String() string {
	var args []string
	for _, e := range e.Args {
		args = append(args, e.String())
	}
	if len(args) > 0 && e.Ellipsis.IsValid() {
		args[len(args)-1] = args[len(args)-1] + "..."
	}
	opt := ""
	if e.Optional {
		opt = "?"
	}
	return e.Func.String() + opt + "(" + strings.Join(args, ", ") + ")"
}

// CharLit represents a character literal.
type CharLit struct {
	Value    rune
	ValuePos Pos
	Literal  string
}

func (e *CharLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CharLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *CharLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *CharLit) String() string {
	return e.Literal
}

// CondExpr represents a ternary conditional expression.
type CondExpr struct {
	Cond        Expr
	True        Expr
	False       Expr
	QuestionPos Pos
	ColonPos    Pos
}

func (e *CondExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *CondExpr) Pos() Pos {
	return e.Cond.Pos()
}

// End returns the position of first character immediately after the node.
func (e *CondExpr) End() Pos {
	return e.False.End()
}

func (e *CondExpr) String() string {
	return "(" + e.Cond.String() + " ? " + e.True.String() +
		" : " + e.False.String() + ")"
}

// ErrorExpr represents an error expression
type ErrorExpr struct {
	Expr     Expr
	ErrorPos Pos
	LParen   Pos
	RParen   Pos
}

func (e *ErrorExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ErrorExpr) Pos() Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ErrorExpr) End() Pos {
	return e.RParen
}

func (e *ErrorExpr) String() string {
	return "error(" + e.Expr.String() + ")"
}

// FloatLit represents a floating point literal.
type FloatLit struct {
	Value    float64
	ValuePos Pos
	Literal  string
}

func (e *FloatLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FloatLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *FloatLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *FloatLit) String() string {
	return e.Literal
}

// BigFloatLit represents a floating point literal.
type BigFloatLit struct {
	Value    *big.Float
	ValuePos Pos
	Literal  string
}

func (e *BigFloatLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BigFloatLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BigFloatLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BigFloatLit) String() string {
	return e.Literal
}

// ComplexLit represents a floating point literal.
type ComplexLit struct {
	Value    complex128
	ValuePos Pos
	Literal  string
}

func (e *ComplexLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ComplexLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *ComplexLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *ComplexLit) String() string {
	return e.Literal
}

// FuncLit represents a function literal.
type FuncLit struct {
	Type *FuncType
	Body *BlockStmt
}

func (e *FuncLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FuncLit) Pos() Pos {
	return e.Type.Pos()
}

// End returns the position of first character immediately after the node.
func (e *FuncLit) End() Pos {
	return e.Body.End()
}

func (e *FuncLit) String() string {
	return "fn" + e.Type.Params.String() + " " + e.Body.String()
}

// FuncType represents a function type definition.
type FuncType struct {
	FuncPos Pos
	Params  *IdentList
}

func (e *FuncType) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *FuncType) Pos() Pos {
	return e.FuncPos
}

// End returns the position of first character immediately after the node.
func (e *FuncType) End() Pos {
	return e.Params.End()
}

func (e *FuncType) String() string {
	return "fn" + e.Params.String()
}

// Ident represents an identifier.
type Ident struct {
	Name    string
	NamePos Pos
}

func (e *Ident) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *Ident) Pos() Pos {
	return e.NamePos
}

// End returns the position of first character immediately after the node.
func (e *Ident) End() Pos {
	return Pos(int(e.NamePos) + len(e.Name))
}

func (e *Ident) String() string {
	if e != nil {
		return e.Name
	}
	return nullRep
}

// ImmutableExpr represents an immutable expression
type ImmutableExpr struct {
	Expr     Expr
	ErrorPos Pos
	LParen   Pos
	RParen   Pos
}

func (e *ImmutableExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImmutableExpr) Pos() Pos {
	return e.ErrorPos
}

// End returns the position of first character immediately after the node.
func (e *ImmutableExpr) End() Pos {
	return e.RParen
}

func (e *ImmutableExpr) String() string {
	return "immutable(" + e.Expr.String() + ")"
}



type EmbedExpr struct {
	FileSrc    string
	Token      token.Token
	TokenPos   Pos
}

func (e *EmbedExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *EmbedExpr) Pos() Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *EmbedExpr) End() Pos {
	// import("moduleName")
	return Pos(int(e.TokenPos) + 10 + len(e.FileSrc))
}

func (e *EmbedExpr) String() string {
	return `embed("` + e.FileSrc + `")"`
}

// ImportExpr represents an import expression
type ImportExpr struct {
	ModuleName string
	Token      token.Token
	TokenPos   Pos
}

func (e *ImportExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ImportExpr) Pos() Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *ImportExpr) End() Pos {
	// import("moduleName")
	return Pos(int(e.TokenPos) + 10 + len(e.ModuleName))
}

func (e *ImportExpr) String() string {
	return `import("` + e.ModuleName + `")"`
}

// IndexExpr represents an index expression.
type IndexExpr struct {
	Expr   Expr
	LBrack Pos
	Index  Expr
	RBrack Pos
	Optional bool
}

func (e *IndexExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IndexExpr) Pos() Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *IndexExpr) End() Pos {
	return e.RBrack + 1
}

func (e *IndexExpr) String() string {
	var index string
	if e.Index != nil {
		index = e.Index.String()
	}
	opt := ""
	if e.Optional {
		opt = "?."
	}
	return e.Expr.String() + opt + "[" + index + "]"
}

// IntLit represents an integer literal.
type IntLit struct {
	Value    int64
	ValuePos Pos
	Literal  string
}

func (e *IntLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *IntLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *IntLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *IntLit) String() string {
	return e.Literal
}

// BigIntLit represents an integer literal.
type BigIntLit struct {
	Value    *big.Int
	ValuePos Pos
	Literal  string
}

func (e *BigIntLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *BigIntLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *BigIntLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *BigIntLit) String() string {
	return e.Literal
}

// MapElementLit represents a map element.
type MapElementLit struct {
	Key      string
	KeyPos   Pos
	ColonPos Pos
	Value    Expr
}

func (e *MapElementLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapElementLit) Pos() Pos {
	return e.KeyPos
}

// End returns the position of first character immediately after the node.
func (e *MapElementLit) End() Pos {
	return e.Value.End()
}

func (e *MapElementLit) String() string {
	return e.Key + ": " + e.Value.String()
}

// MapLit represents a map literal.
type MapLit struct {
	LBrace   Pos
	Elements []*MapElementLit
	RBrace   Pos
}

func (e *MapLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *MapLit) Pos() Pos {
	return e.LBrace
}

// End returns the position of first character immediately after the node.
func (e *MapLit) End() Pos {
	return e.RBrace + 1
}

func (e *MapLit) String() string {
	var elements []string
	for _, m := range e.Elements {
		elements = append(elements, m.String())
	}
	return "{" + strings.Join(elements, ", ") + "}"
}

// ParenExpr represents a parenthesis wrapped expression.
type ParenExpr struct {
	Expr   Expr
	LParen Pos
	RParen Pos
}

func (e *ParenExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *ParenExpr) Pos() Pos {
	return e.LParen
}

// End returns the position of first character immediately after the node.
func (e *ParenExpr) End() Pos {
	return e.RParen + 1
}

func (e *ParenExpr) String() string {
	return "(" + e.Expr.String() + ")"
}

// TupleLit represents a tuple literal.
type TupleLit struct {
	Elements []Expr
	LParen   Pos
	RParen   Pos
}

func (e *TupleLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *TupleLit) Pos() Pos {
	return e.LParen
}

// End returns the position of first character immediately after the node.
func (e *TupleLit) End() Pos {
	return e.RParen + 1
}

func (e *TupleLit) String() string {
	var elements []string
	for _, m := range e.Elements {
		elements = append(elements, m.String())
	}
	return "(" + strings.Join(elements, ", ") + ")"
}

// SelectorExpr represents a selector expression.
type SelectorExpr struct {
	Expr     Expr
	Sel      Expr
	Optional bool
}

func (e *SelectorExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SelectorExpr) Pos() Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SelectorExpr) End() Pos {
	return e.Sel.End()
}

func (e *SelectorExpr) String() string {
	sep := "."
	if e.Optional {
		sep = "?."
	}
	return e.Expr.String() + sep + e.Sel.String()
}

// SliceExpr represents a slice expression.
type SliceExpr struct {
	Expr   Expr
	LBrack Pos
	Low    Expr
	High   Expr
	RBrack Pos
}

func (e *SliceExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *SliceExpr) Pos() Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *SliceExpr) End() Pos {
	return e.RBrack + 1
}

func (e *SliceExpr) String() string {
	var low, high string
	if e.Low != nil {
		low = e.Low.String()
	}
	if e.High != nil {
		high = e.High.String()
	}
	return e.Expr.String() + "[" + low + ":" + high + "]"
}

// StringLit represents a string literal.
type StringLit struct {
	Value    string
	ValuePos Pos
	Literal  string
}

func (e *StringLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *StringLit) Pos() Pos {
	return e.ValuePos
}

// End returns the position of first character immediately after the node.
func (e *StringLit) End() Pos {
	return Pos(int(e.ValuePos) + len(e.Literal))
}

func (e *StringLit) String() string {
	return e.Literal
}

// TemplateLit represents a template literal with interpolations.
type TemplateLit struct {
    Parts []Expr // alternating string literals and expression nodes
    ValuePos Pos // position of opening backtick
}

func (e *TemplateLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *TemplateLit) Pos() Pos { return e.ValuePos }

// End returns the position of first character immediately after the node.
func (e *TemplateLit) End() Pos {
    if len(e.Parts) > 0 {
        return e.Parts[len(e.Parts)-1].End()
    }
    return e.ValuePos + 1 // length of backtick?
}

func (e *TemplateLit) String() string {
    var parts []string
    for _, p := range e.Parts {
        parts = append(parts, p.String())
    }
    return "`" + strings.Join(parts, "") + "`" 
}

// UnaryExpr represents an unary operator expression.
type UnaryExpr struct {
	Expr     Expr
	Token    token.Token
	TokenPos Pos
}

func (e *UnaryExpr) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *UnaryExpr) Pos() Pos {
	return e.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (e *UnaryExpr) End() Pos {
	return e.Expr.End()
}

func (e *UnaryExpr) String() string {
	return "(" + e.Token.String() + e.Expr.String() + ")"
}

// NullLit represents an null literal.
type NullLit struct {
	TokenPos Pos
}

func (e *NullLit) exprNode() {}

// Pos returns the position of first character belonging to the node.
func (e *NullLit) Pos() Pos {
	return e.TokenPos
}

// End returns the position of first character immediately after the node.
func (e *NullLit) End() Pos {
	return e.TokenPos + 9 // len(null) == 9
}

func (e *NullLit) String() string {
	return "null"
}

// StructField represents a field inside a struct type definition.
type StructField struct {
	Name *Ident
	Type Expr
}

// StructTypeExpr represents a struct type definition expression.
type StructTypeExpr struct {
	StructPos Pos
	LBrace    Pos
	RBrace    Pos
	Fields    []*StructField
}

func (e *StructTypeExpr) exprNode() {}
func (e *StructTypeExpr) Pos() Pos { return e.StructPos }
func (e *StructTypeExpr) End() Pos { return e.RBrace + 1 }
func (e *StructTypeExpr) String() string {
	var fields []string
	for _, f := range e.Fields {
		fields = append(fields, f.Name.Name+" "+f.Type.String())
	}
	return "struct { " + strings.Join(fields, "; ") + " }"
}

// StructKeyValExpr represents a Key: Value element inside struct literals.
type StructKeyValExpr struct {
	Key      *Ident
	Colon    Pos
	Value    Expr
}

func (e *StructKeyValExpr) exprNode() {}
func (e *StructKeyValExpr) Pos() Pos { return e.Key.Pos() }
func (e *StructKeyValExpr) End() Pos { return e.Value.End() }
func (e *StructKeyValExpr) String() string {
	return e.Key.Name + ": " + e.Value.String()
}

// StructLitExpr represents struct literals.
type StructLitExpr struct {
	TypeExpr Expr
	LBrace   Pos
	RBrace   Pos
	Elements []Expr
}

func (e *StructLitExpr) exprNode() {}
func (e *StructLitExpr) Pos() Pos { return e.TypeExpr.Pos() }
func (e *StructLitExpr) End() Pos { return e.RBrace + 1 }
func (e *StructLitExpr) String() string {
	var elements []string
	for _, el := range e.Elements {
		elements = append(elements, el.String())
	}
	return e.TypeExpr.String() + "{" + strings.Join(elements, ", ") + "}"
}


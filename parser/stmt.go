package parser

import (
	"strings"

	"github.com/2dprototype/tender/token"
)

// Stmt represents a statement in the AST.
type Stmt interface {
	Node
	stmtNode()
}

// AssignStmt represents an assignment statement.
type AssignStmt struct {
	LHS      []Expr
	RHS      []Expr
	Token    token.Token
	TokenPos Pos
}

func (s *AssignStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *AssignStmt) Pos() Pos {
	return s.LHS[0].Pos()
}

// End returns the position of first character immediately after the node.
func (s *AssignStmt) End() Pos {
	return s.RHS[len(s.RHS)-1].End()
}

func (s *AssignStmt) String() string {
	var lhs, rhs []string
	for _, e := range s.LHS {
		lhs = append(lhs, e.String())
	}
	for _, e := range s.RHS {
		rhs = append(rhs, e.String())
	}
	return strings.Join(lhs, ", ") + " " + s.Token.String() +
		" " + strings.Join(rhs, ", ")
}

// BadStmt represents a bad statement.
type BadStmt struct {
	From Pos
	To   Pos
}

func (s *BadStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *BadStmt) Pos() Pos {
	return s.From
}

// End returns the position of first character immediately after the node.
func (s *BadStmt) End() Pos {
	return s.To
}

func (s *BadStmt) String() string {
	return "<bad statement>"
}

// BlockStmt represents a block statement.
type BlockStmt struct {
	Stmts  []Stmt
	LBrace Pos
	RBrace Pos
}

func (s *BlockStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *BlockStmt) Pos() Pos {
	return s.LBrace
}

// End returns the position of first character immediately after the node.
func (s *BlockStmt) End() Pos {
	return s.RBrace + 1
}

func (s *BlockStmt) String() string {
	var list []string
	for _, e := range s.Stmts {
		list = append(list, e.String())
	}
	return "{" + strings.Join(list, "; ") + "}"
}

// BranchStmt represents a branch statement.
type BranchStmt struct {
	Token    token.Token
	TokenPos Pos
	Label    *Ident
}

func (s *BranchStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *BranchStmt) Pos() Pos {
	return s.TokenPos
}

// End returns the position of first character immediately after the node.
func (s *BranchStmt) End() Pos {
	if s.Label != nil {
		return s.Label.End()
	}

	return Pos(int(s.TokenPos) + len(s.Token.String()))
}

func (s *BranchStmt) String() string {
	var label string
	if s.Label != nil {
		label = " " + s.Label.Name
	}
	return s.Token.String() + label
}

// EmptyStmt represents an empty statement.
type EmptyStmt struct {
	Semicolon Pos
	Implicit  bool
}

func (s *EmptyStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *EmptyStmt) Pos() Pos {
	return s.Semicolon
}

// End returns the position of first character immediately after the node.
func (s *EmptyStmt) End() Pos {
	if s.Implicit {
		return s.Semicolon
	}
	return s.Semicolon + 1
}

func (s *EmptyStmt) String() string {
	return ";"
}

// ExportStmt represents an export statement.
type ExportStmt struct {
	ExportPos Pos
	Result    Expr
}

func (s *ExportStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ExportStmt) Pos() Pos {
	return s.ExportPos
}

// End returns the position of first character immediately after the node.
func (s *ExportStmt) End() Pos {
	return s.Result.End()
}

func (s *ExportStmt) String() string {
	return "export " + s.Result.String()
}

// ExprStmt represents an expression statement.
type ExprStmt struct {
	Expr Expr
}

func (s *ExprStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ExprStmt) Pos() Pos {
	return s.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (s *ExprStmt) End() Pos {
	return s.Expr.End()
}

func (s *ExprStmt) String() string {
	return s.Expr.String()
}

// ForInStmt represents a for-in statement.
type ForInStmt struct {
	ForPos   Pos
	Key      *Ident
	Value    *Ident
	Iterable Expr
	Body     *BlockStmt
}

func (s *ForInStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ForInStmt) Pos() Pos {
	return s.ForPos
}

// End returns the position of first character immediately after the node.
func (s *ForInStmt) End() Pos {
	return s.Body.End()
}

func (s *ForInStmt) String() string {
	if s.Value != nil {
		return "for " + s.Key.String() + ", " + s.Value.String() +
			" in " + s.Iterable.String() + " " + s.Body.String()
	}
	return "for " + s.Key.String() + " in " + s.Iterable.String() +
		" " + s.Body.String()
}

// ForStmt represents a for statement.
type ForStmt struct {
	ForPos Pos
	Init   Stmt
	Cond   Expr
	Post   Stmt
	Body   *BlockStmt
}

func (s *ForStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ForStmt) Pos() Pos {
	return s.ForPos
}

// End returns the position of first character immediately after the node.
func (s *ForStmt) End() Pos {
	return s.Body.End()
}

func (s *ForStmt) String() string {
	var init, cond, post string
	if s.Init != nil {
		init = s.Init.String()
	}
	if s.Cond != nil {
		cond = s.Cond.String() + " "
	}
	if s.Post != nil {
		post = s.Post.String()
	}

	if init != "" || post != "" {
		return "for " + init + " ; " + cond + " ; " + post + s.Body.String()
	}
	return "for " + cond + s.Body.String()
}



// ImportStmt represents an if statement.
type ImportStmt struct {
	Ident    *Ident
	Expr     *ImportExpr
	Symbols  []*Ident
}

func (s *ImportStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ImportStmt) Pos() Pos {
	return s.Expr.TokenPos
}

// End returns the position of first character immediately after the node.
func (s *ImportStmt) End() Pos {
	return s.Expr.TokenPos
}

func (s *ImportStmt) String() string {
	return "import \"" + s.Expr.ModuleName + "\" as " + s.Ident.Name  
}

// FuncStmt represents an if statement.
type FuncStmt struct {
	Ident    *Ident
	Expr     *FuncLit
}

func (s *FuncStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *FuncStmt) Pos() Pos {
	return s.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (s *FuncStmt) End() Pos {
	return s.Expr.End()
}

func (s *FuncStmt) String() string {
	return "fn " + s.Ident.Name + s.Expr.Type.Params.String() + " " + s.Expr.Body.String()
}




// IfStmt represents an if statement.
type IfStmt struct {
	IfPos Pos
	Init  Stmt
	Cond  Expr
	Body  *BlockStmt
	Else  Stmt // else branch; or nil
}

func (s *IfStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *IfStmt) Pos() Pos {
	return s.IfPos
}

// End returns the position of first character immediately after the node.
func (s *IfStmt) End() Pos {
	if s.Else != nil {
		return s.Else.End()
	}
	return s.Body.End()
}

func (s *IfStmt) String() string {
	var initStmt, elseStmt string
	if s.Init != nil {
		initStmt = s.Init.String() + "; "
	}
	if s.Else != nil {
		elseStmt = " else " + s.Else.String()
	}
	return "if " + initStmt + s.Cond.String() + " " +
		s.Body.String() + elseStmt
}

// IncDecStmt represents increment or decrement statement.
type IncDecStmt struct {
	Expr     Expr
	Token    token.Token
	TokenPos Pos
}

func (s *IncDecStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *IncDecStmt) Pos() Pos {
	return s.Expr.Pos()
}

// End returns the position of first character immediately after the node.
func (s *IncDecStmt) End() Pos {
	return Pos(int(s.TokenPos) + 2)
}

func (s *IncDecStmt) String() string {
	return s.Expr.String() + s.Token.String()
}

// ReturnStmt represents a return statement.
type ReturnStmt struct {
	ReturnPos Pos
	Result    Expr
}

func (s *ReturnStmt) stmtNode() {}

// Pos returns the position of first character belonging to the node.
func (s *ReturnStmt) Pos() Pos {
	return s.ReturnPos
}

// End returns the position of first character immediately after the node.
func (s *ReturnStmt) End() Pos {
	if s.Result != nil {
		return s.Result.End()
	}
	return s.ReturnPos + 6
}

func (s *ReturnStmt) String() string {
	if s.Result != nil {
		return "return " + s.Result.String()
	}
	return "return"
}

// TypeDeclStmt represents a type declaration statement.
type TypeDeclStmt struct {
	TypePos    Pos
	Ident      *Ident
	StructType Expr
}

func (s *TypeDeclStmt) stmtNode() {}
func (s *TypeDeclStmt) Pos() Pos { return s.TypePos }
func (s *TypeDeclStmt) End() Pos { return s.StructType.End() }
func (s *TypeDeclStmt) String() string {
	return "type " + s.Ident.Name + " " + s.StructType.String()
}

// MethodStmt represents a struct method declaration statement.
type MethodStmt struct {
	FuncPos      Pos
	ReceiverName *Ident
	ReceiverType *Ident
	Ident        *Ident
	Expr         *FuncLit
}

func (s *MethodStmt) stmtNode() {}
func (s *MethodStmt) Pos() Pos { return s.FuncPos }
func (s *MethodStmt) End() Pos { return s.Expr.End() }
func (s *MethodStmt) String() string {
	return "fn (" + s.ReceiverName.Name + " " + s.ReceiverType.Name + ") " + s.Ident.Name + s.Expr.Type.Params.String() + " " + s.Expr.Body.String()
}



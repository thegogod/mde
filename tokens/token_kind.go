package tokens

type Kind uint8

const (
	Eof Kind = iota

	Plain // # test

	// singles

	Comma        // ,
	Dot          // .
	Colon        // :
	SemiColon    // ;
	LeftParen    // (
	RightParen   // )
	LeftBrace    // {
	RightBrace   // }
	LeftBracket  // [
	RightBracket // ]
	QuestionMark // ?
	AtSymbol     // @
	HashSymbol   // #
	Underscore   // _
	BackTick     // `

	// doubles

	DoubleColon // ::

	// arithmetic

	Plus  // +
	Minus // -
	Star  // *
	Slash // /

	// logical

	Not   // !
	NotEq // !=
	Eq    // =
	EqEq  // ==
	Gt    // >
	GtEq  // >=
	Lt    // <
	LtEq  // <=
	And   // &&
	Or    // ||

	// literals

	Identifier // test
	LString    // "hello world"
	LByte      // 'h'
	LInt       // 1
	LFloat     // 1.1
	Nil        // nil

	// keywords

	If    // @if
	Else  // @else
	For   // @for
	Use   // @use
	Let   // let
	True  // true
	False // false

	// types

	Type   // type
	String // string
	Byte   // byte
	Int    // int
	Float  // float
	Bool   // bool
	Map    // map
)

var Keywords = map[string]Kind{
	"@use":   Use,
	"@if":    If,
	"@else":  Else,
	"@for":   For,
	"let":    Let,
	"true":   True,
	"false":  False,
	"string": Type,
	"byte":   Type,
	"int":    Type,
	"float":  Type,
	"bool":   Type,
	"map":    Type,
}

var Types = map[string]Kind{
	"string": String,
	"byte":   Byte,
	"int":    Int,
	"float":  Float,
	"bool":   Bool,
	"map":    Map,
}

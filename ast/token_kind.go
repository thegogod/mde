package ast

type TokenKind uint8

const (
	Eof TokenKind = iota

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

	// doubles

	DoubleColon // ::

	// arithmetic

	Plus    // +
	PlusEq  // +=
	Minus   // -
	MinusEq // -=
	Star    // *
	StarEq  // *=
	Slash   // /
	SlashEq // /=

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

	If    // if
	Else  // else
	For   // for
	Let   // let
	Self  // self
	Use   // use
	True  // true
	False // false

	// types

	Type
	String // string
	Byte   // byte
	Int    // int
	Float  // float
	Bool   // bool
	Map    // map
)

var Keywords = map[string]TokenKind{
	"@use":   Use,
	"@if":    If,
	"@else":  Else,
	"@for":   For,
	"let":    Let,
	"self":   Self,
	"true":   True,
	"false":  False,
	"string": Type,
	"byte":   Type,
	"int":    Type,
	"float":  Type,
	"bool":   Type,
	"map":    Type,
}

var Types = map[string]TokenKind{
	"string": String,
	"byte":   Byte,
	"int":    Int,
	"float":  Float,
	"bool":   Bool,
	"map":    Map,
}

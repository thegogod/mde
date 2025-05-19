package tokens

type Kind uint8

const (
	Eof Kind = iota
	PlainText

	// whitespace

	Space
	NewLine

	// elements

	H1 // # test
	H2 // ## test
	H3 // ### test
	H4 // #### test
	H5 // ##### test
	H6 // ###### test

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
	At           // @
	Hash         // #
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

func (self Kind) String() string {
	switch self {
	case Eof:
		return "eof"
	case PlainText:
		return "plain-text"
	case Space:
		return "space"
	case NewLine:
		return "newline"
	case H1:
		return "h1"
	case H2:
		return "h2"
	case H3:
		return "h3"
	case H4:
		return "h4"
	case H5:
		return "h5"
	case H6:
		return "h6"
	case Comma:
		return "comma"
	case Dot:
		return "dot"
	case Colon:
		return "colon"
	case SemiColon:
		return "semi-colon"
	case LeftParen:
		return "left-paren"
	case RightParen:
		return "right-paren"
	case LeftBrace:
		return "left-brace"
	case RightBrace:
		return "right-brace"
	case LeftBracket:
		return "left-bracket"
	case RightBracket:
		return "right-bracket"
	case QuestionMark:
		return "question-mark"
	case At:
		return "at"
	case Hash:
		return "hash"
	case Underscore:
		return "underscore"
	case BackTick:
		return "back-tick"
	case DoubleColon:
		return "double-colon"
	case Plus:
		return "plus"
	case Minus:
		return "minus"
	case Star:
		return "star"
	case Slash:
		return "slash"
	case Not:
		return "not"
	case NotEq:
		return "not-eq"
	case Eq:
		return "eq"
	case EqEq:
		return "eq-eq"
	case Gt:
		return "gt"
	case GtEq:
		return "gt-eq"
	case Lt:
		return "lt"
	case LtEq:
		return "lt-eq"
	case And:
		return "and"
	case Or:
		return "or"
	case Identifier:
		return "identifier"
	case LString:
		return "literal-string"
	case LByte:
		return "literal-byte"
	case LInt:
		return "literal-int"
	case LFloat:
		return "literal-float"
	case Nil:
		return "nil"
	case If:
		return "if"
	case Else:
		return "else"
	case For:
		return "for"
	case Use:
		return "use"
	case Let:
		return "let"
	case True:
		return "true"
	case False:
		return "false"
	case Type:
		return "type"
	case String:
		return "string"
	case Byte:
		return "byte"
	case Int:
		return "int"
	case Float:
		return "float"
	case Bool:
		return "bool"
	case Map:
		return "map"
	default:
		panic("unsupported token type")
	}
}

//line parser/tdoc.y:2
package parser

import __yyfmt__ "fmt"

//line parser/tdoc.y:3
import (
	"fmt"
	"github.com/iwalz/tdoc/elements"
)

var program elements.Element

const debug = false

//line parser/tdoc.y:21
type TdocSymType struct {
	yys     int
	val     string
	pos     int
	line    int
	token   int
	element elements.Element
}

const SCOPEIN = 57346
const SCOPEOUT = 57347
const COMPONENT = 57348
const TEXT = 57349
const ERROR = 57350
const IDENTIFIER = 57351
const ALIAS = 57352

var TdocToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SCOPEIN",
	"SCOPEOUT",
	"COMPONENT",
	"TEXT",
	"ERROR",
	"IDENTIFIER",
	"ALIAS",
}
var TdocStatenames = [...]string{}

const TdocEofCode = 1
const TdocErrCode = 2
const TdocInitialStackSize = 16

//line parser/tdoc.y:99

/* Start of the program */

func (p *TdocParserImpl) AST() elements.Element {
	return program
}

//line yacctab:1
var TdocExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const TdocNprod = 8
const TdocPrivate = 57344

var TdocTokenNames []string
var TdocStates []string

const TdocLast = 13

var TdocAct = [...]int{

	10, 8, 3, 12, 5, 6, 7, 11, 7, 1,
	9, 4, 2,
}
var TdocPact = [...]int{

	-2, -1000, -2, 4, -1000, -8, 4, -2, -10, 2,
	-4, -1000, -1000,
}
var TdocPgo = [...]int{

	0, 12, 2, 11, 9,
}
var TdocR1 = [...]int{

	0, 4, 1, 1, 2, 2, 3, 3,
}
var TdocR2 = [...]int{

	0, 1, 1, 2, 4, 1, 2, 4,
}
var TdocChk = [...]int{

	-1000, -4, -1, -2, -3, 6, -2, 4, 9, -2,
	10, 5, 7,
}
var TdocDef = [...]int{

	0, -2, 1, 2, 5, 0, 3, 0, 6, 0,
	0, 4, 7,
}
var TdocTok1 = [...]int{

	1,
}
var TdocTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10,
}
var TdocTok3 = [...]int{
	0,
}

var TdocErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	TdocDebug        = 0
	TdocErrorVerbose = false
)

type TdocLexer interface {
	Lex(lval *TdocSymType) int
	Error(s string)
}

type TdocParser interface {
	Parse(TdocLexer) int
	Lookahead() int
}

type TdocParserImpl struct {
	lval  TdocSymType
	stack [TdocInitialStackSize]TdocSymType
	char  int
}

func (p *TdocParserImpl) Lookahead() int {
	return p.char
}

func TdocNewParser() TdocParser {
	return &TdocParserImpl{}
}

const TdocFlag = -1000

func TdocTokname(c int) string {
	if c >= 1 && c-1 < len(TdocToknames) {
		if TdocToknames[c-1] != "" {
			return TdocToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func TdocStatname(s int) string {
	if s >= 0 && s < len(TdocStatenames) {
		if TdocStatenames[s] != "" {
			return TdocStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func TdocErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !TdocErrorVerbose {
		return "syntax error"
	}

	for _, e := range TdocErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + TdocTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := TdocPact[state]
	for tok := TOKSTART; tok-1 < len(TdocToknames); tok++ {
		if n := base + tok; n >= 0 && n < TdocLast && TdocChk[TdocAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if TdocDef[state] == -2 {
		i := 0
		for TdocExca[i] != -1 || TdocExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; TdocExca[i] >= 0; i += 2 {
			tok := TdocExca[i]
			if tok < TOKSTART || TdocExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if TdocExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += TdocTokname(tok)
	}
	return res
}

func Tdoclex1(lex TdocLexer, lval *TdocSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = TdocTok1[0]
		goto out
	}
	if char < len(TdocTok1) {
		token = TdocTok1[char]
		goto out
	}
	if char >= TdocPrivate {
		if char < TdocPrivate+len(TdocTok2) {
			token = TdocTok2[char-TdocPrivate]
			goto out
		}
	}
	for i := 0; i < len(TdocTok3); i += 2 {
		token = TdocTok3[i+0]
		if token == char {
			token = TdocTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = TdocTok2[1] /* unknown char */
	}
	if TdocDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", TdocTokname(token), uint(char))
	}
	return char, token
}

func TdocParse(Tdoclex TdocLexer) int {
	return TdocNewParser().Parse(Tdoclex)
}

func (Tdocrcvr *TdocParserImpl) Parse(Tdoclex TdocLexer) int {
	var Tdocn int
	var TdocVAL TdocSymType
	var TdocDollar []TdocSymType
	_ = TdocDollar // silence set and not used
	TdocS := Tdocrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Tdocstate := 0
	Tdocrcvr.char = -1
	Tdoctoken := -1 // Tdocrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Tdocstate = -1
		Tdocrcvr.char = -1
		Tdoctoken = -1
	}()
	Tdocp := -1
	goto Tdocstack

ret0:
	return 0

ret1:
	return 1

Tdocstack:
	/* put a state and value onto the stack */
	if TdocDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", TdocTokname(Tdoctoken), TdocStatname(Tdocstate))
	}

	Tdocp++
	if Tdocp >= len(TdocS) {
		nyys := make([]TdocSymType, len(TdocS)*2)
		copy(nyys, TdocS)
		TdocS = nyys
	}
	TdocS[Tdocp] = TdocVAL
	TdocS[Tdocp].yys = Tdocstate

Tdocnewstate:
	Tdocn = TdocPact[Tdocstate]
	if Tdocn <= TdocFlag {
		goto Tdocdefault /* simple state */
	}
	if Tdocrcvr.char < 0 {
		Tdocrcvr.char, Tdoctoken = Tdoclex1(Tdoclex, &Tdocrcvr.lval)
	}
	Tdocn += Tdoctoken
	if Tdocn < 0 || Tdocn >= TdocLast {
		goto Tdocdefault
	}
	Tdocn = TdocAct[Tdocn]
	if TdocChk[Tdocn] == Tdoctoken { /* valid shift */
		Tdocrcvr.char = -1
		Tdoctoken = -1
		TdocVAL = Tdocrcvr.lval
		Tdocstate = Tdocn
		if Errflag > 0 {
			Errflag--
		}
		goto Tdocstack
	}

Tdocdefault:
	/* default state action */
	Tdocn = TdocDef[Tdocstate]
	if Tdocn == -2 {
		if Tdocrcvr.char < 0 {
			Tdocrcvr.char, Tdoctoken = Tdoclex1(Tdoclex, &Tdocrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if TdocExca[xi+0] == -1 && TdocExca[xi+1] == Tdocstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Tdocn = TdocExca[xi+0]
			if Tdocn < 0 || Tdocn == Tdoctoken {
				break
			}
		}
		Tdocn = TdocExca[xi+1]
		if Tdocn < 0 {
			goto ret0
		}
	}
	if Tdocn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Tdoclex.Error(TdocErrorMessage(Tdocstate, Tdoctoken))
			Nerrs++
			if TdocDebug >= 1 {
				__yyfmt__.Printf("%s", TdocStatname(Tdocstate))
				__yyfmt__.Printf(" saw %s\n", TdocTokname(Tdoctoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Tdocp >= 0 {
				Tdocn = TdocPact[TdocS[Tdocp].yys] + TdocErrCode
				if Tdocn >= 0 && Tdocn < TdocLast {
					Tdocstate = TdocAct[Tdocn] /* simulate a shift of "error" */
					if TdocChk[Tdocstate] == TdocErrCode {
						goto Tdocstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if TdocDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", TdocS[Tdocp].yys)
				}
				Tdocp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if TdocDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", TdocTokname(Tdoctoken))
			}
			if Tdoctoken == TdocEofCode {
				goto ret1
			}
			Tdocrcvr.char = -1
			Tdoctoken = -1
			goto Tdocnewstate /* try again in the same state */
		}
	}

	/* reduction by production Tdocn */
	if TdocDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Tdocn, TdocStatname(Tdocstate))
	}

	Tdocnt := Tdocn
	Tdocpt := Tdocp
	_ = Tdocpt // guard against "declared and not used"

	Tdocp -= TdocR2[Tdocn]
	// Tdocp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Tdocp+1 >= len(TdocS) {
		nyys := make([]TdocSymType, len(TdocS)*2)
		copy(nyys, TdocS)
		TdocS = nyys
	}
	TdocVAL = TdocS[Tdocp+1]

	/* consult goto table to find next state */
	Tdocn = TdocR1[Tdocn]
	Tdocg := TdocPgo[Tdocn]
	Tdocj := Tdocg + TdocS[Tdocp].yys + 1

	if Tdocj >= TdocLast {
		Tdocstate = TdocAct[Tdocg]
	} else {
		Tdocstate = TdocAct[Tdocj]
		if TdocChk[Tdocstate] != -Tdocn {
			Tdocstate = TdocAct[Tdocg]
		}
	}
	// dummy call; replaced with literal code
	switch Tdocnt {

	case 1:
		TdocDollar = TdocS[Tdocpt-1 : Tdocpt+1]
		//line parser/tdoc.y:33
		{
			if debug {
				fmt.Println("program")
			}
			TdocVAL.element = program
		}
	case 2:
		TdocDollar = TdocS[Tdocpt-1 : Tdocpt+1]
		//line parser/tdoc.y:41
		{
			if debug {
				fmt.Println("statement_list single")
			}
			TdocDollar[1].element.Root().Add(TdocDollar[1].element)
		}
	case 3:
		TdocDollar = TdocS[Tdocpt-2 : Tdocpt+1]
		//line parser/tdoc.y:49
		{
			if debug {
				fmt.Println("statement_list multi")
			}
			TdocDollar[2].element.Parent(TdocDollar[1].element.Root())
			TdocDollar[2].element.Root().Add(TdocDollar[2].element)
		}
	case 4:
		TdocDollar = TdocS[Tdocpt-4 : Tdocpt+1]
		//line parser/tdoc.y:58
		{
			if debug {
				fmt.Println("Declaration scope")
			}
			TdocDollar[3].element.Parent(TdocDollar[1].element)
			TdocDollar[3].element.Root().Add(TdocDollar[3].element)
		}
	case 6:
		TdocDollar = TdocS[Tdocpt-2 : Tdocpt+1]
		//line parser/tdoc.y:70
		{
			if debug {
				fmt.Println("Component", TdocDollar[1].val, TdocDollar[2].val)
			}
			TdocVAL.element = elements.NewComponent(nil, nil, TdocDollar[1].val, TdocDollar[2].val, "")
			if program == nil {
				program = elements.NewMatrix(nil)
			}

			if TdocVAL.element.Root() == nil {
				TdocVAL.element.Parent(program)
			}
		}
	case 7:
		TdocDollar = TdocS[Tdocpt-4 : Tdocpt+1]
		//line parser/tdoc.y:84
		{
			if debug {
				fmt.Println("alias")
			}
			TdocVAL.element = elements.NewComponent(nil, nil, TdocDollar[1].val, TdocDollar[2].val, TdocDollar[4].val)
			if program == nil {
				program = elements.NewMatrix(nil)
			}

			if TdocVAL.element.Root() == nil {
				TdocVAL.element.Parent(program)
			}
		}
	}
	goto Tdocstack /* stack new state and value */
}

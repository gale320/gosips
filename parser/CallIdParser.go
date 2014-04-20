package parser

import (
	"gosip/core"
	"gosip/header"
	"strings"
)

/** Parser for CALL ID header.
 */
type CallIDParser struct {
	HeaderParserImpl
}

/** Creates new CallIDParser
 * @param String callID message to parse to set
 */
func NewCallIDParser(callId string) *CallIDParser {
	this := &CallIDParser{}

	this.HeaderParserImpl.super(callId)

	return this
}

func (this *CallIDParser) super(callId string) {
	this.HeaderParserImpl.super(callId)
}

/** Constructor
 * @param Lexer lexer to set
 */
func NewCallIDParserFromLexer(lexer core.Lexer) *CallIDParser {
	this := &CallIDParser{}

	this.HeaderParserImpl.superFromLexer(lexer)

	return this
}

/** parse the String message
 * @return SIPHeader (CallID object)
 * @throws ParseException if the message does not respect the spec.
 */
func (this *CallIDParser) Parse() (sh header.SIPHeader, ParseException error) {
	if core.Debug.Debug {
		this.Dbg_enter("parse")
		defer this.Dbg_leave("parse")
	}
	//try {
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_CALL_ID)
	//print(lexer.GetRest());
	lexer.SPorHT()
	//print(lexer.GetRest());
	lexer.Match(':')
	//print(lexer.GetRest());
	lexer.SPorHT()
	//print(lexer.GetRest());
	rest := strings.TrimSpace(lexer.GetRest())

	if callID, err := header.NewCallID(rest); err != nil {
		return nil, err
	} else {
		return callID, nil
	}
	//}finally {
	//if (debug) dbg_leave("parse");
	//}
}
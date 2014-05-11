package parser

import (
	"gosips/core"
	"gosips/sip/header"
)

/** Parser for CSeq headers.
 */
type CSeqParser struct {
	HeaderParser
}

func NewCSeqParser(cseq string) *CSeqParser {
	this := &CSeqParser{}
	this.HeaderParser.super(cseq)
	return this
}

func (this *CSeqParser) super(cseq string) {
	this.HeaderParser.super(cseq)
}

func NewCSeqParserFromLexer(lexer core.Lexer) *CSeqParser {
	this := &CSeqParser{}
	this.HeaderParser.superFromLexer(lexer)
	return this
}

func (this *CSeqParser) Parse() (sh header.Header, ParseException error) {
	//try {
	//c:= header.NewCSeq();
	lexer := this.GetLexer()
	lexer.Match(TokenTypes_CSEQ)
	lexer.SPorHT()
	lexer.Match(':')
	lexer.SPorHT()

	//println(lexer.GetRest());

	var number int
	var method string
	var err error
	if number, err = lexer.Number(); err != nil {
		return nil, err
	}
	//c.SetSequenceNumber(Integer.parseInt(number));
	lexer.SPorHT()

	if method, err = this.Method(); err != nil {
		return nil, err
	}
	//c.SetMethod(m);
	lexer.SPorHT()
	lexer.Match('\n')

	c := header.NewCSeq(number, method)

	return c, nil
	/*      }
	              catch (NumberFormatException ex) {
	                   Debug.printStackTrace(ex);
			   throw createParseException("Number format exception");
	              } catch (InvalidArgumentException ex) {
	                  Debug.printStackTrace(ex);
	                  throw createParseException(ex.getMessage());
	              }*/
}
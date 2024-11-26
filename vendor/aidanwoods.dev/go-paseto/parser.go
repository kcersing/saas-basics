package paseto

import (
	"time"
)

// Parser is used to verify or decrypt a token, and can be provided with
// a set of rules.
type Parser struct {
	rules []Rule
}

// NewParser returns a parser with NotExpired rule preloaded.
func NewParser() Parser {
	return Parser{[]Rule{NotExpired()}}
}

// NewParserWithoutExpiryCheck returns a parser with no currently set rules.
func NewParserWithoutExpiryCheck() Parser {
	return Parser{nil}
}

// NewParserForValidNow returns a parser that will require parsed tokens to be
// valid "now".
func NewParserForValidNow() Parser {
	return Parser{[]Rule{ValidAt(time.Now())}}
}

// MakeParser allows a parser to be constructed with a specified set of rules.
func MakeParser(rules []Rule) Parser {
	return Parser{rules}
}

// ParseV2Local will parse and decrypt a v2 local paseto and validate against
// any parser rules. Error if parsing, decryption, or any rule fails.
func (p Parser) ParseV2Local(key V2SymmetricKey, tainted string) (*Token, error) {
	message, err := newMessage(V2Local, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v2Decrypt(key)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// ParseV2Public will parse and verify a v2 public paseto and validate against
// any parser rules. Error if parsing, verification, or any rule fails.
func (p Parser) ParseV2Public(key V2AsymmetricPublicKey, tainted string) (*Token, error) {
	message, err := newMessage(V2Public, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v2Verify(key)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// ParseV3Local will parse and decrypt a v3 local paseto and validate against
// any parser rules. Error if parsing, decryption, or any rule fails.
func (p Parser) ParseV3Local(key V3SymmetricKey, tainted string, implicit []byte) (*Token, error) {
	message, err := newMessage(V3Local, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v3Decrypt(key, implicit)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// ParseV3Public will parse and verify a v3 public paseto and validate against
// any parser rules. Error if parsing, verification, or any rule fails.
func (p Parser) ParseV3Public(key V3AsymmetricPublicKey, tainted string, implicit []byte) (*Token, error) {
	message, err := newMessage(V3Public, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v3Verify(key, implicit)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// ParseV4Local will parse and decrypt a v4 local paseto and validate against
// any parser rules. Error if parsing, decryption, or any rule fails.
func (p Parser) ParseV4Local(key V4SymmetricKey, tainted string, implicit []byte) (*Token, error) {
	message, err := newMessage(V4Local, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v4Decrypt(key, implicit)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// ParseV4Public will parse and verify a v4 public paseto and validate against
// any parser rules. Error if parsing, verification, or any rule fails.
func (p Parser) ParseV4Public(key V4AsymmetricPublicKey, tainted string, implicit []byte) (*Token, error) {
	message, err := newMessage(V4Public, tainted)
	if err != nil {
		return nil, err
	}

	token, err := message.v4Verify(key, implicit)
	if err != nil {
		return nil, err
	}

	return p.validate(*token)
}

// UnsafeParseFooter returns the footer of a Paseto message. Beware that this
// footer is not cryptographically verified at this stage, nor are any claims
// validated.
func (p Parser) UnsafeParseFooter(protocol Protocol, tainted string) ([]byte, error) {
	message, err := newMessage(protocol, tainted)
	if err != nil {
		return nil, err
	}

	return message.unsafeFooter(), nil
}

// SetRules will overwrite any currently set rules with those specified.
func (p *Parser) SetRules(rules []Rule) {
	p.rules = rules
}

// AddRule will add the given rule(s) to any already specified.
func (p *Parser) AddRule(rule ...Rule) {
	p.rules = append(p.rules, rule...)
}

func (p Parser) validate(token Token) (*Token, error) {
	for _, rule := range p.rules {
		if err := rule(token); err != nil {
			return nil, &RuleError{err}
		}
	}

	return &token, nil
}

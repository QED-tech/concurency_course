package compute

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(query string) ([]string, error) {
	st := newParseStateMachine()

	tokens, err := st.parse(query)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

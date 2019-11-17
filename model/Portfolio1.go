package model

// Describes the characteristics of a portfolio.
type Portfolio1 struct {

	// Additional information related to the portfolio.
	PortfolioInformation []*Max350Text `xml:"PrtflInf,omitempty"`
}

func (p *Portfolio1) AddPortfolioInformation(value string) {
	p.PortfolioInformation = append(p.PortfolioInformation, (*Max350Text)(&value))
}

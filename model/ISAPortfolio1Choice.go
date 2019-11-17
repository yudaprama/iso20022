package model

// Choice to provide additional portfolio information or individual savings account information (UK government scheme provided by UK based financial institutions only).
type ISAPortfolio1Choice struct {

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Individual Savings Account (ISA) is provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue4 `xml:"ISA"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`
}

func (i *ISAPortfolio1Choice) AddISA() *ISAYearsOfIssue4 {
	i.ISA = new(ISAYearsOfIssue4)
	return i.ISA
}

func (i *ISAPortfolio1Choice) AddPortfolio() *Portfolio1 {
	i.Portfolio = new(Portfolio1)
	return i.Portfolio
}

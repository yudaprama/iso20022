package model

// Describes the type of product and the assets to be transferred.
type PEPISATransfer8 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue1 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument11 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer8) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetTransferIdentification(value string) {
	p.TransferIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetCancellationReference(value string) {
	p.CancellationReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer8) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer8) AddISA() *ISAYearsOfIssue1 {
	p.ISA = new(ISAYearsOfIssue1)
	return p.ISA
}

func (p *PEPISATransfer8) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer8) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer8) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument11 {
	newValue := new(FinancialInstrument11)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

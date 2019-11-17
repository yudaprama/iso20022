package model

// Describes the type of product and the assets to be transferred.
type PEPISATransfer4 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the old plan manager to the transfer of account.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification received by the old plan manager and assigned by the new plan manager to the transfer of account.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCashIndicator *YesNoIndicator `xml:"RsdlCshInd"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	ISA *ISAYearsOfIssue3 `xml:"ISA"`

	// UK government schemes to encourage individuals to invest in securities based unit and investment trusts, offering certain tax benefits. These are not investment in their own right but are tax exempt wrappers in which individuals can hold equities, bonds and funds to shelter them from income and capital gains tax.
	//
	// The Personal Equity Plan (PEP) and the Individual Savings Account (ISA) are provided only by UK based financial institutions.
	PEP *PreviousYearChoice `xml:"PEP"`

	// Wrapper for a specific product or a specific sub-product owned by a set of beneficial owners.
	Portfolio *Portfolio1 `xml:"Prtfl"`

	// Specifies the underlying assets for the PEP, ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument11 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (p *PEPISATransfer4) SetMasterReference(value string) {
	p.MasterReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetTransferConfirmationIdentification(value string) {
	p.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetTransferInstructionReference(value string) {
	p.TransferInstructionReference = (*Max35Text)(&value)
}

func (p *PEPISATransfer4) SetActualTransferDate(value string) {
	p.ActualTransferDate = (*ISODate)(&value)
}

func (p *PEPISATransfer4) SetResidualCashIndicator(value string) {
	p.ResidualCashIndicator = (*YesNoIndicator)(&value)
}

func (p *PEPISATransfer4) AddISA() *ISAYearsOfIssue3 {
	p.ISA = new(ISAYearsOfIssue3)
	return p.ISA
}

func (p *PEPISATransfer4) AddPEP() *PreviousYearChoice {
	p.PEP = new(PreviousYearChoice)
	return p.PEP
}

func (p *PEPISATransfer4) AddPortfolio() *Portfolio1 {
	p.Portfolio = new(Portfolio1)
	return p.Portfolio
}

func (p *PEPISATransfer4) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument11 {
	newValue := new(FinancialInstrument11)
	p.FinancialInstrumentAssetForTransfer = append(p.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

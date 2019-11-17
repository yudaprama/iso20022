package model

// Describes the type of product and the assets to be transferred.
type ISATransfer5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio3Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument27 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer5) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer5) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer5) AddPortfolio() *ISAPortfolio3Choice {
	i.Portfolio = new(ISAPortfolio3Choice)
	return i.Portfolio
}

func (i *ISATransfer5) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer5) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument27 {
	newValue := new(FinancialInstrument27)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

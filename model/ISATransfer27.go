package model

// Describes the type of product and the assets to be transferred.
type ISATransfer27 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio3Choice `xml:"Prtfl,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument50 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer27) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer27) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer27) AddPortfolio() *ISAPortfolio3Choice {
	i.Portfolio = new(ISAPortfolio3Choice)
	return i.Portfolio
}

func (i *ISATransfer27) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer27) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument50 {
	newValue := new(FinancialInstrument50)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

package model

// Describes the type of product and the assets to be transferred.
type ISATransfer14 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument37 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer14) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer14) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer14) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer14) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer14) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument37 {
	newValue := new(FinancialInstrument37)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

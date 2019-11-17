package model

// Describes the type of product and the assets to be transferred.
type ISATransfer21 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl,omitempty"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument40 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer21) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer21) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer21) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer21) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer21) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer21) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument40 {
	newValue := new(FinancialInstrument40)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

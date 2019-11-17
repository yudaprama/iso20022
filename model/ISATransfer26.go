package model

// Describes the type of product and the assets to be transferred.
type ISATransfer26 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferInstructionReference *Max35Text `xml:"TrfInstrRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date when the transfer instruction was executed.
	ActualTransferDate *ISODate `xml:"ActlTrfDt,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio2Choice `xml:"Prtfl,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument48 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer26) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer26) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer26) SetTransferInstructionReference(value string) {
	i.TransferInstructionReference = (*Max35Text)(&value)
}

func (i *ISATransfer26) AddClientReference() *AdditionalReference7 {
	i.ClientReference = new(AdditionalReference7)
	return i.ClientReference
}

func (i *ISATransfer26) AddCounterpartyReference() *AdditionalReference7 {
	i.CounterpartyReference = new(AdditionalReference7)
	return i.CounterpartyReference
}

func (i *ISATransfer26) SetBusinessFlowType(value string) {
	i.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (i *ISATransfer26) SetActualTransferDate(value string) {
	i.ActualTransferDate = (*ISODate)(&value)
}

func (i *ISATransfer26) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer26) AddPortfolio() *ISAPortfolio2Choice {
	i.Portfolio = new(ISAPortfolio2Choice)
	return i.Portfolio
}

func (i *ISATransfer26) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer26) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument48 {
	newValue := new(FinancialInstrument48)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

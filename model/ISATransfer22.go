package model

// Describes the type of product and the assets to be transferred.
type ISATransfer22 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned to the transfer of asset, typically assigned by the transferee.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Identification of the confirmation assigned by the transferor to the transfer.
	TransferConfirmationIdentification *Max35Text `xml:"TrfConfId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Requested date at which the assets should be transferred.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies whether all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *AllOtherCash1Code `xml:"AllOthrCsh,omitempty"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument46 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer22) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer22) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer22) SetTransferConfirmationIdentification(value string) {
	i.TransferConfirmationIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer22) AddClientReference() *AdditionalReference7 {
	i.ClientReference = new(AdditionalReference7)
	return i.ClientReference
}

func (i *ISATransfer22) AddCounterpartyReference() *AdditionalReference7 {
	i.CounterpartyReference = new(AdditionalReference7)
	return i.CounterpartyReference
}

func (i *ISATransfer22) SetBusinessFlowType(value string) {
	i.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (i *ISATransfer22) AddRequestedTransferDate() *DateFormat1Choice {
	i.RequestedTransferDate = new(DateFormat1Choice)
	return i.RequestedTransferDate
}

func (i *ISATransfer22) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer22) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer22) SetAllOtherCash(value string) {
	i.AllOtherCash = (*AllOtherCash1Code)(&value)
}

func (i *ISATransfer22) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument46 {
	newValue := new(FinancialInstrument46)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}

package model

// Parameters applied to the settlement of a security transfer.
type ReceiveInformation16 struct {

	// Party that receives (transferee) securities from the delivering agent (transferor).
	Transferee *PartyIdentification70Choice `xml:"Trfee,omitempty"`

	// Account into which the securities are to be received.
	TransfereeRegisteredAccount *Account19 `xml:"TrfeeRegdAcct,omitempty"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *ReceivingPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge29 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission23 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax28 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms26 `xml:"FXDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`
}

func (r *ReceiveInformation16) AddTransferee() *PartyIdentification70Choice {
	r.Transferee = new(PartyIdentification70Choice)
	return r.Transferee
}

func (r *ReceiveInformation16) AddTransfereeRegisteredAccount() *Account19 {
	r.TransfereeRegisteredAccount = new(Account19)
	return r.TransfereeRegisteredAccount
}

func (r *ReceiveInformation16) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	r.IntermediaryInformation = append(r.IntermediaryInformation, newValue)
	return newValue
}

func (r *ReceiveInformation16) SetRequestedSettlementDate(value string) {
	r.RequestedSettlementDate = (*ISODate)(&value)
}

func (r *ReceiveInformation16) SetSettlementAmount(value, currency string) {
	r.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation16) SetStampDuty(value string) {
	r.StampDuty = (*StampDutyType2Code)(&value)
}

func (r *ReceiveInformation16) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *ReceiveInformation16) AddSettlementPartiesDetails() *ReceivingPartiesAndAccount13 {
	r.SettlementPartiesDetails = new(ReceivingPartiesAndAccount13)
	return r.SettlementPartiesDetails
}

func (r *ReceiveInformation16) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *ReceiveInformation16) SetPhysicalTransfer(value string) {
	r.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (r *ReceiveInformation16) AddPhysicalTransferDetails() *DeliveryParameters4 {
	r.PhysicalTransferDetails = new(DeliveryParameters4)
	return r.PhysicalTransferDetails
}

func (r *ReceiveInformation16) AddClientReference() *AdditionalReference7 {
	r.ClientReference = new(AdditionalReference7)
	return r.ClientReference
}

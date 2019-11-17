package model

// Parameters applied to the settlement of a security transfer.
type DeliverInformation17 struct {

	// Party that delivers (transferor) securities to the receiving agent (transferee).
	Transferor *PartyIdentification70Choice `xml:"Trfr,omitempty"`

	// Account from which the securities are to be delivered.
	TransferorRegisteredAccount *Account19 `xml:"TrfrRegdAcct,omitempty"`

	// Identification of a related party or intermediary.
	IntermediaryInformation []*Intermediary34 `xml:"IntrmyInf,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge29 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission23 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax28 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms26 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount13 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation17) AddTransferor() *PartyIdentification70Choice {
	d.Transferor = new(PartyIdentification70Choice)
	return d.Transferor
}

func (d *DeliverInformation17) AddTransferorRegisteredAccount() *Account19 {
	d.TransferorRegisteredAccount = new(Account19)
	return d.TransferorRegisteredAccount
}

func (d *DeliverInformation17) AddIntermediaryInformation() *Intermediary34 {
	newValue := new(Intermediary34)
	d.IntermediaryInformation = append(d.IntermediaryInformation, newValue)
	return newValue
}

func (d *DeliverInformation17) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation17) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation17) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation17) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation17) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation17) AddChargeDetails() *Charge29 {
	newValue := new(Charge29)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddCommissionDetails() *Commission23 {
	newValue := new(Commission23)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddTaxDetails() *Tax28 {
	newValue := new(Tax28)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddForeignExchangeDetails() *ForeignExchangeTerms26 {
	newValue := new(ForeignExchangeTerms26)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation17) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount13 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount13)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation17) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation17) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation17) AddClientReference() *AdditionalReference7 {
	d.ClientReference = new(AdditionalReference7)
	return d.ClientReference
}

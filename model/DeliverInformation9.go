package model

// Parameters applied to the settlement of a security transfer.
type DeliverInformation9 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount8 `xml:"SttlmPtiesDtls,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission12 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax15 `xml:"TaxDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`
}

func (d *DeliverInformation9) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation9) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation9) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation9) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation9) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount8 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount8)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation9) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) AddCommissionDetails() *Commission12 {
	newValue := new(Commission12)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) AddTaxDetails() *Tax15 {
	newValue := new(Tax15)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation9) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation9) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

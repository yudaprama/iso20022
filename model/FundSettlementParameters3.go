package model

// Parameters applied to the settlement of a security transfer.
type FundSettlementParameters3 struct {

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Place where the settlement of transaction will take place. In the context of the investment funds,  the place of settlement is the transfer agent, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SettlementPlace *PartyIdentification2Choice `xml:"SttlmPlc"`

	// Place requested as the place of safekeeping.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc,omitempty"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystemIdentification *Max35Text `xml:"SctiesSttlmSysId,omitempty"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	ReceivingSideDetails *ReceivingPartiesAndAccount3 `xml:"RcvgSdDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	DeliveringSideDetails *DeliveringPartiesAndAccount3 `xml:"DlvrgSdDtls"`
}

func (f *FundSettlementParameters3) SetSettlementDate(value string) {
	f.SettlementDate = (*ISODate)(&value)
}

func (f *FundSettlementParameters3) AddSettlementPlace() *PartyIdentification2Choice {
	f.SettlementPlace = new(PartyIdentification2Choice)
	return f.SettlementPlace
}

func (f *FundSettlementParameters3) AddSafekeepingPlace() *PartyIdentification2Choice {
	f.SafekeepingPlace = new(PartyIdentification2Choice)
	return f.SafekeepingPlace
}

func (f *FundSettlementParameters3) SetSecuritiesSettlementSystemIdentification(value string) {
	f.SecuritiesSettlementSystemIdentification = (*Max35Text)(&value)
}

func (f *FundSettlementParameters3) AddReceivingSideDetails() *ReceivingPartiesAndAccount3 {
	f.ReceivingSideDetails = new(ReceivingPartiesAndAccount3)
	return f.ReceivingSideDetails
}

func (f *FundSettlementParameters3) AddDeliveringSideDetails() *DeliveringPartiesAndAccount3 {
	f.DeliveringSideDetails = new(DeliveringPartiesAndAccount3)
	return f.DeliveringSideDetails
}

package model

// Parameters applied to the settlement of a security.
type FundSettlementParameters4 struct {

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Place where the settlement of transaction will take place. In the context of the investment funds,  the place of settlement is the transfer agent, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SettlementPlace *PartyIdentification2Choice `xml:"SttlmPlc"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository or an International Central Securities Depository.
	SafekeepingPlace *PartyIdentification2Choice `xml:"SfkpgPlc,omitempty"`

	// Identification of a specific system or set of rules and/or processes to be applied at the settlement place.
	SecuritiesSettlementSystemIdentification *Max35Text `xml:"SctiesSttlmSysId,omitempty"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	ReceivingSideDetails *ReceivingPartiesAndAccount3 `xml:"RcvgSdDtls"`

	// Chain of parties involved in the settlement of a transaction resulting in the movement of a security from one account to another.
	DeliveringSideDetails *DeliveringPartiesAndAccount3 `xml:"DlvrgSdDtls,omitempty"`
}

func (f *FundSettlementParameters4) SetSettlementDate(value string) {
	f.SettlementDate = (*ISODate)(&value)
}

func (f *FundSettlementParameters4) AddSettlementPlace() *PartyIdentification2Choice {
	f.SettlementPlace = new(PartyIdentification2Choice)
	return f.SettlementPlace
}

func (f *FundSettlementParameters4) AddSafekeepingPlace() *PartyIdentification2Choice {
	f.SafekeepingPlace = new(PartyIdentification2Choice)
	return f.SafekeepingPlace
}

func (f *FundSettlementParameters4) SetSecuritiesSettlementSystemIdentification(value string) {
	f.SecuritiesSettlementSystemIdentification = (*Max35Text)(&value)
}

func (f *FundSettlementParameters4) AddReceivingSideDetails() *ReceivingPartiesAndAccount3 {
	f.ReceivingSideDetails = new(ReceivingPartiesAndAccount3)
	return f.ReceivingSideDetails
}

func (f *FundSettlementParameters4) AddDeliveringSideDetails() *DeliveringPartiesAndAccount3 {
	f.DeliveringSideDetails = new(DeliveringPartiesAndAccount3)
	return f.DeliveringSideDetails
}

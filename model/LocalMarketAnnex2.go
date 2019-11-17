package model

// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
type LocalMarketAnnex2 struct {

	// Country in which the processing characteristic applies.
	Country []*CountryCode `xml:"Ctry"`

	// Organisation established primarily to provide financial services.
	LocalOrderDesk *ContactAttributes1 `xml:"LclOrdrDsk"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	SubscriptionProcessingCharacteristics *ProcessingCharacteristics2 `xml:"SbcptPrcgChrtcs"`

	// Processing characteristics linked to the instrument, ie, not to  the market.
	RedemptionProcessingCharacteristics *ProcessingCharacteristics3 `xml:"RedPrcgChrtcs"`

	// Account to or from which a cash entry is made.
	SettlementDetails []*CashAccount22 `xml:"SttlmDtls"`
}

func (l *LocalMarketAnnex2) AddCountry(value string) {
	l.Country = append(l.Country, (*CountryCode)(&value))
}

func (l *LocalMarketAnnex2) AddLocalOrderDesk() *ContactAttributes1 {
	l.LocalOrderDesk = new(ContactAttributes1)
	return l.LocalOrderDesk
}

func (l *LocalMarketAnnex2) AddSubscriptionProcessingCharacteristics() *ProcessingCharacteristics2 {
	l.SubscriptionProcessingCharacteristics = new(ProcessingCharacteristics2)
	return l.SubscriptionProcessingCharacteristics
}

func (l *LocalMarketAnnex2) AddRedemptionProcessingCharacteristics() *ProcessingCharacteristics3 {
	l.RedemptionProcessingCharacteristics = new(ProcessingCharacteristics3)
	return l.RedemptionProcessingCharacteristics
}

func (l *LocalMarketAnnex2) AddSettlementDetails() *CashAccount22 {
	newValue := new(CashAccount22)
	l.SettlementDetails = append(l.SettlementDetails, newValue)
	return newValue
}

package model

// Net position of a segregated holding of a single security within the overall position held in the securities account, eg, sub-balance per status.
type HoldingBalance3 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountChoice `xml:"Bal,omitempty"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormatChoice `xml:"SfkpgPlc,omitempty"`

	// Date of the entitlement.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (h *HoldingBalance3) AddBalance() *UnitOrFaceAmountChoice {
	h.Balance = new(UnitOrFaceAmountChoice)
	return h.Balance
}

func (h *HoldingBalance3) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance3) AddSafekeepingPlace() *SafekeepingPlaceFormatChoice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormatChoice)
	return h.SafekeepingPlace
}

func (h *HoldingBalance3) SetDate(value string) {
	h.Date = (*ISODate)(&value)
}

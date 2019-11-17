package model

// Net position of a segregated holding of a single security within the overall position held in the securities account,for example, sub-balance per status.
type HoldingBalance8 struct {

	// Total quantity of financial instrument for the referenced holding.
	Balance *UnitOrFaceAmountOrCode1Choice `xml:"Bal"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	BalanceType *SecuritiesEntryType2Code `xml:"BalTp,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`
}

func (h *HoldingBalance8) AddBalance() *UnitOrFaceAmountOrCode1Choice {
	h.Balance = new(UnitOrFaceAmountOrCode1Choice)
	return h.Balance
}

func (h *HoldingBalance8) SetBalanceType(value string) {
	h.BalanceType = (*SecuritiesEntryType2Code)(&value)
}

func (h *HoldingBalance8) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	h.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return h.SafekeepingPlace
}

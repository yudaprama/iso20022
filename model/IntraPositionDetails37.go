package model

// Details of the intra-position movement.
type IntraPositionDetails37 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType8Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails12 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails37) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails37) AddBalanceFrom() *SecuritiesBalanceType8Choice {
	i.BalanceFrom = new(SecuritiesBalanceType8Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails37) AddIntraPositionMovement() *IntraPositionMovementDetails12 {
	newValue := new(IntraPositionMovementDetails12)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

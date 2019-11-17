package model

// Details of the intra-position movement.
type IntraPositionDetails32 struct {

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Balance from which the securities were moved.
	BalanceFrom *SecuritiesBalanceType6Choice `xml:"BalFr"`

	// Intra-position movement(s) having been performed.
	IntraPositionMovement []*IntraPositionMovementDetails11 `xml:"IntraPosMvmnt"`
}

func (i *IntraPositionDetails32) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	i.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return i.SafekeepingPlace
}

func (i *IntraPositionDetails32) AddBalanceFrom() *SecuritiesBalanceType6Choice {
	i.BalanceFrom = new(SecuritiesBalanceType6Choice)
	return i.BalanceFrom
}

func (i *IntraPositionDetails32) AddIntraPositionMovement() *IntraPositionMovementDetails11 {
	newValue := new(IntraPositionMovementDetails11)
	i.IntraPositionMovement = append(i.IntraPositionMovement, newValue)
	return newValue
}

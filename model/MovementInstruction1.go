package model

// Details of the movement instructions.
type MovementInstruction1 struct {

	// Provides general information about the movement.
	MovementGeneralInformation *CorporateActionMovement1 `xml:"MvmntGnlInf"`

	// Provides information about the underlying securities movement.
	UnderlyingSecuritiesMovementDetails []*UnderlyingSecurityMovement1 `xml:"UndrlygSctiesMvmntDtls,omitempty"`

	// Provides information about the underlying cash movement.
	UnderlyingCashMovementDetails []*CashMovement2 `xml:"UndrlygCshMvmntDtls,omitempty"`

	// Provides information about the proceeds, ie, outturned resources.
	ProceedsMovementDetails []*ProceedsMovement1 `xml:"PrcdsMvmntDtls,omitempty"`
}

func (m *MovementInstruction1) AddMovementGeneralInformation() *CorporateActionMovement1 {
	m.MovementGeneralInformation = new(CorporateActionMovement1)
	return m.MovementGeneralInformation
}

func (m *MovementInstruction1) AddUnderlyingSecuritiesMovementDetails() *UnderlyingSecurityMovement1 {
	newValue := new(UnderlyingSecurityMovement1)
	m.UnderlyingSecuritiesMovementDetails = append(m.UnderlyingSecuritiesMovementDetails, newValue)
	return newValue
}

func (m *MovementInstruction1) AddUnderlyingCashMovementDetails() *CashMovement2 {
	newValue := new(CashMovement2)
	m.UnderlyingCashMovementDetails = append(m.UnderlyingCashMovementDetails, newValue)
	return newValue
}

func (m *MovementInstruction1) AddProceedsMovementDetails() *ProceedsMovement1 {
	newValue := new(ProceedsMovement1)
	m.ProceedsMovementDetails = append(m.ProceedsMovementDetails, newValue)
	return newValue
}

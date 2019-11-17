package model

// Details of the proceeds movements.
type ProceedsMovement1 struct {

	// Provides information about the movement of the cash proceeds.
	CashProceedsMovementDetails []*CashProceeds1 `xml:"CshPrcdsMvmntDtls,omitempty"`

	// Provides information about the movement of the securities proceeds.
	SecuritiesProceedsMovementDetails []*SecuritiesProceeds1 `xml:"SctiesPrcdsMvmntDtls,omitempty"`

	// Provides information about the tax voucher.
	TaxDetails *TaxVoucher1 `xml:"TaxDtls,omitempty"`
}

func (p *ProceedsMovement1) AddCashProceedsMovementDetails() *CashProceeds1 {
	newValue := new(CashProceeds1)
	p.CashProceedsMovementDetails = append(p.CashProceedsMovementDetails, newValue)
	return newValue
}

func (p *ProceedsMovement1) AddSecuritiesProceedsMovementDetails() *SecuritiesProceeds1 {
	newValue := new(SecuritiesProceeds1)
	p.SecuritiesProceedsMovementDetails = append(p.SecuritiesProceedsMovementDetails, newValue)
	return newValue
}

func (p *ProceedsMovement1) AddTaxDetails() *TaxVoucher1 {
	p.TaxDetails = new(TaxVoucher1)
	return p.TaxDetails
}

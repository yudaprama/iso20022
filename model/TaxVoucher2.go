package model

// Specifies tax vouchers in the framework of a corporate action event.
type TaxVoucher2 struct {

	// Unique reference for the tax voucher required by the relevant tax authorities to ensure that any claim relating to this particular tax voucher cannot be duplicated.
	Identification *Max35Text `xml:"Id"`

	// Date on which a dividend reinvestment purchase was completed. If there is only one bargain involved, the time it was struck needs to be included.
	BargainDate *DateAndDateTimeChoice `xml:"BrgnDt,omitempty"`

	// Settlement date of the dividend reinvestment purchase transaction.
	BargainSettlementDate *DateAndDateTimeChoice `xml:"BrgnSttlmDt,omitempty"`
}

func (t *TaxVoucher2) SetIdentification(value string) {
	t.Identification = (*Max35Text)(&value)
}

func (t *TaxVoucher2) AddBargainDate() *DateAndDateTimeChoice {
	t.BargainDate = new(DateAndDateTimeChoice)
	return t.BargainDate
}

func (t *TaxVoucher2) AddBargainSettlementDate() *DateAndDateTimeChoice {
	t.BargainSettlementDate = new(DateAndDateTimeChoice)
	return t.BargainSettlementDate
}

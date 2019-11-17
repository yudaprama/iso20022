package model

// Specifies the cash account elements of a parent cash account.
type ParentCashAccount2 struct {

	// Defines the parent account level within a hierarchy.
	Level *AccountLevel1Code `xml:"Lvl,omitempty"`

	// Unique and unambiguous identification for the parent account between the parent account owner and the parent account servicer.
	Identification *CashAccount24 `xml:"Id"`

	// Financial institution in which the parent account resides.
	Servicer *BranchAndFinancialInstitutionIdentification5 `xml:"Svcr,omitempty"`
}

func (p *ParentCashAccount2) SetLevel(value string) {
	p.Level = (*AccountLevel1Code)(&value)
}

func (p *ParentCashAccount2) AddIdentification() *CashAccount24 {
	p.Identification = new(CashAccount24)
	return p.Identification
}

func (p *ParentCashAccount2) AddServicer() *BranchAndFinancialInstitutionIdentification5 {
	p.Servicer = new(BranchAndFinancialInstitutionIdentification5)
	return p.Servicer
}

package model

// Set of elements to fully identify a proprietary bank transaction code.
type ProprietaryBankTransactionCodeStructure1 struct {

	// Proprietary bank transaction code to identify the underlying transaction.
	Code *Max35Text `xml:"Cd"`

	// Identification of the issuer of the proprietary bank transaction code.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (p *ProprietaryBankTransactionCodeStructure1) SetCode(value string) {
	p.Code = (*Max35Text)(&value)
}

func (p *ProprietaryBankTransactionCodeStructure1) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

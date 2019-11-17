package model

// Information specifying the Mandate.
type OperationMandate1 struct {

	// Unique and unambiguous identification of the mandate.
	Identification *Max35Text `xml:"Id"`

	// Number of required and necessary signatures by the mandate.
	RequiredSignatureNumber *Max15PlusSignedNumericText `xml:"ReqrdSgntrNb"`

	// Indicator whether a certain order of signatures has to be respected or not.
	SignatureOrderIndicator *YesNoIndicator `xml:"SgntrOrdrInd"`

	// Holder of the mandate.
	MandateHolder []*PartyAndCertificate1 `xml:"MndtHldr,omitempty"`

	// Bank operation allowed by a mandate.
	BankOperation []*BankTransactionCodeStructure4 `xml:"BkOpr"`

	// Is the date when the mandate becomes valid.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Is the date when the mandate stops to be valid.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (o *OperationMandate1) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OperationMandate1) SetRequiredSignatureNumber(value string) {
	o.RequiredSignatureNumber = (*Max15PlusSignedNumericText)(&value)
}

func (o *OperationMandate1) SetSignatureOrderIndicator(value string) {
	o.SignatureOrderIndicator = (*YesNoIndicator)(&value)
}

func (o *OperationMandate1) AddMandateHolder() *PartyAndCertificate1 {
	newValue := new(PartyAndCertificate1)
	o.MandateHolder = append(o.MandateHolder, newValue)
	return newValue
}

func (o *OperationMandate1) AddBankOperation() *BankTransactionCodeStructure4 {
	newValue := new(BankTransactionCodeStructure4)
	o.BankOperation = append(o.BankOperation, newValue)
	return newValue
}

func (o *OperationMandate1) SetStartDate(value string) {
	o.StartDate = (*ISODate)(&value)
}

func (o *OperationMandate1) SetEndDate(value string) {
	o.EndDate = (*ISODate)(&value)
}

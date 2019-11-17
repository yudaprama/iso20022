package model

// Information specifying the Mandate.
type OperationMandate3 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Unique and unambiguous identification of the mandate.
	Identification *Max35Text `xml:"Id"`

	// Channel for which the operation mandate is valid. If ApplicableChannel equals Fax, this means that a bank operation instruction sent by fax will be processed according to the mandates exchanged in this message.
	ApplicableChannel []*Channel2Choice `xml:"AplblChanl"`

	// Number of required and necessary signatures by the mandate.
	RequiredSignatureNumber *Max15PlusSignedNumericText `xml:"ReqrdSgntrNb"`

	// Indicator whether a certain order of signatures has to be respected or not.
	SignatureOrderIndicator *YesNoIndicator `xml:"SgntrOrdrInd"`

	// Holder of the mandate.
	MandateHolder []*PartyAndAuthorisation3 `xml:"MndtHldr,omitempty"`

	// Bank operation allowed by a mandate.
	BankOperation []*BankTransactionCodeStructure4 `xml:"BkOpr"`

	// Is the date when the mandate becomes valid.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Is the date when the mandate stops to be valid.
	EndDate *ISODate `xml:"EndDt,omitempty"`
}

func (o *OperationMandate3) SetModificationCode(value string) {
	o.ModificationCode = (*Modification1Code)(&value)
}

func (o *OperationMandate3) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OperationMandate3) AddApplicableChannel() *Channel2Choice {
	newValue := new(Channel2Choice)
	o.ApplicableChannel = append(o.ApplicableChannel, newValue)
	return newValue
}

func (o *OperationMandate3) SetRequiredSignatureNumber(value string) {
	o.RequiredSignatureNumber = (*Max15PlusSignedNumericText)(&value)
}

func (o *OperationMandate3) SetSignatureOrderIndicator(value string) {
	o.SignatureOrderIndicator = (*YesNoIndicator)(&value)
}

func (o *OperationMandate3) AddMandateHolder() *PartyAndAuthorisation3 {
	newValue := new(PartyAndAuthorisation3)
	o.MandateHolder = append(o.MandateHolder, newValue)
	return newValue
}

func (o *OperationMandate3) AddBankOperation() *BankTransactionCodeStructure4 {
	newValue := new(BankTransactionCodeStructure4)
	o.BankOperation = append(o.BankOperation, newValue)
	return newValue
}

func (o *OperationMandate3) SetStartDate(value string) {
	o.StartDate = (*ISODate)(&value)
}

func (o *OperationMandate3) SetEndDate(value string) {
	o.EndDate = (*ISODate)(&value)
}

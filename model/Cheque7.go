package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque7 struct {

	// Specifies the type of cheque to be issued.
	ChequeType *ChequeType2Code `xml:"ChqTp,omitempty"`

	// Unique and unambiguous identifier for a cheque as assigned by the agent.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Identifies the party that ordered the issuance of the cheque.
	ChequeFrom *NameAndAddress10 `xml:"ChqFr,omitempty"`

	// Specifies the delivery method of the cheque by the debtor's agent.
	DeliveryMethod *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`

	// Party to whom the debtor's agent needs to send the cheque.
	DeliverTo *NameAndAddress10 `xml:"DlvrTo,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the payment instruction to apply to the processing of the payment instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Date when the draft becomes payable and the debtor's account is debited.
	ChequeMaturityDate *ISODate `xml:"ChqMtrtyDt,omitempty"`

	// Identifies, in a coded form, the cheque layout, company logo and digitised signature to be used to print the cheque, as agreed between the initiating party and the debtor's agent.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Information that needs to be printed on a cheque, used by the payer to add miscellaneous information.
	MemoField []*Max35Text `xml:"MemoFld,omitempty"`

	// Regional area in which the cheque can be cleared, when a country has no nation-wide cheque clearing organisation.
	RegionalClearingZone *Max35Text `xml:"RgnlClrZone,omitempty"`

	// Specifies the print location of the cheque.
	PrintLocation *Max35Text `xml:"PrtLctn,omitempty"`

	// Signature to be used by the cheque servicer on a specific cheque to be printed.
	Signature []*Max70Text `xml:"Sgntr,omitempty"`
}

func (c *Cheque7) SetChequeType(value string) {
	c.ChequeType = (*ChequeType2Code)(&value)
}

func (c *Cheque7) SetChequeNumber(value string) {
	c.ChequeNumber = (*Max35Text)(&value)
}

func (c *Cheque7) AddChequeFrom() *NameAndAddress10 {
	c.ChequeFrom = new(NameAndAddress10)
	return c.ChequeFrom
}

func (c *Cheque7) AddDeliveryMethod() *ChequeDeliveryMethod1Choice {
	c.DeliveryMethod = new(ChequeDeliveryMethod1Choice)
	return c.DeliveryMethod
}

func (c *Cheque7) AddDeliverTo() *NameAndAddress10 {
	c.DeliverTo = new(NameAndAddress10)
	return c.DeliverTo
}

func (c *Cheque7) SetInstructionPriority(value string) {
	c.InstructionPriority = (*Priority2Code)(&value)
}

func (c *Cheque7) SetChequeMaturityDate(value string) {
	c.ChequeMaturityDate = (*ISODate)(&value)
}

func (c *Cheque7) SetFormsCode(value string) {
	c.FormsCode = (*Max35Text)(&value)
}

func (c *Cheque7) AddMemoField(value string) {
	c.MemoField = append(c.MemoField, (*Max35Text)(&value))
}

func (c *Cheque7) SetRegionalClearingZone(value string) {
	c.RegionalClearingZone = (*Max35Text)(&value)
}

func (c *Cheque7) SetPrintLocation(value string) {
	c.PrintLocation = (*Max35Text)(&value)
}

func (c *Cheque7) AddSignature(value string) {
	c.Signature = append(c.Signature, (*Max70Text)(&value))
}

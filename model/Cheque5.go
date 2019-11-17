package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque5 struct {

	// Specifies the type of cheque to be issued by the first agent.
	ChequeType *ChequeType2Code `xml:"ChqTp,omitempty"`

	// Identifies the cheque number.
	ChequeNumber *Max35Text `xml:"ChqNb,omitempty"`

	// Identifies the party that ordered the issuance of the cheque.
	ChequeFrom *NameAndAddress3 `xml:"ChqFr,omitempty"`

	// Specifies the delivery method of the cheque by the debtor's agent.
	DeliveryMethod *ChequeDeliveryMethod1Choice `xml:"DlvryMtd,omitempty"`

	// Identifies the party to whom the debtor's agent should send the cheque.
	DeliverTo *NameAndAddress3 `xml:"DlvrTo,omitempty"`

	// Urgency or order of importance that the originator would like the recipient of the payment instruction to apply to the processing of the payment instruction.
	InstructionPriority *Priority2Code `xml:"InstrPrty,omitempty"`

	// Date when the draft becomes payable and the debtor's account is debited.
	ChequeMaturityDate *ISODate `xml:"ChqMtrtyDt,omitempty"`

	// Code agreed between the initiating party and the debtor's agent, that specifies the cheque layout, company logo and digitised signature to be used to print the cheque.
	FormsCode *Max35Text `xml:"FrmsCd,omitempty"`

	// Information that needs to be printed on a cheque, used by the payer to add miscellaneous information.
	MemoField *Max35Text `xml:"MemoFld,omitempty"`

	// Regional area in which the cheque can be cleared, when a country has no nation-wide cheque clearing organisation.
	RegionalClearingZone *Max35Text `xml:"RgnlClrZone,omitempty"`

	// Specifies the print location of the cheque.
	PrintLocation *Max35Text `xml:"PrtLctn,omitempty"`
}

func (c *Cheque5) SetChequeType(value string) {
	c.ChequeType = (*ChequeType2Code)(&value)
}

func (c *Cheque5) SetChequeNumber(value string) {
	c.ChequeNumber = (*Max35Text)(&value)
}

func (c *Cheque5) AddChequeFrom() *NameAndAddress3 {
	c.ChequeFrom = new(NameAndAddress3)
	return c.ChequeFrom
}

func (c *Cheque5) AddDeliveryMethod() *ChequeDeliveryMethod1Choice {
	c.DeliveryMethod = new(ChequeDeliveryMethod1Choice)
	return c.DeliveryMethod
}

func (c *Cheque5) AddDeliverTo() *NameAndAddress3 {
	c.DeliverTo = new(NameAndAddress3)
	return c.DeliverTo
}

func (c *Cheque5) SetInstructionPriority(value string) {
	c.InstructionPriority = (*Priority2Code)(&value)
}

func (c *Cheque5) SetChequeMaturityDate(value string) {
	c.ChequeMaturityDate = (*ISODate)(&value)
}

func (c *Cheque5) SetFormsCode(value string) {
	c.FormsCode = (*Max35Text)(&value)
}

func (c *Cheque5) SetMemoField(value string) {
	c.MemoField = (*Max35Text)(&value)
}

func (c *Cheque5) SetRegionalClearingZone(value string) {
	c.RegionalClearingZone = (*Max35Text)(&value)
}

func (c *Cheque5) SetPrintLocation(value string) {
	c.PrintLocation = (*Max35Text)(&value)
}

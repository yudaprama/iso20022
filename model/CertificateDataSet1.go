package model

// Formal document used to record a fact and used as proof of the fact, in the context of a commercial trade transaction.
type CertificateDataSet1 struct {

	// Identifies the certificate data set.
	DataSetIdentification *DocumentIdentification1 `xml:"DataSetId"`

	// Specifies the type of the certificate.
	CertificateType *TradeCertificateType1Code `xml:"CertTp"`

	// Specifies if the certificate data set is required in relation to specific line items, and which line items.
	LineItem []*LineItemAndPOIdentification1 `xml:"LineItm,omitempty"`

	// Characteristics of the goods that are certified, in the context of a commercial trade transaction.
	CertifiedCharacteristics *CertifiedCharacteristics1Choice `xml:"CertfdChrtcs"`

	// Issue date of the document.
	IssueDate *ISODate `xml:"IsseDt"`

	// Place where the certificate was issued.
	PlaceOfIssue *PostalAddress5 `xml:"PlcOfIsse,omitempty"`

	// Issuer of the certificate, typically the inspection company or its agent.
	Issuer *PartyIdentification26 `xml:"Issr"`

	// Date(s) at which inspection of the goods took place.
	InspectionDate *DatePeriodDetails `xml:"InspctnDt,omitempty"`

	// Indicates that the inspection has been performed by an authorised inspector.
	AuthorisedInspectorIndicator *YesNoIndicator `xml:"AuthrsdInspctrInd,omitempty"`

	// Unique identifier of the document.
	CertificateIdentification *Max35Text `xml:"CertId"`

	// Transport information relative to the goods that are covered by the certificate.
	Transport *SingleTransport3 `xml:"Trnsprt,omitempty"`

	// Information about the goods and/or services of a trade transaction.
	GoodsDescription *Max70Text `xml:"GoodsDesc,omitempty"`

	// Party responsible for dispatching the goods.
	Consignor *PartyIdentification26 `xml:"Consgnr,omitempty"`

	// Party to whom the goods (which are the subject of the certificate) must be delivered.
	Consignee *PartyIdentification26 `xml:"Consgn,omitempty"`

	// Manufacturer of the goods which are the subject of the certificate.
	Manufacturer *PartyIdentification26 `xml:"Manfctr,omitempty"`

	// Additional and important information that could not be captured by structured fields.
	AdditionalInformation []*Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CertificateDataSet1) AddDataSetIdentification() *DocumentIdentification1 {
	c.DataSetIdentification = new(DocumentIdentification1)
	return c.DataSetIdentification
}

func (c *CertificateDataSet1) SetCertificateType(value string) {
	c.CertificateType = (*TradeCertificateType1Code)(&value)
}

func (c *CertificateDataSet1) AddLineItem() *LineItemAndPOIdentification1 {
	newValue := new(LineItemAndPOIdentification1)
	c.LineItem = append(c.LineItem, newValue)
	return newValue
}

func (c *CertificateDataSet1) AddCertifiedCharacteristics() *CertifiedCharacteristics1Choice {
	c.CertifiedCharacteristics = new(CertifiedCharacteristics1Choice)
	return c.CertifiedCharacteristics
}

func (c *CertificateDataSet1) SetIssueDate(value string) {
	c.IssueDate = (*ISODate)(&value)
}

func (c *CertificateDataSet1) AddPlaceOfIssue() *PostalAddress5 {
	c.PlaceOfIssue = new(PostalAddress5)
	return c.PlaceOfIssue
}

func (c *CertificateDataSet1) AddIssuer() *PartyIdentification26 {
	c.Issuer = new(PartyIdentification26)
	return c.Issuer
}

func (c *CertificateDataSet1) AddInspectionDate() *DatePeriodDetails {
	c.InspectionDate = new(DatePeriodDetails)
	return c.InspectionDate
}

func (c *CertificateDataSet1) SetAuthorisedInspectorIndicator(value string) {
	c.AuthorisedInspectorIndicator = (*YesNoIndicator)(&value)
}

func (c *CertificateDataSet1) SetCertificateIdentification(value string) {
	c.CertificateIdentification = (*Max35Text)(&value)
}

func (c *CertificateDataSet1) AddTransport() *SingleTransport3 {
	c.Transport = new(SingleTransport3)
	return c.Transport
}

func (c *CertificateDataSet1) SetGoodsDescription(value string) {
	c.GoodsDescription = (*Max70Text)(&value)
}

func (c *CertificateDataSet1) AddConsignor() *PartyIdentification26 {
	c.Consignor = new(PartyIdentification26)
	return c.Consignor
}

func (c *CertificateDataSet1) AddConsignee() *PartyIdentification26 {
	c.Consignee = new(PartyIdentification26)
	return c.Consignee
}

func (c *CertificateDataSet1) AddManufacturer() *PartyIdentification26 {
	c.Manufacturer = new(PartyIdentification26)
	return c.Manufacturer
}

func (c *CertificateDataSet1) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max350Text)(&value))
}

package auth

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document02200101 struct {
	XMLName xml.Name                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.022.001.01 Document"`
	Message *ContractRegistrationStatementV01 `xml:"CtrctRegnStmt"`
}

func (d *Document02200101) AddMessage() *ContractRegistrationStatementV01 {
	d.Message = new(ContractRegistrationStatementV01)
	return d.Message
}

// The ContractRegistrationStatement message is sent by the registration agent to the reporting party, in response to a request or at a pre-agreed date, to send a statement of the operations related to the registered contract subject to currency control.
type ContractRegistrationStatementV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *model.CurrencyControlHeader2 `xml:"GrpHdr"`

	// Provides the contract registration statement, which includes all journals on the activities related to the contract.
	Statement []*model.ContractRegistrationStatement1 `xml:"Stmt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationStatementV01) AddGroupHeader() *model.CurrencyControlHeader2 {
	c.GroupHeader = new(model.CurrencyControlHeader2)
	return c.GroupHeader
}

func (c *ContractRegistrationStatementV01) AddStatement() *model.ContractRegistrationStatement1 {
	newValue := new(model.ContractRegistrationStatement1)
	c.Statement = append(c.Statement, newValue)
	return newValue
}

func (c *ContractRegistrationStatementV01) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

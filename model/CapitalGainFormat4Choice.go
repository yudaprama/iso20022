package model

// Choice between a standard code or proprietary code to specify the type of capital gain.
type CapitalGainFormat4Choice struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June) for an income realised upon sale, a refund or redemption of shares and units etc.
	Code *EUCapitalGain2Code `xml:"Cd"`

	// Proprietary identification of the type of capital gain.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CapitalGainFormat4Choice) SetCode(value string) {
	c.Code = (*EUCapitalGain2Code)(&value)
}

func (c *CapitalGainFormat4Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}

package model

// Choice between tax calculation methods A, B or D.
type BillingMethod1Choice struct {

	// Tax values are based on tax calculation method A.
	MethodA *BillingMethod1 `xml:"MtdA"`

	// Tax values are based on tax calculation method B.
	MethodB *BillingMethod2 `xml:"MtdB"`

	// Tax values are based on tax calculation method D.
	MethodD *BillingMethod3 `xml:"MtdD"`
}

func (b *BillingMethod1Choice) AddMethodA() *BillingMethod1 {
	b.MethodA = new(BillingMethod1)
	return b.MethodA
}

func (b *BillingMethod1Choice) AddMethodB() *BillingMethod2 {
	b.MethodB = new(BillingMethod2)
	return b.MethodB
}

func (b *BillingMethod1Choice) AddMethodD() *BillingMethod3 {
	b.MethodD = new(BillingMethod3)
	return b.MethodD
}

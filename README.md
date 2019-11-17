# ISO20022

This repository contains a full set of Golang parse to the ISO-20022 data catalogs [ISO-20022 specifications](https://www.iso20022.org/full_catalogue.page)

## Usage

See `examples/` directory for an example of usage

```go
import (
	"encoding/xml"
	"github.com/yudaprama/iso20022/pacs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	messages, err := ioutil.ReadFile("./example-message.xml")
	if err != nil {
		log.Fatalf("Unable to read file:  %v", err)
		os.Exit(1)
	}

	var messageParsed pacs.Document00800106
	err = xml.Unmarshal(messages, &messageParsed)
	if err != nil {
		log.Fatalf("Unable to parse file:  %v", err)
		os.Exit(1)
	}
	
	log.Printf("Interbank Settlement Date:  %v", messageParsed.Message.GroupHeader.InterbankSettlementDate)
}

```

## Message Types

Message types covers ISO-20022 messages:

* [ACMT](https://github.com/yudaprama/iso20022/tree/master/acmt)
* [ADMI](https://github.com/yudaprama/iso20022/tree/master/admi)
* [AUTH](https://github.com/yudaprama/iso20022/tree/master/auth)
* [CAAA](https://github.com/yudaprama/iso20022/tree/master/caaa)
* [CAAM](https://github.com/yudaprama/iso20022/tree/master/caam)
* [CAIN](https://github.com/yudaprama/iso20022/tree/master/cain)
* [CAMT](https://github.com/yudaprama/iso20022/tree/master/camt)
* [CATM](https://github.com/yudaprama/iso20022/tree/master/catm)
* [CATP](https://github.com/yudaprama/iso20022/tree/master/catp)
* [COLR](https://github.com/yudaprama/iso20022/tree/master/colr)
* [FXTR](https://github.com/yudaprama/iso20022/tree/master/fxtr)
* [HEAD](https://github.com/yudaprama/iso20022/tree/master/head)
* [PACS](https://github.com/yudaprama/iso20022/tree/master/pacs)
* [PAIN](https://github.com/yudaprama/iso20022/tree/master/pain)
* [REDA](https://github.com/yudaprama/iso20022/tree/master/reda)
* [REMT](https://github.com/yudaprama/iso20022/tree/master/remt)
* [SECL](https://github.com/yudaprama/iso20022/tree/master/secl)
* [SEEV](https://github.com/yudaprama/iso20022/tree/master/seev)
* [SEMT](https://github.com/yudaprama/iso20022/tree/master/semt)
* [SESE](https://github.com/yudaprama/iso20022/tree/master/sese)
* [SETR](https://github.com/yudaprama/iso20022/tree/master/setr)
* [SUPL](https://github.com/yudaprama/iso20022/tree/master/supl)
* [TREA](https://github.com/yudaprama/iso20022/tree/master/trea)
* [TSIN](https://github.com/yudaprama/iso20022/tree/master/tsin)
* [TSMT](https://github.com/yudaprama/iso20022/tree/master/tsmt)
* [TSRV](https://github.com/yudaprama/iso20022/tree/master/tsrv)

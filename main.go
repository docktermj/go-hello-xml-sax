/*

References:
  - https://pkg.go.dev/encoding/xml
  - https://eli.thegreenplace.net/2019/faster-xml-stream-processing-in-go/
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {

	// Open the XML file.

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse file.

	decoder := xml.NewDecoder(file)
	for {
		xmlToken, err := decoder.Token()
		if xmlToken == nil || err == io.EOF {
			break // EOF means we're done.
		} else if err != nil {
			log.Fatalf("Error decoding token: %s", err)
		}

		switch typedXmlToken := xmlToken.(type) {
		case xml.CharData:
			aString := strings.TrimSpace(string(typedXmlToken))
			if len(aString) > 0 {
				fmt.Printf("CharData: %s\n", aString)
			}
		case xml.Comment:
			fmt.Printf("Comment: %s\n", typedXmlToken)
		case xml.Directive:
			fmt.Printf("Directive: %s\n", typedXmlToken)
		case xml.EndElement:
			fmt.Printf("EndElement: %s\n", typedXmlToken)
		case xml.ProcInst:
			fmt.Printf("ProcInst: %s\n", typedXmlToken)
		case xml.StartElement:
			fmt.Printf("StartElement: Token: %s; Name.Space: %s; Name.Local: %s\n", typedXmlToken, typedXmlToken.Name.Space, typedXmlToken.Name.Local)
		default:
			fmt.Printf(">>>>> Add case for `%s`: %v\n", reflect.TypeOf(typedXmlToken), xmlToken)
		}
	}
}

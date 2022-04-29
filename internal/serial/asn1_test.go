package serial

import (
	"encoding/asn1"
	"fmt"
	"unicode/utf8"
)

func ASN1Marshal(value interface{}) ([]byte, string) {

	m, _ := asn1.Marshal(value)

	var asnType string
	switch uint(m[0]) {
	case asn1.TagBoolean:
		asnType = "Boolean"
	case asn1.TagInteger:
		asnType = "Integer"
	case asn1.TagBitString:
		asnType = "BitString"
	case asn1.TagOctetString:
		asnType = "OctetString"
	case asn1.TagNull:
		asnType = "Null"
	case asn1.TagOID:
		asnType = "OID"
	case asn1.TagEnum:
		asnType = "Enum"
	case asn1.TagUTF8String:
		asnType = "UTF8String"
	case asn1.TagSequence:
		asnType = "Sequence"
	case asn1.TagSet:
		asnType = "Set"
	case asn1.TagNumericString:
		asnType = "NumericString"
	case asn1.TagPrintableString:
		asnType = "PrintableString"
	case asn1.TagT61String:
		asnType = "T61String"
	case asn1.TagIA5String:
		asnType = "IA5String"
	case asn1.TagUTCTime:
		asnType = "UTCTime"
	case asn1.TagGeneralizedTime:
		asnType = "GeneralizedTime"
	case asn1.TagGeneralString:
		asnType = "GeneralString"
	case asn1.TagBMPString:
		asnType = "BMPString"
	}

	result := fmt.Sprintf("Go Type: %T ASN.1 Type: %s Length: %v Value: %v", value, asnType, m[2], m[2:])

	return m, result
}

func Example_marshal() {
	encoded, description := ASN1Marshal("HE")
	fmt.Printf("Encoding: %v | %v\n", encoded, description)

	encoded, description = ASN1Marshal(rune('H'))
	fmt.Printf("Encoding: %v | %v\n", encoded, description)

	utf8Char := make([]byte, 3)
	utf8.EncodeRune(utf8Char, '世')
	encoded, description = ASN1Marshal(utf8Char)
	fmt.Printf("Encoding: %v | %v", encoded, description)

	// Output:
	// Encoding: [19 2 72 69] | Go Type: string ASN.1 Type: PrintableString Length: 72 Value: [72 69]
	// Encoding: [2 1 72] | Go Type: int32 ASN.1 Type: Integer Length: 72 Value: [72]
	// Encoding: [4 3 228 184 150] | Go Type: []uint8 ASN.1 Type: OctetString Length: 228 Value: [228 184 150]
}

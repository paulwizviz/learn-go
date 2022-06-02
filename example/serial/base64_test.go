package serial

import (
	"encoding/base64"
	"fmt"
)

func Example_std() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println("Source: ", data)

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoding: ", sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Decoding: ", string(sDec))

	// Output:
	// Source:  abc123!?$*&()'-=@~
	// Encoding:  YWJjMTIzIT8kKiYoKSctPUB+
	// Decoding:  abc123!?$*&()'-=@~

}

func Example_url() {
	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("STD Encoding: ", sEnc)

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("URL Encoding: ", uEnc)

	// Output:
	// STD Encoding:  YWJjMTIzIT8kKiYoKSctPUB+
	// URL Encoding:  YWJjMTIzIT8kKiYoKSctPUB-
}

func Example_raw() {
	data := "abc123!?$*&()'-=@~"

	rStdEnc := base64.RawStdEncoding.EncodeToString([]byte(data))
	fmt.Println("Raw STD encoding: ", rStdEnc)

	rURLEnc := base64.RawURLEncoding.EncodeToString([]byte(data))
	fmt.Println("Raw URL encoding: ", rURLEnc)

	// Output:
	// Raw STD encoding:  YWJjMTIzIT8kKiYoKSctPUB+
	// Raw URL encoding:  YWJjMTIzIT8kKiYoKSctPUB-
}

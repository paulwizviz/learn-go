package client

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const (
	soapEnv   = "http://schemas.xmlsoap.org/soap/envelope/"
	wsseXMLNS = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
)

type SoapHeader struct {
	XMLName xml.Name `xml:"soapenv:Header"`
}

type SoapBody struct {
	XMLName xml.Name `xml:"soap"`
}

type SoapEnv struct {
	XMLName      xml.Name `xml:"soap:Envelope"`
	XMLNSSoapEnv string   `xml:"xmlns:soapenv,attr"`
	Header       SoapHeader
	Body         SoapBody
}

type SOAPClient struct{}

func (s *SOAPClient) SoapCall(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/xml")
	req.Header.Set("SOAPAction", soapEnv)
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rBody, nil

}

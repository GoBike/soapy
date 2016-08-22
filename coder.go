package soapy

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

// Coder wraps header, provides methods to encode/decode soap-envolope.
type Coder struct {
	header interface{}
}

// NewCoder instantiate Coder instance.
func NewCoder(header interface{}) *Coder {
	return &Coder{
		header: header,
	}
}

// Encode encodes request to a readily bytes reader, which can be assigned to
// a HTTP body.
func (c *Coder) Encode(request interface{}) (*bytes.Buffer, error) {
	envelope := SOAPEnvelope{}

	if c.header != nil {
		envelope.Header = &SOAPHeader{Header: c.header}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return nil, fmt.Errorf("soapy: can't encode %v", err)
	}

	if err := encoder.Flush(); err != nil {
		return nil, fmt.Errorf("soapy: can't flush %v", err)
	}

	return buffer, nil
}

// Decode decodes raw bytes into response struct.
func (c *Coder) Decode(raw []byte, response interface{}) (err error) {
	if len(raw) == 0 {
		return errors.New("soapy: zero bytes to decode")
	}

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}

	err = xml.Unmarshal(raw, respEnvelope)
	if err != nil {
		return fmt.Errorf("soapy: unmarshal error %v", err)
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fmt.Errorf("soapy: fault returned %v", fault)
	}

	return nil
}

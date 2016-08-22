package soapy

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

// An Encoder writes SOAP values to an output stream.
type Encoder struct {
	w   io.Writer
	err error

	header interface{}
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode writes the SOAP encoding of v to the stream
func (enc *Encoder) Encode(v interface{}) error {

	if enc.err != nil {
		return enc.err
	}

	envelope := SOAPEnvelope{}

	if enc.header != nil {
		envelope.Header = &SOAPHeader{Header: enc.header}
	}

	envelope.Body.Content = v
	xmlenc := xml.NewEncoder(enc.w)
	//encoder.Indent("  ", "    ")

	if err := xmlenc.Encode(envelope); err != nil {
		return fmt.Errorf("soapy: can't encode %v", err)
	}

	if err := xmlenc.Flush(); err != nil {
		return fmt.Errorf("soapy: can't flush %v", err)
	}

	return nil
}

// A Decoder reads and decodes SOAP values from an input stream.
type Decoder struct {
	r   io.Reader
	buf []byte
	err error
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may
// read data from r beyond the SOAP values requested.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode reads from output stream and stores it in the value pointed to by v.
func (dec *Decoder) Decode(v interface{}) error {

	if dec.err != nil {
		return dec.err
	}

	rawbody, err := ioutil.ReadAll(dec.r) // careful with this.
	if err != nil {
		return err
	}

	if len(rawbody) == 0 {
		return fmt.Errorf("soapy: zero byte length to decode")
	}

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: v}

	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return fmt.Errorf("soapy: unmarshal error %v", err)
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fmt.Errorf("soapy: fault returned %v", fault)
	}

	return nil
}

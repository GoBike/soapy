package soapy

// SOAPEnvelope is the message exchange format between SOAP client & server.
import "encoding/xml"

// SOAPEnvelope is the message exchange format between SOAP client & server.
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

// SOAPHeader is a header part in SOAP envelope.
type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Header  interface{}
}

// SOAPBody is a body part in SOAP envelope.
type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

// SOAPFault encompases in SOAP envolope when error occurs, bounded by transaction.
type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

// BasicAuth is struct to user-credentials on services.
type BasicAuth struct {
	Login    string
	Password string
}

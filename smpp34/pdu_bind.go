package smpp34

import (
	"bytes"
)

var (
	reqBindFields = []string{
		SYSTEM_ID,
		PASSWORD,
		SYSTEM_TYPE,
		INTERFACE_VERSION,
		ADDR_TON,
		ADDR_NPI,
		ADDRESS_RANGE,
	}
)

type Bind struct {
	*Header
	mandatoryFields map[string]Field
	tlvFields       map[uint16]*TLVField
}

func NewBind(hdr *Header, b []byte) (*Bind, error) {
	r := bytes.NewBuffer(b)

	fields, _, err := create_pdu_fields(reqBindFields, r)

	if err != nil {
		return nil, err
	}

	s := &Bind{Header: hdr, mandatoryFields: fields}

	return s, nil
}

func (s *Bind) GetField(f string) Field {
	return s.mandatoryFields[f]
}

func (s *Bind) Fields() map[string]Field {
	return s.mandatoryFields
}

func (s *Bind) MandatoryFieldsList() []string {
	return reqBindFields
}

func (s *Bind) Ok() bool {
	return true
}

func (s *Bind) GetHeader() *Header {
	return s.Header
}

func (s *Bind) SetField(f string, v interface{}) error {
	if s.validate_field(f, v) {
		field := NewField(f, v)

		if field != nil {
			s.mandatoryFields[f] = field

			return nil
		}
	}

	return FieldValueErr
}

func (s *Bind) SetSeqNum(i uint32) {
	s.Header.Sequence = i
}

func (s *Bind) SetTLVField(t, l int, v []byte) error {
	return TLVFieldPduErr
}

func (s *Bind) validate_field(f string, v interface{}) bool {
	if included_check(s.MandatoryFieldsList(), f) && validate_pdu_field(f, v) {
		return true
	}
	return false
}

func (s *Bind) TLVFields() map[uint16]*TLVField {
	return s.tlvFields
}

func (s *Bind) writeFields() []byte {
	b := []byte{}

	for _, i := range s.MandatoryFieldsList() {
		v := s.mandatoryFields[i].ByteArray()
		b = append(b, v...)
	}

	return b
}

func (s *Bind) Writer() []byte {
	b := s.writeFields()
	h := packUi32(uint32(len(b) + 16))
	h = append(h, packUi32(uint32(s.Header.Id))...)
	h = append(h, packUi32(uint32(s.Header.Status))...)
	h = append(h, packUi32(s.Header.Sequence)...)
	return append(h, b...)
}

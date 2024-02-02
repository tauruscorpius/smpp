package smpp34

import (
	"bytes"
)

var (
	reqSSMRespFields = []string{MESSAGE_ID}
)

type SubmitSmResp struct {
	*Header
	mandatoryFields map[string]Field
	tlvFields       map[uint16]*TLVField
}

func NewSubmitSmResp(hdr *Header, b []byte) (*SubmitSmResp, error) {
	r := bytes.NewBuffer(b)

	fields, _, err := create_pdu_fields(reqSSMRespFields, r)

	if err != nil {
		return nil, err
	}

	s := &SubmitSmResp{Header: hdr, mandatoryFields: fields}

	return s, nil
}

func (s *SubmitSmResp) GetField(f string) Field {
	return s.mandatoryFields[f]
}

func (s *SubmitSmResp) Fields() map[string]Field {
	return s.mandatoryFields
}

func (s *SubmitSmResp) MandatoryFieldsList() []string {
	return reqSSMRespFields
}

func (s *SubmitSmResp) Ok() bool {
	if s.Header.Status == ESME_ROK {
		return true
	}

	return false
}

func (s *SubmitSmResp) GetHeader() *Header {
	return s.Header
}

func (s *SubmitSmResp) SetField(f string, v interface{}) error {
	if s.validate_field(f, v) {
		field := NewField(f, v)

		if field != nil {
			s.mandatoryFields[f] = field

			return nil
		}
	}

	return FieldValueErr
}

func (s *SubmitSmResp) SetSeqNum(i uint32) {
	s.Header.Sequence = i
}

func (s *SubmitSmResp) SetTLVField(t, l int, v []byte) error {
	return TLVFieldPduErr
}

func (s *SubmitSmResp) validate_field(f string, v interface{}) bool {
	if included_check(s.MandatoryFieldsList(), f) && validate_pdu_field(f, v) {
		return true
	}
	return false
}

func (s *SubmitSmResp) TLVFields() map[uint16]*TLVField {
	return s.tlvFields
}

func (s *SubmitSmResp) writeFields() []byte {
	b := []byte{}

	for _, i := range s.MandatoryFieldsList() {
		v := s.mandatoryFields[i].ByteArray()
		b = append(b, v...)
	}

	return b
}

func (s *SubmitSmResp) Writer() []byte {
	b := s.writeFields()
	h := packUi32(uint32(len(b) + 16))
	h = append(h, packUi32(uint32(SUBMIT_SM_RESP))...)
	h = append(h, packUi32(uint32(s.Header.Status))...)
	h = append(h, packUi32(s.Header.Sequence)...)

	return append(h, b...)
}

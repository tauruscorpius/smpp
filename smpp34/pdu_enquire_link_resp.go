package smpp34

var (
	reqELRespFields = []string{}
)

type EnquireLinkResp struct {
	*Header
	mandatoryFields map[string]Field
	tlvFields       map[uint16]*TLVField
}

func NewEnquireLinkResp(hdr *Header) (*EnquireLinkResp, error) {
	s := &EnquireLinkResp{Header: hdr}

	return s, nil
}

func (s *EnquireLinkResp) GetField(f string) Field {
	return nil
}

func (s *EnquireLinkResp) SetField(f string, v interface{}) error {
	return FieldValueErr
}

func (s *EnquireLinkResp) SetSeqNum(i uint32) {
	s.Header.Sequence = i
}

func (s *EnquireLinkResp) SetTLVField(t, l int, v []byte) error {
	return TLVFieldPduErr
}

func (s *EnquireLinkResp) Fields() map[string]Field {
	return s.mandatoryFields
}

func (s *EnquireLinkResp) MandatoryFieldsList() []string {
	return reqELRespFields
}

func (s *EnquireLinkResp) Ok() bool {
	if s.Header.Status == ESME_ROK {
		return true
	}

	return false
}

func (s *EnquireLinkResp) GetHeader() *Header {
	return s.Header
}

func (s *EnquireLinkResp) TLVFields() map[uint16]*TLVField {
	return s.tlvFields
}

func (s *EnquireLinkResp) writeFields() []byte {
	return []byte{}
}

func (s *EnquireLinkResp) Writer() []byte {
	b := s.writeFields()
	h := packUi32(uint32(len(b) + 16))
	h = append(h, packUi32(uint32(ENQUIRE_LINK_RESP))...)
	h = append(h, packUi32(uint32(s.Header.Status))...)
	h = append(h, packUi32(s.Header.Sequence)...)
	return append(h, b...)
}

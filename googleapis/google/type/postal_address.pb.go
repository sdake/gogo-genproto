// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/postal_address.proto

package google_type

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Represents a postal address, e.g. for postal delivery or payments addresses.
// Given a postal address, a postal service can deliver items to a premise, P.O.
// Box or similar.
// It is not intended to model geographical locations (roads, towns,
// mountains).
//
// In typical usage an address would be created via user input or from importing
// existing data, depending on the type of process.
//
// Advice on address input / editing:
//  - Use an i18n-ready address widget such as
//    https://github.com/google/libaddressinput)
// - Users should not be presented with UI elements for input or editing of
//   fields outside countries where that field is used.
//
// For more guidance on how to use this schema, please see:
// https://support.google.com/business/answer/6397478
type PostalAddress struct {
	// The schema revision of the `PostalAddress`.
	// All new revisions **must** be backward compatible with old revisions.
	Revision int32 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. CLDR region code of the country/region of the address. This
	// is never inferred and it is up to the user to ensure the value is
	// correct. See http://cldr.unicode.org/ and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if
	// known). This is often the UI language of the input form or is expected
	// to match one of the languages used in the address' country/region, or their
	// transliterated equivalents.
	// This can affect formatting in certain countries, but is not critical
	// to the correctness of the data and will never affect any validation or
	// other non-formatting related operations.
	//
	// If this value is not known, it should be omitted (rather than specifying a
	// possibly incorrect default).
	//
	// Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require
	// postal codes to be present, but where they are used, they may trigger
	// additional validation with other parts of the address (e.g. state/zip
	// validation in the U.S.A.).
	PostalCode string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used
	// in most regions. Where it is used, the value is either a string like
	// "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	// alone, representing the "sector code" (Jamaica), "delivery area indicator"
	// (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode string `protobuf:"bytes,5,opt,name=sorting_code,json=sortingCode,proto3" json:"sorting_code,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal
	// addresses of a country or region.
	// For example, this can be a state, a province, an oblast, or a prefecture.
	// Specifically, for Spain this is the province and not the autonomous
	// community (e.g. "Barcelona" and not "Catalonia").
	// Many countries don't use an administrative area in postal addresses. E.g.
	// in Switzerland this should be left unpopulated.
	AdministrativeArea string `protobuf:"bytes,6,opt,name=administrative_area,json=administrativeArea,proto3" json:"administrative_area,omitempty"`
	// Optional. Generally refers to the city/town portion of the address.
	// Examples: US city, IT comune, UK post town.
	// In regions of the world where localities are not well defined or do not fit
	// into this structure well, leave locality empty and use address_lines.
	Locality string `protobuf:"bytes,7,opt,name=locality,proto3" json:"locality,omitempty"`
	// Optional. Sublocality of the address.
	// For example, this can be neighborhoods, boroughs, districts.
	Sublocality string `protobuf:"bytes,8,opt,name=sublocality,proto3" json:"sublocality,omitempty"`
	// Unstructured address lines describing the lower levels of an address.
	//
	// Because values in address_lines do not have type information and may
	// sometimes contain multiple values in a single field (e.g.
	// "Austin, TX"), it is important that the line order is clear. The order of
	// address lines should be "envelope order" for the country/region of the
	// address. In places where this can vary (e.g. Japan), address_language is
	// used to make it explicit (e.g. "ja" for large-to-small ordering and
	// "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	// an address can be selected based on the language.
	//
	// The minimum permitted structural representation of an address consists
	// of a region_code with all remaining information placed in the
	// address_lines. It would be possible to format such an address very
	// approximately without geocoding, but no semantic reasoning could be
	// made about any of the address components until it was at least
	// partially resolved.
	//
	// Creating an address only containing a region_code and address_lines, and
	// then geocoding is the recommended way to handle completely unstructured
	// addresses (as opposed to guessing which parts of the address should be
	// localities or administrative areas).
	AddressLines []string `protobuf:"bytes,9,rep,name=address_lines,json=addressLines,proto3" json:"address_lines,omitempty"`
	// Optional. The recipient at the address.
	// This field may, under certain circumstances, contain multiline information.
	// For example, it might contain "care of" information.
	Recipients []string `protobuf:"bytes,10,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Optional. The name of the organization at the address.
	Organization string `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (m *PostalAddress) Reset()      { *m = PostalAddress{} }
func (*PostalAddress) ProtoMessage() {}
func (*PostalAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_17c0e9bca935790c, []int{0}
}
func (m *PostalAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostalAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostalAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostalAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostalAddress.Merge(m, src)
}
func (m *PostalAddress) XXX_Size() int {
	return m.Size()
}
func (m *PostalAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PostalAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PostalAddress proto.InternalMessageInfo

func (m *PostalAddress) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *PostalAddress) GetRegionCode() string {
	if m != nil {
		return m.RegionCode
	}
	return ""
}

func (m *PostalAddress) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *PostalAddress) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *PostalAddress) GetSortingCode() string {
	if m != nil {
		return m.SortingCode
	}
	return ""
}

func (m *PostalAddress) GetAdministrativeArea() string {
	if m != nil {
		return m.AdministrativeArea
	}
	return ""
}

func (m *PostalAddress) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *PostalAddress) GetSublocality() string {
	if m != nil {
		return m.Sublocality
	}
	return ""
}

func (m *PostalAddress) GetAddressLines() []string {
	if m != nil {
		return m.AddressLines
	}
	return nil
}

func (m *PostalAddress) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *PostalAddress) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*PostalAddress)(nil), "google.type.PostalAddress")
}

func init() { proto.RegisterFile("google/type/postal_address.proto", fileDescriptor_17c0e9bca935790c) }

var fileDescriptor_17c0e9bca935790c = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xb1, 0x6e, 0xe2, 0x30,
	0x18, 0xc7, 0x63, 0x72, 0x70, 0xf0, 0x05, 0x74, 0x92, 0x6f, 0x89, 0x6e, 0x30, 0x39, 0x6e, 0x61,
	0x82, 0xe1, 0x9e, 0x00, 0x3a, 0x74, 0xe9, 0x10, 0xa1, 0xee, 0x91, 0x49, 0xac, 0xc8, 0x52, 0xb0,
	0x23, 0xdb, 0x20, 0xd1, 0xa9, 0x2f, 0x50, 0xa9, 0xcf, 0xd0, 0xa9, 0x8f, 0xd2, 0x91, 0x91, 0xb1,
	0x84, 0xa5, 0x23, 0x8f, 0x50, 0xc5, 0x4e, 0x29, 0x8c, 0xdf, 0xef, 0xfb, 0xc9, 0xf6, 0xf7, 0xfd,
	0x0d, 0x51, 0x2e, 0x65, 0x5e, 0xb0, 0xa9, 0xd9, 0x96, 0x6c, 0x5a, 0x4a, 0x6d, 0x68, 0x91, 0xd0,
	0x2c, 0x53, 0x4c, 0xeb, 0x49, 0xa9, 0xa4, 0x91, 0x38, 0x70, 0xc6, 0xa4, 0x36, 0x46, 0x4f, 0x3e,
	0x0c, 0x62, 0x6b, 0xcd, 0x9c, 0x84, 0xff, 0x40, 0x57, 0xb1, 0x0d, 0xd7, 0x5c, 0x8a, 0x10, 0x45,
	0x68, 0xdc, 0x5e, 0x9c, 0x6b, 0x3c, 0x84, 0x40, 0xb1, 0x9c, 0x4b, 0x91, 0xa4, 0x32, 0x63, 0x61,
	0x2b, 0x42, 0xe3, 0xde, 0x02, 0x1c, 0xba, 0x91, 0x19, 0xc3, 0xff, 0x60, 0x50, 0x50, 0x91, 0xaf,
	0x69, 0xce, 0x9c, 0xe2, 0x5b, 0xa5, 0xff, 0x05, 0xad, 0x34, 0x84, 0xa0, 0x79, 0x98, 0x55, 0x7e,
	0xb8, 0x53, 0x1c, 0xb2, 0xc2, 0x5f, 0xe8, 0x6b, 0xa9, 0x0c, 0x17, 0xb9, 0x33, 0xda, 0xd6, 0x08,
	0x1a, 0x66, 0x95, 0x29, 0xfc, 0xa6, 0xd9, 0x8a, 0x0b, 0xae, 0x8d, 0xa2, 0x86, 0x6f, 0x58, 0x42,
	0x15, 0xa3, 0x61, 0xc7, 0x9a, 0xf8, 0xba, 0x35, 0x53, 0x8c, 0xd6, 0x63, 0x15, 0x32, 0xa5, 0x05,
	0x37, 0xdb, 0xf0, 0xa7, 0xb5, 0xce, 0x35, 0x8e, 0x20, 0xd0, 0xeb, 0xe5, 0xb9, 0xdd, 0x6d, 0xae,
	0xfb, 0x46, 0xf5, 0x5c, 0xcd, 0x12, 0x93, 0x82, 0x0b, 0xa6, 0xc3, 0x5e, 0xe4, 0xd7, 0x73, 0x35,
	0xf0, 0xae, 0x66, 0x98, 0x00, 0x28, 0x96, 0xf2, 0x92, 0x33, 0x61, 0x74, 0x08, 0xd6, 0xb8, 0x20,
	0x78, 0x04, 0x7d, 0xa9, 0x72, 0x2a, 0xf8, 0x03, 0x35, 0xf5, 0x76, 0x03, 0xb7, 0x9b, 0x4b, 0x36,
	0xa7, 0xbb, 0x03, 0xf1, 0xf6, 0x07, 0xe2, 0x9d, 0x0e, 0x04, 0x3d, 0x56, 0x04, 0xbd, 0x56, 0x04,
	0xbd, 0x55, 0x04, 0xed, 0x2a, 0x82, 0xde, 0x2b, 0x82, 0x3e, 0x2a, 0xe2, 0x9d, 0x2a, 0x82, 0x9e,
	0x8f, 0xc4, 0xdb, 0x1d, 0x89, 0xb7, 0x3f, 0x12, 0x0f, 0x7e, 0xa5, 0x72, 0x35, 0xb9, 0x88, 0x75,
	0x8e, 0xaf, 0x32, 0x8d, 0xeb, 0xdc, 0x63, 0xf4, 0xd2, 0xf2, 0x6f, 0xef, 0xe3, 0x65, 0xc7, 0x7e,
	0x83, 0xff, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x56, 0x56, 0x16, 0x2a, 0x02, 0x00, 0x00,
}

func (this *PostalAddress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PostalAddress)
	if !ok {
		that2, ok := that.(PostalAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Revision != that1.Revision {
		return false
	}
	if this.RegionCode != that1.RegionCode {
		return false
	}
	if this.LanguageCode != that1.LanguageCode {
		return false
	}
	if this.PostalCode != that1.PostalCode {
		return false
	}
	if this.SortingCode != that1.SortingCode {
		return false
	}
	if this.AdministrativeArea != that1.AdministrativeArea {
		return false
	}
	if this.Locality != that1.Locality {
		return false
	}
	if this.Sublocality != that1.Sublocality {
		return false
	}
	if len(this.AddressLines) != len(that1.AddressLines) {
		return false
	}
	for i := range this.AddressLines {
		if this.AddressLines[i] != that1.AddressLines[i] {
			return false
		}
	}
	if len(this.Recipients) != len(that1.Recipients) {
		return false
	}
	for i := range this.Recipients {
		if this.Recipients[i] != that1.Recipients[i] {
			return false
		}
	}
	if this.Organization != that1.Organization {
		return false
	}
	return true
}
func (this *PostalAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 15)
	s = append(s, "&google_type.PostalAddress{")
	s = append(s, "Revision: "+fmt.Sprintf("%#v", this.Revision)+",\n")
	s = append(s, "RegionCode: "+fmt.Sprintf("%#v", this.RegionCode)+",\n")
	s = append(s, "LanguageCode: "+fmt.Sprintf("%#v", this.LanguageCode)+",\n")
	s = append(s, "PostalCode: "+fmt.Sprintf("%#v", this.PostalCode)+",\n")
	s = append(s, "SortingCode: "+fmt.Sprintf("%#v", this.SortingCode)+",\n")
	s = append(s, "AdministrativeArea: "+fmt.Sprintf("%#v", this.AdministrativeArea)+",\n")
	s = append(s, "Locality: "+fmt.Sprintf("%#v", this.Locality)+",\n")
	s = append(s, "Sublocality: "+fmt.Sprintf("%#v", this.Sublocality)+",\n")
	s = append(s, "AddressLines: "+fmt.Sprintf("%#v", this.AddressLines)+",\n")
	s = append(s, "Recipients: "+fmt.Sprintf("%#v", this.Recipients)+",\n")
	s = append(s, "Organization: "+fmt.Sprintf("%#v", this.Organization)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPostalAddress(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PostalAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostalAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Revision != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(m.Revision))
	}
	if len(m.RegionCode) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.RegionCode)))
		i += copy(dAtA[i:], m.RegionCode)
	}
	if len(m.LanguageCode) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.LanguageCode)))
		i += copy(dAtA[i:], m.LanguageCode)
	}
	if len(m.PostalCode) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.PostalCode)))
		i += copy(dAtA[i:], m.PostalCode)
	}
	if len(m.SortingCode) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.SortingCode)))
		i += copy(dAtA[i:], m.SortingCode)
	}
	if len(m.AdministrativeArea) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.AdministrativeArea)))
		i += copy(dAtA[i:], m.AdministrativeArea)
	}
	if len(m.Locality) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.Locality)))
		i += copy(dAtA[i:], m.Locality)
	}
	if len(m.Sublocality) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.Sublocality)))
		i += copy(dAtA[i:], m.Sublocality)
	}
	if len(m.AddressLines) > 0 {
		for _, s := range m.AddressLines {
			dAtA[i] = 0x4a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Recipients) > 0 {
		for _, s := range m.Recipients {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Organization) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintPostalAddress(dAtA, i, uint64(len(m.Organization)))
		i += copy(dAtA[i:], m.Organization)
	}
	return i, nil
}

func encodeVarintPostalAddress(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PostalAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovPostalAddress(uint64(m.Revision))
	}
	l = len(m.RegionCode)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.LanguageCode)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.PostalCode)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.SortingCode)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.AdministrativeArea)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.Locality)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	l = len(m.Sublocality)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	if len(m.AddressLines) > 0 {
		for _, s := range m.AddressLines {
			l = len(s)
			n += 1 + l + sovPostalAddress(uint64(l))
		}
	}
	if len(m.Recipients) > 0 {
		for _, s := range m.Recipients {
			l = len(s)
			n += 1 + l + sovPostalAddress(uint64(l))
		}
	}
	l = len(m.Organization)
	if l > 0 {
		n += 1 + l + sovPostalAddress(uint64(l))
	}
	return n
}

func sovPostalAddress(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPostalAddress(x uint64) (n int) {
	return sovPostalAddress(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PostalAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PostalAddress{`,
		`Revision:` + fmt.Sprintf("%v", this.Revision) + `,`,
		`RegionCode:` + fmt.Sprintf("%v", this.RegionCode) + `,`,
		`LanguageCode:` + fmt.Sprintf("%v", this.LanguageCode) + `,`,
		`PostalCode:` + fmt.Sprintf("%v", this.PostalCode) + `,`,
		`SortingCode:` + fmt.Sprintf("%v", this.SortingCode) + `,`,
		`AdministrativeArea:` + fmt.Sprintf("%v", this.AdministrativeArea) + `,`,
		`Locality:` + fmt.Sprintf("%v", this.Locality) + `,`,
		`Sublocality:` + fmt.Sprintf("%v", this.Sublocality) + `,`,
		`AddressLines:` + fmt.Sprintf("%v", this.AddressLines) + `,`,
		`Recipients:` + fmt.Sprintf("%v", this.Recipients) + `,`,
		`Organization:` + fmt.Sprintf("%v", this.Organization) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPostalAddress(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PostalAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPostalAddress
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PostalAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostalAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegionCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LanguageCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LanguageCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostalCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostalCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortingCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SortingCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdministrativeArea", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdministrativeArea = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locality", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locality = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sublocality", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sublocality = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressLines", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressLines = append(m.AddressLines, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipients", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipients = append(m.Recipients, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPostalAddress
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organization = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPostalAddress(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPostalAddress
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPostalAddress(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPostalAddress
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPostalAddress
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPostalAddress
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPostalAddress
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPostalAddress
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPostalAddress(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPostalAddress
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPostalAddress = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPostalAddress   = fmt.Errorf("proto: integer overflow")
)

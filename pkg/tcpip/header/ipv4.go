// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package header

import (
	"encoding/binary"
	"errors"
	"fmt"

	"gvisor.dev/gvisor/pkg/tcpip"
)

// RFC 971 defines the fields of the IPv4 header on page 11 using the following
// diagram: ("Figure 4")
//    0                   1                   2                   3
//    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |Version|  IHL  |Type of Service|          Total Length         |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |         Identification        |Flags|      Fragment Offset    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |  Time to Live |    Protocol   |         Header Checksum       |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                       Source Address                          |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Destination Address                        |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//   |                    Options                    |    Padding    |
//   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
const (
	versIHL = 0
	tos     = 1
	// IPv4TotalLenOffset is the offset of the total length field in the
	// IPv4 header.
	IPv4TotalLenOffset = 2
	id                 = 4
	flagsFO            = 6
	ttl                = 8
	protocol           = 9
	checksum           = 10
	srcAddr            = 12
	dstAddr            = 16
	options            = 20
)

// IPv4Fields contains the fields of an IPv4 packet. It is used to describe the
// fields of a packet that needs to be encoded. The IHL field is not here as
// it is totally defined by the size of the options.
type IPv4Fields struct {
	// TOS is the "type of service" field of an IPv4 packet.
	TOS uint8

	// TotalLength is the "total length" field of an IPv4 packet.
	TotalLength uint16

	// ID is the "identification" field of an IPv4 packet.
	ID uint16

	// Flags is the "flags" field of an IPv4 packet.
	Flags uint8

	// FragmentOffset is the "fragment offset" field of an IPv4 packet.
	FragmentOffset uint16

	// TTL is the "time to live" field of an IPv4 packet.
	TTL uint8

	// Protocol is the "protocol" field of an IPv4 packet.
	Protocol uint8

	// Checksum is the "checksum" field of an IPv4 packet.
	Checksum uint16

	// SrcAddr is the "source ip address" of an IPv4 packet.
	SrcAddr tcpip.Address

	// DstAddr is the "destination ip address" of an IPv4 packet.
	DstAddr tcpip.Address

	// Options is between 0 and 40 bytes or nil if empty.
	Options IPv4Options
}

// IPv4 is an IPv4 header.
// Most of the methods of IPv4 access to the underlying slice without
// checking the boundaries and could panic because of 'index out of range'.
// Always call IsValid() to validate an instance of IPv4 before using other
// methods.
type IPv4 []byte

const (
	// IPv4MinimumSize is the minimum size of a valid IPv4 packet;
	// i.e. a packet header with no options.
	IPv4MinimumSize = 20

	// IPv4MaximumHeaderSize is the maximum size of an IPv4 header. Given
	// that there are only 4 bits (max 0xF (15)) to represent the header length
	// in 32-bit (4 byte) units, the header cannot exceed 15*4 = 60 bytes.
	IPv4MaximumHeaderSize = 60

	// IPv4MaximumOptionsSize is the largest size the IPv4 options can be.
	IPv4MaximumOptionsSize = IPv4MaximumHeaderSize - IPv4MinimumSize

	// IPv4MaximumPayloadSize is the maximum size of a valid IPv4 payload.
	//
	// Linux limits this to 65,515 octets (the max IP datagram size - the IPv4
	// header size). But RFC 791 section 3.2 discusses the design of the IPv4
	// fragment "allows 2**13 = 8192 fragments of 8 octets each for a total of
	// 65,536 octets. Note that this is consistent with the datagram total
	// length field (of course, the header is counted in the total length and not
	// in the fragments)."
	IPv4MaximumPayloadSize = 65536

	// MinIPFragmentPayloadSize is the minimum number of payload bytes that
	// the first fragment must carry when an IPv4 packet is fragmented.
	MinIPFragmentPayloadSize = 8

	// IPv4AddressSize is the size, in bytes, of an IPv4 address.
	IPv4AddressSize = 4

	// IPv4ProtocolNumber is IPv4's network protocol number.
	IPv4ProtocolNumber tcpip.NetworkProtocolNumber = 0x0800

	// IPv4Version is the version of the IPv4 protocol.
	IPv4Version = 4

	// IPv4AllSystems is the all systems IPv4 multicast address as per
	// IANA's IPv4 Multicast Address Space Registry. See
	// https://www.iana.org/assignments/multicast-addresses/multicast-addresses.xhtml.
	IPv4AllSystems tcpip.Address = "\xe0\x00\x00\x01"

	// IPv4Broadcast is the broadcast address of the IPv4 procotol.
	IPv4Broadcast tcpip.Address = "\xff\xff\xff\xff"

	// IPv4Any is the non-routable IPv4 "any" meta address.
	IPv4Any tcpip.Address = "\x00\x00\x00\x00"

	// IPv4MinimumProcessableDatagramSize is the minimum size of an IP
	// packet that every IPv4 capable host must be able to
	// process/reassemble.
	IPv4MinimumProcessableDatagramSize = 576

	// IPv4MinimumMTU is the minimum MTU required by IPv4, per RFC 791,
	// section 3.2:
	//   Every internet module must be able to forward a datagram of 68 octets
	//   without further fragmentation.  This is because an internet header may be
	//   up to 60 octets, and the minimum fragment is 8 octets.
	IPv4MinimumMTU = 68
)

// Flags that may be set in an IPv4 packet.
const (
	IPv4FlagMoreFragments = 1 << iota
	IPv4FlagDontFragment
)

// IPv4EmptySubnet is the empty IPv4 subnet.
var IPv4EmptySubnet = func() tcpip.Subnet {
	subnet, err := tcpip.NewSubnet(IPv4Any, tcpip.AddressMask(IPv4Any))
	if err != nil {
		panic(err)
	}
	return subnet
}()

// IPVersion returns the version of IP used in the given packet. It returns -1
// if the packet is not large enough to contain the version field.
func IPVersion(b []byte) int {
	// Length must be at least offset+length of version field.
	if len(b) < versIHL+1 {
		return -1
	}
	return int(b[versIHL] >> ipVersionShift)
}

// RFC 791 page 11 shows the header length (IHL) is in the lower 4 bits
// of the first byte, and is counted in multiples of 4 bytes.
//
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |Version|  IHL  |Type of Service|          Total Length         |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//      (...)
//     Version:  4 bits
//       The Version field indicates the format of the internet header.  This
//       document describes version 4.
//
//     IHL:  4 bits
//       Internet Header Length is the length of the internet header in 32
//       bit words, and thus points to the beginning of the data.  Note that
//       the minimum value for a correct header is 5.
const (
	ipVersionShift = 4
	ipIHLMask      = 0x0f
	IPv4IHLStride  = 4
)

// HeaderLength returns the value of the "header length" field of the IPv4
// header. The length returned is in bytes.
func (b IPv4) HeaderLength() uint8 {
	return (b[versIHL] & ipIHLMask) * IPv4IHLStride
}

// SetHeaderLength sets the value of the "Internet Header Length" field.
func (b IPv4) SetHeaderLength(hdrLen uint8) {
	if hdrLen > IPv4MaximumHeaderSize {
		panic(fmt.Sprintf("got IPv4 Header size = %d, want <= %d", hdrLen, IPv4MaximumHeaderSize))
	}
	b[versIHL] = (IPv4Version << ipVersionShift) | ((hdrLen / IPv4IHLStride) & ipIHLMask)
}

// ID returns the value of the identifier field of the IPv4 header.
func (b IPv4) ID() uint16 {
	return binary.BigEndian.Uint16(b[id:])
}

// Protocol returns the value of the protocol field of the IPv4 header.
func (b IPv4) Protocol() uint8 {
	return b[protocol]
}

// Flags returns the "flags" field of the IPv4 header.
func (b IPv4) Flags() uint8 {
	return uint8(binary.BigEndian.Uint16(b[flagsFO:]) >> 13)
}

// More returns whether the more fragments flag is set.
func (b IPv4) More() bool {
	return b.Flags()&IPv4FlagMoreFragments != 0
}

// TTL returns the "TTL" field of the IPv4 header.
func (b IPv4) TTL() uint8 {
	return b[ttl]
}

// FragmentOffset returns the "fragment offset" field of the IPv4 header.
func (b IPv4) FragmentOffset() uint16 {
	return binary.BigEndian.Uint16(b[flagsFO:]) << 3
}

// TotalLength returns the "total length" field of the IPv4 header.
func (b IPv4) TotalLength() uint16 {
	return binary.BigEndian.Uint16(b[IPv4TotalLenOffset:])
}

// Checksum returns the checksum field of the IPv4 header.
func (b IPv4) Checksum() uint16 {
	return binary.BigEndian.Uint16(b[checksum:])
}

// SourceAddress returns the "source address" field of the IPv4 header.
func (b IPv4) SourceAddress() tcpip.Address {
	return tcpip.Address(b[srcAddr : srcAddr+IPv4AddressSize])
}

// DestinationAddress returns the "destination address" field of the IPv4
// header.
func (b IPv4) DestinationAddress() tcpip.Address {
	return tcpip.Address(b[dstAddr : dstAddr+IPv4AddressSize])
}

// IPv4Options is a buffer that holds all the raw IP options.
type IPv4Options []byte

// SizeWithPadding implements stack.NetOptions.
// It reports the size to allocate for the Options. RFC 791 page 23 (end of
// section 3.1) says of the padding at the end of the options:
//    The internet header padding is used to ensure that the internet
//    header ends on a 32 bit boundary.
func (o IPv4Options) SizeWithPadding() int {
	return (len(o) + IPv4IHLStride - 1) & ^(IPv4IHLStride - 1)
}

// Options returns a buffer holding the options or nil.
func (b IPv4) Options() IPv4Options {
	hdrLen := b.HeaderLength()
	if hdrLen > IPv4MinimumSize {
		return IPv4Options(b[options:hdrLen:hdrLen])
	}
	return nil
}

// TransportProtocol implements Network.TransportProtocol.
func (b IPv4) TransportProtocol() tcpip.TransportProtocolNumber {
	return tcpip.TransportProtocolNumber(b.Protocol())
}

// Payload implements Network.Payload.
func (b IPv4) Payload() []byte {
	return b[b.HeaderLength():][:b.PayloadLength()]
}

// PayloadLength returns the length of the payload portion of the IPv4 packet.
func (b IPv4) PayloadLength() uint16 {
	return b.TotalLength() - uint16(b.HeaderLength())
}

// TOS returns the "type of service" field of the IPv4 header.
func (b IPv4) TOS() (uint8, uint32) {
	return b[tos], 0
}

// SetTOS sets the "type of service" field of the IPv4 header.
func (b IPv4) SetTOS(v uint8, _ uint32) {
	b[tos] = v
}

// SetTTL sets the "Time to Live" field of the IPv4 header.
func (b IPv4) SetTTL(v byte) {
	b[ttl] = v
}

// SetTotalLength sets the "total length" field of the IPv4 header.
func (b IPv4) SetTotalLength(totalLength uint16) {
	binary.BigEndian.PutUint16(b[IPv4TotalLenOffset:], totalLength)
}

// SetChecksum sets the checksum field of the IPv4 header.
func (b IPv4) SetChecksum(v uint16) {
	binary.BigEndian.PutUint16(b[checksum:], v)
}

// SetFlagsFragmentOffset sets the "flags" and "fragment offset" fields of the
// IPv4 header.
func (b IPv4) SetFlagsFragmentOffset(flags uint8, offset uint16) {
	v := (uint16(flags) << 13) | (offset >> 3)
	binary.BigEndian.PutUint16(b[flagsFO:], v)
}

// SetID sets the identification field.
func (b IPv4) SetID(v uint16) {
	binary.BigEndian.PutUint16(b[id:], v)
}

// SetSourceAddress sets the "source address" field of the IPv4 header.
func (b IPv4) SetSourceAddress(addr tcpip.Address) {
	copy(b[srcAddr:srcAddr+IPv4AddressSize], addr)
}

// SetDestinationAddress sets the "destination address" field of the IPv4
// header.
func (b IPv4) SetDestinationAddress(addr tcpip.Address) {
	copy(b[dstAddr:dstAddr+IPv4AddressSize], addr)
}

// CalculateChecksum calculates the checksum of the IPv4 header.
func (b IPv4) CalculateChecksum() uint16 {
	return Checksum(b[:b.HeaderLength()], 0)
}

// Encode encodes all the fields of the IPv4 header.
func (b IPv4) Encode(i *IPv4Fields) {
	// The size of the options defines the size of the whole header and thus the
	// IHL field. Options are rare and this is a heavily used function so it is
	// worth a bit of optimisation here to keep the copy out of the fast path.
	hdrLen := IPv4MinimumSize
	if len(i.Options) != 0 {
		// SizeWithPadding is always >= len(i.Options).
		aLen := i.Options.SizeWithPadding()
		hdrLen += aLen
		if hdrLen > len(b) {
			panic(fmt.Sprintf("encode received %d bytes, wanted >= %d", len(b), hdrLen))
		}
		opts := b[options:]
		// This avoids bounds checks on the next line(s) which would happen even
		// if there's no work to do.
		if n := copy(opts, i.Options); n != aLen {
			padding := opts[n:][:aLen-n]
			for i := range padding {
				padding[i] = 0
			}
		}
	}
	b.SetHeaderLength(uint8(hdrLen))
	b[tos] = i.TOS
	b.SetTotalLength(i.TotalLength)
	binary.BigEndian.PutUint16(b[id:], i.ID)
	b.SetFlagsFragmentOffset(i.Flags, i.FragmentOffset)
	b[ttl] = i.TTL
	b[protocol] = i.Protocol
	b.SetChecksum(i.Checksum)
	copy(b[srcAddr:srcAddr+IPv4AddressSize], i.SrcAddr)
	copy(b[dstAddr:dstAddr+IPv4AddressSize], i.DstAddr)
}

// EncodePartial updates the total length and checksum fields of IPv4 header,
// taking in the partial checksum, which is the checksum of the header without
// the total length and checksum fields. It is useful in cases when similar
// packets are produced.
func (b IPv4) EncodePartial(partialChecksum, totalLength uint16) {
	b.SetTotalLength(totalLength)
	checksum := Checksum(b[IPv4TotalLenOffset:IPv4TotalLenOffset+2], partialChecksum)
	b.SetChecksum(^checksum)
}

// IsValid performs basic validation on the packet.
func (b IPv4) IsValid(pktSize int) bool {
	if len(b) < IPv4MinimumSize {
		return false
	}

	hlen := int(b.HeaderLength())
	tlen := int(b.TotalLength())
	if hlen < IPv4MinimumSize || hlen > tlen || tlen > pktSize {
		return false
	}

	if IPVersion(b) != IPv4Version {
		return false
	}

	return true
}

// IsV4MulticastAddress determines if the provided address is an IPv4 multicast
// address (range 224.0.0.0 to 239.255.255.255). The four most significant bits
// will be 1110 = 0xe0.
func IsV4MulticastAddress(addr tcpip.Address) bool {
	if len(addr) != IPv4AddressSize {
		return false
	}
	return (addr[0] & 0xf0) == 0xe0
}

// IsV4LoopbackAddress determines if the provided address is an IPv4 loopback
// address (belongs to 127.0.0.0/8 subnet). See RFC 1122 section 3.2.1.3.
func IsV4LoopbackAddress(addr tcpip.Address) bool {
	if len(addr) != IPv4AddressSize {
		return false
	}
	return addr[0] == 0x7f
}

// ========================= Options ==========================

// An IPv4OptionType can hold the valuse for the Type in an IPv4 option.
type IPv4OptionType byte

// These constants are needed to identify individual options in the option list.
// While RFC 791 (page 31) says "Every internet module must be able to act on
// every option." This has not generally been adhered to and some options have
// very low rates of support. We do not support options other than those shown
// below.

const (
	// IPv4OptionListEndType is the option type for the End Of Option List
	// option. Anything following is ignored.
	IPv4OptionListEndType IPv4OptionType = 0

	// IPv4OptionNOPType is the No-Operation option. May appear between other
	// options and may appear multiple times.
	IPv4OptionNOPType IPv4OptionType = 1

	// IPv4OptionRecordRouteType is used by each router on the path of the packet
	// to record its path. It is carried over to an Echo Reply.
	IPv4OptionRecordRouteType IPv4OptionType = 7

	// IPv4OptionTimestampType is the option type for the Timestamp option.
	IPv4OptionTimestampType IPv4OptionType = 68

	// ipv4OptionTypeOffset is the offset in an option of its type field.
	ipv4OptionTypeOffset = 0

	// IPv4OptionLengthOffset is the offset in an option of its length field.
	IPv4OptionLengthOffset = 1
)

// Potential errors when parsing generic IP options.
var (
	ErrIPv4OptZeroLength   = errors.New("zero length IP option")
	ErrIPv4OptDuplicate    = errors.New("duplicate IP option")
	ErrIPv4OptInvalid      = errors.New("invalid IP option")
	ErrIPv4OptMalformed    = errors.New("malformed IP option")
	ErrIPv4OptionTruncated = errors.New("truncated IP option")
	ErrIPv4OptionAddress   = errors.New("bad IP option address")
)

// IPv4Option is an interface representing various option types.
type IPv4Option interface {
	// Type returns the type identifier of the option.
	Type() IPv4OptionType

	// Size returns the size of the option in bytes.
	Size() uint8

	// Contents returns a slice holding the contents of the option.
	Contents() []byte
}

var _ IPv4Option = (*IPv4OptionGeneric)(nil)

// IPv4OptionGeneric is an IPv4 Option of unknown type.
type IPv4OptionGeneric []byte

// Type implements IPv4Option.
func (o *IPv4OptionGeneric) Type() IPv4OptionType {
	return IPv4OptionType((*o)[ipv4OptionTypeOffset])
}

// Size implements IPv4Option.
func (o *IPv4OptionGeneric) Size() uint8 { return uint8(len(*o)) }

// Contents implements IPv4Option.
func (o *IPv4OptionGeneric) Contents() []byte { return []byte(*o) }

// IPv4OptionIterator is an iterator pointing to a specific IP option
// at any point of time. It also holds information as to a new options buffer
// that we are building up to hand back to the caller.
type IPv4OptionIterator struct {
	options IPv4Options
	// ErrCursor is where we are while parsing options. It is exported as any
	// resulting ICMP packet is supposed to have a pointer to the byte within
	// the IP packet where the error was detected.
	ErrCursor     uint8
	nextErrCursor uint8
	newOptions    [IPv4MaximumOptionsSize]byte
	writePoint    int
}

// MakeIterator sets up and returns an iterator of options. It also sets up the
// building of a new option set.
func (o IPv4Options) MakeIterator() IPv4OptionIterator {
	return IPv4OptionIterator{
		options:       o,
		nextErrCursor: IPv4MinimumSize,
	}
}

// RemainingBuffer returns the remaining (unused) part of the new option buffer,
// into which a new option may be written.
func (i *IPv4OptionIterator) RemainingBuffer() IPv4Options {
	return IPv4Options(i.newOptions[i.writePoint:])
}

// ConsumeBuffer marks a portion of the new buffer as used.
func (i *IPv4OptionIterator) ConsumeBuffer(size int) {
	i.writePoint += size
}

// PushNOPOrEnd puts one of the single byte options onto the new options.
// Only values 0 or 1 (ListEnd or NOP) are valid input.
func (i *IPv4OptionIterator) PushNOPOrEnd(val IPv4OptionType) {
	if val > IPv4OptionNOPType {
		panic(fmt.Sprintf("invalid option type %d pushed onto option build buffer", val))
	}
	i.newOptions[i.writePoint] = byte(val)
	i.writePoint++
}

// Finalize returns the completed replacement options buffer padded
// as needed.
func (i *IPv4OptionIterator) Finalize() IPv4Options {
	// RFC 791 page 31 says:
	//     The options might not end on a 32-bit boundary.  The internet header
	//     must be filled out with octets of zeros.  The first of these would
	//     be interpreted as the end-of-options option, and the remainder as
	//     internet header padding.
	// Since the buffer is already zero filled we just need to step the write
	// pointer up to the next multiple of 4.
	options := IPv4Options(i.newOptions[:(i.writePoint+0x3) & ^0x3])
	// Poison the write pointer.
	i.writePoint = len(i.newOptions)
	return options
}

// Next returns the next IP option in the buffer/list of IP options.
// It returns
// - A slice of bytes holding the next option or nil if there is error.
// - A boolean which is true if parsing of all the options is complete.
// - An error which is non-nil if an error condition was encountered.
func (i *IPv4OptionIterator) Next() (IPv4Option, bool, error) {
	// The opts slice gets shorter as we process the options. When we have no
	// bytes left we are done.
	if len(i.options) == 0 {
		return nil, true, nil
	}

	i.ErrCursor = i.nextErrCursor

	optType := IPv4OptionType(i.options[ipv4OptionTypeOffset])

	if optType == IPv4OptionNOPType || optType == IPv4OptionListEndType {
		optionBody := i.options[:1]
		i.options = i.options[1:]
		i.nextErrCursor = i.ErrCursor + 1
		retval := IPv4OptionGeneric(optionBody)
		return &retval, false, nil
	}

	// There are no more single byte options defined.  All the rest have a length
	// field so we need to sanity check it.
	if len(i.options) == 1 {
		return nil, true, ErrIPv4OptMalformed
	}

	optLen := i.options[IPv4OptionLengthOffset]

	if optLen == 0 {
		i.ErrCursor++
		return nil, true, ErrIPv4OptZeroLength
	}

	if optLen == 1 {
		i.ErrCursor++
		return nil, true, ErrIPv4OptMalformed
	}

	if optLen > uint8(len(i.options)) {
		i.ErrCursor++
		return nil, true, ErrIPv4OptionTruncated
	}

	optionBody := i.options[:optLen]
	i.nextErrCursor = i.ErrCursor + optLen
	i.options = i.options[optLen:]

	// Check the length of some option types that we know.
	switch optType {
	case IPv4OptionTimestampType:
		if optLen < IPv4OptionTimestampHdrLength {
			i.ErrCursor++
			return nil, true, ErrIPv4OptMalformed
		}
		retval := IPv4OptionTimestamp(optionBody)
		return &retval, false, nil

	case IPv4OptionRecordRouteType:
		if optLen < IPv4OptionRecordRouteHdrLength {
			i.ErrCursor++
			return nil, true, ErrIPv4OptMalformed
		}
		retval := IPv4OptionRecordRoute(optionBody)
		return &retval, false, nil
	}
	retval := IPv4OptionGeneric(optionBody)
	return &retval, false, nil
}

//
// IP Timestamp option - RFC 791 page 22.
// +--------+--------+--------+--------+
// |01000100| length | pointer|oflw|flg|
// +--------+--------+--------+--------+
// |         internet address          |
// +--------+--------+--------+--------+
// |             timestamp             |
// +--------+--------+--------+--------+
// |                ...                |
//
// Type = 68
//
// The Option Length is the number of octets in the option counting
// the type, length, pointer, and overflow/flag octets (maximum
// length 40).
//
// The Pointer is the number of octets from the beginning of this
// option to the end of timestamps plus one (i.e., it points to the
// octet beginning the space for next timestamp).  The smallest
// legal value is 5.  The timestamp area is full when the pointer
// is greater than the length.
//
// The Overflow (oflw) [4 bits] is the number of IP modules that
// cannot register timestamps due to lack of space.
//
// The Flag (flg) [4 bits] values are
//
//   0 -- time stamps only, stored in consecutive 32-bit words,
//
//   1 -- each timestamp is preceded with internet address of the
//        registering entity,
//
//   3 -- the internet address fields are prespecified.  An IP
//        module only registers its timestamp if it matches its own
//        address with the next specified internet address.
//
// Timestamps are defined in RFC 791 page 22 as milliseconds since midnight UTC.
//
//        The Timestamp is a right-justified, 32-bit timestamp in
//        milliseconds since midnight UT.  If the time is not available in
//        milliseconds or cannot be provided with respect to midnight UT
//        then any time may be inserted as a timestamp provided the high
//        order bit of the timestamp field is set to one to indicate the
//        use of a non-standard value.

// IPv4OptTSFlags sefines the values expected in the Timestamp
// option Flags field.
type IPv4OptTSFlags uint8

//
// Timestamp option specific related constants.
const (
	// IPv4OptionTimestampHdrLength is the length of the timestamp option header.
	IPv4OptionTimestampHdrLength = 4

	// IPv4OptionTimestampSize is the size of an IP timestamp.
	IPv4OptionTimestampSize = 4

	// IPv4OptionTimestampWithAddrSize is the size of an IP timestamp + Address.
	IPv4OptionTimestampWithAddrSize = IPv4AddressSize + IPv4OptionTimestampSize

	// IPv4OptionTimestampMaxSize is limited by space for options
	IPv4OptionTimestampMaxSize = IPv4MaximumOptionsSize

	// IPv4OptionTimestampOnlyFlag is a flag indicating that only timestamp
	// is present.
	IPv4OptionTimestampOnlyFlag IPv4OptTSFlags = 0

	// IPv4OptionTimestampWithIPFlag is a flag indicating that both timestamps and
	// IP are present.
	IPv4OptionTimestampWithIPFlag IPv4OptTSFlags = 1

	// IPv4OptionTimestampWithPredefinedIPFlag is a flag indicating that
	// predefined IP is present.
	IPv4OptionTimestampWithPredefinedIPFlag IPv4OptTSFlags = 3
)

// ipv4TimestampTime provides the current time as specified in RFC 791.
func ipv4TimestampTime(clock tcpip.Clock) uint32 {
	const millisecondsPerDay = 24 * 3600 * 1000
	const nanoPerMilli = 1000000
	return uint32((clock.NowNanoseconds() / nanoPerMilli) % millisecondsPerDay)
}

// IP Timestamp option fields.
const (
	// IPv4OptTSPointerOffset is the offset of the Timestamp pointer field.
	IPv4OptTSPointerOffset = 2

	// IPv4OptTSPointerOffset is the offset of the combined Flag and Overflow
	// fields, (each being 4 bits).
	IPv4OptTSOFLWAndFLGOffset = 3
	// These constants define the sub byte fields of the Flag and OverFlow field.
	ipv4OptionTimestampOverflowshift      = 4
	ipv4OptionTimestampFlagsMask     byte = 0x0f
)

var _ IPv4Option = (*IPv4OptionTimestamp)(nil)

// IPv4OptionTimestamp is a Timestamp option from RFC 791.
type IPv4OptionTimestamp []byte

// Type implements IPv4Option.Type().
func (ts *IPv4OptionTimestamp) Type() IPv4OptionType { return IPv4OptionTimestampType }

// Size implements IPv4Option.
func (ts *IPv4OptionTimestamp) Size() uint8 { return uint8(len(*ts)) }

// Contents implements IPv4Option.
func (ts *IPv4OptionTimestamp) Contents() []byte { return []byte(*ts) }

// Pointer returns the pointer field in the IP Timestamp option.
func (ts *IPv4OptionTimestamp) Pointer() uint8 {
	return (*ts)[IPv4OptTSPointerOffset]
}

// Flags returns the flags field in the IP Timestamp option.
func (ts *IPv4OptionTimestamp) Flags() IPv4OptTSFlags {
	return IPv4OptTSFlags((*ts)[IPv4OptTSOFLWAndFLGOffset] & ipv4OptionTimestampFlagsMask)
}

// Overflow returns the Overflow field in the IP Timestamp option.
func (ts *IPv4OptionTimestamp) Overflow() uint8 {
	return (*ts)[IPv4OptTSOFLWAndFLGOffset] >> ipv4OptionTimestampOverflowshift
}

// IncOverflow increments the Overflow field in the IP Timestamp option. It
// returns the incremented value. If the return value is 0 then the field
// overflowed.
func (ts *IPv4OptionTimestamp) IncOverflow() uint8 {
	(*ts)[IPv4OptTSOFLWAndFLGOffset] += 1 << ipv4OptionTimestampOverflowshift
	return ts.Overflow()
}

// UpdateTimestamp updates the fields of the next free timestamp slot.
func (ts *IPv4OptionTimestamp) UpdateTimestamp(addr tcpip.Address, clock tcpip.Clock) {
	slot := (*ts)[ts.Pointer()-1:]

	switch ts.Flags() {
	case IPv4OptionTimestampOnlyFlag:
		binary.BigEndian.PutUint32(slot, ipv4TimestampTime(clock))
		(*ts)[IPv4OptTSPointerOffset] += IPv4OptionTimestampSize
	case IPv4OptionTimestampWithIPFlag:
		if n := copy(slot, addr); n != IPv4AddressSize {
			panic(fmt.Sprintf("copied %d bytes, expected %d bytes", n, IPv4AddressSize))
		}
		binary.BigEndian.PutUint32(slot[IPv4AddressSize:], ipv4TimestampTime(clock))
		(*ts)[IPv4OptTSPointerOffset] += IPv4OptionTimestampWithAddrSize
	case IPv4OptionTimestampWithPredefinedIPFlag:
		if tcpip.Address(slot[:IPv4AddressSize]) == addr {
			binary.BigEndian.PutUint32(slot[IPv4AddressSize:], ipv4TimestampTime(clock))
			(*ts)[IPv4OptTSPointerOffset] += IPv4OptionTimestampWithAddrSize
		}
	}
}

// RecordRoute option specific related constants.
//
// from RFC 791 page 20:
//   Record Route
//
//         +--------+--------+--------+---------//--------+
//         |00000111| length | pointer|     route data    |
//         +--------+--------+--------+---------//--------+
//           Type=7
//
//         The record route option provides a means to record the route of
//         an internet datagram.
//
//         The option begins with the option type code.  The second octet
//         is the option length which includes the option type code and the
//         length octet, the pointer octet, and length-3 octets of route
//         data.  The third octet is the pointer into the route data
//         indicating the octet which begins the next area to store a route
//         address.  The pointer is relative to this option, and the
//         smallest legal value for the pointer is 4.
const (
	// IPv4OptionRecordRouteHdrLength is the length of the Record Route option
	// header.
	IPv4OptionRecordRouteHdrLength = 3

	// IPv4OptRRPointerOffset is the offset to the pointer field in an RR
	// option, which points to the next free slot in the list of addresses.
	IPv4OptRRPointerOffset = 2
)

var _ IPv4Option = (*IPv4OptionRecordRoute)(nil)

// IPv4OptionRecordRoute is an IPv4 RecordRoute option defined by RFC 791.
type IPv4OptionRecordRoute []byte

// Pointer returns the pointer field in the IP RecordRoute option.
func (rr *IPv4OptionRecordRoute) Pointer() uint8 {
	return (*rr)[IPv4OptRRPointerOffset]
}

// StoreAddress stores the given IPv4 address into the next free slot.
func (rr *IPv4OptionRecordRoute) StoreAddress(addr tcpip.Address) {
	start := rr.Pointer() - 1 // A one based number.
	// start and room checked by caller.
	if n := copy((*rr)[start:], addr); n != IPv4AddressSize {
		panic(fmt.Sprintf("copied %d bytes, expected %d bytes", n, IPv4AddressSize))
	}
	(*rr)[IPv4OptRRPointerOffset] += IPv4AddressSize
}

// Type implements IPv4Option.
func (rr *IPv4OptionRecordRoute) Type() IPv4OptionType { return IPv4OptionRecordRouteType }

// Size implements IPv4Option.
func (rr *IPv4OptionRecordRoute) Size() uint8 { return uint8(len(*rr)) }

// Contents implements IPv4Option.
func (rr *IPv4OptionRecordRoute) Contents() []byte { return []byte(*rr) }

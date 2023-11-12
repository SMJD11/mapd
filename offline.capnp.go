// Code generated by capnpc-go. DO NOT EDIT.

package main

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	math "math"
)

type Way capnp.Struct

// Way_TypeID is the unique identifier for the type Way.
const Way_TypeID = 0xa4b9c59286b69600

func NewWay(s *capnp.Segment) (Way, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 56, PointerCount: 4})
	return Way(st), err
}

func NewRootWay(s *capnp.Segment) (Way, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 56, PointerCount: 4})
	return Way(st), err
}

func ReadRootWay(msg *capnp.Message) (Way, error) {
	root, err := msg.Root()
	return Way(root.Struct()), err
}

func (s Way) String() string {
	str, _ := text.Marshal(0xa4b9c59286b69600, capnp.Struct(s))
	return str
}

func (s Way) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Way) DecodeFromPtr(p capnp.Ptr) Way {
	return Way(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Way) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Way) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Way) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Way) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Way) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Way) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Way) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Way) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Way) Ref() (string, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.Text(), err
}

func (s Way) HasRef() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Way) RefBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return p.TextBytes(), err
}

func (s Way) SetRef(v string) error {
	return capnp.Struct(s).SetText(1, v)
}

func (s Way) MaxSpeed() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(0))
}

func (s Way) SetMaxSpeed(v float64) {
	capnp.Struct(s).SetUint64(0, math.Float64bits(v))
}

func (s Way) MinLat() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(8))
}

func (s Way) SetMinLat(v float64) {
	capnp.Struct(s).SetUint64(8, math.Float64bits(v))
}

func (s Way) MinLon() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(16))
}

func (s Way) SetMinLon(v float64) {
	capnp.Struct(s).SetUint64(16, math.Float64bits(v))
}

func (s Way) MaxLat() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(24))
}

func (s Way) SetMaxLat(v float64) {
	capnp.Struct(s).SetUint64(24, math.Float64bits(v))
}

func (s Way) MaxLon() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(32))
}

func (s Way) SetMaxLon(v float64) {
	capnp.Struct(s).SetUint64(32, math.Float64bits(v))
}

func (s Way) Nodes() (Coordinates_List, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return Coordinates_List(p.List()), err
}

func (s Way) HasNodes() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Way) SetNodes(v Coordinates_List) error {
	return capnp.Struct(s).SetPtr(2, v.ToPtr())
}

// NewNodes sets the nodes field to a newly
// allocated Coordinates_List, preferring placement in s's segment.
func (s Way) NewNodes(n int32) (Coordinates_List, error) {
	l, err := NewCoordinates_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Coordinates_List{}, err
	}
	err = capnp.Struct(s).SetPtr(2, l.ToPtr())
	return l, err
}
func (s Way) Lanes() uint8 {
	return capnp.Struct(s).Uint8(40)
}

func (s Way) SetLanes(v uint8) {
	capnp.Struct(s).SetUint8(40, v)
}

func (s Way) AdvisorySpeed() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(48))
}

func (s Way) SetAdvisorySpeed(v float64) {
	capnp.Struct(s).SetUint64(48, math.Float64bits(v))
}

func (s Way) Hazard() (string, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return p.Text(), err
}

func (s Way) HasHazard() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Way) HazardBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return p.TextBytes(), err
}

func (s Way) SetHazard(v string) error {
	return capnp.Struct(s).SetText(3, v)
}

// Way_List is a list of Way.
type Way_List = capnp.StructList[Way]

// NewWay creates a new list of Way.
func NewWay_List(s *capnp.Segment, sz int32) (Way_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 56, PointerCount: 4}, sz)
	return capnp.StructList[Way](l), err
}

// Way_Future is a wrapper for a Way promised by a client call.
type Way_Future struct{ *capnp.Future }

func (f Way_Future) Struct() (Way, error) {
	p, err := f.Future.Ptr()
	return Way(p.Struct()), err
}

type Coordinates capnp.Struct

// Coordinates_TypeID is the unique identifier for the type Coordinates.
const Coordinates_TypeID = 0x922b57c60c6a46d1

func NewCoordinates(s *capnp.Segment) (Coordinates, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Coordinates(st), err
}

func NewRootCoordinates(s *capnp.Segment) (Coordinates, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return Coordinates(st), err
}

func ReadRootCoordinates(msg *capnp.Message) (Coordinates, error) {
	root, err := msg.Root()
	return Coordinates(root.Struct()), err
}

func (s Coordinates) String() string {
	str, _ := text.Marshal(0x922b57c60c6a46d1, capnp.Struct(s))
	return str
}

func (s Coordinates) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Coordinates) DecodeFromPtr(p capnp.Ptr) Coordinates {
	return Coordinates(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Coordinates) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Coordinates) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Coordinates) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Coordinates) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Coordinates) Latitude() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(0))
}

func (s Coordinates) SetLatitude(v float64) {
	capnp.Struct(s).SetUint64(0, math.Float64bits(v))
}

func (s Coordinates) Longitude() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(8))
}

func (s Coordinates) SetLongitude(v float64) {
	capnp.Struct(s).SetUint64(8, math.Float64bits(v))
}

// Coordinates_List is a list of Coordinates.
type Coordinates_List = capnp.StructList[Coordinates]

// NewCoordinates creates a new list of Coordinates.
func NewCoordinates_List(s *capnp.Segment, sz int32) (Coordinates_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return capnp.StructList[Coordinates](l), err
}

// Coordinates_Future is a wrapper for a Coordinates promised by a client call.
type Coordinates_Future struct{ *capnp.Future }

func (f Coordinates_Future) Struct() (Coordinates, error) {
	p, err := f.Future.Ptr()
	return Coordinates(p.Struct()), err
}

type Offline capnp.Struct

// Offline_TypeID is the unique identifier for the type Offline.
const Offline_TypeID = 0xcb5ff253617678e0

func NewOffline(s *capnp.Segment) (Offline, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1})
	return Offline(st), err
}

func NewRootOffline(s *capnp.Segment) (Offline, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1})
	return Offline(st), err
}

func ReadRootOffline(msg *capnp.Message) (Offline, error) {
	root, err := msg.Root()
	return Offline(root.Struct()), err
}

func (s Offline) String() string {
	str, _ := text.Marshal(0xcb5ff253617678e0, capnp.Struct(s))
	return str
}

func (s Offline) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Offline) DecodeFromPtr(p capnp.Ptr) Offline {
	return Offline(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Offline) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Offline) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Offline) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Offline) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Offline) MinLat() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(0))
}

func (s Offline) SetMinLat(v float64) {
	capnp.Struct(s).SetUint64(0, math.Float64bits(v))
}

func (s Offline) MinLon() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(8))
}

func (s Offline) SetMinLon(v float64) {
	capnp.Struct(s).SetUint64(8, math.Float64bits(v))
}

func (s Offline) MaxLat() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(16))
}

func (s Offline) SetMaxLat(v float64) {
	capnp.Struct(s).SetUint64(16, math.Float64bits(v))
}

func (s Offline) MaxLon() float64 {
	return math.Float64frombits(capnp.Struct(s).Uint64(24))
}

func (s Offline) SetMaxLon(v float64) {
	capnp.Struct(s).SetUint64(24, math.Float64bits(v))
}

func (s Offline) Ways() (Way_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Way_List(p.List()), err
}

func (s Offline) HasWays() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Offline) SetWays(v Way_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewWays sets the ways field to a newly
// allocated Way_List, preferring placement in s's segment.
func (s Offline) NewWays(n int32) (Way_List, error) {
	l, err := NewWay_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Way_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// Offline_List is a list of Offline.
type Offline_List = capnp.StructList[Offline]

// NewOffline creates a new list of Offline.
func NewOffline_List(s *capnp.Segment, sz int32) (Offline_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1}, sz)
	return capnp.StructList[Offline](l), err
}

// Offline_Future is a wrapper for a Offline promised by a client call.
type Offline_Future struct{ *capnp.Future }

func (f Offline_Future) Struct() (Offline, error) {
	p, err := f.Future.Ptr()
	return Offline(p.Struct()), err
}

const schema_da3a0d9284ca402f = "x\xda\xa4\xd3\xddj\x13[\x14\x07\xf0\xff\x7f\xefI\xd2" +
	"\xef\x9e0syJ\xce9\x9cB[[Mm\x85\x1a" +
	"*F\xfc@\x8a`\xb7#\x14\xa1b7fb#\xe9" +
	"LIb\x9b\x8a\x05\x85*-X\xd0\\\xf8\x04}\x02" +
	"/\x04\x1f@/\xf4J/\xbd\xf2\x19|\x00G\xd6\x94" +
	"|\xd4[\xef\xf6\xfe\xed5km\xd6\xda\x93\xdfe\xd1" +
	"\x99\x1d\x8e\x14\x94\xf9;\x95\x8e\xbf\\{8\xf4q\xe5" +
	"T\x0bf\x84*>S\xfc\xb4\xd7\x1a.|\x83\x93\x01" +
	"\xdc1~u')\xabqn\x83?\xdf\xbc{\xd1\xfa" +
	"\xf0\xfe\xc8\x8c0\xd3\x8dL%\xa1\x07l\xb9\xaf%t" +
	"\xee\x90\x91\x02\xe3\xef\xcd-\xeb\xff\xb8\xf7Y\xf2:=" +
	"\xd1I\xba1\xe7\xad;.\xdf\xcd\xfd\xeb\xe4\x88\x998" +
	"*\x97\xab\x9508\xad\xee\xdb\xcdp\xb3p9\x8aj" +
	"\xa5Jh\x1b\x01\xeb\xcb\xa4\xe9\xd3\x0e\xe0\x10\xc8N." +
	"\x01fB\xd3\xcc+fI\x8f\x82\xb3\xb7\x00\x93\xd74" +
	"\x8b\x8aq\xd56*\x8dG\xa5\x00\x00\x07\xa18\x08\xc6" +
	"\xd5(| \x08\x06\x1dk\x97\xe4q\xc9\x15\xcb\x1d)" +
	"5\xdd.\xe5\x8es\x0a\xf0\xff\xa1\xa6?\xcdn5w" +
	"\x92\xff\x01\xfe\xff\xe2y*RyT\x80;\xc3%\xc0" +
	"\x9f\x16^\x90pM\x8f\x1ap\xcf\xb1\x00\xf8y\xf1E" +
	"qGyt\x00\xf7|\xe2\xf3\xe2E\xf1\x94\xf6\x98\x02" +
	"\xdc\x0b\x89/\x88_\x11O;\x1e\xd3\x80{)\xf1E" +
	"\xf1\xeb\xe2\x19\xe5%\xbd\xbc\xca\xb3\x80_\x14_\x15\xef" +
	"\x9b\xf0\xd8\x07\xb8w\x12\xbf-\xbe&\xde\x9f\xf6\xd8\x0f" +
	"\xb8wY\x03\xfcU\xf1u\xf1\x01\xedq\x00p\x83$" +
	"\xff\x9ax\x95\x8a\xa3\xa1\xdd\x088\x04\xc5!0S\x0b" +
	"\xca\xedu\xbca\x9b\xfef\x10\x94z\x1a|q\xa3\x12" +
	"\xde\xb0\x8d\x13\xdb(\xecnm\xf3\xc4\xa9m\xf6\x9c\xe6" +
	"\xc2\xa8\x14\xd49\x02.k\xf2\xaf\xee\x83\x04\x05sU" +
	"\x1b\x06u\xa6\xa1\x98\x06c[\xda\xaa\xd4\xa3\xda\x0er" +
	"\xc9\x1d:9\xd7\xedc[+u\xee\xf8\xdbpo\x96" +
	"s\xc9^\x06\xecu\xde\xd2n\x010MM\xb3\xd7\xf3" +
	"\x96\x9e\x09>\xd14\xfb\x8aYu<\xdb\xecs\xc1\xa7" +
	"\x9a\xe6\xa5\x0cV'\x83\xcd\x1e\x08\xeei\x9aW\x8at" +
	"\x92\xa1f\x0f\xa7\x00\xb3\xafi\x8e\xd4\x1f5et\xdb" +
	"\xeet{\xd2\xfe\xf1\x8e;\xf2+\x00\x00\xff\xff\xb97" +
	"\xac\xee"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_da3a0d9284ca402f,
		Nodes: []uint64{
			0x922b57c60c6a46d1,
			0xa4b9c59286b69600,
			0xcb5ff253617678e0,
		},
		Compressed: true,
	})
}

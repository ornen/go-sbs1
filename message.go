package sbs1

import (
	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
	"time"
)

type MessageType uint8

const (
	MessageTypeSelectionChange MessageType = iota
	MessageTypeNewId
	MessageTypeNewAircraft
	MessageTypeStatusAircraft
	MessageTypeClick
	MessageTypeTransmission
)

type TransmissionType uint8

const (
	TransmissionTypeESIdentAndCategory TransmissionType = iota
	TransmissionTypeESSurfacePos
	TransmissionTypeESAirbornePos
	TransmissionTypeESAirborneVel
	TransmissionTypeSurveillanceAlt
	TransmissionTypeSurveillanceId
	TransmissionTypeAirToAir
	TransmissionTypeAllCallReply
)

type Message struct {
	MessageType      MessageType
	TransmissionType TransmissionType
	SessionId        string
	AircraftId       string
	HexId            string
	FlightId         string
	Generated        time.Time
	Logged           time.Time
	Callsign         string
	Altitude         int32
	GroundSpeed      int32
	Track            s1.Angle
	Coordinates      s2.LatLng
	VerticalRate     int16
	Squawk           string
}

package sbs1

import (
	"time"

	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

type MessageType string

const (
	MessageTypeSelectionChange MessageType = "SEL"
	MessageTypeNewId           MessageType = "ID"
	MessageTypeNewAircraft     MessageType = "AIR"
	MessageTypeStatusAircraft  MessageType = "STA"
	MessageTypeClick           MessageType = "CLK"
	MessageTypeTransmission    MessageType = "MSG"
)

type TransmissionType uint8

const (
	TransmissionTypeESIdentAndCategory TransmissionType = 1
	TransmissionTypeESSurfacePos       TransmissionType = 2
	TransmissionTypeESAirbornePos      TransmissionType = 3
	TransmissionTypeESAirborneVel      TransmissionType = 4
	TransmissionTypeSurveillanceAlt    TransmissionType = 5
	TransmissionTypeSurveillanceId     TransmissionType = 6
	TransmissionTypeAirToAir           TransmissionType = 7
	TransmissionTypeAllCallReply       TransmissionType = 8
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

package freeGeoIP

import (
	"encoding/json"
	"errors"
	"net"
	"time"
)

// IP is a wrapper for net.IP
type IP net.IP

func ParseIP(ip string) IP {
	return IP(net.ParseIP(ip))
}

func (ip IP) Net() net.IP {
	return net.IP(ip)
}

func (ip IP) String() string {
	if len(ip) == 0 {
		return ""
	}
	return net.IP(ip).String()
}

func (ip *IP) UnmarshalJSON(b []byte) error {
	var _ip string
	if err := json.Unmarshal(b, &_ip); err != nil {
		return wrapError("decoder", err)
	}
	*ip = ParseIP(_ip)
	return nil
}

func (ip *IP) MarshalJSON() ([]byte, error) {
	return json.Marshal(ip.String())
}

// Location is a wrapper for *time.Location
type Location time.Location

func LocationF(zone *time.Location) *Location {
	if zone == nil {
		return nil
	}
	val := Location(*zone)
	return &val
}

func (tz *Location) Time() *time.Location {
	if tz == nil {
		return nil
	}
	val := time.Location(*tz)
	return &val
}

func (tz *Location) String() string {
	if tz == nil {
		return ""
	}
	val := time.Location(*tz)
	return (&val).String()
}

func (tz *Location) UnmarshalJSON(b []byte) error {
	var zoneStr string
	if err := json.Unmarshal(b, &zoneStr); err != nil {
		return wrapError("decoder", err)
	}
	zone, err := time.LoadLocation(zoneStr)
	if err != nil {
		return wrapError("decoder", err)
	}
	if zone == nil {
		return wrapError("decoder", errors.New("invalid timezone string"))
	}
	*tz = Location(*zone)
	return nil
}

func (tz *Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(tz.String())
}

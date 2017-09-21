package models

import (
	"net"
	"time"
)

var hexDigit = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func Hexaddr(addr net.IP) string {
	b := addr.To4()
	s := make([]byte, len(b)*2)
	for i, tn := range b {
		s[i*2], s[i*2+1] = hexDigit[tn>>4], hexDigit[tn&0xf]
	}
	return string(s)
}

// swagger:model
type Lease struct {
	Validation
	Access
	MetaData
	// Addr is the IP address that the lease handed out.
	//
	// required: true
	// swagger:strfmt ipv4
	Addr net.IP
	// Token is the unique token for this lease based on the
	// Strategy this lease used.
	//
	// required: true
	Token string
	// ExpireTime is the time at which the lease expires and is no
	// longer valid The DHCP renewal time will be half this, and the
	// DHCP rebind time will be three quarters of this.
	//
	// required: true
	// swagger:strfmt date-time
	ExpireTime time.Time
	// Strategy is the leasing strategy that will be used determine what to use from
	// the DHCP packet to handle lease management.
	//
	// required: true
	Strategy string
	// State is the current state of the lease.  This field is for informational
	// purposes only.
	//
	// read only: true
	// required: true
	State string
}

func (l *Lease) Prefix() string {
	return "leases"
}

func (l *Lease) Key() string {
	return Hexaddr(l.Addr)
}

func (l *Lease) AuthKey() string {
	return l.Key()
}

type Leases []*Lease

func (s Leases) Elem() Model {
	return &Lease{}
}

func (s Leases) Items() []Model {
	res := make([]Model, len(s))
	for i, m := range s {
		res[i] = m
	}
	return res
}
func (s Leases) Fill(m []Model) {
	q := make([]*Lease, len(m))
	for i, obj := range m {
		q[i] = obj.(*Lease)
	}
	s = q[:]
}

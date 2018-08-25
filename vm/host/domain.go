package host

import (
	"strings"
)

// const ...
const (
	DHCPTable      = "dhcp_table"
	DHCPRTable     = "dhcp_revert_table"
	DHCPOwnerTable = "dhcp_owner_table"
)

// DHCP ...
type DHCP struct {
	h *Host
}

// NewDHCP ...
func NewDHCP(h *Host) DHCP {
	return DHCP{
		h: h,
	}
}

// ContractID .
func (d *DHCP) ContractID(url string) string {
	cid, _ := d.h.MapGet(DHCPTable, url)
	if s, ok := cid.(string); ok {
		return s
	}
	return ""
}

// URLOwner .
func (d *DHCP) URLOwner(url string) string {
	owner, _ := d.h.MapGet(DHCPOwnerTable, url)
	if s, ok := owner.(string); ok {
		return s
	}
	return ""
}

// URLTransfer .
func (d *DHCP) URLTransfer(url, to string) {
	d.h.MapPut(DHCPOwnerTable, url, to)
}

// URL .
func (d *DHCP) URL(cid string) string {
	domain, _ := d.h.MapGet(DHCPRTable, cid)
	if s, ok := domain.(string); ok {
		return s
	}
	return ""
}

// IsDomain .
func (d *DHCP) IsDomain(s string) bool {
	if strings.HasPrefix(s, "Contract") || strings.HasPrefix(s, "iost.") {
		return false
	}
	return true
}

func (d *DHCP) WriteLink(url, cid, owner string) {
	d.h.MapPut(DHCPTable, url, cid)
	d.h.MapPut(DHCPRTable, cid, url)
	d.h.MapPut(DHCPOwnerTable, url, owner)
}

func (d *DHCP) RemoveLink(url, cid string) {
	d.h.MapDel(DHCPRTable, cid)
	d.h.MapDel(DHCPTable, url)
	d.h.MapDel(DHCPOwnerTable, url)
}
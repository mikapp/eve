// Copyright (c) 2017 Zededa, Inc.
// All rights reserved.

// dnsmasq configlets for overlay and underlay interfaces towards domU

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// XXX TODO move ipset to be ACL dependent
// XXX need configdir for that to have one ipset file for each set?
// XXX would still need to restart when configDir content changes AFAIR
// XXX keep ipset's separately in status? (can delete on last ref)
const dnsmasqOverlayStatic = `
# Automatically generated by zedrouter
except-interface=lo
bind-interfaces
log-queries
log-dhcp
no-hosts
no-ping
bogus-priv
stop-dns-rebind
rebind-localhost-ok
neg-ttl=10
# XXX SHOULD be derived from underlay ACL.
# Needed here for underlay since queries for A RRs might come over IPv6
ipset=/google.com/ipv4.google.com,ipv6.google.com
ipset=/zededa.net/ipv4.zededa.net,ipv6.zededa.net
dhcp-range=::,static,0,infinite
`

// XXX TODO move ipset to be ACL dependent
// XXX need configdir for that to have one ipset file for each set?
const dnsmasqUnderlayStatic = `
# Automatically generated by zedrouter
except-interface=lo
bind-interfaces
log-queries
log-dhcp
no-hosts
no-ping
bogus-priv
stop-dns-rebind
rebind-localhost-ok
neg-ttl=10
# XXX SHOULD be derived from underlay ACL
ipset=/google.com/ipv4.google.com,ipv6.google.com
ipset=/zededa.net/ipv4.zededa.net,ipv6.zededa.net
dhcp-range=172.27.0.0,static,255.255.0.0,infinite
`

// Create the dnsmasq configuration for the the overlay interface
// Would be more polite to return an error then to Fatal
func createDnsmasqOverlayConfiglet(cfgPathname string, olIfname string,
	olAddr1 string, olAddr2 string, olMac string, hostsDir string,
	hostName string) {
	fmt.Printf("createDnsmasqOverlayConfiglen: %s\n", olIfname)
	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("os.Create for ", cfgPathname, err)
	}
	defer file.Close()
	file.WriteString(dnsmasqOverlayStatic)
	file.WriteString(fmt.Sprintf("pid-file=/var/run/dnsmasq.%s.pid\n",
		olIfname))
	file.WriteString(fmt.Sprintf("interface=%s\n", olIfname))
	file.WriteString(fmt.Sprintf("listen-address=%s\n", olAddr1))
	file.WriteString(fmt.Sprintf("dhcp-host=%s,[%s],%s\n",
		olMac, olAddr2, hostName))
	file.WriteString(fmt.Sprintf("hostsdir=%s\n", hostsDir))
}

// Create the dnsmasq configuration for the the underlay interface
// Would be more polite to return an error then to Fatal
func createDnsmasqUnderlayConfiglet(cfgPathname string, ulIfname string,
	ulAddr1 string, ulAddr2 string, ulMac string, hostName string) {
	fmt.Printf("createDnsmasqUnderlayConfiglen: %s\n", ulIfname)
	file, err := os.Create(cfgPathname)
	if err != nil {
		log.Fatal("os.Create for ", cfgPathname, err)
	}
	defer file.Close()
	file.WriteString(dnsmasqUnderlayStatic)
	file.WriteString(fmt.Sprintf("pid-file=/var/run/dnsmasq.%s.pid\n",
		ulIfname))
	file.WriteString(fmt.Sprintf("interface=%s\n", ulIfname))
	file.WriteString(fmt.Sprintf("listen-address=%s\n", ulAddr1))
	file.WriteString(fmt.Sprintf("dhcp-host=%s,id:*,%s,%s\n",
		ulMac, ulAddr2, hostName))
}

func deleteDnsmasqConfiglet(cfgPathname string) {
	fmt.Printf("deleteDnsmasqOverlayConfiglen: %s\n", cfgPathname)
	if err := os.Remove(cfgPathname); err != nil {
		log.Println("Remove ", cfgPathname, err)
	}
}

// Run this:
//    DMDIR=/home/nordmark/dnsmasq-2.75/src
//    ${DMDIR}/dnsmasq --conf-file=/etc/dnsmasq.${OLIFNAME}.conf
// or
//    ${DMDIR}/dnsmasq --conf-file=/etc/dnsmasq.${ULIFNAME}.conf
func startDnsmasq(cfgPathname string) {
	fmt.Printf("startDnsmasq: %s\n", cfgPathname)
	cmd := "nohup"
	args := []string{
		"/home/nordmark/dnsmasq-2.75/src/dnsmasq",
		"-C",
		cfgPathname,
	}
	go exec.Command(cmd, args...).Output()
}

//    pkill -u nobody -f dnsmasq.${IFNAME}.conf
func stopDnsmasq(cfgFilename string, printOnError bool) {
	fmt.Printf("stopDnsmasq: %s\n", cfgFilename)
	pkillUserArgs("nobody", cfgFilename, printOnError)
}

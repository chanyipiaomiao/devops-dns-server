package server

import (
	"devops-dns-server/config"
	"devops-dns-server/source"
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

type DNSHandler struct{}

func (d *DNSHandler) ServeDNS(respWriter dns.ResponseWriter, messageFromClient *dns.Msg) {
	var (
		domain             string
		respMessage        dns.Msg
		address            string
		respFromNameServer *dns.Msg //从上游nameserver返回的响应消息
	)

	respMessage = dns.Msg{}
	respMessage.SetReply(messageFromClient)

	domain = strings.ToLower(messageFromClient.Question[0].Name)

	switch messageFromClient.Question[0].Qtype {
	case dns.TypeA:
		respMessage.Authoritative = true
		address = source.GetIP(domain)
		if address != "" {
			respMessage.Answer = append(respMessage.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(address),
			})
			_ = respWriter.WriteMsg(&respMessage)
		} else {
			client := new(dns.Client)
			client.Net = "udp"
			respFromNameServer, _, _ = client.Exchange(messageFromClient, config.GetConfig().String("server::nameserver"))
			respFromNameServer.Question[0].Name = strings.ToLower(respFromNameServer.Question[0].Name)
			for i := 0; i < len(respFromNameServer.Answer); i++ {
				respFromNameServer.Answer[i].Header().Name = strings.ToLower(respFromNameServer.Answer[i].Header().Name)
			}
			_ = respWriter.WriteMsg(respFromNameServer)
		}

	}
}

func Listen() {
	source.WatchFile()
	srv := &dns.Server{Addr: config.GetConfig().String("server::listen"), Net: "udp"}
	srv.Handler = &DNSHandler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}

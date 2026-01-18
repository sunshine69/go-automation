package main

import (
	"fmt"
	"regexp"

	"github.com/go-acme/lego/v4/challenge/dns01"
	u "github.com/sunshine69/golang-tools/utils"
)

type MaraDNSProvider struct {
	ZoneFilePath string
	Vars         map[string]any
}

func NewMaraDNSProvider(p *MaraDNSProvider) *MaraDNSProvider {
	p.ZoneFilePath = u.MapLookup(p.Vars, "maradns_config_file", "/etc/maradns/maradns.org").(string)
	return p
}

// Present creates the TXT record in the MaraDNS csv2 zone file.
func (p *MaraDNSProvider) Present(domain, token, keyAuth string) error {
	// 1. Get the official challenge info (FQDN and hashed value)
	info := dns01.GetChallengeInfo(domain, keyAuth)

	// MaraDNS csv2 format for TXT: name. TXT 'value' ~. Mine one does not have teh ~ though so remove it
	// Note: MaraDNS requires a trailing dot for FQDNs and the '~' record separator
	recordLine := fmt.Sprintf("%s TXT '%s'", info.EffectiveFQDN, info.Value)
	safeFQDN := regexp.QuoteMeta(info.EffectiveFQDN)
	recordLinePtn := fmt.Sprintf(`%s TXT '[^\s']+'`, safeFQDN)
	println("recordLine: " + recordLine)
	println("recordLinePtn: " + recordLinePtn)
	// 2. Use your 'lineinfile' tool to append the record
	if err, _ := u.LineInFile(p.ZoneFilePath, u.NewLineInfileOpt(
		&u.LineInfileOpt{
			Line:   recordLine,
			Regexp: recordLinePtn,
			Backup: true,
		},
	)); err == nil {
		return reloadMaraDNS(p.Vars)
	} else {
		return fmt.Errorf("[ERROR] in lineinfile - %s", err.Error())
	}
}

// CleanUp removes the TXT record from the zone file.
func (p *MaraDNSProvider) CleanUp(domain, token, keyAuth string) error {
	info := dns01.GetChallengeInfo(domain, keyAuth)
	recordLine := fmt.Sprintf("%s TXT '%s'", info.EffectiveFQDN, info.Value)
	safeFQDN := regexp.QuoteMeta(info.EffectiveFQDN)
	recordLinePtn := fmt.Sprintf(`%s TXT '[^\s']+'`, safeFQDN)
	// 2. Use your 'lineinfile' tool to append the record
	if err, _ := u.LineInFile(p.ZoneFilePath, u.NewLineInfileOpt(
		&u.LineInfileOpt{
			Line:   recordLine,
			Regexp: recordLinePtn,
			Backup: true,
			State:  "absent",
		},
	)); err == nil {
		return reloadMaraDNS(p.Vars)
	} else {
		return fmt.Errorf("[ERROR] in lineinfile - %s", err.Error())
	}
}

func reloadMaraDNS(Vars map[string]any) error {
	// Signal MaraDNS to reload (often via SIGHUP or restarting the service)
	if o, err := u.RunSystemCommandV2(u.MapLookup(Vars, "maradns_reload_cmd", "systemctl reload maradns").(string), true); err != nil {
		println("[ERROR] " + err.Error() + "\nOut: " + o)
		return err
	}
	return nil
}

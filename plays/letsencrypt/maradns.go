package main

import (
	"fmt"
	"os/exec"

	"github.com/go-acme/lego/v4/challenge/dns01"
	u "github.com/sunshine69/golang-tools/utils"
)

type MaraDNSProvider struct {
	ZoneFilePath string
}

// Present creates the TXT record in the MaraDNS csv2 zone file.
func (p *MaraDNSProvider) Present(domain, token, keyAuth string) error {
	// 1. Get the official challenge info (FQDN and hashed value)
	info := dns01.GetChallengeInfo(domain, keyAuth)

	// MaraDNS csv2 format for TXT: name. TXT 'value' ~. Mine one does not have teh ~ though so remove it
	// Note: MaraDNS requires a trailing dot for FQDNs and the '~' record separator
	recordLine := fmt.Sprintf("%s. TXT '%s'", info.EffectiveFQDN, info.Value)
	recordLinePtn := fmt.Sprintf(`%s\. TXT '[^\s']+'`, info.EffectiveFQDN)
	// 2. Use your 'lineinfile' tool to append the record
	if err, changed := u.LineInFile(p.ZoneFilePath, u.NewLineInfileOpt(
		&u.LineInfileOpt{
			Line:   recordLine,
			Regexp: recordLinePtn,
			Backup: true,
		},
	)); err == nil && changed {
		return reloadMaraDNS()
	} else {
		return fmt.Errorf("[ERROR] in lineinfile - %s", err.Error())
	}
}

// CleanUp removes the TXT record from the zone file.
func (p *MaraDNSProvider) CleanUp(domain, token, keyAuth string) error {
	info := dns01.GetChallengeInfo(domain, keyAuth)
	recordLine := fmt.Sprintf("%s. TXT '%s'", info.EffectiveFQDN, info.Value)
	recordLinePtn := fmt.Sprintf(`%s\. TXT '[^\s']+'`, info.EffectiveFQDN)
	// 2. Use your 'lineinfile' tool to append the record
	if err, changed := u.LineInFile(p.ZoneFilePath, u.NewLineInfileOpt(
		&u.LineInfileOpt{
			Line:   recordLine,
			Regexp: recordLinePtn,
			Backup: true,
			State:  "absent",
		},
	)); err == nil && changed {
		return reloadMaraDNS()
	} else {
		return fmt.Errorf("[ERROR] in lineinfile - %s", err.Error())
	}
}

func reloadMaraDNS() error {
	// Signal MaraDNS to reload (often via SIGHUP or restarting the service)
	return exec.Command("systemctl", "reload", "maradns").Run()
}

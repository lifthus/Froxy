package config

import (
	"crypto/tls"

	"github.com/lifthus/froxy/internal/froxyfile"
	"github.com/lifthus/froxy/pkg/helper"
)

// ReverseFroxy holds each reverse proxy's config
type ReverseFroxy struct {
	Name     string
	Port     string
	Insecure bool
	Proxy    []*ReverseProxySet
}

func (rf *ReverseFroxy) GetTLSConfig() *tls.Config {
	if rf.Insecure {
		return nil
	}
	certs := make([]tls.Certificate, len(rf.Proxy))
	for _, p := range rf.Proxy {
		certs = append(certs, p.Certificate)
	}
	return &tls.Config{Certificates: certs}
}

type ReverseProxySet struct {
	Host        string
	Certificate tls.Certificate
	Target      []struct {
		Path string
		To   []string
	}
}

func configReverseProxyList(ff []froxyfile.ReverseFroxy) (rfs []*ReverseFroxy, err error) {
	rfs = make([]*ReverseFroxy, len(ff))
	for i, f := range ff {
		rf := &ReverseFroxy{}
		rf.Name = f.Name
		rf.Port = f.Port
		rf.Insecure = f.Insecure
		rf.Proxy, err = setReverseProxies(f.Insecure, f.Proxy)
		if err != nil {
			return nil, err
		}
		rfs[i] = rf
	}
	return rfs, nil
}

func setReverseProxies(insecure bool, rpfs []struct {
	Host string `yaml:"host"`
	TLS  *struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"tls"`
	Target []struct {
		Path string   `yaml:"path"`
		To   []string `yaml:"to"`
	} `yaml:"target"`
}) ([]*ReverseProxySet, error) {
	var err error
	rpss := make([]*ReverseProxySet, len(rpfs))

	for i, rpf := range rpfs {
		rps := &ReverseProxySet{}
		if !insecure && !isKeyPairGiven(&rpf) {
			rps.Certificate, err = helper.SignTLSCertSelf()
		} else if !insecure && isKeyPairGiven(&rpf) {
			rps.Certificate, err = helper.LoadTLSCert(rpf.TLS.Cert, rpf.TLS.Key)
		}
		if err != nil {
			return nil, err
		}

		rps.Host = rpf.Host
		rps.Target = []struct {
			Path string
			To   []string
		}(rpf.Target)

		rpss[i] = rps
	}
	return rpss, nil
}

func isKeyPairGiven(p *struct {
	Host string `yaml:"host"`
	TLS  *struct {
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"tls"`
	Target []struct {
		Path string   `yaml:"path"`
		To   []string `yaml:"to"`
	} `yaml:"target"`
}) bool {
	return p.TLS != nil
}

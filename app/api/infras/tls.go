package infras

import (
	"crypto/tls"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitTLS() *tls.Config {
	cfg := &tls.Config{
		MinVersion:         tls.VersionTLS10,
		InsecureSkipVerify: true,
	}

	cert, err := tls.LoadX509KeyPair(
		"./app/api/cert/server.crt",
		"./app/api/cert/server.key",
	)

	if err != nil {
		hlog.Fatal("tls  failed", err)
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	return cfg
}

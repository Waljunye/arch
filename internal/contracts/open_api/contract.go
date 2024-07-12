package open_api

import (
	v1 "arch/internal/contracts/open_api/v1"
	"crypto/tls"
	"log"
)

func New(
	logger *log.Logger,
	cert *tls.Certificate,
	v1 *v1.Contract) /*common.PortListener*/ {
	contractListener := &contractListener{
		v1: v1,
	}

	// creating gin handler
	//handler := gin.New()

	// routing
	// handler.GET("/some_domain_object", contractListener.v1.DoBuisnessLogic)

	// in return: portListener creating
	//return http_server.New(handler, cert,...)
}

type contractListener struct {
	v1 *v1.Contract
}

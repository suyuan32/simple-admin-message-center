package config

import (
	"crypto/tls"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/smtp"
)

type EmailConf struct {
	AuthType  string `json:",default=plain,options=[plain,CRAMMD5],env=EMAIL_AUTH_TYPE"`
	EmailAddr string `json:",env=EMAIL_ADDR"`
	Password  string `json:",optional,env=EMAIL_PASSWORD"`
	HostName  string `json:",env=EMAIL_HOST_NAME"`
	Identify  string `json:",optional,env=EMAIL_IDENTIFY"`
	Secret    string `json:",optional,env=EMAIL_SECRET"`
	Port      int    `json:",optional,default=25,env=EMAIL_PORT"`
	TLS       bool   `json:",default=false,env=EMAIL_TLS"`
}

// NewAuth creates the auth from config
func (e EmailConf) NewAuth() *smtp.Auth {
	var auth smtp.Auth
	switch e.AuthType {
	case "plain":
		auth = smtp.PlainAuth(e.Identify, e.EmailAddr, e.Password, e.HostName)
	case "CRAMMD5":
		auth = smtp.CRAMMD5Auth(e.EmailAddr, e.Secret)
	}
	return &auth
}

func (e EmailConf) NewClient(auth *smtp.Auth) *smtp.Client {
	hostAddress := fmt.Sprintf("%s:%d", e.HostName, e.Port)
	if e.TLS == true {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         hostAddress,
		}

		// initialize connection
		conn, err := tls.Dial("tcp", hostAddress, tlsconfig)
		logx.Must(err)

		// get client
		c, err := smtp.NewClient(conn, e.HostName)
		logx.Must(err)

		err = c.Auth(*auth)
		logx.Must(err)

		return c
	} else {
		// initialize connection
		c, err := smtp.Dial(hostAddress)
		logx.Must(err)

		err = c.Auth(*auth)
		logx.Must(err)

		return c
	}
}

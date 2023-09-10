// Package email implements utility routines for email convert
package email

import "github.com/suyuan32/simple-admin-message-center/internal/enum/emailauthtype"

func ConvertAuthTypeToString(data uint8) string {
	switch data {
	case emailauthtype.Plain:
		return "plain"
	case emailauthtype.CRAMMD5:
		return "CRAMMD5"
	}

	return "plain"
}

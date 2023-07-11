package smssdk

import sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

type SmsConf struct {
	SecretId  string
	SecretKey string
	Provider  string
	Region    string
}

type SmsGroup struct {
	TencentSmsClient *sms.Client
}

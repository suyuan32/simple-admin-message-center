package smssdk

import sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

type SmsClient interface {
	SendSms() (any, error)
}

type SmsConf struct {
	SecretId  string
	SecretKey string
	Provider  string `json:",default=tencent,options=[tencent]"`
	Region    string `json:",optional"`
}

type SmsGroup struct {
	TencentSmsClient *sms.Client
}

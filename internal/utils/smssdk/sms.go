package smssdk

import sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

type SmsClient interface {
	SendSms() (any, error)
}

type SmsConf struct {
	SecretId  string `json:",env=SMS_SECRET_ID"`
	SecretKey string `json:",env=SMS_SECRET_KEY"`
	Provider  string `json:",default=tencent,options=[tencent],env=SMS_PROVIDER"`
	Region    string `json:",optional,env=SMS_REGION"`
}

type SmsGroup struct {
	TencentSmsClient *sms.Client
}

package sms

// AliyunSms 公共参数
type AliyunSms struct {
	Format           string //返回值的类型，支持JSON与XML
	Version          string //API版本号，为日期形式：YYYY-MM-DD，本版本对应为2017-03-21
	AccessKeyID      string `url:"AccessKeyId"` //阿里云颁发给用户的访问服务所用的密钥ID
	SignatureMethod  string //签名方式，目前支持HMAC-SHA1
	Timestamp        string //请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ例如，2017-3-29T12:00:00Z(为北京时间2017年3月29日的20点0分0秒
	SignatureVersion string //签名算法版本，目前版本是1.0
	SignatureNonce   string //唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值
	AccessSecret     string `url:"-"`
}

// SendSmsResposeEntity 发送短信返回
type SendSmsResposeEntity struct {
	RequestID string `json:"RequestId"`
	Code      string
	Message   string
	BizID     string `json:"BizId"`
}

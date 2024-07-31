package tools

// SecretId 监控密钥
const SecretID = "AAAA0000bbbb1111CCCC2222dddd3333EEEE"
const SecretKey = "abcdefg0123456hijklmn7890123opqr"

// TEST

// ServerAddr 日志汇上报地址
const ServerAddr = "report.test.com:12345"

// Topic zhiyan-log接入信息
const Topic = "sdk-abcdefgh12345678"

// Limit 1次分页查询数量
const Limit = 100

// DefaultOffset 分页查询的起始偏移量
const DefaultOffset = 0

const Goroutines = 1

type InstanceInfo struct {
	InstanceId   string `json:"InstanceId"`
	InstanceName string `json:"InstanceName"`
}

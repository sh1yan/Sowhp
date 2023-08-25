package scripts

import "regexp"

// IsIPAddress 判断字符串是否是IP地址的格式
func IsIPAddress(str string) bool {
	ipPattern := `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`
	match, _ := regexp.MatchString(ipPattern, str)
	return match
}

// IsIPAddressWithPort 判断字符串是否是IP地址加端口号的格式
func IsIPAddressWithPort(str string) bool {
	ipPortPattern := `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d+$`
	match, _ := regexp.MatchString(ipPortPattern, str)
	return match
}

// IsDomainName 判断字符串是否是域名
func IsDomainName(str string) bool {
	domainPattern := `^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(domainPattern, str)
	return match
}

// IsDomainNameWithPort 判断字符串是否是域名加端口形式
func IsDomainNameWithPort(str string) bool {
	domainPortPattern := `^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}:\d+$`
	match, _ := regexp.MatchString(domainPortPattern, str)
	return match
}

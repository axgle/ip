package ip

import "testing"

func init() {
	Load("17monipdb.dat")
}
func TestFind(t *testing.T) {
	ip_addresses := map[string]string{
		"8.8.8.8":        "GOOGLE\tGOOGLE",
		"8.8.4.4":        "GOOGLE\tGOOGLE",
		"202.106.195.68": "中国\t北京",
		"123.118.91.246": "中国\t北京",
		"202.115.128.64": "中国\t四川",
		"110.106.46.151": "中国\t辽宁",
		"11.106.46.151":  "美国\t美国",
		"20.106.46.151":  "美国\t美国",
		"1.106.46.151":   "韩国\t韩国",
		"99.106.46.151":  "加拿大\t加拿大",
		"10.106.46.151":  "局域网\t局域网",
	}

	for ip, address := range ip_addresses {
		found_adress := Find(ip)
		if found_adress != address {
			t.Error(ip, "should match", address, "but", found_adress, "found")
		}
	}
}

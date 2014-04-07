package main

import "github.com/axgle/ip"

func main() {
	ip.Load("../17monipdb.dat")

	address := ip.Find("8.8.8.8") //address: GOOGLE\tGOOGLE\t
	println(address)

	address = ip.Find("202.106.46.151") //address: 中国\t北京\t
	println(address)

	address = ip.Find("202.115.128.64") //address: 中国\t四川\t成都理工大学
	println(address)

}

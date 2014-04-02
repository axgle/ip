17MonIP Golang Lib
======

IP search based on 17monipdb, the IP database parser for china with golang.


install
--------

	 go get github.com/axgle/ip

example
-------

	package main
	
	import "github.com/axgle/ip"
	
	func main() {
		ip.Load("../17monipdb.dat")
	
		address := ip.Find("8.8.8.8") //address: GOOGLE\tGOOGLE
		println(address)
	
		address = ip.Find("202.106.46.151") //address: 中国\t北京
		println(address)
	
	}


## License

BSD

## Author

* axgle
## Version

0.0.2 remove binary package
0.0.1 init

## Thanks

* Paul Gao: for his 17monipdb.dat data.
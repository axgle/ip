package ip

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

var buffer []byte
var ip_offset uint32
var ip_index []byte

func Load(ipBinaryFilePath string) {
	var err error
	buffer, err = ioutil.ReadFile(ipBinaryFilePath)
	if err != nil {
		log.Fatal(err)
	}
}

func Find(ip_address string) string {

	ipdot := strings.Split(ip_address, ".")
	pre_ip, _ := strconv.Atoi(ipdot[0])
	tmp_offset := pre_ip * 4

	ip_offset = bytesBigEndianToUint32(buffer[:4])
	ip_index = buffer[4:ip_offset]
	nip := ip2long(ip_address)

	start_len := binary.LittleEndian.Uint32(ip_index[tmp_offset : tmp_offset+4])

	var index_offset uint32 = 0
	var index_length uint32 = 0

	for start := start_len*8 + 1024; start < ip_offset-1028; start += 8 {
		if bytesBigEndianToUint32(ip_index[start:start+4]) >= nip {
			index_length = 0xFF & uint32(ip_index[start+7])
			tmp := ip_index[start+4 : start+7]
			tmp = append(tmp, 0x0)
			index_offset = binary.LittleEndian.Uint32(tmp)
			break
		}
	}
	if index_offset == 0 {
		return "N/A\tN/A"
	}

	pos := index_offset + ip_offset - 1024

	//fmt.Printf("%s", ip_index[pos-4:pos+index_length-4])

	return string(ip_index[pos-4 : pos+index_length-4])

}

func bytesBigEndianToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return bytesBigEndianToUint32(ip)
}

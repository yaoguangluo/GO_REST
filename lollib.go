package lib

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func Aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func Mtos(ipMask net.IP) int {
	mask := net.IPMask(ipMask.To4())
	size, _ := mask.Size()
	return size
}

func Stom(size int) net.IP {
	divisor := size / 8
	remainder := size % 8
	var mask string
	var b bool
	for i := 1; i <= 4; i++ {
		if b {
			mask += "0."
			continue
		} else if i <= divisor {
			mask += "255."
		} else {
			mask += fmt.Sprint(2<<7-2<<(7-uint(remainder))) + "."
			b = true
		}
	}
	return net.ParseIP(mask[:len(mask)-1])
}

// base input 1.2.3.4/24 return 1.2.3.0
func Base(block string) net.IP {
	_, ipnet, err := net.ParseCIDR(block)
	if err != nil {
		fmt.Println(err.Error())
	}
	return ipnet.IP
}

// mask input 1.2.3.4/24 return 24
func GetMask(block string) int {
	_, ipnet, err := net.ParseCIDR(block)
	if err != nil {
		fmt.Println(err.Error())
	}
	ms, _ := ipnet.Mask.Size()
	return ms
}

// Len input 1.2.3.4/24 return 256
func Len(block string) int64 {
	i := GetMask(block)
	if i >= 32 {
		return 0
	}
	return 2 << uint(32-i-1)
}

// get all ip in block, input 1.2.3.4/24 return all ip
func GetAllIP(block string) []net.IP {
	var ips []net.IP
	base := Aton(Base(block))
	end := Len(block)
	if end == 0 {
		return []net.IP{Base(block)}
	}
	var i int64
	for i = 0; i < end; i++ {
		ip := Ntoa(base + i)
		ips = append(ips, ip)
	}
	return ips
}

// the nth ip of block, input 1.2.3.4/24 1 return 1.2.3.1
func Nth(block string, i int64) net.IP {
	l := Len(block)
	if l == 0 {
		return Base(block)
	}
	if i > l {
		fmt.Println("cannot do it")
	}
	ip, _, err := net.ParseCIDR(block)
	if err != nil {
		fmt.Println(err.Error())
	}
	resnum := Aton(ip) + i - 1
	return Ntoa(resnum)
}

//parser ip from string
func ParseIPInt(ip net.IP) [4]int {
	var i [4]int
	token := strings.Split(ip.String(), ".")
	for k, v := range token {
		i[k], _ = strconv.Atoi(v)
	}
	return i
}

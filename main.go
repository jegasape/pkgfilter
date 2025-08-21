package main

import "fmt"

type PackageTypes int

const (
	TCP PackageTypes = iota
	UPD
	ICMP
	HTTP
	HTTPS
	SSH
)

type Package struct {
	Type PackageTypes
}

func (p Package) String() string {
	return [...]string{"TCP", "UPD", "ICMP", "HTTP", "HTTPS", "SSH"}[p.Type]
}

func main() {
	pkg := Package{Type: SSH}
	fmt.Println(pkg)
}

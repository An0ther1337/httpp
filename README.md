# HttpFloodProtect
Flood protection for your site on linux server

# Dependencies
Golang version 1.7+<br>
Iptables<br>
Libs: golang.org/x/net/ipv4, github.com/google/gopacket

# Files
httpp - compiled file<br>
httpp.go - source file in Golang

# Run
Type ./httpp \<port\> \<limit of requests\><br>
Example: ./httpp 80 1000

# Compile
Type "go build" in folder with httpp.go file

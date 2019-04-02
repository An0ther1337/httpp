package main

import(
    "fmt"
    "net"
    "golang.org/x/net/ipv4"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "os"
    "strconv"
    "strings"
    "time"
    "os/exec"
)

func main(){
    
    if len(os.Args) < 3{
        fmt.Println("Error: Not enough data. Enter port and limit.\nExample: ./httpp 80 10000");
        return
    }
    
    port, err := strconv.Atoi(os.Args[1])
    if err != nil{
        fmt.Println("Error: Unreal port");
        return
    }
    
    limit, err := strconv.Atoi(os.Args[2])
    if err != nil{
        fmt.Println("Error: Unreal limit");
        return
    }
    
    ips := make(map[string]int)
    sec := int64(0)
    
    c, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
    if err != nil{
        fmt.Println("Error: ", err)
    }
    defer c.Close()
    
    r, err := ipv4.NewRawConn(c)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    
    fmt.Println("\n\tHttpProtect enabled!\n\tProtected port: ", port, "\n\tRequests/sec limit: ", limit, "\n\tAuthor: github.com/An0ther1337\n")
    
    for{
        if time.Now().Unix() != sec{
            sec = time.Now().Unix()
            ips = make(map[string]int)
        }
        
        buf := make([]byte, 128)
        r.ReadFrom(buf)
        packet := gopacket.NewPacket(buf, layers.LayerTypeIPv4, gopacket.Default)
        if packet != nil{
            ip := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
            tcp := packet.Layer(layers.LayerTypeTCP).(*layers.TCP)
            app := packet.ApplicationLayer()
            if tcp.DstPort == layers.TCPPort(port) && app != nil{
                ips[ip.SrcIP.String()] += 1;
                //fmt.Println(ips)
                fmt.Println("New TCP (HTTP) packet!\nSender's IP: ", ip.SrcIP, "\nYour IP: ", ip.DstIP, "\nSender's port: ", tcp.SrcPort, "\nYour port: ", tcp.DstPort, "\nProtocol: ", ip.Protocol, "\nPayload: ", strings.Replace(string(app.Payload()), "\n", "\\n", -1), "\n")
                
                if ips[string(ip.SrcIP)] > limit{
                    exec.Command("iptables", "-I", "INPUT", "-s", ip.SrcIP.String(), "-j", "DROP")
                    
                }
            }
        }
    }
}

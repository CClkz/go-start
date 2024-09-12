package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	protocolICMP = 1 // ICMP 协议号
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: myping <hostname>")
		return
	}

	hostname := os.Args[1]
	addr, err := net.ResolveIPAddr("ip4", hostname)
	if err != nil {
		fmt.Printf("Failed to resolve address: %v\n", err)
		return
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		fmt.Printf("Failed to create connection: %v\n", err)
		return
	}
	defer conn.Close()

	// 构建 ICMP Echo Request
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, // Echo Request 类型
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,      // 使用进程 ID 作为 ID
			Seq:  1,                         // 序列号
			Data: []byte("HELLO-R-U-THERE"), // 发送的数据
		},
	}

	// 将消息序列化为字节
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		fmt.Printf("Failed to marshal ICMP message: %v\n", err)
		return
	}

	// 记录发送时间
	start := time.Now()

	// 发送 ICMP Echo Request
	_, err = conn.Write(msgBytes)
	if err != nil {
		fmt.Printf("Failed to send ICMP message: %v\n", err)
		return
	}

	// 接收 ICMP Echo Reply
	reply := make([]byte, 1500)
	n, err := conn.Read(reply)
	if err != nil {
		fmt.Printf("Failed to receive ICMP reply: %v\n", err)
		return
	}

	// 记录接收时间
	duration := time.Since(start)

	// 解析 ICMP 回复消息
	rm, err := icmp.ParseMessage(protocolICMP, reply[:n])
	if err != nil {
		fmt.Printf("Failed to parse ICMP reply: %v\n", err)
		return
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		// 获取 ICMP Echo Reply 的具体内容
		fmt.Printf("Ping to %s successful! RTT: %v\n", hostname, duration)
	default:
		fmt.Printf("Got unexpected ICMP message: %v\n", rm)
	}
}

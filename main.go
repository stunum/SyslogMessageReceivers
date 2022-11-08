package main

import (
	"flag"
	"fmt"

	"gopkg.in/mcuadros/go-syslog.v2"
)

var facilityMap = map[interface{}]string{
	0:  "Kernel messages",
	1:  "User-level messages",
	2:  "Mail system",
	3:  "System daemons",
	4:  "Security/authorization messages",
	5:  "Messages generated internally by Syslogd",
	6:  "Line printer subsystem",
	7:  "Network news subsystem",
	8:  "UUCP subsystem",
	9:  "Clock daemon",
	10: "Security/authorization messages",
	11: "FTP daemon",
	12: "NTP subsystem",
	13: "Log audit",
	14: "Log alert",
	15: "Clock daemon",
	16: "Local use 0 (local0)",
	17: "Local use 1 (local1)",
	18: "Local use 2 (local2)",
	19: "Local use 3 (local3)",
	20: "Local use 4 (local4)",
	21: "Local use 5 (local5)",
	22: "Local use 6 (local6)",
	23: "Local use 7 (local7)",
}
var severityMap = map[interface{}]string{
	0: "Emergency",
	1: "Alert",
	2: "Critical",
	3: "Error",
	4: "Warning",
	5: "Notice",
	6: "Informational",
	7: "Debug",
}

var proto string
var address string

func init() {
	flag.StringVar(&proto, "t", "UDP", "Syslog协议: TCP or UDP")
	flag.StringVar(&address, "l", "0.0.0.0:514", "Syslog监听地址,例: 127.0.0.1:514")
	flag.Parse()
}

func main() {
	msChan := make(syslog.LogPartsChannel, 10000)
	handler := syslog.NewChannelHandler(msChan)
	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	switch proto {
	case "TCP":
		server.ListenTCP(address)
	case "UDP":
		server.ListenUDP(address)
	}
	go func(msgchan *syslog.LogPartsChannel) {
		for msg := range *msgchan {
			fmt.Println("====================================")
			fmt.Printf("优先级(priority): %d\n", msg["priority"])
			fmt.Printf("设备编号(facility): %s\n", facilityMap[msg["facility"]])
			fmt.Printf("日志级别(severity): %s\n", severityMap[msg["severity"]])
			fmt.Printf("来源(client): %s\n", msg["client"])
			fmt.Printf("内容(content): %s\n", msg["content"])
			fmt.Printf("时间(timestamp): %s\n", msg["timestamp"])
			fmt.Println("************************************")
		}
	}(&msChan)
	fmt.Printf("开始接受syslog请求，监听地址：%s 协议：%s \n", address, proto)
	server.Boot()
	server.Wait()
}

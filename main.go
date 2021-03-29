package main

import (
	"fmt"
	"github.com/allanpk716/Premote-Plugin-SmartProtocol/Model"
	gpd "github.com/allanpk716/go-protocol-detector"
	"os"
	"strings"
	"time"
)


func main() {
	/*
		0	this app path
		1	protocol name
		2	SP://address
		out	return 0
			worked address
			return 1
			error info
	*/
	// check and load args
	if len(os.Args) != 3 {
		fmt.Println("input arg error")
		os.Exit(Model.ExitCode)
	}
	protocolName := strings.ToUpper(os.Args[1])
	CPAddress := strings.ToUpper(os.Args[2])
	if strings.Contains(CPAddress, Model.SmartProtocolPrefix) == false {
		fmt.Println("`SP://` not found in this address", CPAddress)
		os.Exit(Model.ExitCode)
	}
	// init config
	smartPMap, err := Model.InitConfigure()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(Model.ExitCode)
	}
	if sp, ok := smartPMap[CPAddress]; ok {

		if protocolName != sp.ProtocolName {
			fmt.Println(Model.ErrInputProtocolNameNotFitConfigProtocolName.Error())
			os.Exit(Model.ExitCode)
		}
		detect := gpd.NewDetector(time.Duration(sp.TimeOut) * time.Millisecond)
		outAddressAndPort, err := Model.CheckAll(detect, sp)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(Model.ExitCode)
		}
		fmt.Println(outAddressAndPort)
	} else {
		fmt.Println(CPAddress ,"not found")
		os.Exit(Model.ExitCode)
	}
}
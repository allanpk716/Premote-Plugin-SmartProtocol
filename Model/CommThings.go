package Model

import "errors"

const (
	SmartProtocolPrefix = "SP://"
	SPConfigsField = "SPConfigs"
	SPConfigsName = "SPConfigs."
	MultiAddressName = ".MultiAddress"
	ProtocolName = ".ProtocolName"
	TimeOutName = ".TimeOut"
	UseCommonPortCheckName = ".UseCommonPortCheck"

	TimeOutDefault = 1000

	ExitCode = 100


	ProtocolNameRDP = "RDP"
	ProtocolNameFTP = "FTP"
	ProtocolNameSFTP = "SFTP"
	ProtocolNameSSH = "SSH"
	ProtocolNameVNC = "VNC"
	ProtocolNameTelnet = "Telnet"
)

var (
	ErrInputProtocolNameNotFitConfigProtocolName        = errors.New("input protocol name not fit config protocol name")
	ErrInitConfigReadAddressError                       = errors.New("init config read address error")
	ErrProtocolNotSupport                               = errors.New("protocol not support")
	ErrInitConfigSmartProtocolNameIsDuplicateDefinition = errors.New("init config smart protocol name is duplicate definition")
	ErrProtocolNameIsEmpty                              = errors.New("protocol name is empty")
	ErrCheckProtocolError                               = errors.New("check protocol error")
)

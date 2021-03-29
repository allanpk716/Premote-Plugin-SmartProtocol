package Model

import (
	gpd "github.com/allanpk716/go-protocol-detector"
	"sort"
	"time"
)

func timeCost(start time.Time, addInfo *AddressInfo) {
	tc := time.Since(start)
	addInfo.CostTime = tc
}

func checkOne(check func(host string, port string) error, sp *SmartProtocol) (string, error) {
	//wg := &sync.WaitGroup{}
	var ch = make(chan AddressInfo, len(sp.MultiAddressInfo))
	defer close(ch)
	for _, info := range sp.MultiAddressInfo {
		tmpInfo := info
		//wg.Add(1)

		go func(info AddressInfo) {
			nowTime := time.Now()
			defer func() {
				timeCost(nowTime, &tmpInfo)
				//wg.Done()
				ch <- tmpInfo
			}()

			err := check(tmpInfo.IP, tmpInfo.Port)
			if err != nil {
				tmpInfo.Worked = false
			} else {
				tmpInfo.Worked = true
			}
		}(info)
	}
	//wg.Wait()
	for i := 0; i < len(sp.MultiAddressInfo); i++ {
		nowInfo := <- ch
		for index := range sp.MultiAddressInfo {
			if sp.MultiAddressInfo[index].IP == nowInfo.IP && sp.MultiAddressInfo[index].Port == nowInfo.Port {
				sp.MultiAddressInfo[index] = nowInfo
			}
		}
	}
	// 找到有效的，且低延迟的
	sort.Sort(sort.Reverse(AddressInfoSlice(sp.MultiAddressInfo)))
	for i, info := range sp.MultiAddressInfo {
		if info.Worked == true {
			return sp.MultiAddressInfo[i].IP + ":" +sp.MultiAddressInfo[i].Port, nil
		}
	}
	return "", ErrCheckProtocolError
}

func CheckAll(detect *gpd.Detector, sp SmartProtocol) (string, error) {
	var err error
	outAddressAndPort := ""
	switch sp.ProtocolName {
	case ProtocolNameRDP:
		outAddressAndPort, err = checkOne(detect.RDPCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	case ProtocolNameFTP:
		outAddressAndPort, err = checkOne(detect.FTPCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	case ProtocolNameSFTP:
		// SFTP Check need to be improved
		outAddressAndPort, err = checkOne(detect.SSHCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	case ProtocolNameSSH:
		outAddressAndPort, err = checkOne(detect.SSHCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	case ProtocolNameVNC:
		outAddressAndPort, err = checkOne(detect.VNCCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	case ProtocolNameTelnet:
		outAddressAndPort, err = checkOne(detect.TelnetCheck, &sp)
		if err != nil {
			return "", err
		}
		return outAddressAndPort, nil
	default:
		return "", ErrProtocolNotSupport
	}
}

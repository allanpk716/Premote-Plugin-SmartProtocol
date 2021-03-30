package Model

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type SmartProtocol struct {
	ProtocolName 		string
	TimeOut				int
	UseCommonPortCheck	bool
	MultiAddressInfo	[]AddressInfo
}

type AddressInfo struct {
	IP 			string
	Port 		string
	CostTime 	time.Duration
	Worked		bool
}

type AddressInfoSlice [] AddressInfo

func (a AddressInfoSlice) Len() int {
	return len(a)
}
func (a AddressInfoSlice) Swap(i, j int) {
	a[i], a[j] =  a[j], a[i]
}
func (a AddressInfoSlice) Less(i, j int) bool {
	return a[j].CostTime < a[i].CostTime
}

func initConfigure() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("SPConfig")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.New("error reading config:" + err.Error())
	}
	return v, nil
}

func InitConfigure() (map[string]SmartProtocol, error) {
	spMap := map[string]SmartProtocol{}
	config, err := initConfigure()
	if err != nil {
		return nil, err
	}
	sps := config.GetStringMapString(SPConfigsField)
	for spName := range sps {
		oneSP := SmartProtocol{}
		oneSP.ProtocolName = strings.ToUpper(config.GetString(SPConfigsName + spName + ProtocolName))
		oneSP.TimeOut = config.GetInt(SPConfigsName + spName + TimeOutName)
		oneSP.UseCommonPortCheck = config.GetBool(SPConfigsName + spName + UseCommonPortCheckName)
		adds := config.GetStringSlice(SPConfigsName + spName + MultiAddressName)
		for _, onAddress := range adds{
			tmp := strings.Split(onAddress, ":")
			if len(tmp) != 2 {
				return nil, ErrInitConfigReadAddressError
			}
			oneSP.MultiAddressInfo = append(oneSP.MultiAddressInfo, AddressInfo{
				IP: tmp[0],
				Port: tmp[1],
			})
		}
		if oneSP.ProtocolName == "" {
			return nil, ErrProtocolNameIsEmpty
		}
		if oneSP.TimeOut <= 0 {
			oneSP.TimeOut = TimeOutDefault
		}
		// to upper word
		upperSpName := strings.ToUpper(SmartProtocolPrefix + spName)
		// check Duplicate Definition
		if _, ok := spMap[upperSpName]; ok {
			return nil, ErrInitConfigSmartProtocolNameIsDuplicateDefinition
		}
		spMap[upperSpName] = oneSP
	}
	return spMap, nil
}

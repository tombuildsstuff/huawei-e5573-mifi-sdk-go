package mifi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Status struct {
	CurrentBatteryPercentage     int
	BatteryCharging              bool
	MaxSignalBars                int
	CurrentSignalBars            int
	NumberOfUsersConnectedToWifi int
}

type statusResponse struct {
	BatteryPercent  int  `xml="BatteryPercent"`
	CurrentWifiUser int  `xml="CurrentWifiUser"`
	SignalIcon      int  `xml="SignalIcon"`
	BatteryStatus   bool `xml="BatteryStatus"` // 0 or 1 is auto-converted to bool
}

func (m Mifi) CurrentStatus() (*Status, error) {
	endpoint := fmt.Sprintf("%s/api/monitoring/status", m.Endpoint)
	resp, err := m.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sr := &statusResponse{}
	err = xml.Unmarshal(responseData, sr)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling XML: %+v", err)
	}

	if sr != nil {
		result := Status{
			CurrentBatteryPercentage:     sr.BatteryPercent,
			CurrentSignalBars:            sr.SignalIcon,
			MaxSignalBars:                5,
			NumberOfUsersConnectedToWifi: sr.CurrentWifiUser,
			BatteryCharging:              sr.BatteryStatus,
		}
		return &result, nil
	}

	return nil, fmt.Errorf("XML wasn't valid: %s", string(responseData))
}

// TODO
// <ConnectionStatus>901</ConnectionStatus>1
// <WifiConnectionStatus>902</WifiConnectionStatus>
// <SignalStrength></SignalStrength>
// <CurrentNetworkType>9</CurrentNetworkType>
// <CurrentServiceDomain>3</CurrentServiceDomain>i
// <RoamingStatus>0</RoamingStatus>
// <BatteryLevel>4</BatteryLevel>
// <simlockStatus>0</simlockStatus>
// <PrimaryDns>139.7.30.126</PrimaryDns>
// <SecondaryDns>139.7.30.125</SecondaryDns>
// <PrimaryIPv6Dns></PrimaryIPv6Dns>
// <SecondaryIPv6Dns></SecondaryIPv6Dns>
// <TotalWifiUser>16</TotalWifiUser>
// <currenttotalwifiuser>16</currenttotalwifiuser>
// <ServiceStatus>2</ServiceStatus>
// <SimStatus>1</SimStatus>
// <WifiStatus>1</WifiStatus>
// <CurrentNetworkTypeEx>46</CurrentNetworkTypeEx>
// <WanPolicy>0</WanPolicy>
// <maxsignal>5</maxsignal>
// <wifiindooronly>0</wifiindooronly>
// <wififrequence>0</wififrequence>
// <classify>mobile-wifi</classify>
// <flymode>0</flymode>
// <cellroam>1</cellroam>
// <ltecastatus>0</ltecastatus>

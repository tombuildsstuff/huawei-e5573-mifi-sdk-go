## Go SDK for the Huawei E5573 Mifi

This is a bare-bones Go SDK for [the Huawei E5573 Mifi](https://consumer.huawei.com/in/support/mobile-broadband/e5573/).

## Example Usage

```
m := mifi.Mifi{
	Endpoint: "http://192.168.1.1",
}

err := m.ParseCookie()
if err != nil {
	return fmt.Errorf("Error obtaining authentication cookie for Mifi: %+v", err)
}

wifiSettings, err := m.WifiSettings()
if err != nil {
	return fmt.Errorf("Error getting Wifi Settings from the Mifi: %+v", err)
}

log.Printf("SSID: %q", wifiSettings.SSID)
log.Printf("Country: %q", wifiSettings.Country)

carrier, err := m.CarrierDetails()
if err != nil {
	return fmt.Errorf("Error getting Carrier Details from the Mifi: %+v", err)
}
log.Printf("Carrier: %q (%q / ID %d)", carrier.FullName, carrier.ShortName, carrier.CarrierID)

status, err := m.CurrentStatus()
if err != nil {
	return fmt.Errorf("Error getting Status from the Mifi: %+v", err)
}

log.Printf("Battery Percent: %d", status.CurrentBatteryPercentage)
log.Printf("Current Signal: %d / %d", status.CurrentSignalBars, status.MaxSignalBars)
log.Printf("Users Connected: %d", status.NumberOfUsersConnectedToWifi)

traffic, err := m.TrafficStatistics()
if err != nil {
	return fmt.Errorf("Error retrieving Traffic Statistics: %+v", err)
}

log.Printf("Connected for %d seconds..", traffic.SecondsConnectedToNetwork)
log.Printf("Total Downloaded: %.2f MB", traffic.DownloadedMB)
log.Printf("Total Uploaded %.2f MB", traffic.UploadedMB)
```

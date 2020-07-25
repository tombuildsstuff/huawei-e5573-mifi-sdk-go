package mifi_test

import (
	"net/http"

	"github.com/tombuildsstuff/huawei-e5573-mifi-sdk-go/mifi"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Status", func() {
	var (
		err    error
		myfi   mifi.Mifi
		status *mifi.Status
	)

	BeforeEach(func() {
		httpmock.RegisterResponder(
			"GET",
			"/html/index.html",
			httpmock.ResponderFromResponse(&http.Response{
				Header: http.Header{
					"Set-Cookie": []string{"SessionID=foobar42"},
				},
			}),
		)

		httpmock.RegisterResponder("GET", "/api/monitoring/status",
			httpmock.NewStringResponder(200, `<?xml version="1.0" encoding="UTF-8"?>
			<response>
			  <BatteryPercent>42</BatteryPercent>
			  <BatteryStatus>1</BatteryStatus>
				<CurrentWifiUser>11</CurrentWifiUser>
				<SignalIcon>5</SignalIcon>
			</response>
		`))

		myfi = mifi.Mifi{Endpoint: "http://192.168.8.1"}
	})

	JustBeforeEach(func() {
		status, err = myfi.CurrentStatus()
	})

	It("should not error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("has the battery percentage", func() {
		Expect(status.CurrentBatteryPercentage).To(Equal(42))
	})

	It("has the battery charging status", func() {
		Expect(status.BatteryCharging).To(Equal(true))
	})

	It("has the number of WiFi users", func() {
		Expect(status.NumberOfUsersConnectedToWifi).To(Equal(11))
	})

	It("has the number of signal bars", func() {
		Expect(status.CurrentSignalBars).To(Equal(5))
	})
})

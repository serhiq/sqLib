package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"path/filepath"
	"strings"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Testing Unexported methods
//___________________________________

func getTestDataPath() string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, ".testdata")
}

func printOutput(resp *resty.Response, err error) {
	fmt.Println(resp, err)
}

func RawResponse(resp *resty.Response) string {
	var s strings.Builder
	s.WriteString("Response Info:")
	s.WriteString("  Status     :" + resp.Status())
	s.WriteString("  Proto      :" + resp.Proto())
	s.WriteString("  Time       :" + resp.Time().String())
	s.WriteString("  Received At:" + resp.ReceivedAt().String())
	s.WriteString("  Body       \n" + resp.String())

	return s.String()
}

func checkLife() {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("http://localhost:7777")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func CreateAccessToken(token string) string {
	return "Bearer " + string(token)
}

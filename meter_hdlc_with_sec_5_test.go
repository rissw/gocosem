package gocosem

import (
	"fmt"
	"testing"
	"time"
)

func init_TestMeterHdlc_with_sec_5() {
	testMeterIp = "127.0.0.1"
	testHdlcResponseTimeout = time.Duration(1) * time.Hour
	testHdlcCosemWaitTime = time.Duration(5000) * time.Millisecond
	testHdlcSnrmTimeout = time.Duration(45) * time.Second
	testHdlcDiscTimeout = time.Duration(45) * time.Second
}

func TestMeterHdlc_with_sec_5_TcpConnect(t *testing.T) {
	init_TestMeterHdlc_with_sec_5()
	dconn, err := TcpConnect(testMeterIp, 4059)
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("transport connected")
	defer dconn.Close()
}

func TestMeterHdlc_with_sec_5_HdlcConnect(t *testing.T) {
	init_TestMeterHdlc_with_sec_5()
	physicalDeviceId := uint16(37)
	serverAddressLength := int(4)
	dconn, err := HdlcConnect(testMeterIp, 4059, 3, 1, &physicalDeviceId, &serverAddressLength, testHdlcResponseTimeout, &testHdlcCosemWaitTime, testHdlcSnrmTimeout, testHdlcDiscTimeout)
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("transport connected")
	defer dconn.Close()
}

/*
func TestMeterHdlc_with_sec_5_AppConnect_no_security(t *testing.T) {
	init_TestMeterHdlc_with_sec_5()
	physicalDeviceId := uint16(37)
	serverAddressLength := int(4)
	dconn, err := HdlcConnect(testMeterIp, 4059, 33, 1, &physicalDeviceId, &serverAddressLength, testHdlcResponseTimeout, &testHdlcCosemWaitTime, testHdlcSnrmTimeout, testHdlcDiscTimeout)
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("transport connected")
	defer dconn.Close()

	var aarq AARQapdu

	aarq.applicationContextName = tAsn1ObjectIdentifier([]uint32{2, 16, 756, 5, 8, 1, 1})

	var buf bytes.Buffer
	err := encode_AARQapdu(&buf, &aarq)
	if nil != err {
		t.Fatalf("encode_AARQapdu() failed")
	}
}
*/

func TestMeterHdlc_with_sec_5_AppConnect(t *testing.T) {
	init_TestMeterHdlc_with_sec_5()

	physicalDeviceId := uint16(37)
	serverAddressLength := int(4)
	dconn, err := HdlcConnect(testMeterIp, 4059, 3, 1, &physicalDeviceId, &serverAddressLength, testHdlcResponseTimeout, &testHdlcCosemWaitTime, testHdlcSnrmTimeout, testHdlcDiscTimeout)
	if nil != err {
		t.Fatal(err)
	}
	t.Logf("transport connected")
	defer dconn.Close()

	aconn, err := dconn.AppConnectWithSecurity5(3, 1, 0x30, []byte{0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDE, 0xDF}, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, []uint32{2, 16, 756, 5, 8, 1, 3}, []byte{0x4D, 0x45, 0x4C, 0x00, 0x00, 0x00, 0x00, 0x00}, "ZDXO2;66", []byte{0x21, 0x1F, 0x30, 0x00, 0x00, 0x00, 0x59, 0x36, 0x43, 0x91, 0x44, 0x1B, 0x6C, 0xE5, 0x3C, 0x29, 0x2A, 0x9D, 0x02, 0xD8, 0xDE, 0xA3, 0x76, 0xC9, 0xA2, 0xC6, 0x69, 0xCC, 0xD8, 0x1A, 0x8E, 0x69, 0x7F})
	//	aconn, err := dconn.AppConnectWithSecurity5(3, 1, 0x30, []byte{0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7, 0xD8, 0xD9, 0xDA, 0xDB, 0xDC, 0xDE, 0xDF}, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, []uint32{2, 16, 756, 5, 8, 1, 3}, []byte{0x4D, 0x45, 0x4C, 0x00, 0x00, 0x00, 0x00, 0x01}, ")HB+0F04", []byte{0x28, 0x1F, 0x30, 0x00, 0x00, 0x00, 0x2F, 0xF9, 0xF1, 0x4F, 0x54, 0x98, 0xBD, 0x2A, 0x0B, 0xB0, 0x00, 0x7F, 0xDB, 0x93, 0x18, 0xB7, 0x79, 0x77, 0x48, 0x5F, 0x54, 0xC4, 0xEE, 0x12, 0x10, 0x1B, 0xB1})
	if nil != err {
		t.Fatalf(fmt.Sprintf("%s\n", err))
	}
	t.Logf("application connected")
	defer aconn.Close()
}

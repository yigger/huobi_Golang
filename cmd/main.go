package main

import (
	"github.com/yigger/huobi_golang/cmd/accountclientexample"
	"github.com/yigger/huobi_golang/cmd/accountwebsocketclientexample"
	"github.com/yigger/huobi_golang/cmd/commonclientexample"
	"github.com/yigger/huobi_golang/cmd/crossmarginclientexample"
	"github.com/yigger/huobi_golang/cmd/etfclientexample"
	"github.com/yigger/huobi_golang/cmd/isolatedmarginclientexample"
	"github.com/yigger/huobi_golang/cmd/marketclientexample"
	"github.com/yigger/huobi_golang/cmd/marketwebsocketclientexample"
	"github.com/yigger/huobi_golang/cmd/orderclientexample"
	"github.com/yigger/huobi_golang/cmd/orderwebsocketclientexample"
	"github.com/yigger/huobi_golang/cmd/walletclientexample"
	"github.com/yigger/huobi_golang/logging/perflogger"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
	marketwebsocketclientexample.RunAllExamples()
	accountwebsocketclientexample.RunAllExamples()
	orderwebsocketclientexample.RunAllExamples()
}

// Run performance test
func runPerfTest() {
	perflogger.Enable(true)

	commonclientexample.RunAllExamples()
	accountclientexample.RunAllExamples()
	orderclientexample.RunAllExamples()
	marketclientexample.RunAllExamples()
	isolatedmarginclientexample.RunAllExamples()
	crossmarginclientexample.RunAllExamples()
	walletclientexample.RunAllExamples()
	etfclientexample.RunAllExamples()
}

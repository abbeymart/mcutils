// @Author: abbeymart | Abi Akindele | @Created: 2020-12-05 | @Updated: 2020-12-05
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcutils

import (
	"github.com/abbeymart/mcresponse"
	"github.com/abbeymart/mctest"
	"testing"
)

func TestStat(t *testing.T) {
	// test-data
	const msgType = "success"
	const msgType2 = "checkError"
	options := mcresponse.ResponseMessageOptions{
		Message: "",
		Value:   []string{"a", "b", "c"},
	}

	res := mcresponse.ResponseMessage{
		Code:       "success",
		ResCode:    200,
		ResMessage: "OK",
		Value:      "",
		Message:    "Request completed successfully",
	}
	//const unAuthMsg = "unAuthorized"

	mctest.McTest(mctest.OptionValue{
		Name: "should return success code for success-message",
		TestFunc: func() {
			req := mcresponse.GetResMessage(msgType, options)
			mctest.AssertEquals(t, req.Code, res.Code, "response-code should be: " + res.Code)
			mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	mctest.PostTestResult()
}


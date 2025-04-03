package model

import (
	"fmt"
)

type WangSuSDKError struct {
	Code      string
	Message   string
	RequestId string
}

func (e *WangSuSDKError) Error() string {
	if e.RequestId == "" {
		return fmt.Sprintf("[WangSuSDKError] Code=%s, Message=%s", e.Code, e.Message)
	}
	return fmt.Sprintf("[WangSuSDKError] Code=%s, Message=%s, RequestId=%s", e.Code, e.Message, e.RequestId)
}

func NewWangSuSDKError(code, message, requestId string) error {
	return &WangSuSDKError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (e *WangSuSDKError) GetCode() string {
	return e.Code
}

func (e *WangSuSDKError) GetMessage() string {
	return e.Message
}

func (e *WangSuSDKError) GetRequestId() string {
	return e.RequestId
}

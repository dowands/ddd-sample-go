package expection

import . "go-ddd/pkg/expection"

type expection struct {
	EMPTY Exception
	CAN_NOT_ASSIGN Exception
	NOT_FOUND Exception
}

var Expections = expection{
	EMPTY: Exception{},
	CAN_NOT_ASSIGN: Exception{
		Code:    100405,
		Message: "不允许分配客户",
	},
	NOT_FOUND: Exception{
		Code:    100404,
		Message: "客户不存在",
	},
}
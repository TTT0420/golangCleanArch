package pkg

const (
	OK               = "OK" //　正常時のmessageパラメーター
	NG               = "NG" //　失敗時のmessageパラメーター
	FAILED_ID        = -1   // 失敗時のIDパラメーター ex) { "message": "NG","id": -1 }
	INVALID_REQ_MSG  = "不正なリクエストです"
	SERVER_ERROR_MSG = "予期しないエラーが発生しました"
)

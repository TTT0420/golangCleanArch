package pkg

const (
	ResOK                   = "OK"                    //　正常時のmessageパラメーター
	ResNG                   = "NG"                    //　失敗時のmessageパラメーター
	ResMsg                  = "message"               //　レスポンス時のパラメーター ex) { "message": "OK" }
	ResErr                  = "error"                 //　レスポンス時のエラーパラメーター ex) { "error": "OK" }
	ResID                   = "id"                    // レスポンス時のIDパラメーター ex) { "message": "OK","id": 1 }
	ResPosts                = "posts"                 // レスポンス時の投稿パラメーター ex) { "message": "OK","posts": [ {entity.Post} ] }
	ResMsgInternalServerErr = "Internal server error" // 予期しないエラーの場合のメッセージ
	FailedID                = -1                      // 失敗時のIDパラメーター ex) { "message": "NG","id": -1 }
)

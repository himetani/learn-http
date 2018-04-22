# Memo
- testしたいのは、dispatchされるhandlerの挙動が正しいかどうか
- input(=request) -> output(=ResponseRecorder型)
- 実際のサーバ(=TCPサーバ)を立てる必要はなくて、requestに対するresponseを検証できればよい
- responseはhttp.ResponseWriter interfaceを満たせばよいので、ResponseRecorder型をhandlerに対してDIすればテストできる

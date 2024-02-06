package result

type Result struct {
	code int
	msg  string
	data any
}

func Ok(data any) Result {
	return Result{code: 200, msg: "ok", data: data}
}

func Error(code int, msg string) Result {
	return Result{code: code, msg: msg, data: nil}
}

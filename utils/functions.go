package utils

type Res struct {
	Code int // struct 变量一定大写
	Msg string
	Data interface{} // FIXME 这个返回是通用类型吗
}

func ReturnData(code int,msg string,data interface{}) Res {
	var rs Res
	rs.Msg = msg
	rs.Data = data
	rs.Code = code
	return rs
}

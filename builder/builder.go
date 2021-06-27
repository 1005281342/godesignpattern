package builder

// 用于创建参数复杂的对象
// Go 常用的参数传递方法
// 建造者模式

type Student struct {
	Age     int32
	CodeAge int32
	Name    string
}

type StudentInfosOptFunc func(option *Student)

func NewStudent(opts ...StudentInfosOptFunc) *Student {
	var s = &Student{}
	for _, f := range opts {
		f(s)
	}
	return s
}

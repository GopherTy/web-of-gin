package module

// IRegister 每个模块要实现该接口
type IRegister interface {
	// 模块注册
	Regist()
}

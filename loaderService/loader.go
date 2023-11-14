package loaderService

// Loader loader需要实现的若干功能
// loader是一个加载本地函数的服务 它有若干个不同的实现
type Loader interface {
	// LoadHexPackage 根据路径加载二进制包并返回句柄
	// 传入：路径
	// 传出：二进制执行包
	LoadHexPackage(path string) (*HexPackage, error)
	// ReleasePackage 释放dll包
	// 传入：二进制执行包
	// 传出：无
	ReleasePackage(hexPackage *HexPackage) error
}

// HexPackage 二进制可执行包
/*
每个包都有一个全局唯一的name
而每个函数也是全局唯一的 其通过name和function name来断定
函数调用直接传入和传出指针数组 其没有显式的类型检查
*/
type HexPackage interface {
	// GetName 获取名字
	// 传入：无
	// 传出：名字
	GetName() string
	// GetID 获取ID
	// 传入：无
	// 传出：ID
	GetID() int
	// GetFunctions 获取支持的函数列表
	// 传入：无
	// 传出：支持的函数列表
	GetFunctions() []string
	// GetInfo 获取别的信息
	// 传入：key
	// 传出：value
	GetInfo(key string) string

	// Execute 执行函数
	// 传入：方法名，参数
	// 传出：返回值
	Execute(method string, args []uintptr) ([]uintptr, error)
}

// HexInfo 二进制包信息j
type HexInfo struct {
	// 支持的函数名
	Functions []string
	// info
	Info map[string]string
}

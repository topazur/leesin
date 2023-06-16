package cmd

/**
// 参数可绑定变量，因此没有返回值
StringVar 用于定义一个只有短选项（如 -n）的字符串类型命令行参数，同时它也会把参数绑定到一个变量上；
StringVarP 用于定义一个既有短选项（如 -n）又有长选项（如 --name）的字符串类型命令行参数，同时它也会把参数绑定到一个变量上；

// 有返回值，可自定义复制实现绑定变量，针对不同类型参数更加灵活
String 用于定义一个只有短选项（如 -n）的字符串类型命令行参数，但是不会绑定到任何变量上；
StringP 用于定义一个既有短选项（如 -n）又有长选项（如 --name）的字符串类型命令行参数，但是不会绑定到任何变量上。
*/

/**
// eg: 短选项（如 -n）、长选项（如 --name）
leesin server --config=test
leesin server --config test
leesin server -c=test
leesin server -c test
*/

/**
PersistentFlags 则表示一个命令/子命令的全局可用标志，即这些标志可以被该命令及其所有子命令所使用。
与 Flags 相比，PersistentFlags 的主要作用是定义全局模式标志，例如调试模式、静默模式等。
*/

// customFlag 该结构体用以传入pflag的参数，统一变量定义，不会造成变量定义混乱冗余
type customFlag struct {
	name      string
	shorthand string
	value     interface{}
	usage     string
}

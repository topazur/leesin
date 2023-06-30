package dao

// QuerierMerge 整合 满足Querier接口的结构体方法 和 自定义的结构体方法 成一个新接口。
// 在Go语言中，把方法绑定到结构体上，也就是给结构体定义方法的过程，称之为“结构体方法”（Struct Method）。
type QuerierMerge interface {
	Querier
	// Add custom queries at 👇🏼
}

var _ QuerierMerge = (*Queries)(nil)

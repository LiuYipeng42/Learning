package memo1

import "sync"

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}
 
// 使用 重复抑制的方式解决数据竞态和重复的执行
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 初始化 e 指针，初始化过程较快，加上锁也不会对速度有太大影响。
		// 初始化完成之后，e 指针不为 nil，即使 e 指针指向的数据还为空，所以，
		// 若在运行 get 函数发起请求时，有同样 key 的 goroutine 被运行，就会进入 else，
		// 然后等待 e.ready 通道被关闭，即意味着请求已经完毕，然后就可以返回数据，
		// 而不是再次请求同样的 url 链接
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready)
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()

		// 通道被关闭时阻塞会停止，此时 e 指针上的数据就已经被初始化完毕
		<-e.ready
	}
	return e.res.value, e.res.err
}

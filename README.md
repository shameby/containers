### 一个Go的并发安全基本容器组

### 使用

```
import ds "github.com/shameby/containers"
```

### 支持的数据结构有

- Queue 队列(包括lockFree队列)
- Set 集合
- Stack 栈
- DequeArr 底层为数组的双端队列
- DequeList 底层为链表的双端队列
- BinaryTree 二叉树
- PriorityQueue 底层为堆的优先队列

### Set
初始化

```
NewSet(locker) Set
```

初始化参数

|参数名|类型|说明|
|:---|:----- |----|
|locker|RWLocker|locker为实现了接口RWLocker类型(见文底)的引用，如果不需要并发，传nil|

### RWLocker

```
type RWLocker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}
```
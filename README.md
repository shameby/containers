### 一个Go的并发安全基本容器组

### 使用

```
import ds "github.com/shameby/containers"
```

### 支持的数据结构有

- Set 集合
- Stack 栈
- DequeArr 底层为数组的双端队列
- DequeList 底层为链表的双端队列
- BinaryTree 二叉树

### Set
初始化

```
NewSet(isConcurrency bool) Set
```

初始化参数

|参数名|类型|说明|
|:---|:----- |----|
|isConcurrency|bool|true为并发安全，反之则不是|

>觉得数据结构和算法难的你并不孤独

这里会记录作为一个小白学习数据结构的过程，会分不同的阶段来学。有时候不必死磕，慢慢来，比较快

>本部分主要是数据结构相关的内容，算法的内容见：

#### 目标
不是要成为算法大神、参加算法竞赛之类的
* 能在实际的业务当中使用到
* 就算是大厂的算法题，也能比较轻松搞定
* 能看懂一些源码中所使用到的数据结构和算法

#### 学习方式
* 分四个不同的阶段来学习
* 看经典的数据结构和算法书、看专栏、看高质量的文章
* 刷题（并不是大量的盲目刷题，刷真实的场景题及经典题）
* 如果有现阶段也打算学习数据结构和算法的兄弟，一起讨论学习是很不错的

#### 第一阶段
该阶段涉及到的数据结构必须达到可以轻松实现。这些都是后边复杂数据结构及算法的基础

* [时间复杂度和空间复杂度](https://mp.weixin.qq.com/s?__biz=MzU5MjA1MzcyMA==&mid=2247484405&idx=1&sn=30f2ae319ca0e7b10e821f6cd76bb5f0&chksm=fe24d742c9535e54316013decdc9ea9da239b816d372d270f8eb1b7c8ca3770fd5a0dcd40275&token=1596559759&lang=zh_CN#rd)
   * 时间复杂度分析
   * 空间复杂度分析
* [数组](https://mp.weixin.qq.com/s?__biz=MzU5MjA1MzcyMA==&mid=2247484982&idx=1&sn=1b7b2c50c5bce6c4eaaf04b9d7d7ce50&chksm=fe24d281c9535b97f34a68dbd8875d9f58bad917a7f21e114f78486e76d5a18a903d254be8a0&token=1596559759&lang=zh_CN#rd)
   * 数据移动
   * 数组扩容（数据搬移）
* [栈](https://mp.weixin.qq.com/s/-qFPqV34Go47_nJ7dw0eTw)
   * [顺序栈](https://github.com/Rain-Life/data-structure/blob/master/Stack/arrayStack/arrayStack.go)
   * [链式栈](https://github.com/Rain-Life/data-structure/blob/master/Stack/linkListStack/linkListStack.go)
* [队列](https://mp.weixin.qq.com/s/-qFPqV34Go47_nJ7dw0eTw)
   * [顺序队列](https://github.com/Rain-Life/data-structure/blob/master/Queue/arrayQueue/arrayQueue.go)
   * [链式队列](https://github.com/Rain-Life/data-structure/blob/master/Queue/linkListQueue/linkListQueue.go)
   * [循环队列](https://github.com/Rain-Life/data-structure/blob/master/Queue/loopQueue/loopQueue.go)
* [链表](https://mp.weixin.qq.com/s/kILIFX22Djdil7hWoGMTOQ)
   * [单链表（增删改查、遍历）](https://github.com/Rain-Life/data-structure/blob/master/LinkList/baseList.go)
   * [循环链表](https://github.com/Rain-Life/data-structure/blob/master/LinkList/loopLinkList/loopLinkedList.go)
   * [双向链表](https://github.com/Rain-Life/data-structure/blob/master/LinkList/doublyLinkedList/doublyLinkedList.go)
   * [双向循环链表](https://github.com/Rain-Life/data-structure/blob/master/LinkList/doublyLoopLinkedList/doublyLoopLinkedList.go)
   * [单链表反转](https://mp.weixin.qq.com/s/GFDEOpF7tFcKPB8Q-UNdXA)
   * [循环链表环的检测](https://mp.weixin.qq.com/s/GFDEOpF7tFcKPB8Q-UNdXA)
   * [单链表合并](https://mp.weixin.qq.com/s/GFDEOpF7tFcKPB8Q-UNdXA)
   * [更多](https://mp.weixin.qq.com/s/6amum0iNRxgitXTPG47ExQ)
* [递归](https://mp.weixin.qq.com/s/ku9wPF3V4Y8SECw8eQxnfQ)
   * 递归思想
   * 递归时间复杂度分析 
* [排序](https://mp.weixin.qq.com/s/POoYx0E-lN2CwDOI8Rtdyw)
   * O(n^2)：[冒泡排序、插入排序、选择排序](https://mp.weixin.qq.com/s/POoYx0E-lN2CwDOI8Rtdyw)
   * O(nlogn)：[快速排序、归并排序](https://mp.weixin.qq.com/s/G84aBHEMa5sSr36HRLmBHQ)
   * O(n)：桶排序、计数排序、计数排序（跳过就行）
* [二分查找](https://mp.weixin.qq.com/s/JLw7m1fFGeu6j141q3zN5w)
   * [非递归实现](https://github.com/Rain-Life/data-structure/blob/master/BinarySearch/implement.go)
   * [递归实现](https://github.com/Rain-Life/data-structure/blob/master/BinarySearch/implement.go)
   * [递归实现查找第一个值等于给定值的元素](https://github.com/Rain-Life/data-structure/blob/master/BinarySearch/implement.go)

#### 第二阶段

* [跳表](https://mp.weixin.qq.com/s?__biz=MzU5MjA1MzcyMA==&mid=2247485195&idx=1&sn=d44e1cbf26cd9c53578f721b0db3f74e&chksm=fe24d3bcc9535aaab2e278edb4709a4c4e5d2da932a50b7f59c63121a072b4f9c078aa030239&token=1596559759&lang=zh_CN#rd)
   * 跳表基本操作
   * 跳表动态更新
* **[散列表](https://mp.weixin.qq.com/s/BpgaNMTH6bZxNZ0-IIQYzQ)**
   * [散列冲突](https://github.com/Rain-Life/data-structure/tree/master/hashTable)
   * 散列表扩容
   * [实现LRU](https://github.com/Rain-Life/data-structure/blob/master/hashTable/hash/LRU.go)
* **二叉树**
   * 二叉树遍历
   * 二叉查找树
   * 完全二叉树
   * 满二叉树
   * 平衡二叉树
   * 递归树
* **堆和堆排序**
   * 堆的基本操作
   * 堆排序
   * 堆应用（优先级队列、求TOP K）
* BF/RK字符串匹配算法
   * BF算法（Brute Force）
   * RK算法（Rabin-Karp）
* **Trie树**
   * Trie树实现（字典树）
   * Trie实现简单的搜索关键词提示功能
* **图的表示**
   * 图的存储方式（邻接矩阵、邻接表）
* **深度优先（BFS）和广度优先（BFS）**
   * 深度有限搜索（DFS）
   * 广度优先搜索（BFS）
   * 两种搜索算法的应用（迷宫）

#### 第三阶段
这个阶段可以说是**最最最重要的部分**，也是最难的部分。有了前边两个阶段的基础，再看这里可能稍微轻松一些

* 贪心算法
* 分治算法
* 回溯算法
* 动态规划

#### 第四阶段
* BM、KMP、AC自动机（字符串匹配）
   * BM算法（）
* 拓扑算法、最短路径（Dijkstra算法）、A*算法
* B+数
* 位图
* 红黑树
* 哈希算法

## LeetCode专题练习
具体代码，移步这里（README.md中有索引）：
* 数组
* 堆
* 栈
* 队列
* 链表
* 递归
* 排序
* 哈希表
* 树
   * 二叉树
   * 二叉搜索树
   * 字典树
   * 线段树
* 查找算法
   * 二分查找
   * 字符串查找
   * 广度优先搜索
   * 深度优先搜索
* 字符串
* 图
* 分治算法
* 回溯算法
* 贪心算法
* 动态规划


每一个数据结构都有沉淀文章，如果对你有帮助，别忘了给个star和关注哦

![image](https://github.com/Rain-Life/data-structure/blob/master/photos/%E5%85%AC%E4%BC%97%E5%8F%B7.jpg)

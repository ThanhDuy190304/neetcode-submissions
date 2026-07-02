type node struct {
    key, value int
    prev, next *node
}

type LRUCache struct {
    capacity   int
    cache      map[int]*node
    head, tail *node
}

func Constructor(capacity int) LRUCache {
    head := &node{}
    tail := &node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*node),
        head:     head,
        tail:     tail,
    }
}

func (l *LRUCache) remove(n *node) {
    n.prev.next = n.next
    n.next.prev = n.prev
}

func (l *LRUCache) insertFront(n *node) {
    n.next = l.head.next
    n.prev = l.head
    l.head.next.prev = n
    l.head.next = n
}

func (l *LRUCache) Get(key int) int {
    n, ok := l.cache[key]
    if !ok {
        return -1
    }
    l.remove(n)
    l.insertFront(n)
    return n.value
}

func (l *LRUCache) Put(key int, value int) {
    if n, ok := l.cache[key]; ok {
        n.value = value
        l.remove(n)
        l.insertFront(n)
        return
    }
    if len(l.cache) == l.capacity {
        lru := l.tail.prev
        l.remove(lru)
        delete(l.cache, lru.key)
    }
    n := &node{key: key, value: value}
    l.cache[key] = n
    l.insertFront(n)
}
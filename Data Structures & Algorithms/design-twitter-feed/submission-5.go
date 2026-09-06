
type Tweet struct {
	tweetId   int
	timestamp int
}

type User struct {
	userId    int
	following map[int]bool
	tweets    []Tweet
}

type Twitter struct {
	users     map[int]*User
	timestamp int
}

func Constructor() Twitter {
	return Twitter{
		users:     make(map[int]*User),
		timestamp: 0,
	}
}

func (t *Twitter) getOrCreateUser(userId int) *User {
	u, ok := t.users[userId]
	if !ok {
		u = &User{
			userId:    userId,
			following: map[int]bool{userId: true},
			tweets:    []Tweet{},
		}
		t.users[userId] = u
	}
	return u
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	u := t.getOrCreateUser(userId)
	t.timestamp++
	u.tweets = append(u.tweets, Tweet{tweetId: tweetId, timestamp: t.timestamp})
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	follower := t.getOrCreateUser(followerId)
	t.getOrCreateUser(followeeId)
	follower.following[followeeId] = true
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower := t.getOrCreateUser(followerId)
	delete(follower.following, followeeId)
}

type mergeEntry struct {
	tweetId   int
	timestamp int
	userIdx   int
	pos       int
}

type MaxHeapByTime []mergeEntry

func (h MaxHeapByTime) Len() int            { return len(h) }
func (h MaxHeapByTime) Less(i, j int) bool  { return h[i].timestamp > h[j].timestamp }
func (h MaxHeapByTime) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeapByTime) Push(x interface{}) { *h = append(*h, x.(mergeEntry)) }
func (h *MaxHeapByTime) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	u := t.getOrCreateUser(userId)
	h := &MaxHeapByTime{}

	followees := []int{}
	for id := range u.following {
		followees = append(followees, id)
	}
	for _, id := range followees {
		tw := t.users[id].tweets
		if len(tw) > 0 {
			heap.Push(h, mergeEntry{tweetId: tw[len(tw)-1].tweetId, timestamp: tw[len(tw)-1].timestamp, userIdx: id, pos: len(tw) - 1})
		}
	}

	result := []int{}
	for h.Len() > 0 && len(result) < 10 {
		e := heap.Pop(h).(mergeEntry)
		result = append(result, e.tweetId)
		if e.pos > 0 {
			tw := t.users[e.userIdx].tweets
			heap.Push(h, mergeEntry{tweetId: tw[e.pos-1].tweetId, timestamp: tw[e.pos-1].timestamp, userIdx: e.userIdx, pos: e.pos - 1})
		}
	}
	return result
}

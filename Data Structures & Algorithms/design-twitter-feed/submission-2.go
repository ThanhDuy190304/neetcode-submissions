
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

type entry struct {
	tweetId   int
	timestamp int
}

type MaxHeapByTime []entry

func (h MaxHeapByTime) Len() int            { return len(h) }
func (h MaxHeapByTime) Less(i, j int) bool  { return h[i].timestamp > h[j].timestamp }
func (h MaxHeapByTime) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeapByTime) Push(x interface{}) { *h = append(*h, x.(entry)) }
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

	for followeeId := range u.following {
		followee := t.users[followeeId]
		tweets := followee.tweets
		count := 0
		for i := len(tweets) - 1; i >= 0 && count < 10; i-- {
			heap.Push(h, entry{tweetId: tweets[i].tweetId, timestamp: tweets[i].timestamp})
			count++
		}
	}

	result := []int{}
	for h.Len() > 0 && len(result) < 10 {
		e := heap.Pop(h).(entry)
		result = append(result, e.tweetId)
	}
	return result
}
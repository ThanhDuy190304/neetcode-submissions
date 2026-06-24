type entry struct {
    value     string
    timestamp int
}

type TimeMap struct {
    store map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]entry)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    if val, exists := this.store[key]; exists {
        if len(val) == 0 || val[len(val)-1].value != value {
            this.store[key] = append(val, entry{value, timestamp})
        }
    } else {
        this.store[key] = append(this.store[key], entry{value, timestamp})
    }
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if val, exists := this.store[key]; exists{
		l, r := 0, len(val) - 1
		for l < r{
			mid := (l + r + 1) / 2
			if timestamp > val[mid].timestamp{
				l = mid
			}else if timestamp < val[mid].timestamp{
				r = mid - 1
			}else{	
				return val[mid].value
			}
		}
		if timestamp >= val[l].timestamp{
			return val[l].value
		}
	}
	return ""
}

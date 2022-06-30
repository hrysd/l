type RecentCounter struct {
	Requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Requests: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	start := t - 3000
	end := t
	count := 0
	this.Requests = append(this.Requests, t)
	for _, r := range this.Requests {
		if start <= r && r <= end {
			count += 1
		}
	}
	return count
}

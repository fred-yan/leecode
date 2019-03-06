package recentCounter

type RecentCounter struct {
	list []int
}


func Constructor() RecentCounter {
	return RecentCounter{
		list: []int{},
	}

}

func (this *RecentCounter) Ping(t int) int {
	this.list = append(this.list, t)

	for (t - this.list[0]) > 3000 {
		this.list = this.list[1:]
	}
	return len(this.list)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
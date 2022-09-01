type MyHashSet struct {
	M map[int]bool
}

func Constructor() MyHashSet {
	a := MyHashSet{}
	a.M = map[int]bool{}

	return a
}

func (this *MyHashSet) Add(key int) {
	this.M[key] = true
}

func (this *MyHashSet) Remove(key int) {
	delete(this.M, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.M[key]

	return ok
}

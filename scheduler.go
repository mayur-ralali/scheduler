package scheduler

// Scheduler model
type Scheduler struct {
	batchSize int
	hash      map[int]*Val
}

// Val contains the data scheduler store
type Val struct {
	metadata []Meta
	c        int
}

// Meta type of data is stored by scheduler
type Meta interface {
	GetSeq() int
}

// NewScheduler will create new scheduler
func NewScheduler(batchSize int) *Scheduler {
	return &Scheduler{
		batchSize: batchSize,
		hash:      make(map[int]*Val),
	}
}

// Process will process the given input
// It will return data only if all seq are arrived of that batch
func (s *Scheduler) Process(ip Meta) (data []Meta) {
	k := s.getKey(ip)
	if _, exists := s.hash[k]; !exists {
		s.hash[k] = newVal(s.batchSize)
	}
	s.hash[k].add(ip)
	data = s.batchProcess(k)
	return
}

// batchProcess will process only if given key hash if full
// i.e. all seq are available in that hash map
func (s *Scheduler) batchProcess(k int) (data []Meta) {
	if s.hash[k].c == s.batchSize {
		data = s.hash[k].metadata
		delete(s.hash, k)
	}
	return
}

// getKey will return the hashed key for given input
func (s *Scheduler) getKey(ip Meta) int {
	return int(ip.GetSeq() / s.batchSize)
}

// newVal will create new value object
func newVal(size int) *Val {
	return &Val{
		metadata: make([]Meta, size),
	}
}

// add new element in metadata
func (v *Val) add(i Meta) {
	v.c++
	v.metadata[i.GetSeq()%len(v.metadata)] = i
}

package scheduler

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Sample input
type Sample struct {
	ID   int
	Seq  int
	Data string
}

//GetSeq will implement Meta interface
func (s Sample) GetSeq() int {
	return s.Seq
}

func TestSimpleData(t *testing.T) {

	size := 50000
	batch := 100
	input := getSampleInput(size)
	s := NewScheduler(batch)

	assert := assert.New(t)
	for _, ip := range input {
		if data := s.Process(ip); data != nil {
			assert.Equal(len(data), batch, "Invalid batch size")
			assert.True(isSorted(data), fmt.Sprintf("data should be sorted %v", getSeq(data)))
			fmt.Println(seq)
		}
	}
}

func isSorted(data []Meta) (sorted bool) {
	if len(data) > 0 {
		prev := data[0].GetSeq()
		sorted = true
		for i := 1; i < len(data); i++ {
			if prev+1 != data[i].GetSeq() {
				sorted = false
				break
			}
			prev = data[i].GetSeq()
		}
	}
	return
}

func getSeq(data []Meta) (seq []int) {
	for _, v := range data {
		seq = append(seq, v.GetSeq())
	}
	return
}

func getSampleInput(size int) (input []Sample) {
	rand.Seed(time.Now().Unix())
	input = make([]Sample, size)
	vis := make(map[int]struct{})
	for i := 0; i < size; i++ {
		no := getRandNo(vis, size)
		input[i] = Sample{ID: no, Seq: no, Data: fmt.Sprintf("Sampel %v data", no)}
	}
	return
}

func getRandNo(vis map[int]struct{}, size int) (no int) {
	for {
		no = rand.Intn(size)
		if _, exists := vis[no]; !exists {
			vis[no] = struct{}{}
			break
		}
	}
	return
}

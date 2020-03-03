# scheduler
batch process scheduler


### TO test
make test

### Algo:
I have maintainted an array of each batch and as once all the sequece of a batch is received, It will return the current batch.  

You need to implement the Meta interface with you struct.

### Example

```
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

func main() {
    batch := 100
    s := NewScheduler(batch)
    for _, ip := range getInput() {
        if data := s.Process(ip); data != nil {
            fmt.Println(data)
        }
    }
}
```
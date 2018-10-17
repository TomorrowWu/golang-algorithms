package lock_free_queue

//const dataLen = 10000
//
//var q *Queue
//var data []string
//var mdata map[string]bool
//
//func isIn(data []string, datum string) (found bool) {
//	for _, d := range data {
//		if d == datum {
//			found = true
//			return
//		}
//	}
//
//	return
//}
//
//func TestInit(t *testing.T) {
//	mdata = make(map[string]bool)
//}
//
//func TestSingleGoroutine(t *testing.T) {
//	q = NewQueue()
//
//	v1 := "scryner"
//	v2 := "skdul"
//
//	q.Enqueue(v1)
//	q.Enqueue(v2)
//
//	w1, ok := q.Dequeue()
//	if !ok || w1 != v1 {
//		t.Errorf("not same: %v, %v", v1, w1)
//
//	}
//
//	w2, ok := q.Dequeue()
//	if !ok || w2 != v2 {
//		t.Errorf("not same: %v, %v", v2, w2)
//	}
//}
//
//func TestConcurrentWrite(t *testing.T) {
//
//	for i := 0; i < dataLen; i++ {
//		datum := fmt.Sprintf("data_%d", i)
//		data = append(data, datum)
//	}
//
//	var wg sync.WaitGroup
//
//	for i := 0; i < dataLen; i++ {
//		datum := data[i]
//
//		wg.Add(1)
//
//		go func() {
//			q.Enqueue(datum)
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}
//
//type result struct {
//	retval string
//}
//
//func TestConcurrentRead(t *testing.T) {
//	var wg sync.WaitGroup
//
//	var results []*result
//	for i := 0; i < dataLen; i++ {
//		ret := new(result)
//		results = append(results, ret)
//
//		wg.Add(1)
//
//		go func(ret *result) {
//			datum, _ := q.Dequeue()
//			datum2 := datum.(string)
//			ret.retval = datum2
//
//			wg.Done()
//		}(ret)
//	}
//
//	wg.Wait()
//
//	for i := 0; i < dataLen; i++ {
//		ret := results[i]
//
//		if !isIn(data, ret.retval) {
//			t.Errorf("datum is not in data: %v", ret.retval)
//			return
//		}
//
//		if _, ok := mdata[ret.retval]; ok {
//			t.Errorf("redundant retrieval: %v", ret.retval)
//			return
//		}
//
//		mdata[ret.retval] = true
//	}
//
//	_, ok := q.Dequeue()
//	if ok {
//		t.Errorf("queue must be empty")
//		return
//	}
//}
//
//func TestConcurrentReadWrite(t *testing.T) {
//	succEnq := make(chan int)
//	succDeq := make(chan int)
//
//	for i := 0; i < dataLen*2; i++ {
//		go func(i int) {
//			if i%2 == 0 {
//				// enqueue
//				datum := data[i/2]
//				q.Enqueue(datum)
//				succEnq <- 1
//
//				if (i/2)+1 == dataLen {
//					succEnq <- -1
//				}
//			} else {
//				// dequeue
//				_, ok := q.Dequeue()
//				if ok {
//					succDeq <- 1
//				} else {
//					succDeq <- 0
//				}
//
//				if (i/2)+1 == dataLen {
//					succDeq <- -1
//				}
//			}
//		}(i)
//	}
//
//	var enqSuccess, deqSuccess int
//	var endEnq, endDeq bool
//
//	for {
//		if endEnq && endDeq {
//			break
//		}
//
//		select {
//		case i := <-succEnq:
//			if i < 0 {
//				endEnq = true
//			} else {
//				enqSuccess += i
//			}
//		case i := <-succDeq:
//			if i < 0 {
//				endDeq = true
//			} else {
//				deqSuccess += i
//			}
//		}
//	}
//
//	if enqSuccess != dataLen {
//		t.Errorf("some %d enqueing operations is wrong", dataLen-enqSuccess)
//		return
//	}
//
//	retry := dataLen - deqSuccess
//
//	for i := 0; i < retry; i++ {
//		_, ok := q.Dequeue()
//		if !ok {
//			t.Errorf("retry dequeue failed: total %d retry but, dequeue has %d", retry, i)
//			return
//		}
//	}
//
//	if _, ok := q.Dequeue(); ok {
//		t.Errorf("queue must be empty")
//		return
//	}
//}
//
//func TestIter(t *testing.T) {
//	v1 := "test1"
//	v2 := "test2"
//
//	q.Enqueue(v1)
//	q.Enqueue(v2)
//
//	var data []interface{}
//
//	for datum := range q.Iter() {
//		data = append(data, datum)
//	}
//
//	for i, datum := range data {
//		datum2 := datum.(string)
//
//		var s string
//
//		switch i {
//		case 0:
//			s = "test1"
//		case 1:
//			s = "test2"
//		}
//
//		if s != datum2 {
//			t.Errorf("not matched: %v, %v", s, datum2)
//		}
//	}
//}
//
//func TestWatchIterator(t *testing.T) {
//	watchIterator := q.NewWatchIterator()
//	iter := watchIterator.Iter()
//	defer watchIterator.Close()
//
//	timeout := time.After(time.Second * 1)
//	var timeouted bool
//
//	select {
//	case <-iter:
//		timeouted = false
//	case <-timeout:
//		timeouted = true
//	}
//
//	if !timeouted {
//		t.Errorf("it must be timeouted, because queue may be empty")
//		return
//	}
//
//	var poped []int
//	var wg sync.WaitGroup
//
//	// consumer
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		until := time.After(time.Second * 1)
//
//		for {
//			select {
//			case <-until:
//				return
//			case v := <-iter:
//				poped = append(poped, v.(int))
//			}
//		}
//	}()
//
//	// producer
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		interval := time.Millisecond * 300
//
//		for i := 0; i < 5; i++ {
//			q.Enqueue(i)
//			time.Sleep(interval)
//		}
//	}()
//
//	wg.Wait()
//
//	for i := 0; i < 3; i++ {
//		if i != poped[i] {
//			t.Errorf("poped elements invalid: %d (expected %d)", poped[i], i)
//			return
//		}
//	}
//
//	timeout = time.After(time.Millisecond * 200)
//	select {
//	case <-timeout:
//	case v := <-iter:
//		poped = append(poped, v.(int))
//	}
//
//	for i := 0; i < 5; i++ {
//		if i != poped[i] {
//			t.Errorf("poped elements invalid: %d (expected %d)", poped[i], i)
//			return
//		}
//	}
//}

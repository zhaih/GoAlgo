package parallel

// basic parallel algorithms, scan and reduce
// restrict to int operation to avoid using []interface{}

const parallelThreshold = 2048

func Reduce(ary []int, f func(int, int) int, init int) int {
	if len(ary) == 0 {
		return init
	}
	var inner func([]int, *chan int)
	inner = func(ar []int, c *chan int) {
		if len(ar) <= parallelThreshold {
			r := init
			for _, v := range ar {
				r = f(r, v)
			}
			*c <- r
			return
		}
		nextc := make(chan int)
		go inner(ar[:len(ar)/2], &nextc)
		go inner(ar[len(ar)/2:], &nextc)
		fst := <-nextc
		snd := <-nextc
		*c <- f(fst, snd)
	}
	retc := make(chan int)
	go inner(ary, &retc)
	return <-retc
}

// modified from exercise of book Introduction to Algorithms
// slow because of too many recursive function calls
func Scan(ary []int, f func(int, int) int, init int) []int {
	if len(ary) == 0 {
		return []int{}
	}
	ret := make([]int, len(ary))
	copy(ret, ary)
	var reduce func([]int, *chan int)
	reduce = func(ar []int, c *chan int) {
		if len(ar) <= parallelThreshold {
			if len(ar) <= 1 {
				return
			}
			reduce(ar[:len(ar)/2], nil)
			reduce(ar[len(ar)/2:], nil)
			ar[len(ar)-1] = f(ar[len(ar)/2-1], ar[len(ar)-1])
			if c != nil {
				*c <- 1
			}
			return
		}
		nextc := make(chan int)
		go reduce(ar[:len(ar)/2], &nextc)
		go reduce(ar[len(ar)/2:], &nextc)
		<-nextc
		<-nextc
		ar[len(ar)-1] = f(ar[len(ar)/2-1], ar[len(ar)-1])
		*c <- 1
		return
	}
	fin := make(chan int)
	go reduce(ret, &fin)
	<-fin
	// first step done
	var sweep func(int, []int, *chan int)
	sweep = func(pad int, ar []int, c *chan int) {
		// invariant 1: pad should contain all cumulative result before ar's first element
		// invariant 2: ar's last element is always in final result (need no more calculation)
		if len(ar) <= parallelThreshold {
			if len(ar) == 2 {
				ar[0] = f(pad, ar[0])
				return
			} else if len(ar) == 1 {
				return
			}
			ar[len(ar)/2-1] = f(pad, ar[len(ar)/2-1])
			sweep(pad, ar[:len(ar)/2], nil)
			sweep(ar[len(ar)/2-1], ar[len(ar)/2:], nil)
			if c != nil {
				*c <- 0
			}
			return
		}
		ar[len(ar)/2-1] = f(pad, ar[len(ar)/2-1])
		go sweep(pad, ar[:len(ar)/2], c)
		go sweep(ar[len(ar)/2-1], ar[len(ar)/2:], c)
	}
	go sweep(init, ret, &fin)
	for i := 0; i < len(ret)/parallelThreshold+1; i++ {
		<-fin
	}
	return ret
}

// Similar idea as Scan, but use loop when under threshold
func Scan2(ary []int, f func(int, int) int, init int) []int {
	if len(ary) == 0 {
		return []int{}
	}
	ret := make([]int, len(ary))
	copy(ret, ary)
	var reduce func([]int, *chan int)
	reduce = func(ar []int, c *chan int) {
		if len(ar) <= parallelThreshold {
			for i := 1; i < len(ar); i++ {
				ar[i] = f(ar[i-1], ar[i])
			}
			*c <- 1
			return
		}
		nextc := make(chan int)
		go reduce(ar[:len(ar)/2], &nextc)
		go reduce(ar[len(ar)/2:], &nextc)
		<-nextc
		<-nextc
		ar[len(ar)-1] = f(ar[len(ar)/2-1], ar[len(ar)-1])
		*c <- 1
		return
	}
	fin := make(chan int)
	go reduce(ret, &fin)
	<-fin
	// first step done
	var sweep func(int, []int, *chan int)
	sweep = func(pad int, ar []int, c *chan int) {
		// invariant 1: pad should contain all cumulative result before ar's first element
		// invariant 2: ar's last element is always in final result (need no more calculation)
		// invariant 3: if len(ar) < threshold, then ar should be a local scanned array except the last element
		if len(ar) <= parallelThreshold {
			for i := 0; i < len(ar)-1; i++ {
				ar[i] = f(pad, ar[i])
			}
			*c <- 0
			return
		}
		ar[len(ar)/2-1] = f(pad, ar[len(ar)/2-1])
		go sweep(pad, ar[:len(ar)/2], c)
		go sweep(ar[len(ar)/2-1], ar[len(ar)/2:], c)
	}
	go sweep(init, ret, &fin)
	for i := 0; i < len(ret)/parallelThreshold+1; i++ {
		<-fin
	}
	return ret
}

func Scan3(ary []int, f func(int, int) int, init int, ps int) []int {
	if len(ary) == 0 {
		return []int{}
	}
	ret := make([]int, len(ary))
	copy(ret, ary)
	mid := make([]int, ps)
	fin := make(chan bool)
	var seg int
	seg = len(ary) / ps
	localScan := func(ar []int, p int) {
		for i := 1; i < len(ar); i++ {
			ar[i] = f(ar[i-1], ar[i])
		}
		mid[p] = ar[len(ar)-1]
		fin <- true
	}
	addUp := func(ar []int, p int) {
		if p == 0 {
			fin <- true
			return
		}
		for i := 0; i < len(ar); i++ {
			ar[i] = f(mid[p-1], ar[i])
		}
		fin <- true
	}
	for i := 0; i < ps; i++ {
		if i != ps-1 {
			go localScan(ret[i*seg:(i+1)*seg], i)
		} else {
			go localScan(ret[i*seg:], i)
		}
	}
	// barrier
	for i := 0; i < ps; i++ {
		<-fin
	}
	for i := 1; i < ps; i++ {
		mid[i] = f(mid[i-1], mid[i])
	}
	for i := 0; i < ps; i++ {
		if i != ps-1 {
			go addUp(ret[i*seg:(i+1)*seg], i)
		} else {
			go addUp(ret[i*seg:], i)
		}
	}
	// barrier
	for i := 0; i < ps; i++ {
		<-fin
	}
	return ret
}

func Scan4(ary []int, f func(int, int) int, init int, ps int) []int {
	if len(ary) == 0 {
		return []int{}
	}
	ret := make([]int, len(ary))
	copy(ret, ary)
	mid := make([]int, ps)
	notification := make([]chan bool, ps)
	for i := 0; i < ps; i++ {
		notification[i] = make(chan bool)
	}
	fin := make(chan bool)
	var seg int
	seg = len(ary) / ps
	localScan := func(ar []int, p int) {
		for i := 1; i < len(ar); i++ {
			ar[i] = f(ar[i-1], ar[i])
		}
		if p != 0 {
			<-notification[p]
			mid[p] = f(mid[p-1], ar[len(ar)-1])
		} else {
			mid[p] = ar[len(ar)-1]
		}
		if p != ps-1 {
			notification[p+1] <- true
		}
		if p != 0 {
			for i := 0; i < len(ar); i++ {
				ar[i] = f(mid[p-1], ar[i])
			}
		}
		fin <- true
	}

	for i := 0; i < ps; i++ {
		if i != ps-1 {
			go localScan(ret[i*seg:(i+1)*seg], i)
		} else {
			go localScan(ret[i*seg:], i)
		}
	}
	// barrier
	for i := 0; i < ps; i++ {
		<-fin
	}
	return ret
}

package mergesort

//import (
//  "fmt"
//)

// Parallel multiway merge sort.

var S [][]int
var channels []chan []int

func parallelMultiwayMergesort(in []int, processor_count int) []int{
  out := in
  if len(out) <= processor_count {
    return Sort(out)
  }
  seq_len := len(out)/processor_count
  S = make([][]int, processor_count)
  channels = make([]chan []int, processor_count)
  for i:=0; i<processor_count; i++ {
    // Prepare payload
    for j:=0; j<seq_len; j++ {
      S[i] = append(S[i], out[seq_len*i + j])
    }
    // Launch subroutine to process the payload
    go func(k int, c* chan []int) {
      // Normal mergesort onto payload
      // then send result through channel.
      (*c) <- Sort(S[k])
    }(i, &channels[i])
  }
  // Payload sent to processors, now wait for
  // results.
  for i:=0; i<processor_count; i++ {
    S[i] = <-channels[i]
    for j:=seq_len*i; j<seq_len*(i+1); j++ {
      out[j] = S[i][j-seq_len*i]
    }
  }
  // Global merge of partitions
  return Sort(out)
}

func SortParallel(in []int) []int {
  return parallelMultiwayMergesort(in, 4)
}

package mergesort

import (
  "sync"
)

/**
* MergeSort(arr[], l, r)
*   If r > l
*     1. Find the middle point to divide the array into two halves.
*       middle m = l + (r-l)/2
*     2. Call mergeSort for the first half:
*       Call mergeSort(arr, l, m)
*     3. Call mergeSort for the second half:
*       Call mergeSort(arr, m+1, r)
*     4. Merge the two halves sorted in step 2 and 3.
*       Call merge (arr, l, m, r)
*/

var mu sync.Mutex
func merge(arr []int, left int, middle int, right int) {
  // Find sizes for each half.
  n1 := middle - left + 1
  n2 := right - middle

  // Create temporal arrays.
  var leftArr []int = []int{}
  var rightArr []int = []int{}

  // Copy data
  mu.Lock()
  defer mu.Unlock()
  for i:=0; i<n1; i++ {
    leftArr = append(leftArr, arr[i + left])
  }
  for j:=0; j<n2; j++ {
    rightArr = append(rightArr, arr[j+middle+1])
  }

  // Merge the temp arrays back into arr[l..r]
  il := 0 // Initial left subarray index
  ir := 0 // Initial right subarray index
  k  := left // Initial merged subarray index

  for (il < n1 && ir < n2) {
    if ( leftArr[il] <= rightArr[ir] ) {
      arr[k] = leftArr[il]
      il++
    } else {
      arr[k] = rightArr[ir]
      ir++
    }
    k++
  }

  // Copy the remaining elements of the left array, if any
  for il < n1 {
    arr[k] = leftArr[il]
    il++
    k++
  }
  // Copy the remaining elements of the right array, if any
  for ir < n2 {
    arr[k] = rightArr[ir]
    ir++
    k++
  }
}

func concurrentMerge(arr []int, left int, middle int, right int) []int{
  here := arr
  // Find sizes for each half.
  n1 := middle - left + 1
  n2 := right - middle

  // Create temporal arrays.
  var leftArr []int = []int{}
  var rightArr []int = []int{}

  // Copy data
  for i:=0; i<n1; i++ {
    leftArr = append(leftArr, here[i + left])
  }
  for j:=0; j<n2; j++ {
    rightArr = append(rightArr, here[j+middle+1])
  }

  // Merge the temp arrays back into arr[l..r]
  il := 0 // Initial left subarray index
  ir := 0 // Initial right subarray index
  k  := left // Initial merged subarray index

  for (il < n1 && ir < n2) {
    if ( leftArr[il] <= rightArr[ir] ) {
      here[k] = leftArr[il]
      il++
    } else {
      here[k] = rightArr[ir]
      ir++
    }
    k++
  }

  // Copy the remaining elements of the left array, if any
  for il < n1 {
    here[k] = leftArr[il]
    il++
    k++
  }
  // Copy the remaining elements of the right array, if any
  for ir < n2 {
    here[k] = rightArr[ir]
    ir++
    k++
  }
  return here;
}

func mergeSort(arr []int, l int, r int) {
  if l >= r {
    return
  }
  // Find middle point
  m := l + (r-l)/2
  mergeSort(arr,l,m)
  mergeSort(arr,m+1,r)
  merge(arr,l,m,r)
}

func concurrentMergeSort(arrr []int, l int, r int) []int {
  arr := arrr
  if l >= r {
    return arr
  }
  // Find middle point
  m := l + (r-l)/2
  channel1 := make(chan []int)
  channel2 := make(chan []int)
  go func(){
    arr = concurrentMergeSort(arr,l,m)
    channel1 <- arr
  }()
  go func(){
    arr := <-channel1
    arr = concurrentMergeSort(arr,m+1,r)
    channel2 <- arr
  }()
  arr = <-channel2
  arr = concurrentMerge(arr,l,m,r)
  return arr
}

// Sort an array
func Sort(in []int) []int {
  var out []int = in
  mergeSort(out, 0, len(out)-1)
  return out
}

func GoSort(in []int) []int {
  var out []int = in
  concurrentMergeSort(out, 0, len(out)-1)
  return out
}

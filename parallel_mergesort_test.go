package mergesort
/*
import (
  "testing"
  "reflect"
  "math/rand"
)

func TestParaSort1(t *testing.T) {
  want := []int{1}
  in := []int{1}
  if got := SortParallel(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestParaSort2(t *testing.T) {
  want := []int{1,2}
  in := []int{2,1}
  if got := SortParallel(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestParaSort3(t *testing.T) {
  want := []int{1,2,3,4}
  in := []int{2,3,1,4}
  if got := SortParallel(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestParaSort4(t *testing.T) {
  want := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  in := []int{2,3,1,4,12,5,6,10,9,8,11,7}
  if got := SortParallel(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}


func BenchmarkParallelSort(b *testing.B) {
  arr := [][]int{}
  out := [][]int{}
  // Fill 10 arrays with random integers
  for i:=0; i<1; i++ {
    arr = append(arr, []int{})
    for j:=0; j<b.N; j++ {
      arr[i] = append(arr[i], rand.Int())
    }
  }
  // Sort those 10 arrays.
  for i:=0; i<1; i++ {
    out = append(out, SortParallel(arr[i]))
  }
}*/

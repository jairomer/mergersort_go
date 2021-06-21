package mergesort

import (
  "testing"
  "fmt"
  "reflect"
  "math/rand"
)
// TODO: Put in README an image of a mergesort tree.
func toStr(in []int) string {
  out := "["
  first := true
  for _, i := range in {
    if first {
      first = false
      out += fmt.Sprint(i)
    } else {
      out += ", " + fmt.Sprint(i)
    }
  }
  out += "]"
  return out
}

func TesttoStr(t *testing.T) {
  in := []int{1,2,3}
  want := "[1, 2, 3]"

  if got := toStr(in); got != want {
    t.Errorf("toStr(): %s, want %s", got, want)
  }
}


func TestSort1(t *testing.T) {
  want := []int{1}
  in := []int{1}
  if got := Sort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("Sort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestSort2(t *testing.T) {
  want := []int{1,2}
  in := []int{2,1}
  if got := Sort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("Sort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestSort3(t *testing.T) {
  want := []int{1,2,3,4}
  in := []int{2,3,1,4}
  if got := Sort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("Sort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestSort4(t *testing.T) {
  want := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  in := []int{2,3,1,4,12,5,6,10,9,8,11,7}
  if got := Sort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("Sort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSort1(t *testing.T) {
  want := []int{1}
  in := []int{1}
  if got := GoSort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSort2(t *testing.T) {
  want := []int{1,2}
  in := []int{2,1}
  if got := GoSort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSort3(t *testing.T) {
  want := []int{1,2,3,4}
  in := []int{2,3,1,4}
  if got := GoSort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSort() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSort4(t *testing.T) {
  want := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  in := []int{2,3,1,4,12,5,6,10,9,8,11,7}
  if got := GoSort(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSort() = %s, want %s", toStr(got), toStr(want))
  }
}



func BenchmarkSort(b *testing.B) {
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
    out = append(out, Sort(arr[i]))
  }
}

func BenchmarkGoSort(b *testing.B) {
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
    out = append(out, GoSort(arr[i]))
  }
}

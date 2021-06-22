package mergesort

import (
  "testing"
  "fmt"
  "reflect"
  "math/rand"
)
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

func TestGoSortA1(t *testing.T) {
  want := []int{1}
  in := []int{1}
  if got := GoSortA(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortA2(t *testing.T) {
  want := []int{1,2}
  in := []int{2,1}
  if got := GoSortA(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortA3(t *testing.T) {
  want := []int{1,2,3,4}
  in := []int{2,3,1,4}
  if got := GoSortA(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortA4(t *testing.T) {
  want := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  in := []int{2,3,1,4,12,5,6,10,9,8,11,7}
  if got := GoSortA(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortA() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortB1(t *testing.T) {
  want := []int{1}
  in := []int{1}
  if got := GoSortB(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortB() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortB2(t *testing.T) {
  want := []int{1,2}
  in := []int{2,1}
  if got := GoSortB(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortB() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortB3(t *testing.T) {
  want := []int{1,2,3,4}
  in := []int{2,3,1,4}
  if got := GoSortB(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortB() = %s, want %s", toStr(got), toStr(want))
  }
}

func TestGoSortB4(t *testing.T) {
  want := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  in := []int{2,3,1,4,12,5,6,10,9,8,11,7}
  if got := GoSortB(in); !reflect.DeepEqual(want, got) {
    t.Errorf("GoSortB() = %s, want %s", toStr(got), toStr(want))
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

func BenchmarkGoSortA(b *testing.B) {
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
    out = append(out, GoSortA(arr[i]))
  }
}

func BenchmarkGoSortB(b *testing.B) {
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
    out = append(out, GoSortB(arr[i]))
  }
}

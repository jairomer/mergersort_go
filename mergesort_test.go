package mergesort

import (
  "testing"
  "fmt"
  "reflect"
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


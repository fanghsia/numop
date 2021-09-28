package numop

import "testing"

func TestAdd(t *testing.T){
  r := Add(1,2,3,4)
  if r != 10 {
    t.Errorf("Add 1, 2, 3, 4 = %v, want %v", r,10)
  }
}

package math

func MinInt( data... int) int {
  var min int
  for i, v := range data {
    if i == 0 {
      min = v
    }else{
      if v < min {
        min = v
      }
    }
  }
  return min
}

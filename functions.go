package main

import "math"

func first(x float64) float64{
  return x*x+math.Log(x)-(2*math.Cos(x))-1
}

func second(x float64) float64{
  return x*x*x-2*x-11
}

func third(x float64) float64{
  return x*x-15
}

func fourth(x float64) float64{
  return 1/(math.Sqrt(x))
}
func systemOne(y float64) float64{
  return -y+3
}
func systemTwo(x float64) float64{
  return math.Sqrt(-x*x+9)
}
func systemFour(x float64) float64{
  return x*x+1
}
func systemThree(y float64) float64{
  return y-1
}

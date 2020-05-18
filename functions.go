package main

import "math"

func first(x float64) float64{
  return /*x*x+math.Log10(x)-(2*math.Cos(x))-1*/math.Log(x)
}

func second(x float64) float64{
  return x*x*x-2*x-11
}

func third(x float64) float64{
  return x*x-15
}

func fourth(x float64) float64{
  return (1/(math.Sqrt(x)))-0.5
}
func systemOne(x float64) float64{
  return x*x+2
}
func systemTwo(x float64) float64{
  return math.Sin(x)
}
func systemFour(x float64) float64{
  return math.Sin(x)
}
func systemThree(x float64) float64{
  return -(x*x*x)+4
}

package main

import(
  "net/http"
  "html/template"
  "fmt"
  "math"
  "strconv"
)

var t = template.Must(template.ParseFiles("./template/main.html"))

func MainPage(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
    t.Execute(w,nil)
  } else if r.Method == "POST" {
    var method string = r.FormValue("type")
    function := r.FormValue("f")
    errorString := r.FormValue("error")
    leftString := r.FormValue("right")
    rightString := r.FormValue("left")
    var f func(float64) float64
    var fs []func(float64) float64
    error, left, right := formatData(errorString, leftString, rightString)

    if method == "secant" || method == "fixedPoint" {
      switch function {
      case "1":
        f = first
      case "2":
        f = second
      case "3":
        f = third
      case "4":
        f = fourth
      }
      var answer float64 = 0.0
      switch method {
      case "secant":
        answer = secant(error, left, right, f)
      case "fixedPoint":
        answer = fixedPoint(error, left, right, f)
      }
      fmt.Fprint(w,fmt.Sprintf("%f;0.0",answer))
    }else{
      switch function {
      case "5":
        fs = ([]func(float64) float64{systemOne,systemTwo})
      case "6":
        fs = []func(float64) float64{systemThree,systemFour}
      }
      x,y := fixedPointSys(error,left,right,fs)
      fmt.Fprint(w,fmt.Sprintf("%f;%f",x,y))
    }
  }
}

func main(){
  http.HandleFunc("/",MainPage)
  http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
  http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
  http.ListenAndServe(":1414",nil)
}

//works well as it even can be!
func secant(Uerror, a, b float64, f func(float64) float64) float64{
  i := 0
  error := math.Abs(a-b)
  results := make([]float64,1)
  results[0] = 0
  var x float64
  if f(a)*f(b)<0 {
    for error > Uerror {
      i++
      x = a - (f(a)/(f(b)-f(a)))*(b-a)
      if f(a)*f(x) < 0 {
        b = x
      }
      if f(x)*f(b) < 0 {
        a = x
      }
      results = append(results,x)
      error = math.Abs(results[i]-results[i-1])
    }
  }
  return x
}

func fixedPointSys(error, a, b float64, f []func(float64) float64) (float64, float64){
  var x0, y0, d1, d2 float64 = 0.2,0.2,1.0,1.0
  i := 0
  var x, y float64
  for {
    x,y = f[0](y0), f[1](x0)
    d1, d2 = f[0](x)-x, y-f[1](y)
    x0,y0,i = x,y, i+1
    if !(math.Abs(d1)>error && math.Abs(d2)>error && i < 10000) {break}
  }
  return x,y
}

//finally it's fucking works
func fixedPoint(error, a, b float64, f func(float64) float64) float64{
  n := 0
  _ = b
  var x,y,eps float64 = 0.0000000001,0.0000000001,1000
  x = a
  for eps >= error && n<1000 {
    y = x-0.003*f(x)
    eps = math.Abs(x-y)
    x = y
    n++
  }
  return x
}

func formatData(error, left, right string) (float64,float64,float64) {
  e, _ := strconv.ParseFloat(error,64)
  l, _ := strconv.ParseFloat(left,64)
  r, _ := strconv.ParseFloat(right,64)
  return e,l,r
}

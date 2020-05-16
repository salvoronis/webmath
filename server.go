package main

import(
  "net/http"
  "html/template"
  "fmt"
  "math"
  "strconv"
  //"os"
  "errors"
)

func MainPage(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET" {
    t := template.Must(template.ParseFiles("./template/main.html"))
    t.Execute(w,nil)
  } else if r.Method == "POST" {
    var method string = r.FormValue("type")
    function := r.FormValue("f")
    errorString := r.FormValue("error")
    leftString := r.FormValue("right")
    rightString := r.FormValue("left")
    var f func(float64) float64
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
    }
    var answer float64 = 0.0
    switch method {
    case "secant":
      answer = secant(error, left, right, f)
    case "fixedPoint":
      answer = fixedPoint(error, left, right, f)
    }
    fmt.Fprint(w,answer)
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

func pizdocPoint(x float64) float64{
  return x-0.003*second(x)
}

func suka(x float64) float64{
  return x-0.003*third(x)
}

func calculateError(a,b,e float64) (int, float64, error) {
  if e == 0.0 {
    return 0, 0, errors.New("Деление на 0 или символы")
  }
  floatN := (b-a)/math.Pow(e, 0.25)
  n := int(math.Ceil(floatN))
  newerr := math.Pow(((b-a)/float64(n)),4)
  return n, newerr, nil
}

func formatData(error, left, right string) (float64,float64,float64) {
  e, _ := strconv.ParseFloat(error,64)
  l, _ := strconv.ParseFloat(left,64)
  r, _ := strconv.ParseFloat(right,64)
  return e,l,r
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"math"
)

var tpl1 = template.Must(template.New("page1").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Вибір кабелів для живлення двотрансформаторної підстанції системи внутрішнього електропостачання підприємства напругою 10 кВ</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            width: 350px;
            text-align: center;
        }
        h2 {
            color: #333;
        }
        input[type="text"] {
            width: calc(100% - 20px);
            padding: 8px;
            margin: 5px 0;
            border: 1px solid #ccc;
            border-radius: 5px;
        }
        input[type="submit"] {
            background: #007BFF;
            color: white;
            padding: 10px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            width: 100%;
            margin-top: 10px;
        }
        input[type="submit"]:hover {
            background: #0056b3;
        }
        .results {
            margin-top: 20px;
            background: #e9ecef;
            padding: 10px;
            border-radius: 5px;
            text-align: left;
            white-space: pre-wrap;
        }
        .nav-buttons {
            margin-top: 20px;
        }
        .nav-buttons a {
            display: inline-block;
            margin-top: 10px;
            padding: 10px;
            background-color: #007BFF;
            color: white;
            text-decoration: none;
            border-radius: 5px;
        }
        .nav-buttons a:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Вибір кабелів для живлення двотрансформаторної підстанції системи внутрішнього електропостачання підприємства напругою 10 кВ</h2>
        <form method="post">
			<input type="text" name="amperage" placeholder="Струм КЗ" value="{{.Amperage}}">
			<input type="text" name="time" placeholder="Фіктивний час вимикання струму КЗ" value="{{.Time}}">
			<input type="text" name="density" placeholder="Економічна густина струму" value="{{.Density}}">
            <input type="submit" value="Розрахувати">
        </form>
       {{if .Result}}
			<div class="results">
				<h3>Результати:</h3>
				<pre>{{.Result}}</pre>
			</div>
		{{end}}
        <div class="nav-buttons">
            <a href="/page2">Перейти на другу сторінку</a>
        </div>
		<div class="nav-buttons">
            <a href="/page3">Перейти на третю сторінку</a>
        </div>
    </div>
</body>
</html>`))

var tpl2 = template.Must(template.New("page2").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Визначення струмів КЗ на шинах 10 кВ ГПП</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            width: 350px;
            text-align: center;
        }
        h2 {
            color: #333;
        }
        input[type="text"] {
            width: calc(100% - 20px);
            padding: 8px;
            margin: 5px 0;
            border: 1px solid #ccc;
            border-radius: 5px;
        }
        input[type="submit"] {
            background: #007BFF;
            color: white;
            padding: 10px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            width: 100%;
            margin-top: 10px;
        }
        input[type="submit"]:hover {
            background: #0056b3;
        }
        .results {
            margin-top: 20px;
            background: #e9ecef;
            padding: 10px;
            border-radius: 5px;
            text-align: left;
            white-space: pre-wrap;
        }
        .nav-buttons {
            margin-top: 20px;
        }
        .nav-buttons a {
            display: inline-block;
            margin-top: 10px;
            padding: 10px;
            background-color: #007BFF;
            color: white;
            text-decoration: none;
            border-radius: 5px;
        }
        .nav-buttons a:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Визначення струмів КЗ на шинах 10 кВ ГПП</h2>
        <form method="post">
			<input type="text" name="voltage" placeholder="" value="{{.Voltage}}">
			<input type="text" name="power" placeholder="" value="{{.Power}}">
			<input type="text" name="nomPower" placeholder="" value="{{.NomPower}}">
			<input type="text" name="basicPower" placeholder="" value="{{.BasicPower}}">
			<input type="submit" value="Розрахувати">
        </form>
       {{if .Result}}
			<div class="results">
				<h3>Результати:</h3>
				<pre>{{.Result}}</pre>
			</div>
		{{end}}
        <div class="nav-buttons">
            <a href="/page1">Перейти на першу сторінку</a>
        </div>
		<div class="nav-buttons">
            <a href="/page3">Перейти на третю сторінку</a>
        </div>
    </div>
</body>
</html>`))

var tpl3 = template.Must(template.New("page3").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Визначення струмів КЗ для підстанції Хмельницьких північних електричних мереж</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            width: 350px;
            text-align: center;
        }
        h2 {
            color: #333;
        }
        input[type="text"] {
            width: calc(100% - 20px);
            padding: 8px;
            margin: 5px 0;
            border: 1px solid #ccc;
            border-radius: 5px;
        }
        input[type="submit"] {
            background: #007BFF;
            color: white;
            padding: 10px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            width: 100%;
            margin-top: 10px;
        }
        input[type="submit"]:hover {
            background: #0056b3;
        }
        .results {
            margin-top: 20px;
            background: #e9ecef;
            padding: 10px;
            border-radius: 5px;
            text-align: left;
            white-space: pre-wrap;
        }
        .nav-buttons {
            margin-top: 20px;
        }
        .nav-buttons a {
            display: inline-block;
            margin-top: 10px;
            padding: 10px;
            background-color: #007BFF;
            color: white;
            text-decoration: none;
            border-radius: 5px;
        }
        .nav-buttons a:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Визначення струмів КЗ для підстанції Хмельницьких північних електричних мереж</h2>
        <form method="post">
			<input type="text" name="voltage" placeholder="" value="{{.Voltage}}">
			<input type="text" name="nomPower" placeholder="" value="{{.NomPower}}">
			<input type="submit" value="Розрахувати">
        </form>
       {{if .Result}}
			<div class="results">
				<h3>Результати:</h3>
				<pre>{{.Result}}</pre>
			</div>
		{{end}}
        <div class="nav-buttons">
            <a href="/page1">Перейти на першу сторінку</a>
        </div>
		<div class="nav-buttons">
            <a href="/page2">Перейти на другу сторінку</a>
        </div>
    </div>
</body>
</html>`))

func calculatePage1(amperage, time, density float64) string {
	result_val := math.Sqrt(time) * amperage / density

	result := fmt.Sprintf(`
Необхідне значення розміру перерізу
s ≥ s min: %.2f мм2
`, result_val)
	return result
}

type Data1 struct {
	Result   string
	Amperage float64
	Time     float64
	Density  float64
}

func handlerPage1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		a, _ := strconv.ParseFloat(r.FormValue("amperage"), 64)
		t, _ := strconv.ParseFloat(r.FormValue("time"), 64)
		d, _ := strconv.ParseFloat(r.FormValue("density"), 64)

		data := Data1{
			Result:   calculatePage1(a, t, d),
			Amperage: a,
			Time:     t,
			Density:  d,
		}

		tpl1.Execute(w, data)
		return
	}
	tpl1.Execute(w, nil)
}


func calculatePage2(voltage, power, nomPower, basicPower float64) string {
	xs := math.Pow(voltage, 2) / power
	xt := math.Pow(voltage, 3) / (100 * nomPower)
	xsum := xs + xt

	startAmperage := voltage / (math.Sqrt(3) *  xsum)
	basicAmperage := basicPower / (math.Sqrt(3) * voltage)

	result := fmt.Sprintf(`
Опори елементів заступної схеми:
Xс = %.2f Ом
Xт = %.2f Ом
Сумарний опір для точки К1:
XΣ = %.2f Ом
Початкове діюче значення струму трифазного КЗ:
Iп0 = %.2f кА
Базисне значення струму трифазного КЗ:
Iп0 = %.2f кА
	`, xs, xt, xsum, startAmperage, basicAmperage)
	return result
}

type Data2 struct {
	Result     string
	Voltage    float64
	Power      float64
	NomPower   float64
	BasicPower float64
}

func handlerPage2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
		p, _ := strconv.ParseFloat(r.FormValue("power"), 64)
		n, _ := strconv.ParseFloat(r.FormValue("nomPower"), 64)
		b, _ := strconv.ParseFloat(r.FormValue("basicPower"), 64)

		data := Data2{
			Result:      calculatePage2(v, p, n, b),
			Voltage:     v,
			Power:       p,
			NomPower:    n,
			BasicPower:  b,
		}

		tpl2.Execute(w, data)
		return
	}
	tpl2.Execute(w, nil)
}

func calculatePage3(voltage, nomPower float64) string {
	xt := 11.1 * math.Pow(voltage, 2) / (100 * nomPower)
	xsh := xt + 24.02
	zsh := math.Sqrt(math.Pow(10.65, 2) + math.Pow(xsh, 2))

	xshMin := xt + 65.68
	zshMin := math.Sqrt(math.Pow(34.88, 2) + math.Pow(xshMin, 2))

	i13 := voltage * 1000 / (math.Sqrt(3.0) * zsh)
	i12 := i13 * math.Sqrt(3.0) / 2
	i13Min := voltage * 1000 / (math.Sqrt(3.0) * zshMin)
	i12Min := i13Min * math.Sqrt(3.0) / 2


	var k float64 = math.Pow(11, 2) / math.Pow(115, 2)
	xsh *= k

	zsh = math.Sqrt(math.Pow(10.65 * k, 2) + math.Pow(xsh, 2))
	xshMin *= k
	zshMin = math.Sqrt(math.Pow(34.88 * k, 2) + math.Pow(xshMin, 2))

	i23 := 11 * 1000 / (math.Sqrt(3.0) * zsh)
	i22 := i23 * math.Sqrt(3.0) / 2
	i23Min := 11 * 1000 / (math.Sqrt(3.0) * zshMin)
	i22Min := i23Min * math.Sqrt(3.0) / 2

	rSum := 10.65 * k + 7.91
	xSum := xsh + 4.49
	zSum := math.Sqrt(math.Pow(rSum, 2) + math.Pow(xSum, 2))
	rSumMin := 34.88 * k + 7.91
	xSumMin := xshMin + 4.49
	zSumMin := math.Sqrt(math.Pow(rSumMin, 2) + math.Pow(xSumMin, 2))

	i33 := 11 * 1000 / (math.Sqrt(3.0) * zSum)
	i32 := i33 * math.Sqrt(3.0) / 2
	i33Min := 11 * 1000 / (math.Sqrt(3.0) * zSumMin)
	i32Min := i33Min * math.Sqrt(3.0) / 2

	result := fmt.Sprintf(`
Струм трифазного КЗ (I1-3)
Нормальний: %.2f A | Мінімальний: %.2f A

Струм двофазного КЗ (I1-2)
Нормальний: %.2f A | Мінімальний: %.2f A

Дійсний струм трифазного КЗ (I2-3)
Нормальний: %.2f A | Мінімальний: %.2f A

Дійсний струм двофазного КЗ (I2-2)
Нормальний: %.2f A | Мінімальний: %.2f A

Струм короткого замикання трифазного КЗ (I3-3)
Нормальний: %.2f A | Мінімальний: %.2f A

Струм короткого замикання двофазного КЗ (I3-2)
Нормальний: %.2f A | Мінімальний: %.2f A
	`, i13, i13Min, i12, i12Min, i23, i23Min, i22, i22Min, i33, i33Min, i32, i32Min)
	return result
}

type Data3 struct {
	Result     string
	Voltage    float64
	NomPower   float64
}

func handlerPage3(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		v, _ := strconv.ParseFloat(r.FormValue("voltage"), 64)
		n, _ := strconv.ParseFloat(r.FormValue("nomPower"), 64)

		data := Data2{
			Result:   calculatePage3(v, n),
			Voltage:  v,
			NomPower: n,
		}

		tpl3.Execute(w, data)
		return
	}
	tpl3.Execute(w, nil)
}

func main() {
	http.HandleFunc("/page1", handlerPage1)
	http.HandleFunc("/page2", handlerPage2)
	http.HandleFunc("/page3", handlerPage3)

	http.ListenAndServe(":8080", nil)
}

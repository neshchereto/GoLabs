package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Template for the first page
var tpl = template.Must(template.New("page1").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Калькулятор складу палива</title>
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
        <h2>Калькулятор складу палива</h2>
        <form method="post">
			<input type="text" name="hydrogen" placeholder="Водень (H)" value="{{.Hydrogen}}">
			<input type="text" name="carbon" placeholder="Вуглець (C)" value="{{.Carbon}}">
			<input type="text" name="sulfur" placeholder="Сірка (S)" value="{{.Sulfur}}">
			<input type="text" name="nitrogen" placeholder="Азот (N)" value="{{.Nitrogen}}">
			<input type="text" name="oxygen" placeholder="Кисень (O)" value="{{.Oxygen}}">
			<input type="text" name="ash" placeholder="Зола (A)" value="{{.Ash}}">
			<input type="text" name="wet" placeholder="Вологість (W)" value="{{.Wet}}">
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
    </div>
</body>
</html>`))

// Template for the second page (second calculator)
var tpl2 = template.Must(template.New("page2").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Калькулятор складу палива (Ванадій)</title>
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
        <h2>Калькулятор складу палива</h2>
        <form method="post">
            <input type="text" name="hydrogen" placeholder="Водень (H)" value="{{.Hydrogen}}">
            <input type="text" name="carbon" placeholder="Вуглець (C)" value="{{.Carbon}}">
            <input type="text" name="sulfur" placeholder="Сірка (S)" value="{{.Sulfur}}">
            <input type="text" name="oxygen" placeholder="Кисень (O)" value="{{.Oxygen}}">
            <input type="text" name="ash" placeholder="Зола (A)" value="{{.Ash}}">
            <input type="text" name="wet" placeholder="Вологість (W)" value="{{.Wet}}">
            <input type="text" name="vanadium" placeholder="Ванадій (V)" value="{{.Vanadium}}">
            <input type="text" name="q" placeholder="q" value="{{.Q}}">
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
    </div>
</body>
</html>`))

type FuelComposition struct {
	Result   string
	Hydrogen float64
	Carbon   float64
	Sulfur   float64
	Nitrogen float64
	Oxygen   float64
	Ash      float64
	Wet      float64
	Vanadium float64
	Q        float64
}

func calculateComposition(hydrogen, carbon, sulfur, nitrogen, oxygen, ash, wet float64) string {
	krs := 100 / (100 - wet)
	krg := 100 / (100 - wet - ash)

	hydrogenS := hydrogen * krs
	carbonS := carbon * krs
	sulfurS := sulfur * krs
	nitrogenS := nitrogen * krs
	oxygenS := oxygen * krs
	ashS := ash * krs

	hydrogenG := hydrogen * krg
	carbonG := carbon * krg
	sulfurG := sulfur * krg
	nitrogenG := nitrogen * krg
	oxygenG := oxygen * krg

	q := (339 * carbon + 1030 * hydrogen + 108.8 * (oxygen - sulfur) - 25 * wet) / 1000
	qs := (q + 0.025 * wet) * 100 / (100 - wet)
	qg := (q + 0.025 * wet) * 100 / (100 - wet - ash)

	result := fmt.Sprintf(`
Водень (H): %.2f %% (суха), %.2f %% (горюча)
Вуглець (C): %.2f %% (суха), %.2f %% (горюча)
Сірка (S): %.2f %% (суха), %.2f %% (горюча)
Азот (N): %.2f %% (суха), %.2f %% (горюча)
Кисень (O): %.2f %% (суха), %.2f %% (горюча)
Зола (A): %.2f %% (суха)

Нижча теплота згорання (робоча): %.2f МДж/кг
Нижча теплота згорання (суха): %.2f МДж/кг
Нижча теплота згорання (горюча): %.2f МДж/кг
`, hydrogenS, hydrogenG, carbonS, carbonG, sulfurS, sulfurG, nitrogenS, nitrogenG, oxygenS, oxygenG, ashS, q, qs, qg)

	return result
}

func handlerPage1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h, _ := strconv.ParseFloat(r.FormValue("hydrogen"), 64)
		c, _ := strconv.ParseFloat(r.FormValue("carbon"), 64)
		s, _ := strconv.ParseFloat(r.FormValue("sulfur"), 64)
		n, _ := strconv.ParseFloat(r.FormValue("nitrogen"), 64)
		o, _ := strconv.ParseFloat(r.FormValue("oxygen"), 64)
		a, _ := strconv.ParseFloat(r.FormValue("ash"), 64)
		wet, _ := strconv.ParseFloat(r.FormValue("wet"), 64)

		data := FuelComposition{
			Result:   calculateComposition(h, c, s, n, o, a, wet),
			Hydrogen: h,
			Carbon:   c,
			Sulfur:   s,
			Nitrogen: n,
			Oxygen:   o,
			Ash:      a,
			Wet:      wet,
		}

		tpl.Execute(w, data)
		return
	}
	tpl.Execute(w, nil)
}

func calculateCompositionPage2(hydrogen, carbon, sulfur, oxygen, ash, wet, vanadium, q float64) string {
	krs := (100 - wet) / 100
	krg := (100 - wet - ash) / 100

	hydrogenS := hydrogen * krg
	carbonS := carbon * krg
	sulfurS := sulfur * krg
	oxygenS := oxygen * krg
	ashS := ash * krs
	vanadiumS := vanadium * krs

	result := q * (100 - wet - ash) / 100 - 0.025 * wet

	return fmt.Sprintf(`
Водень (H): %.2f %% 
Вуглець (C): %.2f %% 
Сірка (S): %.2f %% 
Кисень (O): %.2f %% 
Зола (A): %.2f %% 
Ванадій (V): %.2f %% 

Нижча теплота згорання: %.2f МДж/кг`, hydrogenS, carbonS, sulfurS, oxygenS, ashS, vanadiumS, result)
}

func handlerPage2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h, _ := strconv.ParseFloat(r.FormValue("hydrogen"), 64)
		c, _ := strconv.ParseFloat(r.FormValue("carbon"), 64)
		s, _ := strconv.ParseFloat(r.FormValue("sulfur"), 64)
		o, _ := strconv.ParseFloat(r.FormValue("oxygen"), 64)
		a, _ := strconv.ParseFloat(r.FormValue("ash"), 64)
		wet, _ := strconv.ParseFloat(r.FormValue("wet"), 64)
		v, _ := strconv.ParseFloat(r.FormValue("vanadium"), 64)
		q, _ := strconv.ParseFloat(r.FormValue("q"), 64)

		data := FuelComposition{
			Result:   calculateCompositionPage2(h, c, s, o, a, wet, v, q),
			Hydrogen: h,
			Carbon:   c,
			Sulfur:   s,
			Oxygen:   o,
			Ash:      a,
			Wet:      wet,
			Vanadium: v,
			Q:        q,
		}

		tpl2.Execute(w, data)
		return
	}
	tpl2.Execute(w, nil)
}

func main() {
	http.HandleFunc("/page1", handlerPage1)
	http.HandleFunc("/page2", handlerPage2)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Розрахунку прибутку від сонячних електростанцій з встановленою системою прогнозування сонячної потужності</title>
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
        <h2>Розрахунок прибутку від сонячних електростанцій з встановленою системою прогнозування сонячної потужності</h2>
        <form method="post">
			<input type="text" name="power" placeholder="Середньодобова потужність" value="{{.Power}}">
			<input type="text" name="tarif" placeholder="Тариф на електроенергію" value="{{.Tarif}}">
            <input type="submit" value="Розрахувати">
        </form>
       {{if .Result}}
			<div class="results">
				<h3>Результати:</h3>
				<p>{{.Result}}</p>
			</div>
		{{end}}
    </div>
</body>
</html>`))

type Data struct {
	Result string
	Power  float64
	Tarif  float64
}

func calculate(power, tarif float64) string {
	loss   := 0.8 * power * 24 * tarif - 0.2 * power * 24 * tarif
	profit := 0.68 * power * 24 * tarif - 0.32 * power * 24 * tarif

	result := fmt.Sprintf(`
При СКВ в 1 МВт електростанція втрачатиме: %.2f тисяч грн.
При СКВ в 0.25 МВт електростанція зароблятиме:  %.2f тисяч грн.
`, loss, profit)

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		power, _ := strconv.ParseFloat(r.FormValue("power"), 64)
		tarif, _ := strconv.ParseFloat(r.FormValue("tarif"), 64)

		data := Data{
			Result: calculate(power, tarif),
			Power:  power,
			Tarif:  tarif,
		}

		tpl.Execute(w, data)
		return
	}
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
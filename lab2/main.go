package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"math"
)

var tpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Розрахунок валових викидів шкідливих речовин у вигляді суспендованих твердих частинок при спалювання вугілля, мазуту та природного газу</title>
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
        <h2>Розрахунок валових викидів шкідливих речовин у вигляді суспендованих твердих частинок при спалювання вугілля, мазуту та природного газу</h2>
        <form method="post">
			<input type="text" name="coal" placeholder="Вугілля" value="{{.Coal}}">
			<input type="text" name="oil" placeholder="Мазут" value="{{.Oil}}">
			<input type="text" name="gaz" placeholder="Газ" value="{{.Gaz}}">
            <input type="submit" value="Розрахувати">
        </form>
       {{if .Result}}
			<div class="results">
				<h3>Результати:</h3>
				<pre>{{.Result}}</pre>
			</div>
		{{end}}
    </div>
</body>
</html>`))

type Data struct {
	Result string
	Coal   float64
	Oil    float64
	Gaz    float64
}

func calculate(coal, oil float64) string {
	coalRes := math.Pow(10, -6) * 150 * coal * 20.47
	oilRes  := math.Pow(10, -6) * 0.57 * oil * 39.48

	result := fmt.Sprintf(`
Валовий викид при спалюванні вугілля: %.2f т.
Валовий викид при спалюванні мазуту:  %.2f т.
Сумарно:                              %.2f т.
`, coalRes, oilRes, coalRes + oilRes)

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		coal, _ := strconv.ParseFloat(r.FormValue("coal"), 64)
		oil, _  := strconv.ParseFloat(r.FormValue("oil"), 64)
		gaz, _  := strconv.ParseFloat(r.FormValue("gaz"), 64)

		data := Data{
			Result: calculate(coal, oil),
			Coal:   coal,
			Oil:    oil,
			Gaz:    gaz,
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

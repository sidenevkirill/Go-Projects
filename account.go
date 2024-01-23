package main

import (
	"html/template"
	"log"
	"net/http"
)

type PersonData struct {
	Name     string
	Login    string
	Password string
	Token    string
}

type PageData struct {
	PageTitle string
	People    []PersonData
}

func main() {
	http.HandleFunc("/", tmplServer)
	log.Println("Starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func tmplServer(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("tpl").Parse(layout))

	pageData := PageData{
		PageTitle: "Аккаунты ВКонтакте",
		People: []PersonData{
			PersonData{
				Name:     "Женя Конт",
				Login:    "89269021734",
				Password: "margosha",
				Token:    "vk1.a.lA8nxsLmHR6PXT_cDYPJCbdZMm-eNAcIN1tMjja4aWS-pjHkzWafz2sDNtgXKn9jx8wSIc7ycqmBLjxIhc1VOTw9dz0Dza4UZ7m4QFslYQd5FMVcGIs_lPrVe_kXyytBWp6_2V-hxcox_P6moRP4klkzYFYs1zzPBQxfUMOtcpQlHotVRauGhRV2DC1-KZFDhR49dBhSle5w6Zdja18daQ",
			},
			PersonData{
				Name:     "Захар Галка",
				Login:    "89081829910",
				Password: "Peresvet_20.",
				Token:    "vk1.a.gGk7Q1ain7AgZ7Pp_Wm6f41vI65L-6mfpi1mlHaAv512F-l3wOV37vYjaPZyMRJQklbCY7RpCgQetdELYJdaGFzac2EYJripa4nkMHQpp2jmnkKLO5TyQG_RsCs0l2NR0T4ZpqT48x8gtJO5VJuGKvw0B3nmbha8GEv9z4VNrWqgg9cyVrSN2t9G7iwiBtg25FbvN72zrC27XaA80tSXvA",
			},
			PersonData{
				Name:     "Иван Разливалов",
				Login:    "89829659231",
				Password: "?Gettherefast72",
				Token:    "vk1.a.lwMwuwAm3PHcIDV8djqMZCrG7kLvjTQ0NeMDrXm8RUJCS5BqMqNqsgQvreg-zoM2KPKpPMx_QnZvceCYd26O1zwxaaZfZpBDNwvS90uQniHSw_iwVX6St-ubKwVbD_lZ2u0Wt44JC90pyutiDE_tZXlFJvEIsNzdR7MKDkcFN66vu3s8z5zrzVgN6sdH3N0ostfsHqzN-eGCfU28hvTeQw",
			},
			PersonData{
				Name:     "Ирина Кондратьева",
				Login:    "89080804958",
				Password: "ирина120222",
				Token:    "vk1.a.C13G5mA8Q-JYeBrxPr1a0bzszXJQHSgg7suW4waOs2S8Nr_Fnb5PCJiDg8WV8X-cNFSJZY08iJathWEWyxXNHeLWrjwba18Ll0DtEFah6IbfY-J9s6CZ4ULwvg2VLoeULxDh3xD3URbqJDb7s2sPw8NOsfvCFw1ERl1zG-9FpSNmZwTVoOKbhZjbmzjrTKYSXWD6-nM6Wv5Ye8dSdASnhg",
			},
			PersonData{
				Name:     "Азалия Нигматуллина",
				Login:    "89674623683",
				Password: "AzaNi123456789",
				Token:    "vk1.a.sh0e-eBvPv7sE8vEULDsQmJaqF50jC5jJgL0x3l2gWD8ZQkDZg0enJO8aTSFEbTmKc3452lxKnYBs7b2VZ2OeLxWv5Bp-wp0_uTq_X5vjKH_H-zguD7DAA5ArPtp-bRda2MmP8p7wAr8gumGdkj37sTGh6dWz5XMMep2Y0bs4EF8NfZN391g242t_9sldwk_okgRGda2aqvlGlxCLzMnpg",
			},
			PersonData{
				Name:     "Диана Боровик",
				Login:    "2075220d@gmail.com",
				Password: "diana2075220",
				Token:    "vk1.a.Sgj79LmaAG9_sU6voGCr3Okwb308VLYzvIZHE2y7GbFO0SDAyMiPhq3KR2L4KJCIN8oRyBiuPU3z4vVyav4_178D634jEaETbuYlZx06v0HL5cbe5RmIll1DtQ5cQ-j_J1UuZFuLFrTGnTyeTOzgreaACH1CspbB7Pu7VhraVl2LDro75e8-46mLbIZfdfrvkZlw6h0Jse90YZN9hUNaWw",
			},
			PersonData{
				Name:     "Екатерина Кувшинова",
				Login:    "89867337985",
				Password: "Zzzxx7985zzxxx",
				Token:    "vk1.a.8w2-vn3wC0XtNTnqP-rOYPzpZFwXoH2skI4JKVeIwcNQ9P9MuGEqF4QtjJEgCVqUEKps5_cua_iomwrVKxzwbRWI8EB-C4ggDnTlVJ6ARlBtkx4UVYNfGAbauIZUsuYLn1P1wmSgKzl85oxOAA2AW5a67nQi8CbZsgGhIoeT1ekhbAYkknJVMklfIRthlAhEZRRkAwM3hcq7Lv5Brp8Jng",
			},
			PersonData{
				Name:     "Диана Хабибуллина",
				Login:    "89174416112",
				Password: "fqdfh123",
				Token:    "vk1.a.qEh-jPh8HzMAcGia5V8DVkMk2q404g3J1XCkvVysIbV2ORaLILq49nLJ34ISq72euShNMH1yMH88OqAR4_uB3lHOeLNlcUg9WZajZk2XDyf2zyTP9OaoL1xpEMqSHcwpgO0kznl1h1O6JmeLd5FWzCGW7Zuj9uJw2bXrtQZUn1zO4SyvYfmQho5I_Ovpj27uS6I1vp95GW9VP50LXRenhg",
			},
			PersonData{
				Name:     "Илья Чубаров",
				Login:    "+79620972454",
				Password: "Iliya12345",
				Token:    "vk1.a.wt1LWFaoHGVA_gT0ZrBkSlo1zIUqwCj0rDkiVkDrjpcNnv9xuuijtmdmuv9Pbi_OJREwKE73vmPSGM5gXiQQrXXGNaVnn461WROAWnKA8s5avYiuYSa_0PFlOW8fSDxDCnaofY5SNHFxDRq2s70BxLIb6oEM35cUBRvJWtI1_L0lg8sgLGfGQ9C-wzvAlV5PujKDcqhbwzEcjafGY8Mu4g",
			},
			PersonData{
				Name:     "Ольга Болотнова",
				Login:    "89636318707",
				Password: "2M9U1S6iNa*",
				Token:    "vk1.a.ErtVatT-GE4nfc3ReDpNKpLj34OQHGDdG50y_6QDapC-x0uavl-nWM4oEpOBcvzTUz_vIKL17phFaeJR8NaXY8xxNf5DDUYgEGLTtMl_AsZCpF5Ky_oa-yCwDggh9lSaqn_2sqsf29wgBvxmqNevXsm3jj1HeEdcLbZ7d1xo-_roA870M2xlsya34FkaR39_KH433p3KXQkOuJ_cWa6gjg",
			},
			PersonData{
				Name:     "Киря Левинов",
				Login:    "89281993635",
				Password: "Yulia3635_",
				Token:    "vk1.a.UOEZB6j6-NhvApmxsD8K7MlH0hEVkHC4lMQcQu3xDXJWjPsNV27YZHZ38t9xgzhXWoo7oqqIVVE0zx__mcD7uK38SJ4B-xS0IrTAUrNlBoZXj5wwo2qqaLNDHsrDNt9n4fly1U3lr4eaGnxBML-df3J9ZJoDfmwOoi25OJnl9JU5vo_PR1zUlKG1VHet1iXeGtsmhbG8rLIyXSPpFGemyw",
			},
			PersonData{
				Name:     "Анастасия Григорян",
				Login:    "nastya7858@yandex.ru",
				Password: "1216Grigoryash",
				Token:    "vk1.a.9LOl7XVt6getc-q-0GqbwEI5wQiWEkslCeIgJYYEVci1U2f_JvdTYyFJvQP_b6VcwrOpFikLXW_f9wNU0Bsx4HzQullHZMp5797KX4Q7-9w_lAYyg76NdI50l6Nqq-bxAagITdZhRgY3ar_CizHVhmUZBd74p5TH9hcQsDptUGHe-ZE6f5c7oLX8b4N0qLszqHjsQnOwGkdtVAFyLBvv7g",
			},
			PersonData{
				Name:     "Владимир Владимиров",
				Login:    "89521844447",
				Password: "Zaxar3480z",
				Token:    "vk1.a.1tlBOz6sCSFnJ2le3hwyhMoHzNZrieqVpHoyHD-LXotaKPwi_2s9P6MorhyqjQh1gFGigj8mYQ3oeeorG6fgO68QGbEEoBikFykjCUUaP-iRGf828I7G2F9-tzeuwXTvuI-7jMPLP8JUYlixFLTG0dkozQBrH0a45xwS0pcMGJbC_a5NLfQ06ItmXcDMHvRQ9zJMkWRjLK-vVRXX1VLRVA",
			},
			PersonData{
				Name:     "Alice Makeeva",
				Login:    "79118895773",
				Password: "люблюсебя_16",
				Token:    "vk1.a.F1FOhlyYVdFzZMgn-KQLcXpEvKifU3XEEYQbXjAYYCDYKWTr6TsCinIn5aDJlV3c20TEO1TnSPFmusAxqkhaNv6qyL4fbOq7lS8InQ6rTEGs8qh4wZQkn2gXXWHbLFOlIUhXVJeDFaL55My9ujKv4i07zVyJZ-72OSE_LEDvwZp3TTM1pieeeg_Vh1dBBqoEFBV7TvGO7z9Si_iVn-3Z_g",
			},
			PersonData{
				Name:     "Сергей Киреев",
				Login:    "89537705763",
				Password: "1957077a",
				Token:    "vk1.a.iqbRZjNa1y2NMVXyrRVVyrXLkMrSv4IWGXWQwrxps2lrYghZBh0awX5gLt5-e-7bCfITdT9ha-FDYwAPioznOU1_0LLAY2-5vYCaeXdMIb0tJGBD5V6RFX5BOga11HURkj9Fwufq3j50eBYjZaIsgTiR3L58GKlVD46-hzH-4shpV_5R2piW16kGJ2rl-yCa5PpG8s9snI5rznsXvNznjw",
			},
		},
	}

	tmpl.Execute(w, pageData)
}

var layout string = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>{{.PageTitle}}</title>
	<style>
		body {
			margin: 0 auto;
			font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
			font-size: 1rem;
			font-weight: 400;
			line-height: 1.5;
			color: #212529;
			text-align: left;
			background-color: #fff;
			padding: 20px;
		}

		header {
			position: fixed;
			top: 0;
			left: 0;
			width: 100%;
			height: 44px;
			z-index: 10;
			background: rgba(249,249,249,.85);
			border-bottom: 1px solid rgba(0,0,0,.1);
			-webkit-backdrop-filter: blur(20px);
			backdrop-filter: blur(20px);
		}

		h1 {
			margin: 0;
			text-align: center;
			font-size: 18px;
			line-height: 44px;
		}

		hr {
			box-sizing: content-box;
			height: 0;
			overflow: visible;
			margin-top: 1rem;
			margin-bottom: 1rem;
			border: 0;
			border-top: 1px solid rgba(0, 0, 0, 0.1);
		}

		table {
			border-collapse: collapse;
			border-spacing: 0;
			empty-cells: show;
			border: 1px solid #cbcbcb;
		}

		table td,
		table th {
			border-left: 1px solid #cbcbcb;
			border-width: 0 0 0 1px;
			font-size: inherit;
			margin: 0;
			overflow: visible;
			padding: 6px 12px;
		}

		table td:first-child,
		table th:first-child {
			border-left-width: 0;
		}

		table thead {
			background: #e0e0e0;
			color: #000;
			text-align: left;
			vertical-align: bottom;
		}

		table td {
			background-color: transparent;
			border-bottom: 1px solid #cbcbcb;
		}

		table tr:nth-child(2n) td {
			background-color: #f2f2f2;
		}
		table tbody > tr:last-child td{
			border-bottom-width: 0;
		}
	</style>
</head>
<body>
<header>
     <h1>{{.PageTitle}}</h1>
    </header>
	<hr/>
	<table>
		<thead>
			<tr>
				<th>Имя</th>
				<th>Логин</th>
				<th>Пароль</th>
				<th>Токен</th>
			</tr>
		</thead>
		<tbody>
			{{range .People}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.Login}}</td>
				<td>{{.Password}}</td>
				<td>{{.Token}}</td>
			</tr>
			{{end}}
		</tbody>
	</table>
	<hr/>
	<p>Data from <a href="https://qqaazzg.github.io/">QQAAZZ CREAM GROUP</a></p>
</body>
</html>`

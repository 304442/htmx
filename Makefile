run:
	templ generate ; go build -o bin/app cmd/main.go && bin/app serve --http="localhost:8080"

templ:
	templ generate -watch -proxy=http://localhost:8090

tcss:
	tailwindcss --config configs/tailwind.config.js \
		-i configs/input.css \
		-o assets/css/styles.css \
		-w \
		--minify

stripehook:
	stripe listen --forward-to http://127.0.0.1:8080/webhook --forward-connect-to http://127.0.0.1:8080/webhook
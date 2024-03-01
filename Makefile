run:
	templ generate ; go build -o bin/app cmd/main.go && bin/app serve


tailwindcss:

	./tcss --config configs/tailwind.config.js \
		-i configs/input.css \
		-o assets/css/styles.css \
		-w \
		--minify
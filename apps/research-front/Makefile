.PHONY: test
.PHONY: build

main: run

run:
	[[ `docker ps -a | grep web-proxy | wc -l` -eq 1 ]] && docker stop web-proxy && docker rm web-proxy || echo "ok"
	docker run --name web-proxy -v ${PWD}/proxy.conf:/etc/nginx/nginx.conf:ro -p 8080:8080 -d nginx
	webdev serve --no-injected-client --hostname=0.0.0.0 web:18080

proxy:
	docker run --name web-proxy -v ${PWD}/proxy.conf:/etc/nginx/nginx.conf:ro --rm -p 8080:8080 -d nginx

stop-proxy:
	docker stop web-proxy

# JSONをビルドするタスク
build:
	pub run build_runner build --delete-conflicting-outputs

clean:
	pub run build_runner clean

test:
	pub run test

analyze:
	dartanalyzer .

release:
	webdev build -- --delete-conflicting-outputs

run-standalone:
	webdev serve --auto=restart --launch-in-chrome --debug

run-dev:
	webdev serve --no-injected-client --hostname=0.0.0.0 web:18080

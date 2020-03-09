.PHONY: test
.PHONY: build

main: run
run:
	webdev serve --auto=restart --launch-in-chrome --debug

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
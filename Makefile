.PHONY: all assets compile clean

all: clean assets compile

clean:
	rm -rf dist
	mkdir -p dist/assets

assets:
	cp -r assets/* dist/assets/

compile:
	go build -o "dist/mines"

run:
ifdef file
	@go build -o ./.bin/main $(file)
	@./.bin/main
	@rm -r ./.bin
else
	@echo "file is not defined"
endif
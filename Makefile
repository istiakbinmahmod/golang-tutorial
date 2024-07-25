run:
ifdef file
	go build -o ./.bin/main $(file)
	@echo "---Running $(file)---"
	./.bin/main
	@echo "---Removing binary file---"
	sudo rm -r ./.bin
else
	@echo "file is not defined"
endif
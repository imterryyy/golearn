OUR_DIR=build

golearn:
	go build -o $(OUR_DIR)/golearn main.go

run_golearn:
	./build/golearn

golearn_and_run: golearn run_golearn

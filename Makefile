gocc: main.go
	go build -o gocc main.go

test: gocc
	./test.sh

clean:
	rm -rf gocc *.o *~ tmp*

.PHONY: test clean
build:
	find . -name "main.go" | xargs -L 1 dirname | grep -v -E 'test|cmd|vendor' | sort | xargs -L 1 bash -c 'cd "$0" && make build'

run:
	find . -name "main.go" | xargs -L 1 dirname | grep -v -E 'test|cmd|vendor' | sort | xargs -L 1 micro run

kill:
	find . -name "main.go" | xargs -L 1 dirname | grep -v -E 'test|cmd|vendor' | xargs basename | xargs -L 1 micro kill

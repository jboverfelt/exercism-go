all: testall

testall:
	find . -type d -not -path "./.git/*" -not -path "." -not -path "./.git" | xargs -n1 $(MAKE) -C

.PHONY: testall

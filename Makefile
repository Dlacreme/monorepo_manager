SRC=src/main.go
GO = go
NAME = mm
INSTALL_FOLDER=~/.local/bin/

# usage: CMD=init make run
run:
	$(GO) run $(SRC) $$CMD

test:
	$(GO) test -failfast ./...

test.suite:
	$(GO) test ./...

test.coverage:
	$(GO) test -coverpkg=./... ./...

build:
	$(GO) build $(SRC)

release:
	$(GO) build -ldflags "-s -w" -o $(NAME) $(SRC)

install: release
	echo "$(NAME) successfully built. Now installing in $(INSTALL_FOLDER)"
	mv $(NAME) $(INSTALL_FOLDER)
	echo "Done. Thank you for using Monorepo Manager."

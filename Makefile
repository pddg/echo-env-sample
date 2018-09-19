NAME    := "echo-env-sample"
SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -extldflags \"-static\""

.PHONY: build clean;

build: $(NAME);
$(NAME): $(SRCS)
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(NAME) .

clean:
	rm ./$(NAME)

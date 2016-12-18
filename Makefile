# ref. http://postd.cc/auto-documented-makefile/

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

todo: ## List up TODO tasks
	@grep --color=auto -nR TODO src

test: ## Run tests only
	go test $$(glide novendor) $(OPT)

install: ## Install packages for dependencies
	go get github.com/Masterminds/glide
	glide install
	cd .git/hooks && [ -L pre-commit ] || ln -s ../../scripts/git-hooks/pre-commit

update: ## Update packages for dependencies
	glide update

.PHONY: api
# generate api
api:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# generate wire
wire:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: proto
# generate proto
proto:
	find app -maxdepth 2 -mindepth 2 -type d -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) proto'

.PHONY: set-env
set-env:
	docker-compose -f docker-compose.yaml up -d

.PHONY: clean-env
clean-env:
	docker-compose -f docker-compose.yaml down
VERSION = 'v0.2.0'

build:
	@docker build -t kuops/alertmanger-feishu-webhook:$(VERSION) . && \
	docker push kuops/alertmanger-feishu-webhook:$(VERSION)

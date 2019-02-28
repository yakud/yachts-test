.PHONY: all
all: rundb build

build:
	docker build -t yakud/yachts-test .

run:
	docker run -i --name=yachts-rest --link=es-yachts --rm -p 80:80 yakud/yachts-test

log:
	docker logs -f yachts-rest

test:
	go test -v -cover ./yacht/ && \
	go test -v -cover ./gds/

rundb:
	docker run -d \
	--name=es-yachts \
	-p 9200:9200 \
	-p 9300:9300 \
	-e "discovery.type=single-node" \
	--ulimit nofile=65535:65535 \
	--ulimit memlock=-1:-1 \
	-e "bootstrap.memory_lock=true" \
	-e "ES_JAVA_OPTS=-Xms1g -Xmx1g" \
	docker.elastic.co/elasticsearch/elasticsearch:6.6.1


.PHONY: postgres adminer release
postgres:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
		docker run --rm --network host adminer

release:
		go test

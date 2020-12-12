.PHONY: postgres adminer


postgres:
		docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
		docker run --rm -ti --network host adminer

build:
	docker-compose build
upd:
	docker-compose up -d

down:
	docker-compose down

go_app:
	docker-compose exec app /bin/bash

go_db:
	docker-compose exec postgres /bin/bash

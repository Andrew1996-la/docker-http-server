service-deploy:
	docker compose up --build application

service-undeploy:
	docker compose down application
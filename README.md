This is just a test repository.

My problem is, that I have to wait in a test for the elasticsearch indexing to happen.

For test execute (while docker-compose is installed and after docker-machine set the env variables)

	make dev # starts the docker containers
	make test-docker # executes the test on the container

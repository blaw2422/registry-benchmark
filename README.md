# registry-benchmark

Simple tool to benchmark docker-registry using `docker pull` command.

## Usage

```
registry-benchmark <image[:tag]> <iterations>
```

Make sure your shell already has the environment variables such as
`DOCKER_HOST`, `DOCKER_CERT_DIR` (if applicable) set and Docker client
installed.

The program will write logs to stderr, and duration of each `docker pull`
(in seconds) to the stdout.

### How to run in a container

The image for this program exists on [Docker Hub](https://hub.docker.com/r/ahmetalpbalkan/registry-benchmark/):

    docker pull ahmetalpbalkan/registry-benchmark

In order to test pull speed of the Docker engine you are running at:

    docker run -it -v /var/run/docker.sock/:/var/run/docker.sock/ \
        ahmetalpbalkan/registry-benchmark <image> <iterations>

## Author

Ahmet Alp Balkan

## License

See [LICENSE](LICENSE).

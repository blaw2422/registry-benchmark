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

## Author

Ahmet Alp Balkan

## License

See [LICENSE](LICENSE).

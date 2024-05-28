## Excel Reporter for Github PRs in golang


### Deploy Docker Image


#### Load enviroment variables

Replace the environment variables in the file `env.tmpl`

```bash
    cp env.tmpl .env
    
    EMAIL_ADDRESS   # Email Address (e.g. "some@example.com")
    EMAIL_PASSWORD  # Password or secret key (e.g. "abcd defg hijkl mnop")
    API_PORT        # API Port (e.g. "8082")
    
    source .env
```

#### Build Docker image

```bash
    ./docker-build-image.sh
```

#### Deploy Docker image

```bash
    ./docker-deploy-image.sh
```

#### Print API logs
The logs are printed in the docker console. Run the next comand to display the logs of the API.
```bash
    docker logs reporter -f
```
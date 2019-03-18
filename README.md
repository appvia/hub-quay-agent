## **Hub Quay Agent**

The project teams up with the [appvia-hub](https://github.com/appvia/appvia-hub) providing the ability to provision repositories, robots tokens and permissions within the Quay.io registry.

```shell
$ bin/hub-quay-agent --help
NAME:
   hub-quay-agent - A backend agent used to provision resources within quay.io

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

AUTHOR:
   Rohith Jayawardene <gambol99@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --listen INTERFACE       the interface to bind the service to INTERFACE (default: "127.0.0.1") [$LISTEN]
   --http-port PORT         network interface the service should listen on PORT (default: 10080) [$HTTP_PORT]
   --https-port PORT        network interface the service should listen on PORT (default: 10443) [$HTTPS_PORT]
   --tls-cert PATH          the path to the file containing the certificate pem PATH [$TLS_CERT]
   --tls-key PATH           the path to the file containing the private key pem PATH [$TLS_KEY]
   --auth-token TOKEN       authentication token used to verifier the caller TOKEN [$AUTH_TOKEN]
   --quay-endpoint-url URL  the url for the quay.io api URL (default: "https://quay.io") [$QUAY_ENDPOINT_URL]
   --quay-api-token TOKEN   an authentication token used to permit api access TOKEN [$QUAY_API_TOKEN]
   --help, -h               show help
   --version, -v            print the version
```

#### **Quay.io Authentication**

In order to speak to quay.io you need to provision an access token for the organization. Go to Applications in the Organization settings and create a new application. Once done enter the application and generate a token _(bottom tab)_. Ensure the token has right scope to provision repositories.


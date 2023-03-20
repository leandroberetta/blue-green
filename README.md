# Blue-Green

Simple application to test a blue-green deployment.

## Usage

To test the deployment, an environment variable called **BACKGROUND_COLOR** can be used with some colors like **blue**, **green** or any other color you like.

## Build image

```
podman build -t bg .
```

## Run application

```
podman run -p 8080:8080 -e 'BACKGROUND_COLOR=blue' bg:latest
```
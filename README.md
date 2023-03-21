# NucWrapGo

Nucwrap is a wrapper tool around the nuclei scanner which converts the output to a loadable format of JSON and table format.



# Using the Nucwrap

## Installation
If you have Go installed and configured (i.e. with $GOPATH/bin in your $PATH):

```bash
▶ git clone github.com/90r1ll4/NucWrapGo
```

## To build (Normal)
To use the Nucwrap, you can run it with the following command:
```bash
▶ go build -o nucwrapgo main.go
```

### Usage
```bash
Wrapper For Nuclei

Usage:
  ./nucwrapgo [flags]

Flags:
   -u, -url string       url to scan
   -f, -url_file string  list of urls
   -o, -output string    Output in text form[tables][json]
   -json                 Output in json form
   -tables               Output in table form
```
For an overview of all commands use the following command:

```bash
nucwrapgo -h
```


## Installation with Docker
This tool can also be used with [Docker](https://www.docker.com/). To set up the Docker environment, follow these steps (trying using with sudo, if you get any error):

```bash
docker build -t nucwrapgo .
```

## Using the Docker Container

A typical run through Docker would look as follows:

```bash
docker run --rm nucwrapgo -u scanme.nmap.org  -tables -json
```

**NOTE:** Nucwrapgo should be used responsibly and with permission from the target owner.

# Famiphoto

## How to install

You have to set up docker before do follow steps.
Do commands follow steps.

```sh
git clone git@github.com:hiroyky/famiphoto.git
cd famiphoto
./setup_env.sh
make init
```

`make init` executes docker compose build and up.

## Set up for development at local

You have to set up docker before do follow steps.

1. Do `./setup_env.sh`, enter configuration items and `.env` file will be created here.
1. Do `make init`, docker containers will begin.

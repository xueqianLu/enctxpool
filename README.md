# enctxpool

# step
1. install ego-dev 
```shell
sudo snap install ego-dev --classic
```
2. build and sign
```shell
ego-go build ./cmd/enctxpool
ego sign enctxpool

```

3. run in enclave
```shell
ego run ./enctxpool
```
or run in simulate mode.
```shell
OE_SIMULATION=1 ego run ./enctxpool
```
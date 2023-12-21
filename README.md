# Hypersign Load Test

A CLI tool to similute SSI transactions on Hypersign Prajna Testnet

## Usage

### Setup

1. Clone this repository and install the binary (Binary name: `hypersign-load-test`)

```
$ git clone github.com/hypersign-protocol/hypersign-load-test
$ cd hypersign-load-test
$ make install
```

If `make install` doesn't work, then run the following:

```
$ make build
$ sudo cp ./build/hypersign-load-test /usr/local/bin
```

2. We require an account which will submit transactions to Prajna Tesnet Network:

```
hypersign-load-test create-account <name-of-the-account>
```

Fund atleast `2.6 HID` (or `260000uhid`) to this newly created account.

3. To list all created accounts:

```
hypersign-load-test list-accounts
```

### Initiating the Load test

4. Start the Load test (Make sure the account that you just created is funded)

```
hypersign-load-test start --account <name-of-the-account>
```

If the Node you want to connect is running on a different IP and Port, use the `--node` flag:

```
hypersign-load-test start --account <name-of-the-account> --node <RPC address>
```

### Get Smart Contract Owner Address (Golang)

A simple library that will return the address of a smart contract owner in string format. E.g. `0x463cf5545ea6da915cf37483a48a5f36bb7f7845`.

[See the code](vendor/zcontractowner/zcontractowner.go) | [Example](main.go)

Not using Go bindings/abigen. But JSON RPC. :rocket:

This library expects your contract to look like this:

```
contract YourContract {
    // code omitted for brevity

    address public owner;

    // code omitted for brevity
}
```

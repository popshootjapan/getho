# getho  
> This package provides a unified command line interface to [getho](https://getho.io).

<br>
<div align="center">
  <a href="https://github.com/popshootjapan/getho">
    <img width=360px src="https://s3-us-west-2.amazonaws.com/io.getho.public/i/getho-cli-logo.png">
  </a>
</div>
<br>

## Installation  
On MacOS you can install or upgrade to the latest released version with Homebrew:

```
$ brew tap popshootjapan/homebrew-getho
$ brew install getho
```

Or, to install `getho`, simply execute the following command in a terminal from your `$GOPATH`:  

```
go get github.com/popshootjapan/getho
```


## Get started  
When you use `getho`, you'll need to login to your account.  

```
$ getho login
$ email: // type your email.
$ password: // type your password.
$ login completed
```

then, you can intract with your node.  

## Command line  

```
$ getho [command] <options>
```

Command:
- `login`: login to your account.
- `contracts`: get list of contracts you uploaded.
- `nodes`: get list of nodes.
- `upload`: upload your smart contract to a node.

Options:
- `--subdomain` or `-s`: Specify your subdomain which related with a node.
- `--network-id` or `-n`: Specify your network in truffle.js (default: `1010`)
- `--email` or `-e`: Specify your email you use.
- `--password` or `-p`: Specify your password you use.

## Upload smart contract  
To upload your smart contract, you'll need to compile it before.

```
$ truffle compile
```

Then, specify compiled contract json file from `build/contracts` directory and upload.  

```
$ getho upload ./build/contracts/GethoCoin.json --subdomain <SUBDOMAIN> --network-id <NETWORK_ID>
$ Uploaded
```

## License
`getho` is available under the MIT license. See the LICENSE file for more info.

# myhttp

## Overview

Myhttp is a CLI tool that makes a GET request to the given URLs and after if the request is successful it prints the URLs on a separate line with the MD5 hash of the response body of the calls.

### Installation
Before you begin you must have Go installed and configured properly for your computer. Please see [link] https://golang.org/doc/install

### Usage
To be able to use the myhttp CLI, you need to build the binary first. You can build it running the following commands on your favorite shell
```sh
cd path/to/location
git clone https://github.com/brkykan/myhttp.git
cd myhttp
source project
go build
```

After you build it, you can run it with the following command

```sh
./myhttp [-parallel n] google.com https://facebook.com
```

Please note that 

* You don't have to specify the scheme of the URLs you pass in the arguments, it's completely okay if you speficy as well.

* If you want to make the requests in parallel, you can pass the maximum number of parallel requests you want to make as in the `-parallel` flag. The default value for the flag is set to 10. If you don't pass any limit in the flag, by default maximum 10 parallel requests will be made.

You can use the flag as `-parallel=n` or `-parallel n`, whichever way you'd prefer.
### Tests

You can run the tests files with the following command

```sh
cd /path/to/myhttp/source
go test ./...
```

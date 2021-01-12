# PhaserJS!

Learning how to write simple browser based games using the [PhaserJS](https://phaser.io/) framework.

## Why Phaser?

### Setup

Phaser had a great tutorial that was easy to use and easy to get started as well as easy to understand how to integrate it with whatever server backend that you should so desire. I am personally not a node.js expert (yet) so the other game library tutorials that started with instructions like "run `git clone ...` then `npm install` and `npm run serve` and that's it!"; while the setup is "_convenient_" it is wildly unclear to the javascript uninitiated (i.e. me) which of the hundreds of auto generated files I am now looking at are required as part of this framework and which are required as part of the node.js server which I do not plan to use in production (looking at you MelonJS).

### Syntax

Additionally the syntax was super easy to understand and the tutorial did a fantastic job walking through how to start phaser with minimal dependency management as well as how to build a working game again with minimal dependencies.

### Plus

Plus it works and looks nice.

# Using this code

You will need to have [golang](https://golang.org/doc/install) installed. Then just `cd` into the desired directory for either the `src/getting-started/server` or the `src/tutorial/server` and run:

```bash
go run *.go
open http://localhost:8001/
```

If you do not have `golang` or do not use `golang`, then you can spin up your own server of whatever type you prefer and use that to serve the files in the `phaser/` directories. I don't really have any other feedback on that. You cannot just open it in the filesystem given the filesystem is very restricted in terms of executing arbitrary javascript and as such will not let you just view the webpage as that is not secure.

# ejic
json-iterator codec plugin for [erpc](https://github.com/henrylee2cn/erpc)

## Usage
```
import (
    "github.com/sqos/ejic"
    "github.com/henrylee2cn/erpc/v6"
)

peer := erpc.NewPeer(erpc.PeerConfig{DefaultBodyCodec: ejic.JsoniterCodec{}.Name()})
```

***NOTE: client and server should set DefaultBodyCodec with same configure***

## License

ejic is under Apache v2 License. See the [LICENSE](https://github.com/sqos/ejic/LICENSE) file for the full license text

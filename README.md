# imageflux-cli

This is unofficial [ImageFlux](https://www.sakura.ad.jp/services/imageflux/) API client. Official API reference is [here](https://console.imageflux.jp/docs/).

## Lookup cache

```
imageflux-cli cache.lookup -k $url
```

## Delete cache

```
imageflux-cli cache.delete -k $url
```

## Configuration

`imageflux-cli` refers to *~/.imageflux/conf.ini*. *conf.ini* stores API token.

```
token = ffff...
```

## Contribution

Please read the CLA below carefully before submitting your contribution.

https://www.mercari.com/cla/

## License

Copyright 2017 Mercari, Inc.

Licensed under the MIT License.

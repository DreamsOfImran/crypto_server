# crypto-server

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![CircleCI](https://circleci.com/gh/DreamsOfImran/crypto_server.svg?style=svg)](https://circleci.com/gh/DreamsOfImran/crypto_server)
[![Maintainability](https://api.codeclimate.com/v1/badges/dd82c8401273b3d4153c/maintainability)](https://codeclimate.com/github/DreamsOfImran/crypto_server/maintainability)

This application for register, encrypt and decrypt using aes via API calls

## Table of Contents

- [Install](#install)
- [API](#api)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Install

```
# To run the unit test
$ sh bin/test

# To build the application. This needs go version to be >= 1.11 because of go modules dependency
$ sh bin/build

# To run the application
$ sh bin/run

# To start the server from binary
./app
```

## API
* To register User `/register/:id`
* To Encrypt the message `/encrypt_message/:id` with JSON body `{ "message": "Your Message"}`
* To Decrypt the message `/send_message/:id` with JSON body `{"message": "Your Encrypted Message"}`

## Maintainers

[@DreamsOfImran](https://github.com/DreamsOfImran)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2020 [Imran Basha](https://github.com/DreamsOfImran)

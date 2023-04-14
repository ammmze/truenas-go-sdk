# Changelog

## [1.0.0](https://github.com/ammmze/truenas-go-sdk/compare/v0.9.0...v1.0.0) (2023-04-14)


### ⚠ BREAKING CHANGES

* Go module location is changed

### Features

* add bash script for generator ([b30ade4](https://github.com/ammmze/truenas-go-sdk/commit/b30ade44eda7bfc9a0d9de814974397895069b18))
* add create dataset endpoint ([1d2b7cb](https://github.com/ammmze/truenas-go-sdk/commit/1d2b7cb5d531f2c45c26ef6d14736c1f39801339))
* add createVM operation ([d39123a](https://github.com/ammmze/truenas-go-sdk/commit/d39123ae41ab9acdbf726894642f06b356fe33af))
* add cronjob creation endpoint ([9da9c71](https://github.com/ammmze/truenas-go-sdk/commit/9da9c710d4b1a1e6d0b8fd1efbcb4432b88da561))
* add cronjob delete method ([53d0e47](https://github.com/ammmze/truenas-go-sdk/commit/53d0e4750c3f4cf917ffe2fcc154c69ef13d1681))
* add cronjob update endpoint ([aa27a1f](https://github.com/ammmze/truenas-go-sdk/commit/aa27a1f5727dcc8168bb99844f7a83b8441cd592))
* add dataset getter ([ea4b8b6](https://github.com/ammmze/truenas-go-sdk/commit/ea4b8b6b41c6a0063c7f8a4d30781cfbe3c101b3))
* add dataset update ([33630c6](https://github.com/ammmze/truenas-go-sdk/commit/33630c6f0d2ac14dd16ab51dfef35af26de39cd5))
* add delete dataset endpoint ([6eb3563](https://github.com/ammmze/truenas-go-sdk/commit/6eb35638f474b982a829716593f8e60e39965df9))
* add deleteVM operation ([7221395](https://github.com/ammmze/truenas-go-sdk/commit/72213955c8b811153ff6abae05830c2086a4da98))
* add fetch script and update original spec ([c547f32](https://github.com/ammmze/truenas-go-sdk/commit/c547f32c04a23eb713447ecfdd7f2424b789e6de))
* add get request for cronjob ([7499f96](https://github.com/ammmze/truenas-go-sdk/commit/7499f96fd45d1b88a3434bdc8a7a57f428ebb3ff))
* add getter for service ([6aae8fc](https://github.com/ammmze/truenas-go-sdk/commit/6aae8fc3b807752dd6bdb97db7b3311816aa313f))
* add initial bluefin sdk ([d9a7009](https://github.com/ammmze/truenas-go-sdk/commit/d9a7009c486a31544a446bc60072f99783da8b37))
* add list vms endpoint ([2f0d245](https://github.com/ammmze/truenas-go-sdk/commit/2f0d2452772738bf8b393d22ad6fa28cf21bfd8e))
* add network configuration endpoint ([70cc7a0](https://github.com/ammmze/truenas-go-sdk/commit/70cc7a0d74123c345a47b2c9f8578147a9df64af))
* add network summary endpoint ([db69105](https://github.com/ammmze/truenas-go-sdk/commit/db691058c36c00722660cf96f15fba29f08571fe))
* add props to service announcement obj ([8a85704](https://github.com/ammmze/truenas-go-sdk/commit/8a8570440ec27a90edef61ac9eb18fad110cf7dc))
* add sharing API for NFS and SMB ([5fbb3a0](https://github.com/ammmze/truenas-go-sdk/commit/5fbb3a0e914f25e75b4c4e39b23a700fe0e8ec4a))
* add updateVM and deleteVM operations ([1d6c5a3](https://github.com/ammmze/truenas-go-sdk/commit/1d6c5a358de43d998dd8c847404580d6090a7d34))
* add user/group endpoints ([b1ad2c3](https://github.com/ammmze/truenas-go-sdk/commit/b1ad2c385167dd637ce7763f304adc403718dc0b))
* add vm get endpoint ([030726a](https://github.com/ammmze/truenas-go-sdk/commit/030726a5c95e5429766876b04fc6235ebbb8a41a))
* additional properties to pools response ([1160194](https://github.com/ammmze/truenas-go-sdk/commit/1160194322fd8eb445834d7e6b0740752c3e9a67))
* create custom openapi spec with more detailed responses and re-generate apis ([a864257](https://github.com/ammmze/truenas-go-sdk/commit/a864257114ca89615fef9d50960a5a46c1bdfb57))
* few additional properties for datasets ([c96196f](https://github.com/ammmze/truenas-go-sdk/commit/c96196fef4bfb25c7660e63f44e44703b0b62677))
* initial generated sdk ([009cd9c](https://github.com/ammmze/truenas-go-sdk/commit/009cd9c86dc92454977129d7f7c4489cd7a9ce33))
* re-generate code using latest openapi v6.2.1 ([fa9f148](https://github.com/ammmze/truenas-go-sdk/commit/fa9f14820da96fc945be73ae1d29d01b3da841dd))
* start using release-please for managing releases ([10dd5d5](https://github.com/ammmze/truenas-go-sdk/commit/10dd5d5150d36fcf54dd88fe60edf57875d164aa))


### Bug Fixes

* add missing dataset prop ([49a7223](https://github.com/ammmze/truenas-go-sdk/commit/49a7223503a5506da3b3dfa9166440609b0077a0))
* add missing force_size parameter for dataset update ([d4545b6](https://github.com/ammmze/truenas-go-sdk/commit/d4545b60a35fb74270f7c807d2e1a1702b0cf857))
* add missing volblocksize property ([ff51528](https://github.com/ammmze/truenas-go-sdk/commit/ff51528cc166427152dd92c9ba74b2e751f332c3))
* add temporary patch to work with text/plain responses ([b3dec6d](https://github.com/ammmze/truenas-go-sdk/commit/b3dec6d5c97184ffc17910cde570f8107020a012))
* hashicorp gpg action no longer works ([16b70bf](https://github.com/ammmze/truenas-go-sdk/commit/16b70bf957d241b0fb458babac3181cfcb6f74d8))
* pids are ints ([fd3fe60](https://github.com/ammmze/truenas-go-sdk/commit/fd3fe60db0b1302a3d7aa527215841c821977504))
* pool id should be int ([7a39301](https://github.com/ammmze/truenas-go-sdk/commit/7a393010a15091df2c23b1feb882a1f9b41e6ee5))
* switch some integers to proper int64 format ([32d8e5b](https://github.com/ammmze/truenas-go-sdk/commit/32d8e5b887c86326adbd3ab6656d44a85bb5cc1f))
* try skipping build (this is only library) ([5a5fd02](https://github.com/ammmze/truenas-go-sdk/commit/5a5fd02d32c554076bbd4cdc8139b3fa1e0ebc8e))
* update nfs share api for bluefin ([c7f0b5e](https://github.com/ammmze/truenas-go-sdk/commit/c7f0b5e1fb25d093ac945884306f5d5fc59c2664))
* update vm parameters for bluefin sdk ([b81872f](https://github.com/ammmze/truenas-go-sdk/commit/b81872f8ba7416a7f99cb6876e7d2741329364a1))
* used wrong quota type ([f6d24b3](https://github.com/ammmze/truenas-go-sdk/commit/f6d24b351eee1b835192d1fcc87b083df815f174))
* volblocksize is string ([4757f7f](https://github.com/ammmze/truenas-go-sdk/commit/4757f7f79d7391c32f881393ec9f2d16a1ddbb3f))
* wrong option in git workflow ([2811e6f](https://github.com/ammmze/truenas-go-sdk/commit/2811e6f9b57a9ab381a49a91b0c8e6db10925317))


### Code Refactoring

* create subfolder for TrueNAS Angelfish release ([330e867](https://github.com/ammmze/truenas-go-sdk/commit/330e867a05d65cd703e8a29799d3025d208e3708))

## [0.9.0](https://github.com/dariusbakunas/truenas-go-sdk/compare/v0.8.1...v0.9.0) (2022-11-15)


### Features

* add createVM operation ([d39123a](https://github.com/dariusbakunas/truenas-go-sdk/commit/d39123ae41ab9acdbf726894642f06b356fe33af))
* add deleteVM operation ([7221395](https://github.com/dariusbakunas/truenas-go-sdk/commit/72213955c8b811153ff6abae05830c2086a4da98))
* add fetch script and update original spec ([c547f32](https://github.com/dariusbakunas/truenas-go-sdk/commit/c547f32c04a23eb713447ecfdd7f2424b789e6de))
* add updateVM and deleteVM operations ([1d6c5a3](https://github.com/dariusbakunas/truenas-go-sdk/commit/1d6c5a358de43d998dd8c847404580d6090a7d34))
* re-generate code using latest openapi v6.2.1 ([fa9f148](https://github.com/dariusbakunas/truenas-go-sdk/commit/fa9f14820da96fc945be73ae1d29d01b3da841dd))
* start using release-please for managing releases ([10dd5d5](https://github.com/dariusbakunas/truenas-go-sdk/commit/10dd5d5150d36fcf54dd88fe60edf57875d164aa))

## [0.8.1](https://github.com/dariusbakunas/truenas-go-sdk/compare/v0.8.0...v0.8.1) (2022-11-15)


### Features

* start using release-please for managing releases ([10dd5d5](https://github.com/dariusbakunas/truenas-go-sdk/commit/10dd5d5150d36fcf54dd88fe60edf57875d164aa))


### Bug Fixes

* try skipping build (this is only library) ([5a5fd02](https://github.com/dariusbakunas/truenas-go-sdk/commit/5a5fd02d32c554076bbd4cdc8139b3fa1e0ebc8e))

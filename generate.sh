#!/usr/bin/env bash
set -e
OPEN_API_GENERATOR_CLI_VERSION=2.5.2

# gets the latest "clean" version of the truenas OpenAPI because what truenas produces, is not valid OpenAPI spec.
# then replaces the model named new_brick with replacement_brick, because otherwise we end up with a model named
# Brick and NewBrick, but it also produces a "constructor" with "New{Model}", so then in the same scope, we end up
# with a struct named NewBrick and a function named NewBrick, which was no good.
# Finally using a custom built version of the openapi-generator-cli from: https://github.com/OpenAPITools/openapi-generator/pull/13472
curl -s 'https://raw.githubusercontent.com/ammmze/truenas-openapi/main/schemas/clean/scale-latest.yaml' | \
    yq '(.. | select((. | tag == "!!map") and .title == "new_brick").title) = "replacement_brick"' | \
    java -jar ~/projects/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -i /dev/stdin -c cfg/config.yaml -o . -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
# cat './cfg/scale-clean.yaml' | \
#     npx "@openapitools/openapi-generator-cli@${OPEN_API_GENERATOR_CLI_VERSION}" generate -i /dev/stdin -c cfg/config.yaml -o . -g go --git-user-id dariusbakunas --git-repo-id truenas-go-sdk
patch client.go < client.patch
go mod tidy
go fmt ./...

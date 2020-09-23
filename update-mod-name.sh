#! /bin/bash
set -e

# update all self-imports in the module to the new module name
grep -rl --include=*.go \"github.com/tendermint/tendermint . | xargs sed -i '' 's#"github.com/tendermint/tendermint#"github.com/kava-labs/tendermint#'
# update the module name
go mod edit -module github.com/kava-labs/tendermint

# update tm-db imports to use forked version
grep -rl --include=*.go \"github.com/tendermint/tm-db . | xargs sed -i '' 's#"github.com/tendermint/tm-db#"github.com/kava-labs/tm-db#'
# update go mod to refer to forked version
go mod edit -droprequire github.com/tendermint/tm-db -require github.com/kava-labs/tm-db@v0.1.1-kava

# next: fix commands below, run in terminal, go mod tidy, run tests. Reset and run all at once.
# update go-amino imports to use forked version
grep -rl --include=*.go \"github.com/tendermint/go-amino . | xargs sed -i '' 's#"github.com/tendermint/go-amino#"github.com/kava-labs/go-amino#'
# update go mod to refer to forked version
go mod edit -droprequire github.com/tendermint/go-amino -require github.com/kava-labs/go-amino@v0.14.1-kava
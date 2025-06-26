#!/bin/bash
# Original source: https://github.com/Bookshelf-Writer/scripts-for-integration/blob/main/_run/commit-hook.sh
echo "[HOOK]" "Commit"

run_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
values_dir="$run_dir/values"
script_dir="$run_dir/scripts"
root_path=$(cd "$run_dir/.." && pwd)

VERSION=$(bash "$script_dir/sys.sh" -v)
NAME=$(bash "$script_dir/git.sh" -b)

echo -e "$NAME [$VERSION] \n" $(cat "$1") > "$1"
#############################################################################

go test -v ./...

bash "$script_dir/creator_const_Go.sh"

#############################################################################
exit 0


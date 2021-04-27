#!/usr/bin/env bash
set -eu -o pipefail

help_download_sample="Download SHINRA2021-LinkJP task sample data."
download-sample() {
  mkdir -p tmp
  mkdir -p data

  local fileid=1rH-0L2E7Cxd8JIhss6AL1RZdPkSJLyJ-
  local out=tmp/linkjp-sample.zip
  curl -L -o ${out} "https://drive.google.com/uc?export=download&id=${fileid}"
  unzip ${out} -d data/
}

list() {
  declare -F | awk '{print $3}'
}

help_help="Print help text, or detailed help for a task."
help() {
	local item="${1-}"
	if [ -n "${item}" ]; then
		local help_name="help_${item//-/_}"
		echo "${!help_name-}"
		return
	fi

	type -t help-text-intro > /dev/null && help-text-intro
	for item in $(list); do
		local help_name="help_${item//-/_}"
		local text="${!help_name-}"
		[ -n "$text" ] && printf "%-20s\t%s\n" $item "$(echo "$text" | head -1)"
	done
}

case "${1-}" in
  list) list;;
  ""|"help") help "${2-}";;
  *) "$@";;
esac

#!/bin/bash
set -euxo pipefail

ZFS_FILE="$(pwd)/zfs_file"

truncate --size 1GB "${ZFS_FILE}"

sudo zpool create -f \
  -O compression=on \
  -O atime=off \
  -O recordsize=128k \
  -O logbias=throughput \
  -m /var/lib/test/dblab/test_dblab_pool \
  test_dblab_pool \
  "${ZFS_FILE}"

sudo zfs list

#!/bin/bash

# add jurisdiction
cd linux || exit

chmod +x set_path.sh
chmod +x add_all.sh
chmod +x commit.sh
chmod +x push.sh

echo "source_path=$1" > set_path.sh
echo "target_path=$2" >> set_path.sh

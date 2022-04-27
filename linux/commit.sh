source linux/set_path.sh

cd ${source_path}

git commit -m "$1"

cd ${target_path}

git commit -m "$1"


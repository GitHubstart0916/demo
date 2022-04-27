source linux/set_path.sh

# echo $source_path
# echo $target_path

cd ${source_path}
git add .
cp -r ${source_path}/backend/* ${target_path}

cd ${target_path}
git add .
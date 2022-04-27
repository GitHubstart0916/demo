source linux/set_path.sh

if [ $# == 0 ]; then
    # if body
    cd ${source_path}
    git push
    cd ${target_path}
    git push
elif [ $# == 2 ]; then
    # else if body
    cd ${source_path}
    if [[ "$1" == "-c" ]]; then
        git push --set-upstream origin $2
    elif [[ "$1" == "-f" ]]; then
        git push -f $2
    else
        echo "param error, create branch use -c and force to push use -f"
    fi
    cd ${target_path}
    git push
else
   echo "param num error!"
fi

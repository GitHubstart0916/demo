@echo off

cd windows
call base\set_path.bat
:: set source_path and target_path


set num=0
for %%x in (%*) do Set /A num+=1
:: echo %num%
:: echo %source_path%
:: echo %target_path%

if %num% EQU 0 (
    cd %source_path%
    git push
    cd %target_path%
    git push
    goto end
)
if %num% EQU 2 (
    cd %source_path%
    if "%1"=="-c" (
        git push --set-upstream origin %2
    ) else if "%1"=="-f" (
        git push -f %2
    ) else (
        echo param error, create branch use -c and force to push use -f
    )
    cd %target_path%
    git push
) else (
    echo param num error!
)


:end
cd %source_path%


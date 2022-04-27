cd windows

echo   sourcePath = r"%1/backend" >> py/copy.py
echo   targetPath = r"%2" >> py/copy.py
echo   getDirAndCopyFile(sourcePath, targetPath) >> py/copy.py
echo set source_path=%1 >> base\set_path.bat
echo set target_path=%2 >> base\set_path.bat

:: call base\add_at_begin "set source_path=%1" >> add_all.bat

cd ..

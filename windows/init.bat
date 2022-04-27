@echo off

cd windows


echo set source_path=%1> base\set_path.bat
echo set target_path=%2>> base\set_path.bat

:: call base\add_at_begin "set source_path=%1" >> add_all.bat

cd ..

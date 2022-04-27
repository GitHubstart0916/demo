cd windows

call base\set_path.bat
:: set source_path and target_path

cd %source_path%

git add .
:: python windows\py\copy.py
xcopy %source_path%\* %target_path% /e /c /y

cd %target_path%
git add .

cd %source_path%

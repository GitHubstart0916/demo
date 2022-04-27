

copy %2 temp.txt
echo %1 > %2
type temp.txt >> %2
del temp.txt
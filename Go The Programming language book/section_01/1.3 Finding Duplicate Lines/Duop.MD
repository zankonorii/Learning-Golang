## V2

The function `os.Open` returns two values. The first is an open file(*os.File) that is used in subsequent reads by the Scanner. The second resoult of `os.Open` is a value of the built-in error type.
If `err` equals the special built-in value `nil`, the file was opened successfully. After that the file is read, and when the input is rached, `Colse` function closes the file and release any resources. On the other hand, if `err` is not nil, something went wrong.
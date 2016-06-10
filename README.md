#File Signature Extractor (Written in Go)
This Go app will extract the file Signature from your File as an argument like below:
```
#bin/filesig -f FILE_GOES_HERE -b NUMBER_OF_BYTES_TO_READ"
```
***for example***
*To extract only 10 Bytes from a PDF file you use the below command:*

-
*لإستخراج أول عشرة بايت ل
ملف
PDF
نستخدم الأمر التالي*
-

```
#bin/filesig -f ~/Downloads/0912f5080126cc252b000000.pdf -b 10
25 50 44 46 2d 31 2e 34 0a 25
```

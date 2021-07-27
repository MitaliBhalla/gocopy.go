[readme1.odt](https://github.com/MitaliBhalla/gocopy.go/files/6882985/readme1.odt)
(Refer readme1.odt)




command line tool which clones the famous “wc” shell command# gocopy.goSyntax
wcg wc [OPTION]... [FILE]...


The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.
  -c, --bytes        	print the byte counts
  -m, --chars        	print the character counts
  -l, --lines        	print the newline counts
  	--files0-from=F	read input from the files specified by
                       	NULL-terminated names in file F;
                       	If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words        	print the word counts
  	--help 	display this help and exit
  	--version  output version information and exit

Example
Input: 
wc -c apple.txt 
Output:
31 apple.txt

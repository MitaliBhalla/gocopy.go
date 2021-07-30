


# **CLI TOOL USING COBRA AND GOLANG**

## **command line tool which clones the famous “wc” shell command# gocopy.goSyntax : wcg wc [OPTION]... [FILE]...**




## Performs the following tasks:
1. Counts number of characters
2. Counts number of lines
3. Counts number of bytes
4. Max-Line length
5. Help message

**HELP_DOCUMENT:** 
*The options below may be used to select which counts are printed, always in the following order: newline, word, character, byte, maximum line length.
  -c, --bytes        	print the byte counts
  -m, --chars        	print the character counts
  -l, --lines        	print the newline counts
  	--files0-from=F	read input from the files specified by
                       	NULL-terminated names in file F;
                       	If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words        	print the word counts
  	--help 	display this help and exit
  	--version  output version information and exit*
    

**Example**
Input: 
wc -c apple.txt 
Output:
31 apple.txt

![Screenshot from 2021-07-27 11-10-56](https://user-images.githubusercontent.com/74012736/127607622-cd35f607-f48a-4aa7-bc5e-58f126b35475.png)
![Screenshot from 2021-07-27 11-02-01](https://user-images.githubusercontent.com/74012736/127607632-e5ca1141-8686-4d9d-ab03-f258bc9f6dc0.png)


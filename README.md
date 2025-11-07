# yup-join

```
NAME:
   join - join lines of two files on a common field

USAGE:
   join [OPTIONS] FILE1 FILE2

      For each pair of input lines with identical join fields, write a line to
      standard output. The default join field is the first, delimited by blanks.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --field1 value, -1 value         join on this FIELD of file 1 (default: 1)
   --field2 value, -2 value         join on this FIELD of file 2 (default: 1)
   --output-format value, -o value  obey FORMAT while constructing output line
   --empty value, -e value          replace missing input fields with STRING
   --ignore-case, -i                ignore differences in case when comparing fields (default: false)
   --outer-join, -a                 also print unpairable lines (default: false)
   --unpaired-1, --v1               like -a 1, but suppress joined output lines (default: false)
   --unpaired-2, --v2               like -a 2, but suppress joined output lines (default: false)
   --check-order                    check that the input is correctly sorted (default: false)
   --help, -h                       show help
```

# Rwd Code Challenge

This application is a code challenge that reads text files and performs operations such as counting lines, bytes, and words.

## How to Run

To run the application, navigate to the `\rwc\cmd\rwc` folder and use the command below, replacing `[OPTIONS]` with the desired option and `test.txt` with the path to the file you want to process:

```sh
go run . [OPTION] ..\..\tests\test.txt
```

## Avaliable Options

- `-c`: Counts the number of bytes in the file.
- `-l`: Counts the number of lines in the file.
- `-w`: Counts the number of words in the file.
- `-m`: Counts the number of characters in the file.

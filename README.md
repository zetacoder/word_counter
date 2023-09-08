# Word Counter

Word Counter is a program that receives one or more text files from different types of inputs and returns a list of the most repeated consecutive words. It ignores punctuation, end of lines and capitalization. Also can handle concurrently multiples files, being fast and reliable.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Project Structure](#project-structure)
- [Usage](#usage)
- [Testing](#testing)
- [Bonus and Extras](#bonus)
- [Future Iterations and Comments](#future-iterations)


## Getting Started

Before starting, locate the folder with the program in a directory you know and read the README carefully.

### Prerequisites

- You must have Golang installed on your computer. Download it here: https://go.dev/doc/install

- You must have a code editor installed, for example Visual Studio Code (VSC). 
If you don't have it, download it from here: https://code.visualstudio.com/download

- It is strongly recommended to download the VSC extension of Golang. For that, open VSC, go to the left lateral bar and click Extension. Once there type Go in the search bar. Click the first option and install it.

- It is not mandatory, but it is recommended that you have Git installed on your computer. If you don't have it, download it here: https://git-scm.com/downloads

### Installation 

1) Download the folder word-counter-master
2) Locate it in a directory of your preference in your machine
3) Open the folder with Visual Studio Code or your prefered Code Editor
4) Open the console (Git, Windows Shell or the one you use) and move to the root directory of the program using the cd command.
5) Once in the root directory, we need to pull all the necessary packages for the program. In the console put this command: go mod tidy 
6) Once all the packages where downloaded, press Alt + s (if you installed the Go extension in VSC), this will save and format and automatically add the packages in the import statement at the start of each .go file. Otherwise, you must add it manually.

## Project Structure

The folder structure goes as follows with a brief description of each functionality:
word_counter 				--> main folder
	-texts 				--> contains different .txt files you can use as input. Comes with the number of the most repeated words.
	-main.go 				--> initilizes the program
	-word_counter.go 		--> provides the functions to handle each of the different types of inputs as argument or from stdin
	-word_counter_test.go  --> provides differente test scenarios for the word_counter.go functions
	-utils.go 			--> provides tranversal utility functions: merge texts, clean texts, get repeated words, sort and display the result
	-utils_test.go 		--> test the previous functions mentioned
	-word_counter.exe		--> executable we are going to use run the app.

## Usage

We have two types of inputs to handle:
1) From command line arguments. For example in the console: ./executable arg1 arg2 ... argN 
This method will take as argument all of the commands after the executable.

How to try it:
1.2 Open the console (git bash, windows shell, etc). Move to the root directory where is the executable word_counter.exe located.
1.3 Look into the texts directory the file names.
1.2 In console type the command: ./word_counter ./texts/"the_text_name_1" ./texts/"the_text_name_2" and so on...
1.3 In the console will be showed the 100 most repeated 3 words from all the .txt files passed as arguments.

2) From stdin (standar input). For example in the console: cat arg1 arg2 ... argN | ./executable
This method will concatenate the content of the arguments and then pass that content to the executable.

How to try it:
2.1 Same two first steps from the previous method.
2.2 In console type the command: cat ./texts/"the_text_name_1" ./texts/"the_text_name_2" and so on..
2.3 If you want to get all the texts from the same directory you can use: cat ./texts/*
2.4 See the 100 most repeated 3 words from all the .txt files processed from stdin.

## Testing
1) Open VSC, go to the _test.go file you want to test.
2) Look the functions, the descriptions and how it works
3) Click run test above the function you want to try
4) Or click in run file tests at the top of the _test.go file to run all the test of that file. 
5) Check the response from the console.

## Bonus and Extras
- Added gorutines to maximize speed for files as arguments inputs.
- Added configuration variables, so you can set:
	-number of consecutives word you want to consider (default is 3)
	-number of sequences showed in the top list (default 100)

## Future Iterations and comments

Implement a rule on how to handle when repeated sequence of words have the same value and are in conflict to appear in the last top words. 
For example, if the 100th and 101th sequence repeats the same times, define which rule to use to determine which key is considered and which excluded. Currently, I decided to order in alphabetic order for that case, from A to Z for readability.













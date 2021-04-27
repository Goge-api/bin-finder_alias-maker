import os

#//TODO make go version...but python is so much easier to make cmd line scripts
f = open("../binsFound.log", "r")

logArr = f.readline().split(',')

alias = logArr[3]

print(alias)
ali = f'alias{alias}'

# print()
os.system(ali)
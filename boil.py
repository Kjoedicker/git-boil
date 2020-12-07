#!/bin/python3

import os, sys

gitcommand="git remote set-url --add --push origin"

os.system(f"git remote add {sys.argv[1]}") 
for url in range(len(sys.argv)-1):
    os.sytem(f"{gitcommand} {os.argv[url]")
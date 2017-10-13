"""generateDics.py: Generates and compresses a dictionary of passwords and their statistical frequencies."""

__author__      = "Bernardo A. Rodrigues"
__copyright__   = "Copyright 2017, Universidade Federal de Goias, NExT"
__credits__     = ["Bernardo A. Rodrigues", "Wesley P. Calixto"]

__license__     = "GPLv2"
__version__     = "1.0.1"
__mantainer__   = "Bernardo Rodrigues"
__email__       = "bernardoaraujor@gmail.com"
__status__      = "development"

##############################################################
##############################################################

import sys
import getopt
import collections

##############################################################
# loadFile(path)
# description
# inputs:
# returns:


def main(argv):
        inputFile = argv[1]

        d = {}
        for line in open(inputFile, 'r', encoding="utf-8"):
                line = line.split('\t')
                freq = int(line[0])
                password = line[1][0:-1]
                
                d[password] = freq
 
                print(d['crazy1'])



if __name__ == "__main__":
        main(sys.argv)


import subprocess
import gzip
import csv
import pandas as pd

import os

password = "test"
args = ("java", "-jar", "/home/bernardo/IdeaProjects/passfault_corinda2/out/artifacts/passfault_corinda_jar/passfault_corinda2.jar")

popen = subprocess.Popen(args, stdout=subprocess.PIPE, stdin=subprocess.PIPE)

data = pd.read_csv('../training/input/rockyou.csv.gz', compression='gzip',
                   error_bad_lines=False)

for p in data.password:
    out, err = popen.communicate(input='bernardo'.encode())
    popen.wait();
    print(out)
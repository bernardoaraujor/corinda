import threading
import time
import os, sys
import pandas as pd
from model_recognition.model_recognition import ModelRecognition

file_dir = os.path.dirname(__file__)
sys.path.append(file_dir)

recognizer = ModelRecognition()

# threads = []
# vals = []
#
# for i in range(0, 10):
#     t = threading.Thread(target=recognizer.analyzeThread, args=('password', vals, ))
#     threads.append(t)
#     t.start()
#
print("loading input file")
data = pd.read_csv('input/rockyou.csv.gz', compression='gzip',
                   error_bad_lines=False)

n_threads = 1000
threads = []
n_rows, _ = data.shape

model_dict = {}

count_lock = threading.Lock()
count = 0

for i in range(0, n_threads):
    slice_size = n_rows/n_threads

    #TODO: fix slicing
    slice = data.loc[i*slice_size:(i+1)*slice_size]
    t = threading.Thread(target=recognizer.analyzer_thread, args=(slice, model_dict, ))
    threads.append(t)
    t.start()
    print("starting thread number " + str(i))

while(1):
    count = recognizer.get_count()
    time.sleep(1)
    count_ = recognizer.get_count()
    print('total evaluated: ' + str(100*count/n_rows) + '\tspeed: ' + str(100*(count_ - count)/n_rows) + '% / s')
# for index, row in data.iterrows():
#     # if (row['password'] == 'NULL'):
#     #    a = 1
#     print(str(index) + '\t' + str(index / data.size) + '\t' + recognizer.analyze(row['password']))
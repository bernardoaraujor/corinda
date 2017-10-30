from model_recognition.model_recognition import ModelRecognition
import threading
import pandas as pd
import time

recognizer = ModelRecognition()

# threads = []
# vals = []
#
# for i in range(0, 10):
#     t = threading.Thread(target=recognizer.analyzeThread, args=('password', vals, ))
#     threads.append(t)
#     t.start()
#

data = pd.read_csv('input/rockyou.csv.gz', compression='gzip',
                   error_bad_lines=False)

n_threads = 10
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

while(1):
    time.sleep(1)
    print(count)
# for index, row in data.iterrows():
#     # if (row['password'] == 'NULL'):
#     #    a = 1
#     print(str(index) + '\t' + str(index / data.size) + '\t' + recognizer.analyze(row['password']))
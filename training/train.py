from model_recognition.model_recognition import ModelRecognition
import threading
from multiprocessing import Queue

recognizer = ModelRecognition()

threads = []
vals = []

def func(val):
    print(val)

for i in range(0, 10):
    q = Queue()

    t = threading.Thread(target=recognizer.analyzeThread, args=('password', vals, ))
    threads.append(t)
    t.start()

a
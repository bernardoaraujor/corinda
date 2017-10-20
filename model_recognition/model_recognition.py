import subprocess

password = "test"
args = ("java", "-jar", "/home/bernardo/PycharmProjects/corinda/model_recognition/passfault_corinda/out/artifacts/passfault_jar/passfault.jar")

popen = subprocess.Popen(args, stdout=subprocess.PIPE, stdin=subprocess.PIPE)

out, err = popen.communicate(input=''.encode())
print(out)
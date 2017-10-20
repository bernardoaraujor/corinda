import subprocess

password = "test"
args = ("passfault/commandLine/build/install/passfault/bin/passfault")

popen = subprocess.Popen(args, stdout=subprocess.PIPE, stdin=subprocess.PIPE)

out, err = popen.communicate(input='bernardo\n'.encode())
print(out)
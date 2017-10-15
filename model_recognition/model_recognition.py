import subprocess

password = "test"
args = ("passfault/commandLine/build/install/passfault/bin/passfault", "-p", password)

popen = subprocess.Popen(args, stdout=subprocess.PIPE)
popen.wait()
output = popen.stdout.read()
print(output)
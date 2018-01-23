#!/bin/sed -f
#
# tex2latin - script sed que converte acentos latex para latin
# Copyright (C) 2001 Elgio Schlemer
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
#
# Elgio Schlemer  02/Out/2001
# Exemplo:   "\'a"  ->  "�"
#
# Nao remove as { } opcionais.
#             Exemplo:  "{\'a}" ->  "{�}"
#
# Forma de execu��o:
# cat arquivo.tex | tex2latin
# Joga na saida padrao o arquivo convertido
#
# Ou 
# tex2latin arquivo.tex
# Converte arquivo.tex e joga o resultado na saida padrao.
#
# ATENCAO: JAMAIS FA�A:
#           tex2latin arquivo.tex > arquivo.tex
#           cat arquivo.tex | tex2latin > arquivo.tex
# Isto farah voce PERDER o arquivo!!!

# Vogais, acento agudo
s/\\'a/�/g
s/\\'e/�/g
s/\\'\\i/�/g
s/\\'o/�/g
s/\\'u/�/g

s/\\'A/�/g
s/\\'E/�/g
s/\\'I/�/g
s/\\'O/�/g
s/\\'U/�/g

# Vogais, acento crase
s/\\`a/�/g
s/\\`e/�/g
s/\\`i/�/g
s/\\`o/�/g
s/\\`u/�/g

s/\\`A/�/g
s/\\`E/�/g
s/\\`I/�/g
s/\\`O/�/g
s/\\`U/�/g

# Vogais, til
s/\\~a/�/g
s/\\~o/�/g
s/\\~n/�/g

s/\\~A/�/g
s/\\~O/�/g
s/\\~N/�/g

# Cedilha
s/\\c c/�/g
s/\\c{c}/�/g
s/\\c C/�/g
s/\\c{C}/�/g

# Vogais, acento cincunflexo
s/\\^a/�/g
s/\\^e/�/g
s/\\^i/�/g
s/\\^o/�/g
s/\\^u/�/g

s/\\^A/�/g
s/\\^E/�/g
s/\\^I/�/g
s/\\^O/�/g
s/\\^U/�/g


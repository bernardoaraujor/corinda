#!/bin/sed -f
#
# latin2tex - script sed que converte acentos latin para comandos TeX
# Copyright (C) 2002 Rafael �vila
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
# -----------------------------------------------
# latin2tex v1.0
# uso: tex2latin < in.latin > out.tex
# baseado no script `tex2latin' de Elgio Schlemer
#
# a vers�o mais recente de ambos pode ser encontrada em
# http://www.inf.ufrgs.br/utug/download

# Vogais, acento agudo
s/�/{\\'a}/g
s/�/{\\'e}/g
s/�/{\\'\\i}/g
s/�/{\\'o}/g
s/�/{\\'u}/g

s/�/{\\'A}/g
s/�/{\\'E}/g
s/�/{\\'I}/g
s/�/{\\'O}/g
s/�/{\\'U}/g

# Vogais, acento crase
s/�/{\\`a}/g
s/�/{\\`e}/g
s/�/{\\`i}/g
s/�/{\\`o}/g
s/�/{\\`u}/g

s/�/{\\`A}/g
s/�/{\\`E}/g
s/�/{\\`I}/g
s/�/{\\`O}/g
s/�/{\\`U}/g

# Vogais, til
s/�/{\\~a}/g
s/�/{\\~o}/g
s/�/{\\~n}/g

s/�/{\\~A}/g
s/�/{\\~O}/g
s/�/{\\~N}/g

# Cedilha
s/�/{\\c{c}}/g
s/�/{\\c{c}}/g
s/�/{\\c{C}}/g
s/�/{\\c{C}}/g

# Vogais, acento cincunflexo
s/�/{\\^a}/g
s/�/{\\^e}/g
s/�/{\\^i}/g
s/�/{\\^o}/g
s/�/{\\^u}/g

s/�/{\\^A}/g
s/�/{\\^E}/g
s/�/{\\^I}/g
s/�/{\\^O}/g
s/�/{\\^U}/g


#LATEX TO PDF (INPE)
#-------------------
#Makefile para criação de documentos [PDF]
#utilizando o processador de textos LaTeX.
#
#Bruno A.F. Roth                    v1.0 [04/2002] 
#Bruno A.F. Roth (roth@lac.inpe.br) v1.1 [04/2005] +55(12)3945-6836
#Mantido por Alan Wilter S .da Silva v1.3a [29/07/2005] +55(12)3945-6911
#alan@sid.inpe.br


BASENAME=publicacao

PROGRAM_BIB=bibtex
PROGRAM_PDFLATEX=pdflatex
PROGRAM_LATEX=latex
PROGRAM_HTLATEX=htlatex
PROGRAM_MAKEINDEX=makeindex

help:
	@cat LEIAME

pdf:	
	${PROGRAM_PDFLATEX} ${BASENAME}.tex
	${PROGRAM_BIB} ${BASENAME}
	${PROGRAM_MAKEINDEX} ${BASENAME}
	${PROGRAM_PDFLATEX} ${BASENAME}.tex
	${PROGRAM_PDFLATEX} ${BASENAME}.tex
	${PROGRAM_PDFLATEX} ${BASENAME}.tex

dvi:	
	${PROGRAM_LATEX} ${BASENAME}.tex
	${PROGRAM_BIB} ${BASENAME}
	${PROGRAM_MAKEINDEX} ${BASENAME}
	${PROGRAM_LATEX} ${BASENAME}.tex
	${PROGRAM_LATEX} ${BASENAME}.tex
	${PROGRAM_LATEX} ${BASENAME}.tex

html:	
	${PROGRAM_HTLATEX} ${BASENAME}.tex
	${PROGRAM_BIB} ${BASENAME}
	${PROGRAM_MAKEINDEX} ${BASENAME}
	${PROGRAM_HTLATEX} ${BASENAME}.tex

ps:
	make dvi
	dvips ${BASENAME}.dvi -o ${BASENAME}.ps

pdf1:	
	${PROGRAM_PDFLATEX} ${BASENAME}.tex

clean:
	rm -f *.aux *.log *.toc *.lot *.lof *.out *.blg *.bbl *.idx *.ilg *.ind *.brf
	rm -f `find . -name "*.log"`
	rm -f `find . -name "*.aux"`
	rm -f *.tmp *.css *.lg *.idv *.4tc *.xref *.4ct *.png *.html
all:
	make clean
	make pdf
	
